# c++线程中的几种锁

线程之间的锁有：**互斥锁、条件锁、自旋锁、读写锁、递归锁。**一般而言，锁的功能越强大，性能就会越低。

## 1、互斥锁

互斥锁用于控制多个线程对他们之间共享资源互斥访问的一个信号量。也就是说是为了避免多个线程在某一时刻同时操作一个共享资源。例如线程池中的有多个空闲线程和一个任务队列。任何是一个线程都要使用互斥锁互斥访问任务队列，以避免多个线程同时访问任务队列以发生错乱。

在某一时刻，只有一个线程可以获取互斥锁，在释放互斥锁之前其他线程都不能获取该互斥锁。如果其他线程想要获取这个互斥锁，那么这个线程只能以阻塞方式进行等待。

```c++
头文件：<pthread.h>

类型：pthread_mutex_t，

函数：pthread_mutex_init(pthread_mutex_t * mutex, const phtread_mutexattr_t * mutexattr);//动态方式创建锁，相当于new动态创建一个对象

pthread_mutex_destory(pthread_mutex_t *mutex)//释放互斥锁，相当于delete

pthread_mutex_t mutex = PTHREAD_MUTEX_INITIALIZER;//以静态方式创建锁

pthread_mutex_lock(pthread_mutex_t *mutex)//以阻塞方式运行的。如果之前mutex被加锁了，那么程序会阻塞在这里。

pthread_mutex_unlock(pthread_mutex_t *mutex)

int pthread_mutex_trylock(pthread_mutex_t * mutex);//会尝试对mutex加锁。如果mutex之前已经被锁定，返回非0,；如果mutex没有被锁定，则函数返回并锁定mutex

该函数是以非阻塞方式运行了。也就是说如果mutex之前已经被锁定，函数会返回非0，程序继续往下执行。
```

## 2、条件锁

条件锁就是所谓的条件变量，**某一个线程因为某个条件为满足时可以使用条件变量使改程序处于阻塞状态。** 一旦条件满足以“信号量”的方式唤醒一个因为该条件而被阻塞的线程。最为常见就是在线程池中，起初没有任务时任务队列为空，此时线程池中的线程因为“任务队列为空”这个条件处于阻塞状态。一旦有任务进来，就会以信号量的方式唤醒一个线程来处理这个任务。这个过程中就使用到了条件变量pthread_cond_t。

```c++
头文件：<pthread.h>

类型：pthread_cond_t

函数：pthread_cond_init(pthread_cond_t * condtion, const phtread_condattr_t * condattr);//对条件变量进行动态初始化，相当于new创建对象

pthread_cond_destory(pthread_cond_t * condition);//释放动态申请的条件变量，相当于delete释放对象

pthread_cond_t condition = PTHREAD_COND_INITIALIZER;//静态初始化条件变量

pthread_cond_wait(pthread_cond_t * cond, pthread_mutex_t * mutex);//该函数以阻塞方式执行。如果某个线程中的程序执行了该函数，那么这个线程就会以阻塞方式等待，直到收到pthread_cond_signal或者pthread_cond_broadcast函数发来的信号而被唤醒。

注意：pthread_cond_wait函数的语义相当于：首先解锁互斥锁，然后以阻塞方式等待条件变量的信号，收到信号后又会对互斥锁加锁。

为了防止“虚假唤醒”，该函数一般放在while循环体中。例如

pthread_mutex_lock(mutex);//加互斥锁
while(条件不成立)//当前线程中条件变量不成立
{
    pthread_cond_wait(cond, mutex);//解锁，其他线程使条件成立发送信号，加锁。
}
//对进程之间的共享资源进行操作
pthread_mutex_unlock(mutex);//释放互斥锁
pthread_cond_signal(pthread_cond_t * cond);//在另外一个线程中改变线程，条件满足发送信号。唤醒一个等待的线程（可能有多个线程处于阻塞状态），唤醒哪个线程由具体的线程调度策略决定

pthread_cond_broadcast(pthread_cond_t * cond);//以广播形式唤醒所有因为该条件变量而阻塞的所有线程，唤醒哪个线程由具体的线程调度策略决定

pthread_cond_timedwait(pthread_cond_t * cond, pthread_mutex_t * mutex, struct timespec * time);//以阻塞方式等待，如果时间time到了条件还没有满足还是会结束
```

## 3、自旋锁

前面的两种锁是比较常见的锁，也比较容易理解。下面通过比较互斥锁和自旋锁原理的不同，这对于真正理解自旋锁有很大帮助。

假设我们有一个两个处理器core1和core2计算机，现在在这台计算机上运行的程序中有两个线程：T1和T2分别在处理器core1和core2上运行，两个线程之间共享着一个资源。

首先我们说明互斥锁的工作原理，互斥锁是是一种sleep-waiting的锁。假设线程T1获取互斥锁并且正在core1上运行时，此时线程T2也想要获取互斥锁（pthread_mutex_lock），但是由于T1正在使用互斥锁使得T2被阻塞。当T2处于阻塞状态时，T2被放入到等待队列中去，处理器core2会去处理其他任务而不必一直等待（忙等）。也就是说处理器不会因为线程阻塞而空闲着，它去处理其他事务去了。

而自旋锁就不同了，自旋锁是一种busy-waiting的锁。也就是说，如果T1正在使用自旋锁，而T2也去申请这个自旋锁，此时T2肯定得不到这个自旋锁。与互斥锁相反的是，此时运行T2的处理器core2会一直不断地循环检查锁是否可用（自旋锁请求），直到获取到这个自旋锁为止。

从“自旋锁”的名字也可以看出来，如果一个线程想要获取一个被使用的自旋锁，那么它会一致占用CPU请求这个自旋锁使得CPU不能去做其他的事情，直到获取这个锁为止，这就是“自旋”的含义。

**当发生阻塞时，互斥锁可以让CPU去处理其他的任务；而自旋锁让CPU一直不断循环请求获取这个锁。** 通过两个含义的对比可以我们知道“自旋锁”是比较耗费CPU的

```c++
头文件：<linux\spinlock.h>

自旋锁的类型：spinlock_t

相关函数：初始化：spin_lock_init(spinlock_t *x);

spin_lock(x);   //只有在获得锁的情况下才返回，否则一直“自旋”
spin_trylock(x);  //如立即获得锁则返回真，否则立即返回假
释放锁：spin_unlock(x);

spin_is_locked(x)//　　该宏用于判断自旋锁x是否已经被某执行单元保持（即被锁），如果是，   返回真，否则返回假。
```

注意：自旋锁适合于短时间的的轻量级的加锁机制。

## 4、读写锁

说到读写锁我们可以借助于“读者-写者”问题进行理解。首先我们简单说下“读者-写者”问题。

计算机中某些数据被多个进程共享，对数据库的操作有两种：一种是读操作，就是从数据库中读取数据不会修改数据库中内容；另一种就是写操作，写操作会修改数据库中存放的数据。因此可以得到我们允许在数据库上同时执行多个“读”操作，但是某一时刻只能在数据库上有一个“写”操作来更新数据。这就是一个简单的读者-写者模型。