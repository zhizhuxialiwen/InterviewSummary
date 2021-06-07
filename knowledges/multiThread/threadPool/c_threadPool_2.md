# c语言线程池

<!-- GFM-TOC -->
* [1线程池原理剖析](#1线程池原理剖析)
* [2线程池的结构体描述信息](#2线程池的结构体描述信息)
* [3线程池的各个函数解析](#3线程池的各个函数解析)
* [4线程池完整的头文件和实现文件](#4线程池完整的头文件和实现文件)
<!-- GFM-TOC -->

## 1线程池原理剖析

### 1.1 为何需要线程池？

我们知道，多路IO转接是可以支持多个客户端的连接与处理，举例poll解释，他可以一次监听到多个客户端请求返回，并遍历confd数组进行处理，也就是说，Server是直接进行处理的;再看epoll，他是有请求就给你建立连接，并且设置为非阻塞是非常高效的，不需要立马就处理请求。因为它有回调函数帮忙进行处理，但实际上仍然是Server自己处理。如果客户端非常活跃的情况下，导致很多回调函数处理，也是非常耗时的，所以epoll比较支持多连接，少活跃的情况。
那么我们能不能建立连接后的事情处理不放在主线程Server处理呢，所以这就是我们需要使用到线程池的原因。

2 图形帮助理解

![threadPool1](../../../images/threadPool1.PNG)

后面的章节将会分析线程池结构体成员信息、线程池各个函数描述与注释。

## 2、线程池的结构体描述信息


直接看代码，代码里有详细的注释。

### 2.1 描述任务队列的结构体

```c++
typedef struct {
    void *(*function)(void *);          /* 函数指针，回调函数 */
    void *arg;                          /* 上面函数的参数 */
} threadpool_task_t;                    /* 各子线程任务结构体 */
```

### 2.2 描述线程池相关信息 

```c++
struct threadpool_t {
    //两把锁与两个条件变量
    pthread_mutex_t lock;               /* 用于锁住本结构体 */    
    pthread_mutex_t thread_counter;     /* 记录忙状态线程个数的琐 --busy_thr_num */

    pthread_cond_t queue_not_full;      /* 当任务队列满时，添加任务的线程(server主线程)阻塞，等待此条件变量.用于S/C之间 */
    pthread_cond_t queue_not_empty;     /* 任务队列里不为空时，通知等待任务的线程.用于Ser即主线程与各个子线程之间. */

    //线程id
    pthread_t *threads;                 /* 存放线程池中每个线程的tid。数组 */
    pthread_t adjust_tid;               /* 存管理线程tid */

    //任务队列,记录客户端的任务请求
    threadpool_task_t *task_queue;      /* 任务队列 */

    //线程个数的范围、实际存在的线程、正在工作的线程、准备销毁的线程
    int min_thr_num;                    /* 线程池最小线程数 */
    int max_thr_num;                    /* 线程池最大线程数 */
    int live_thr_num;                   /* 当前存活线程个数 */
    int busy_thr_num;                   /* 忙状态线程个数 */
    int wait_exit_thr_num;              /* 要销毁的线程个数 */

    //任务队列的头尾指针、队列的实际大小、队列的上限个数
    int queue_front;                    /* task_queue队头下标 */
    int queue_rear;                     /* task_queue队尾下标 */
    int queue_size;                     /* task_queue队中实际任务数 */
    int queue_max_size;                 /* task_queue队列可容纳任务数上限 */

    int shutdown;                       /* 标志位，线程池使用状态，true或false */
};
```

## 3、线程池的各个函数解析
### 3.1 业务层逻辑分析

main函数中比较简单，就是创建线程池，然后往池里添加任务，等待处理完成后销毁线程池。
下面会根据main的调用顺序来逐步解析各个函数的作用。

```c++
/* 线程池中的线程，模拟处理业务 */
void *process(void *arg)
{
    printf("thread 0x%x working on task %d\n ",(unsigned int)pthread_self(),*(int *)arg);
    sleep(1);
    printf("task %d is end\n",*(int *)arg);

    return NULL;
}

int main(void)
{
    /*threadpool_t *threadpool_create(int min_thr_num, int max_thr_num, int queue_max_size);*/

	//1 先创建线程池
    threadpool_t *thp = threadpool_create(3,100,100);/*创建线程池，池里最小3个线程，最大100，队列最大100*/
    printf("pool inited");

    //int *num = (int *)malloc(sizeof(int)*20);
	//2 模拟客户端请求的任务20个
    int num[20], i;
    for (i = 0; i < 20; i++) {
        num[i]=i;
        printf("add task %d\n",i);
        threadpool_add(thp, process, (void*)&num[i]);     /* 向线程池中添加任务 */
    }
	
	//3 等子线程完成任务
    sleep(10);                                          /* 等子线程完成任务 */
	
	//4 销毁线程池
    threadpool_destroy(thp);

    return 0;
}
```

### 3.2 线程池创建函数的解析

看代码即可，里面有详细的注释。

```c++
threadpool_t *threadpool_create(int min_thr_num, int max_thr_num, int queue_max_size)
{
    int i;
	//1 首先创建线程池;用dowhile代替goto,一旦出错直接threadpool_free(pool);
    threadpool_t *pool = NULL;
    do {
        if((pool = (threadpool_t *)malloc(sizeof(threadpool_t))) == NULL) {  
            printf("malloc threadpool fail");
            break;/*跳出do while*/
        }
		
		//2 对线程相关的int成员赋值
        pool->min_thr_num = min_thr_num;
        pool->max_thr_num = max_thr_num;
        pool->busy_thr_num = 0;
        pool->live_thr_num = min_thr_num;               /* 活着的线程数 初值=最小线程数 */
		//3 对队列相关的int成员赋值
        pool->queue_size = 0;                           /* 有0个产品 */
        pool->queue_max_size = queue_max_size;
        pool->queue_front = 0;
        pool->queue_rear = 0;
        pool->shutdown = false;                         /* 不关闭线程池 */

        /* 根据最大线程上限数， 给工作线程数组开辟空间, 并清零 */
		//4 对线程池的各个线程开辟空间(实际为装着int型的tid),并清零
        pool->threads = (pthread_t *)malloc(sizeof(pthread_t)*max_thr_num); 
        if (pool->threads == NULL) {
            printf("malloc threads fail");
            break;
        }
        memset(pool->threads, 0, sizeof(pthread_t)*max_thr_num);

        /* 队列开辟空间 */
		//5 对线程池的任务队列开辟内存
        pool->task_queue = (threadpool_task_t *)malloc(sizeof(threadpool_task_t)*queue_max_size);
        if (pool->task_queue == NULL) {
            printf("malloc task_queue fail");
            break;
        }

        /* 初始化互斥琐、条件变量 */
		//6 用静态方法初始化两把锁和两个条件变量
        if (pthread_mutex_init(&(pool->lock), NULL) != 0
                || pthread_mutex_init(&(pool->thread_counter), NULL) != 0
                || pthread_cond_init(&(pool->queue_not_empty), NULL) != 0
                || pthread_cond_init(&(pool->queue_not_full), NULL) != 0)
        {
            printf("init the lock or cond fail");
            break;
        }

        /* 启动 min_thr_num 个 work thread */
		//7 依次创建各个线程,并使各个线程统一回调工作线程函数
        for (i = 0; i < min_thr_num; i++) {
            pthread_create(&(pool->threads[i]), NULL, threadpool_thread, (void *)pool);/*pool指向当前线程池*/
            printf("start thread 0x%x...\n", (unsigned int)pool->threads[i]);
        }
		//8 额外开辟管理线程,回调函数为管理函数
        pthread_create(&(pool->adjust_tid), NULL, adjust_thread, (void *)pool);/* 启动管理者线程 */

        return pool;

    } while (0);

	//9 出错则free之前对应的空间
    threadpool_free(pool);      /* 前面代码调用失败时，释放poll存储空间 */

    return NULL;
}
```

## 3.3 往线程池添加任务的函数解析

看代码即可，里面有详细的注释。

```c++
/* 向线程池中 添加一个任务 */
//参数为任务的两个成员
int threadpool_add(threadpool_t *pool, void*(*function)(void *arg), void *arg)
{
	//1 先给线程池上锁
    pthread_mutex_lock(&(pool->lock));

    /* ==为真，队列已经满， 调wait阻塞 */
	//2 队列为满，阻塞用户请求.不管队列为满还是为空,都是通过pool->queue_size判断，从而调用相应条件变量阻塞
    while ((pool->queue_size == pool->queue_max_size) && (!pool->shutdown)) {
        pthread_cond_wait(&(pool->queue_not_full), &(pool->lock));//作用:阻塞，解锁，上锁
    }
	//3 若为真，解锁线程池，其它线程此时有机会在线程池操作(实际上是被诱导主动销毁).下面的任务正常添加.
    if (pool->shutdown) {
        pthread_mutex_unlock(&(pool->lock));
    }

    /* 清空工作线程 调用回调函数、参数arg */
	//4 清空工作线程 调用回调函数、参数arg
    if (pool->task_queue[pool->queue_rear].arg != NULL) {
        //free(pool->task_queue[pool->queue_rear].arg);//不能释放临时变量
        pool->task_queue[pool->queue_rear].arg = NULL;
    }
    /*添加任务到任务队列里*/
	//5 任务入队；与队头操作一样.例如队头=1,队尾=3,入队后+1,队尾=4,人数添加了一个.
    pool->task_queue[pool->queue_rear].function = function;
    pool->task_queue[pool->queue_rear].arg = arg;
    pool->queue_rear = (pool->queue_rear + 1) % pool->queue_max_size;       /* 队尾指针移动, 模拟环形 */
    pool->queue_size++;

    /*添加完任务后，队列不为空，唤醒线程池中 等待处理任务的线程*/
	//6 队列不为空,发信号唤醒空闲的线程工作
    pthread_cond_signal(&(pool->queue_not_empty));
	
	//7 解锁线程池
    pthread_mutex_unlock(&(pool->lock));

    return 0;
}
```

### 3.4 线程池中各个线程的工作函数解析

看代码即可，里面有详细的注释。

```c++
/* 线程池中各个工作线程 */
//处理线程池结构体的每一个成员都需要上锁解锁!!!
void *threadpool_thread(void *threadpool)
{
	//1 先拿到线程池和要处理的单个任务
    threadpool_t *pool = (threadpool_t *)threadpool;
    threadpool_task_t task;

    while (true) {
        /* Lock must be taken to wait on conditional variable */
        /*刚创建出线程，等待任务队列里有任务，否则阻塞等待任务队列里有任务后再唤醒接收任务*/
		//2 锁住线程池
        pthread_mutex_lock(&(pool->lock));

        /*queue_size == 0 说明没有任务，调 wait 阻塞在条件变量上, 若有任务，跳过该while*/
		//3 若队列无任务,则阻塞在queue_not_empty上，此时该线程为空闲线程;管理线程根据算法是否结束该线程
        while ((pool->queue_size == 0) && (!pool->shutdown)) {  
            printf("thread 0x%x is waiting\n", (unsigned int)pthread_self());
            pthread_cond_wait(&(pool->queue_not_empty), &(pool->lock));

            /*清除指定数目的空闲线程，如果要结束的线程个数大于0，结束线程*/
            if (pool->wait_exit_thr_num > 0) {//pool->wait_exit_thr_num--放在下面的if条件上面是因为,始终要保持最低的线程数,
                pool->wait_exit_thr_num--;    //如果存活的线程数等于最低,则不必再exit，但我们需要人为的wait_exit_thr_num--，确保本次清除空闲线程的个数

                /*如果线程池里线程个数大于最小值时可以结束当前线程*/
                if (pool->live_thr_num > pool->min_thr_num) {
                    printf("thread 0x%x is exiting\n", (unsigned int)pthread_self());
                    pool->live_thr_num--;
                    pthread_mutex_unlock(&(pool->lock));
                    pthread_exit(NULL);//所有线程的退出都是主动调用该函数退出
                }
            }
        }

        /*如果指定了true，要关闭线程池里的每个线程，自行退出处理*/
		//4 如果要销毁线程池，解掉线程池的锁，然后销毁该线程(多个线程依次调用，故不能写成for来pthread_exit)
        if (pool->shutdown) {
            pthread_mutex_unlock(&(pool->lock));
            printf("thread 0x%x is exiting\n", (unsigned int)pthread_self());
            pthread_exit(NULL);     /* 线程自行结束 */
        }

        /*从任务队列里获取任务, 是一个出队操作*/
		//5 执行到这里,说明队列有任务,然后线程池取出该任务的回调函数和参数(并未开始处理，只是取出)
        task.function = pool->task_queue[pool->queue_front].function;
        task.arg = pool->task_queue[pool->queue_front].arg;

		//6 出队并且队列减1.出队的理解:例如队头=1(排名第一),加1后取余变成2,即现在队头变成了第二名，参考排队。
        pool->queue_front = (pool->queue_front + 1) % pool->queue_max_size;       /* 出队，模拟环形队列 */
        pool->queue_size--;

        /*通知可以有新的任务添加进来*/
		//7 上面处理完任务后(注：只是处理完任务，真正的事情还未处理,故在下面才加busy_thr_num),通知Server主线程不必再阻塞，可以让客户端请求
        pthread_cond_broadcast(&(pool->queue_not_full));

        /*任务取出后，立即将 线程池琐 释放*/
		//8 任务取出即可释放线程池锁
        pthread_mutex_unlock(&(pool->lock));

        /*执行任务*/ 
		//9 执行任务,真正的任务事情执行忙状态线程数+1，并执行回调
        printf("thread 0x%x start working\n", (unsigned int)pthread_self());
        pthread_mutex_lock(&(pool->thread_counter));                            /*忙状态线程数变量琐*/
        pool->busy_thr_num++;                                                   /*忙状态线程数+1*/
        pthread_mutex_unlock(&(pool->thread_counter));
        (*(task.function))(task.arg);                                           /*执行回调函数任务*/
        //task.function(task.arg);                                              /*执行回调函数任务*/

        /*任务结束处理*/ 
		//10 任务结束处理,忙状态数线程数-1
        printf("thread 0x%x end working\n", (unsigned int)pthread_self());
        pthread_mutex_lock(&(pool->thread_counter));
        pool->busy_thr_num--;                                       /*处理掉一个任务，忙状态数线程数-1*/
        pthread_mutex_unlock(&(pool->thread_counter));
    }

	//11 结束线程,但看本函数逻辑,它实际上是走不出while的
    pthread_exit(NULL);
}
```

### 3.5 线程池中的管理线程函数的解析

看代码即可，里面有详细的注释。

```c++
/* 管理线程 */
//主要按照busy_thr_num与live_thr_num的比例(时间也行)对线程池进行伸缩
void *adjust_thread(void *threadpool)
{
    int i;
	//1 先获取线程池
    threadpool_t *pool = (threadpool_t *)threadpool;
    while (!pool->shutdown) {

		//2 不用每秒都管理各个线程,定时管理即可,一般为10s
        sleep(DEFAULT_TIME);                                    /*定时 对线程池管理*/

		//3 锁住线程池(除了创建线程池时不用,其它访问线程池的基本都要),获取存活线程数和实际队列数
        pthread_mutex_lock(&(pool->lock));
        int queue_size = pool->queue_size;                      /* 关注 任务数 */
        int live_thr_num = pool->live_thr_num;                  /* 存活 线程数 */
        pthread_mutex_unlock(&(pool->lock));

		//4 锁住忙线程，获取忙线程数
        pthread_mutex_lock(&(pool->thread_counter));
        int busy_thr_num = pool->busy_thr_num;                  /* 忙着的线程数 */
        pthread_mutex_unlock(&(pool->thread_counter));

        /* 创建新线程 算法： 任务数大于最小线程池个数, 且存活的线程数少于最大线程个数时 如：30>=10 && 40<100*/
		//5 增加线程,当满足算法时(这个算法可以自己定义)
        if (queue_size >= MIN_WAIT_TASK_NUM && live_thr_num < pool->max_thr_num) {
            pthread_mutex_lock(&(pool->lock));  
            int add = 0;

            /*一次增加 DEFAULT_THREAD 个线程*/
            for (i = 0; i < pool->max_thr_num && add < DEFAULT_THREAD_VARY
                    && pool->live_thr_num < pool->max_thr_num; i++) {
                if (pool->threads[i] == 0 || !is_thread_alive(pool->threads[i])) {
                    pthread_create(&(pool->threads[i]), NULL, threadpool_thread, (void *)pool);
                    add++;
                    pool->live_thr_num++;
                }
            }

            pthread_mutex_unlock(&(pool->lock));
        }

        /* 销毁多余的空闲线程 算法：忙线程X2 小于 存活的线程数 且 存活的线程数 大于 最小线程数时*/
		//6 销毁空闲线程，,当满足算法时(这个算法可以自己定义)；(busy_thr_num * 2) < live_thr_num表示至少有一半的线程在存活中是空闲的.
		//后面的live_thr_num > pool->min_thr_num表示销毁只需要考虑最低线程数,不需要考虑最大线程数(增加时需要)
        if ((busy_thr_num * 2) < live_thr_num  &&  live_thr_num > pool->min_thr_num) {

            /* 一次销毁DEFAULT_THREAD个线程, 隨機10個即可 */
            pthread_mutex_lock(&(pool->lock));
            pool->wait_exit_thr_num = DEFAULT_THREAD_VARY;      /* 要销毁的线程数 设置为10 */
            pthread_mutex_unlock(&(pool->lock));

            for (i = 0; i < DEFAULT_THREAD_VARY; i++) {
                /* 通知处在空闲状态的线程, 他们会自行终止*/
                pthread_cond_signal(&(pool->queue_not_empty));
            }
        }
    }

    return NULL;
}
```

### 3.6 线程池的销毁函数的解析

看代码即可，里面有详细的注释。

```c++
//线程池的销毁函数
int threadpool_destroy(threadpool_t *pool)
{
    int i;
    if (pool == NULL) {
        return -1;
    }
	//1 标志为改为销毁
    pool->shutdown = true;

    /*先销毁管理线程*/
	//2 先销毁额外的管理线程
    pthread_join(pool->adjust_tid, NULL);//该函数与pthread_exit区别是,join会等待tid的子线程释放资源而退出,即在一个线程结束其它线程;而pthread_exit只会退出当前线程.

	//3 通知所有空闲的线程结束，正常工作的不受影响，因为并没有阻塞在该条件变量，等处理完后就会进入shutdown = true被结束掉
    for (i = 0; i < pool->live_thr_num; i++) {
        /*通知所有空闲的线程*/
        pthread_cond_broadcast(&(pool->queue_not_empty));
    }
	//4 这里可以认为对上面仍在工作的线程没有接到信号并没退出,这里重新退出，已经退出的则内部直接返回
    for (i = 0; i < pool->live_thr_num; i++) {
        pthread_join(pool->threads[i], NULL);
    }
	
	//5 最后释放掉线程池自己(即线程池里的成员,看下面的函数)
    threadpool_free(pool);

    return 0;
}
/*
根据销毁函数，总结一下pthread_join,pthread_exit,exit,return的区别:
pthread_join：在本线程中结束其它线程，可以回收子线程资源。
pthread_exit：在本线程调用，就结束本线程(真正意义的结束本线程)。特别强调，主线程调用时，也只会结束主线程，其它线程不会结束。与retuen区分开来。
exit:无论在哪个调用线程调用，都会结束整个进程。程序结束。
return：在main返回时，因为main是主线程，所以整个进程结束,其它线程也被结束(并非真正的结束本线程)。其它线程return时，只会结束return的子线程。
*/
```

### 3.7 释放线程池描述结构体内部成员的函数的解析

看代码即可，里面有详细的注释。

```c++
//释放线程池内部成员的函数
int threadpool_free(threadpool_t *pool)
{
    if (pool == NULL) {
        return -1;
    }
	//1 释放认为队列空间
    if (pool->task_queue) {
        free(pool->task_queue);
    }
	//2 释放线程tid数组和两把锁两个条件变量
    if (pool->threads) {
        free(pool->threads);
        pthread_mutex_lock(&(pool->lock));
        pthread_mutex_destroy(&(pool->lock));
        pthread_mutex_lock(&(pool->thread_counter));
        pthread_mutex_destroy(&(pool->thread_counter));
        pthread_cond_destroy(&(pool->queue_not_empty));
        pthread_cond_destroy(&(pool->queue_not_full));
    }
	//3 最后释放线程池自己
    free(pool);
    pool = NULL;

    return 0;
}
```

### 3.8 其它一些获取线程池信息的函数

```c++
//获取线程池里存活的线程数
int threadpool_all_threadnum(threadpool_t *pool)
{
    int all_threadnum = -1;
    pthread_mutex_lock(&(pool->lock));
    all_threadnum = pool->live_thr_num;
    pthread_mutex_unlock(&(pool->lock));
    return all_threadnum;
}

//获取忙线程数
int threadpool_busy_threadnum(threadpool_t *pool)
{
    int busy_threadnum = -1;
    pthread_mutex_lock(&(pool->thread_counter));
    busy_threadnum = pool->busy_thr_num;
    pthread_mutex_unlock(&(pool->thread_counter));
    return busy_threadnum;
}

//测试某个线程释放存活
int is_thread_alive(pthread_t tid)
{
    int kill_rc = pthread_kill(tid, 0);     //发0号信号，测试线程是否存活
    if (kill_rc == ESRCH) {
        return false;
    }

    return true;
}
```

### 3.9 总结

到此为止，我们已经将线程池的原理、描述结构体、各个函数了解完毕。

## 4、线程池完整的头文件和实现文件(.c)

前提：这里的线程池是基于Linux下C的实现，且在任务的回调函数中只是模拟工作，具体业务需要自己根据实际编写。这里可以举个例子，当客户端请求过来，我们可以将建立好的通信套接字通过任务的void* arg传进来，这样我们就可以在子线程中与客户端进行交互，极大的减少服务器主线程的工作压力。

### 4.1 头文件--threadPool.h

```c++

#ifndef __THREADPOOL_H_
#define __THREADPOOL_H_

typedef struct threadpool_t threadpool_t;

/**
 * @function threadpool_create
 * @descCreates a threadpool_t object.
 * @param thr_num  thread num
 * @param max_thr_num  max thread size
 * @param queue_max_size   size of the queue.
 * @return a newly created thread pool or NULL
 */
threadpool_t *threadpool_create(int min_thr_num, int max_thr_num, int queue_max_size);

/**
 * @function threadpool_add
 * @desc add a new task in the queue of a thread pool
 * @param pool     Thread pool to which add the task.
 * @param function Pointer to the function that will perform the task.
 * @param argument Argument to be passed to the function.
 * @return 0 if all goes well,else -1
 */
int threadpool_add(threadpool_t *pool, void*(*function)(void *arg), void *arg);

/**
 * @function threadpool_destroy
 * @desc Stops and destroys a thread pool.
 * @param pool  Thread pool to destroy.
 * @return 0 if destory success else -1
 */
int threadpool_destroy(threadpool_t *pool);

/**
 * @desc get the thread num
 * @pool pool threadpool
 * @return # of the thread
 */
int threadpool_all_threadnum(threadpool_t *pool);

/**
 * desc get the busy thread num
 * @param pool threadpool
 * return # of the busy thread
 */
int threadpool_busy_threadnum(threadpool_t *pool);

#endif
```

### 4.2 实现文件--threadPool.c

注：若最底下的一些描述文件没有在.h声明，自己可以添加声明即可。

```c++
#include <stdlib.h>
#include <pthread.h>
#include <unistd.h>
#include <assert.h>
#include <stdio.h>
#include <string.h>
#include <signal.h>
#include <errno.h>
#include "threadpool.h"

#define DEFAULT_TIME 10                 /*10s检测一次*/
#define MIN_WAIT_TASK_NUM 10            /*如果queue_size > MIN_WAIT_TASK_NUM 添加新的线程到线程池*/ 
#define DEFAULT_THREAD_VARY 10          /*每次创建和销毁线程的个数*/
#define true 1
#define false 0

typedef struct {
    void *(*function)(void *);          /* 函数指针，回调函数 */
    void *arg;                          /* 上面函数的参数 */
} threadpool_task_t;                    /* 各子线程任务结构体 */

/* 描述线程池相关信息 */
struct threadpool_t {
    pthread_mutex_t lock;               /* 用于锁住本结构体 */    
    pthread_mutex_t thread_counter;     /* 记录忙状态线程个数de琐 -- busy_thr_num */
    pthread_cond_t queue_not_full;      /* 当任务队列满时，添加任务的线程阻塞，等待此条件变量 */
    pthread_cond_t queue_not_empty;     /* 任务队列里不为空时，通知等待任务的线程 */

    pthread_t *threads;                 /* 存放线程池中每个线程的tid。数组 */
    pthread_t adjust_tid;               /* 存管理线程tid */
    threadpool_task_t *task_queue;      /* 任务队列 */

    int min_thr_num;                    /* 线程池最小线程数 */
    int max_thr_num;                    /* 线程池最大线程数 */
    int live_thr_num;                   /* 当前存活线程个数 */
    int busy_thr_num;                   /* 忙状态线程个数 */
    int wait_exit_thr_num;              /* 要销毁的线程个数 */

    int queue_front;                    /* task_queue队头下标 */
    int queue_rear;                     /* task_queue队尾下标 */
    int queue_size;                     /* task_queue队中实际任务数 */
    int queue_max_size;                 /* task_queue队列可容纳任务数上限 */

    int shutdown;                       /* 标志位，线程池使用状态，true或false */
};

/**
 * @function void *threadpool_thread(void *threadpool)
 * @desc the worker thread
 * @param threadpool the pool which own the thread
 */
void *threadpool_thread(void *threadpool);

/**
 * @function void *adjust_thread(void *threadpool);
 * @desc manager thread
 * @param threadpool the threadpool
 */
void *adjust_thread(void *threadpool);

/**
 * check a thread is alive
 */
int is_thread_alive(pthread_t tid);
int threadpool_free(threadpool_t *pool);

threadpool_t *threadpool_create(int min_thr_num, int max_thr_num, int queue_max_size)
{
    int i;
    threadpool_t *pool = NULL;
    do {
        if((pool = (threadpool_t *)malloc(sizeof(threadpool_t))) == NULL) {  
            printf("malloc threadpool fail");
            break;/*跳出do while*/
        }

        pool->min_thr_num = min_thr_num;
        pool->max_thr_num = max_thr_num;
        pool->busy_thr_num = 0;
        pool->live_thr_num = min_thr_num;               /* 活着的线程数 初值=最小线程数 */
        pool->queue_size = 0;                           /* 有0个产品 */
        pool->queue_max_size = queue_max_size;
        pool->queue_front = 0;
        pool->queue_rear = 0;
        pool->shutdown = false;                         /* 不关闭线程池 */

        /* 根据最大线程上限数， 给工作线程数组开辟空间, 并清零 */
        pool->threads = (pthread_t *)malloc(sizeof(pthread_t)*max_thr_num); 
        if (pool->threads == NULL) {
            printf("malloc threads fail");
            break;
        }
        memset(pool->threads, 0, sizeof(pthread_t)*max_thr_num);

        /* 队列开辟空间 */
        pool->task_queue = (threadpool_task_t *)malloc(sizeof(threadpool_task_t)*queue_max_size);
        if (pool->task_queue == NULL) {
            printf("malloc task_queue fail");
            break;
        }

        /* 初始化互斥琐、条件变量 */
        if (pthread_mutex_init(&(pool->lock), NULL) != 0
                || pthread_mutex_init(&(pool->thread_counter), NULL) != 0
                || pthread_cond_init(&(pool->queue_not_empty), NULL) != 0
                || pthread_cond_init(&(pool->queue_not_full), NULL) != 0)
        {
            //cond: condition 条件
            printf("init the lock or cond fail");
            break;
        }

        /* 启动 min_thr_num 个 work thread */
        for (i = 0; i < min_thr_num; i++) {
            pthread_create(&(pool->threads[i]), NULL, threadpool_thread, (void *)pool);/*pool指向当前线程池*/
            printf("start thread 0x%x...\n", (unsigned int)pool->threads[i]);
        }
        pthread_create(&(pool->adjust_tid), NULL, adjust_thread, (void *)pool);/* 启动管理者线程 */

        return pool;

    } while (0);

    threadpool_free(pool);      /* 前面代码调用失败时，释放poll存储空间 */

    return NULL;
}

/* 向线程池中 添加一个任务 */
int threadpool_add(threadpool_t *pool, void*(*function)(void *arg), void *arg)
{
    pthread_mutex_lock(&(pool->lock));

    /* ==为真，队列已经满， 调wait阻塞 */
    while ((pool->queue_size == pool->queue_max_size) && (!pool->shutdown)) {
        pthread_cond_wait(&(pool->queue_not_full), &(pool->lock));
    }
    if (pool->shutdown) {
    	pthread_cond_broadcast(&(pool->queue_not_empty));//唤醒多个线程。这句可以不写。
        pthread_mutex_unlock(&(pool->lock));
        return 0;//与上面的pthread_cond_broadcast可以不写，通过下面的操作进行唤醒，不过最好还是写
    }

    /* 清空 工作线程 调用的回调函数 的参数arg */
    if (pool->task_queue[pool->queue_rear].arg != NULL) {
        //free(pool->task_queue[pool->queue_rear].arg);//不能释放临时变量
        pool->task_queue[pool->queue_rear].arg = NULL;
    }
    /*添加任务到任务队列里*/
    pool->task_queue[pool->queue_rear].function = function;
    pool->task_queue[pool->queue_rear].arg = arg;
    pool->queue_rear = (pool->queue_rear + 1) % pool->queue_max_size;       /* 队尾指针移动, 模拟环形 */
    pool->queue_size++;

    /*添加完任务后，队列不为空，唤醒线程池中 等待处理任务的线程*/
    pthread_cond_signal(&(pool->queue_not_empty));//至少唤醒一个线程。
    pthread_mutex_unlock(&(pool->lock));

    return 0;
}

/* 线程池中各个工作线程 */
void *threadpool_thread(void *threadpool)
{
    threadpool_t *pool = (threadpool_t *)threadpool;
    threadpool_task_t task;

    while (true) {
        /* Lock must be taken to wait on conditional variable */
        /*刚创建出线程，等待任务队列里有任务，否则阻塞等待任务队列里有任务后再唤醒接收任务*/
        pthread_mutex_lock(&(pool->lock));

        /*queue_size == 0 说明没有任务，调 wait 阻塞在条件变量上, 若有任务，跳过该while*/
        while ((pool->queue_size == 0) && (!pool->shutdown)) {  
            printf("thread 0x%x is waiting\n", (unsigned int)pthread_self());
            pthread_cond_wait(&(pool->queue_not_empty), &(pool->lock));

            /*清除指定数目的空闲线程，如果要结束的线程个数大于0，结束线程*/
            if (pool->wait_exit_thr_num > 0) {
                pool->wait_exit_thr_num--;

                /*如果线程池里线程个数大于最小值时可以结束当前线程*/
                if (pool->live_thr_num > pool->min_thr_num) {
                    printf("thread 0x%x is exiting\n", (unsigned int)pthread_self());
                    pool->live_thr_num--;
                    pthread_mutex_unlock(&(pool->lock));
                    pthread_exit(NULL);
                }
            }
        }

        /*如果指定了true，要关闭线程池里的每个线程，自行退出处理*/
        if (pool->shutdown) {
            pthread_mutex_unlock(&(pool->lock));
            printf("thread 0x%x is exiting\n", (unsigned int)pthread_self());
            pthread_exit(NULL);     /* 线程自行结束 */
        }

        /*从任务队列里获取任务, 是一个出队操作*/
        task.function = pool->task_queue[pool->queue_front].function;
        task.arg = pool->task_queue[pool->queue_front].arg;

        pool->queue_front = (pool->queue_front + 1) % pool->queue_max_size;       /* 出队，模拟环形队列 */
        pool->queue_size--;

        /*通知可以有新的任务添加进来*/
        pthread_cond_broadcast(&(pool->queue_not_full));

        /*任务取出后，立即将 线程池琐 释放*/
        pthread_mutex_unlock(&(pool->lock));

        /*执行任务*/ 
        printf("thread 0x%x start working\n", (unsigned int)pthread_self());
        pthread_mutex_lock(&(pool->thread_counter));                            /*忙状态线程数变量琐*/
        pool->busy_thr_num++;                                                   /*忙状态线程数+1*/
        pthread_mutex_unlock(&(pool->thread_counter));
        (*(task.function))(task.arg);                                           /*执行回调函数任务*/
        //task.function(task.arg);                                              /*执行回调函数任务*/

        /*任务结束处理*/ 
        printf("thread 0x%x end working\n", (unsigned int)pthread_self());
        pthread_mutex_lock(&(pool->thread_counter));
        pool->busy_thr_num--;                                       /*处理掉一个任务，忙状态数线程数-1*/
        pthread_mutex_unlock(&(pool->thread_counter));
    }

    pthread_exit(NULL);
}

/* 管理线程 */
void *adjust_thread(void *threadpool)
{
    int i;
    threadpool_t *pool = (threadpool_t *)threadpool;
    while (!pool->shutdown) {

        sleep(DEFAULT_TIME);                                    /*定时 对线程池管理*/

        pthread_mutex_lock(&(pool->lock));
        int queue_size = pool->queue_size;                      /* 关注 任务数 */
        int live_thr_num = pool->live_thr_num;                  /* 存活 线程数 */
        pthread_mutex_unlock(&(pool->lock));

        pthread_mutex_lock(&(pool->thread_counter));
        int busy_thr_num = pool->busy_thr_num;                  /* 忙着的线程数 */
        pthread_mutex_unlock(&(pool->thread_counter));

        /* 创建新线程 算法： 任务数大于最小线程池个数, 且存活的线程数少于最大线程个数时 如：30>=10 && 40<100*/
        if (queue_size >= MIN_WAIT_TASK_NUM && live_thr_num < pool->max_thr_num) {
            pthread_mutex_lock(&(pool->lock));  
            int add = 0;

            /*一次增加 DEFAULT_THREAD 个线程*/
            for (i = 0; i < pool->max_thr_num && add < DEFAULT_THREAD_VARY
                    && pool->live_thr_num < pool->max_thr_num; i++) {
                if (pool->threads[i] == 0 || !is_thread_alive(pool->threads[i])) {
                    pthread_create(&(pool->threads[i]), NULL, threadpool_thread, (void *)pool);
                    add++;
                    pool->live_thr_num++;
                }
            }

            pthread_mutex_unlock(&(pool->lock));
        }

        /* 销毁多余的空闲线程 算法：忙线程X2 小于 存活的线程数 且 存活的线程数 大于 最小线程数时*/
        if ((busy_thr_num * 2) < live_thr_num  &&  live_thr_num > pool->min_thr_num) {

            /* 一次销毁DEFAULT_THREAD个线程, 隨機10個即可 */
            pthread_mutex_lock(&(pool->lock));
            pool->wait_exit_thr_num = DEFAULT_THREAD_VARY;      /* 要销毁的线程数 设置为10 */
            pthread_mutex_unlock(&(pool->lock));

            for (i = 0; i < DEFAULT_THREAD_VARY; i++) {
                /* 通知处在空闲状态的线程, 他们会自行终止*/
                pthread_cond_signal(&(pool->queue_not_empty));
            }
        }
    }

    return NULL;
}

int threadpool_destroy(threadpool_t *pool)
{
    int i;
    if (pool == NULL) {
        return -1;
    }
    pool->shutdown = true;

    /*先销毁管理线程*/
    pthread_join(pool->adjust_tid, NULL);

    for (i = 0; i < pool->live_thr_num; i++) {
        /*通知所有的空闲线程*/
        pthread_cond_broadcast(&(pool->queue_not_empty));
    }
    for (i = 0; i < pool->live_thr_num; i++) {
        pthread_join(pool->threads[i], NULL);
    }
    threadpool_free(pool);

    return 0;
}

int threadpool_free(threadpool_t *pool)
{
    if (pool == NULL) {
        return -1;
    }

    if (pool->task_queue) {
        free(pool->task_queue);
    }
    if (pool->threads) {
        free(pool->threads);
        pthread_mutex_lock(&(pool->lock));
        pthread_mutex_destroy(&(pool->lock));
        pthread_mutex_lock(&(pool->thread_counter));
        pthread_mutex_destroy(&(pool->thread_counter));
        pthread_cond_destroy(&(pool->queue_not_empty));
        pthread_cond_destroy(&(pool->queue_not_full));
    }
    free(pool);
    pool = NULL;

    return 0;
}

int threadpool_all_threadnum(threadpool_t *pool)
{
    int all_threadnum = -1;
    pthread_mutex_lock(&(pool->lock));
    all_threadnum = pool->live_thr_num;
    pthread_mutex_unlock(&(pool->lock));
    return all_threadnum;
}

int threadpool_busy_threadnum(threadpool_t *pool)
{
    int busy_threadnum = -1;
    pthread_mutex_lock(&(pool->thread_counter));
    busy_threadnum = pool->busy_thr_num;
    pthread_mutex_unlock(&(pool->thread_counter));
    return busy_threadnum;
}

int is_thread_alive(pthread_t tid)
{
    int kill_rc = pthread_kill(tid, 0);     //发0号信号，测试线程是否存活
    if (kill_rc == ESRCH) {
        return false;
    }

    return true;
}
```