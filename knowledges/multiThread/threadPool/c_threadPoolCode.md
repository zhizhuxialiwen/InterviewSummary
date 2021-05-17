# Linux下C语言实现线程池


    前面看完C++实现线程池之后，接下来我们看看C语言如何实现线程池，原理是类似的，C++主要是使用线程的类Thread，而C语言的实现主要是调用Linux中的系统函数。

## 1、线程池的具体的实现代码如下：
   
* 文件thread_pool.c：

```c++
#include<stdio.h>
#include<stdlib.h>
#include<unistd.h>
#include<sys/types.h>
#include<pthread.h>
#include<assert.h>
 
//任务队列的结构体定义
typedef struct worker
{
	void *(*process)(void *arg);    //一个返回值为void*的参数为void*的函数指针
	void *arg;                      //函数参数
	struct worker *next;            //指向下一个结构体
}CThread_worker;
 
 
//线程池的结构体定义
typedef struct
{
	pthread_mutex_t queue_lock;     //互斥量
	pthread_cond_t queue_ready;     //条件变量
	CThread_worker *queue_head;     //线程任务队列的链表头指针
	int shutdown;                   //是否销毁这个线程池
	pthread_t *threadid;            //指向线程id的指针
	int max_thread_num;             //最大线程数量
	int cur_queue_size;             //任务队列中的任务数
}CThread_pool;
 
//线程池相关的全局变量
int pool_add_worker(void *(*process)(void *arg), void *arg);   //增加一个任务到任务队列
void *thread_routine(void *arg);                              //用于初始化线程的函数指针
static CThread_pool *pool = NULL;                             //指向线程池的指针
 
//线程池初始化函数
void pool_init(int max_thread_num){
	pool = (CThread_pool *)malloc(sizeof(CThread_pool));     //初始化线程池指针
	pthread_mutex_init(&(pool->queue_lock), NULL);           //初始化互斥量
	pthread_cond_init(&(pool->queue_ready), NULL);           //初始化条件变量
	pool->queue_head = NULL;                                 //初始化任务队列头指针
	pool->max_thread_num = max_thread_num;                   //初始化最大线程数
	pool->cur_queue_size = 0;                                //初始化任务队列当前数量
	pool->shutdown = 0;                                      //初始化是否销毁线程池的标志
	pool->threadid = (pthread_t *)malloc(max_thread_num * sizeof(pthread_t)); //初始化线程id
	int i = 0;
	for (i = 0; i < max_thread_num; i++)
    {
		pthread_create(&(pool->threadid[i]), NULL, thread_routine, NULL);   //创建线程
	}
}
 
//执行任务队列的任务的函数，用于初始化线程
/*
在调用pthread_cond_wait()前必须由本线程加锁（pthread_mutex_lock()），而在更新条件等待队列以前，mutex保持锁定状态，
并在线程挂起进入等待前解锁。在条件满足从而离开pthread_cond_wait()之前，mutex将被重新加锁，
以与进入pthread_cond_wait()前的加锁动作对应。
*/
void *thread_routine(void *arg)
{
	printf("start thread 0x%x\n", pthread_self());
	while (1)
    {
		pthread_mutex_lock(&(pool->queue_lock));                        //互斥量加锁
		while(pool->cur_queue_size == 0 && !pool->shutdown)	            //等待条件发生
        {
			printf("thread 0x%x is waiting\n", pthread_self()); 
			pthread_cond_wait(&(pool->queue_ready), &(pool->queue_lock));
		}
 
		if(pool->shutdown)                                   //如果shutdown为1则线程退出
        {
			pthread_mutex_unlock(&(pool->queue_lock));
			printf("thread 0x%x will exit\n", pthread_self());
			pthread_exit(NULL);
		}
		printf("thread 0x%x is starting to work\n", pthread_self());
		assert(pool->cur_queue_size != 0);
		assert(pool->queue_head != NULL);
		pool->cur_queue_size--;
		CThread_worker *worker = pool->queue_head;                       //将任务加入任务队列
		pool->queue_head = worker->next;
		pthread_mutex_unlock(&(pool->queue_lock));                       // 互斥量解锁
		(*(worker->process))(worker->arg);                               // 执行这个任务
		free(worker);                                                    // 释放任务指针
		worker = NULL;
	}
	pthread_exit(NULL);
}
 
//添加一个任务到任务队列
int pool_add_worker(void *(*process)(void *arg), void *arg)
{
	CThread_worker *newworker = (CThread_worker *)malloc(sizeof(CThread_worker));  //创建队列的一个节点
	newworker->process = process;                                      //对节点进行初始化
	newworker->arg = arg;
	newworker->next = NULL;
	pthread_mutex_lock(&(pool->queue_lock));                           //互斥量上锁
	CThread_worker *member = pool->queue_head;                         //将这个节点插入任务队列中
	if (member != NULL)
    {
		while (member->next != NULL)
        {
            member = member->next;
        }
		member->next = newworker;
	}
	else
    {
        pool->queue_head = newworker;
    }
		
	assert(pool->queue_head != NULL);
	pool->cur_queue_size++;                                            //任务队列个数加一
	pthread_mutex_unlock(&(pool->queue_lock));                         //互斥量解锁
	pthread_cond_signal(&(pool->queue_ready));                         //给条件变量发信号
	return 0;
}
 
//销毁线程
int pool_destroy()
{
	if (pool->shutdown)
    {
        return -1;
    }		
	pool->shutdown = 1;
	pthread_cond_broadcast(&(pool->queue_ready));      //唤醒所有线程
	int i;
	for (i = 0; i < pool->max_thread_num; i++)
    {
		pthread_join(pool->threadid[i], NULL);         //以参与的方式回收线程资源
    }
	free(pool->threadid);                              //释放线程id的指针
	CThread_worker *head = NULL;                     
	while (pool->queue_head != NULL)                   //释放队列中所有的节点
    {
		head = pool->queue_head;
		pool->queue_head = pool->queue_head->next;
		free(head);
	}
	pthread_mutex_destroy(&(pool->queue_lock));        //销毁互斥量
	pthread_cond_destroy(&(pool->queue_ready));        //销毁条件变量
	free(pool);                                        //释放线程池指针
	pool = NULL;                                       //线程池指针置空
 	return 0;
}
```

## 2、线程池的测试
 

* 头文件thread_pool.h: 

```c++
#include <stdio.h>
#include <stdlib.h>
 
void pool_init(int max_thread_num);
int pool_add_worker(void *(*process)(void *arg), void *arg);
int pool_destroy();
```

  
* 测试函数test.c：

```c++
#include<stdio.h>
#include<stdlib.h>
#include<unistd.h>
#include<pthread.h>
#include "thread_pool.h"
 
//任务函数
void *myprocess(void *arg){
	printf("thread is 0x%x, working on task %d\n", pthread_self(), *(int *)arg);
	sleep(1);
	return NULL;
}
 
int main(){
 
	pool_init(3);                                           //初始化线程池
	int *workingnum = (int *)malloc(sizeof(int)*10);        //保存线程id的指针
	int i;
	for(i = 0; i< 10; i++){
		workingnum[i] = i;
		pool_add_worker(myprocess, &workingnum[i]);
	}
	sleep(5);                                               //休眠5秒，保证每个线程启动后再销毁
	pool_destroy();                                         //销毁线程池
	free(workingnum);                                       //释放保存线程id的指针
	return 0;
}
```
  
编译：
 
```s
 g++ -w test.c thread_pool.c -o test -lpthread  # 这里需要加入动态链接库libpthread.so
 

 运行： 

./test
 
输出：
start thread 0x4069700
thread 0x4069700 is starting to work
thread is 0x4069700, working on task 0
start thread 0x3868700
thread 0x3868700 is starting to work
thread is 0x3868700, working on task 1
start thread 0x3067700
thread 0x3067700 is starting to work
thread is 0x3067700, working on task 2
thread 0x4069700 is starting to work
thread is 0x4069700, working on task 3
thread 0x3868700 is starting to work
thread is 0x3868700, working on task 4
thread 0x3067700 is starting to work
thread is 0x3067700, working on task 5
thread 0x4069700 is starting to work
thread is 0x4069700, working on task 6
thread 0x3868700 is starting to work
thread is 0x3868700, working on task 7
thread 0x3067700 is starting to work
thread is 0x3067700, working on task 8
thread 0x4069700 is starting to work
thread is 0x4069700, working on task 9
thread 0x3868700 is waiting
thread 0x3067700 is waiting
thread 0x4069700 is waiting
thread 0x3868700 will exit
thread 0x3067700 will exit
thread 0x4069700 will exit
```
