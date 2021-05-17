# condition_variable

C++11并行编程-条件变量(condition_variable)详细说明
<condition_variable >头文件主要包含有类和函数相关的条件变量。

包括相关类 
std::condition_variable和 std::condition_variable_any，还有枚举类型std::cv_status。另外还包含函数 std::notify_all_at_thread_exit()，以下分别介绍一下以上几种类型。

## 1、std::condition_variable 类介绍

std::condition_variable是条件变量，很多其它有关条件变量的定义參考维基百科。Linux下使用 Pthread库中的 pthread_cond_*() 函数提供了与条件变量相关的功能， Windows 则參考 MSDN。

当 std::condition_variable对象的某个wait 函数被调用的时候，它使用 std::unique_lock(通过 std::mutex) 来锁住当前线程。

当前线程会一直被堵塞。直到另外一个线程在同样的 std::condition_variable 对象上调用了 notification 函数来唤醒当前线程。

std::condition_variable 对象通常使用 std::unique_lock<std::mutex> 来等待，假设须要使用另外的 lockable 类型，能够使用std::condition_variable_any类。本文后面会讲到 std::condition_variable_any 的使用方法。

```c++
#include <iostream>                // std::cout
#include <thread>                // std::thread
#include <mutex>                // std::mutex, std::unique_lock
#include <condition_variable>    // std::condition_variable
 
std::mutex mtx; // 全局相互排斥锁.
std::condition_variable cv; // 全局条件变量.
bool ready = false; // 全局标志位.
 
void do_print_id(int id)
{
    std::unique_lock <std::mutex> lck(mtx);
    while (!ready) // 假设标志位不为 true, 则等待...
        cv.wait(lck); // 当前线程被堵塞, 当全局标志位变为 true 之后,
    // 线程被唤醒, 继续往下运行打印线程编号id.
    std::cout << "thread " << id << '\n';
}
 
void go()
{
    std::unique_lock <std::mutex> lck(mtx);
    ready = true; // 设置全局标志位为 true.
    cv.notify_all(); // 唤醒全部线程.
}
 
int main()
{
    std::thread threads[10];
    // spawn 10 threads:
    for (int i = 0; i < 10; ++i)
        threads[i] = std::thread(do_print_id, i);
 
    std::cout << "10 threads ready to race...\n";
    go(); // go!
 
    for (auto & th:threads)
        th.join();
 
    return 0;
}
```

结果：
```
10 threads ready to race...
thread 1
thread 0
thread 2
thread 3
thread 4
thread 5
thread 6
thread 7
thread 8
thread 9
```

std::condition_variable 的拷贝构造函数被禁用，仅仅提供了默认构造函数。

看看 std::condition_variable 的各个成员函数
std::condition_variable::wait() 介绍:
std::condition_variable提供了两种 wait() 函数。

```c++
void wait (unique_lock<mutex>& lck);
 
template <class Predicate>
void wait (unique_lock<mutex>& lck, Predicate pred);
```

当前线程调用 wait() 后将被堵塞(此时当前线程应该获得了锁（mutex），最好还是设获得锁 lck)，直到另外某个线程调用 notify_* 唤醒了当前线程。

在线程被堵塞时，该函数会自己主动调用 lck.unlock() 释放锁，使得其它被堵塞在锁竞争上的线程得以继续运行。另外，一旦当前线程获得通知(notified，一般是另外某个线程调用 notify_* 唤醒了当前线程)，wait()函数也是自己主动调用 lck.lock()，使得lck的状态和 wait 函数被调用时同样。

在另外一种情况下（即设置了 Predicate）。仅仅有当 pred 条件为false 时调用 wait() 才会堵塞当前线程。而且在收到其它线程的通知后仅仅有当 pred 为 true 时才会被解除堵塞。

因此另外一种情况相似以下代码：

```c++
#include <iostream>                // std::cout
#include <thread>                // std::thread, std::this_thread::yield
#include <mutex>                // std::mutex, std::unique_lock
#include <condition_variable>    // std::condition_variable
 
std::mutex mtx;
std::condition_variable cv;
 
int cargo = 0;
bool shipment_available()
{
    return cargo != 0;
}
 
// 消费者线程.
void consume(int n)
{
    for (int i = 0; i < n; ++i) {
        std::unique_lock <std::mutex> lck(mtx);
        cv.wait(lck, shipment_available);
        std::cout << cargo << '\n';
        cargo = 0;
    }
}
 
int main()
{
    std::thread consumer_thread(consume, 10); // 消费者线程.
 
    // 主线程为生产者线程, 生产 10 个物品.
    for (int i = 0; i < 10; ++i) {
        while (shipment_available())
            std::this_thread::yield();
        std::unique_lock <std::mutex> lck(mtx);
        cargo = i + 1;
        cv.notify_one();
    }
 
    consumer_thread.join();
 
    return 0;
}
```

### 1.1、std::condition_variable::wait_for() 介绍

```c++
template <class Rep, class Period>
  cv_status wait_for (unique_lock<mutex>& lck,
                      const chrono::duration<Rep,Period>& rel_time);
 
template <class Rep, class Period, class Predicate>
       bool wait_for (unique_lock<mutex>& lck,
                      const chrono::duration<Rep,Period>& rel_time, Predicate pred);
```

与std::condition_variable::wait() 相似，只是 wait_for能够指定一个时间段，在当前线程收到通知或者指定的时间 rel_time 超时之前。该线程都会处于堵塞状态。而一旦超时或者收到了其它线程的通知，wait_for返回，剩下的处理步骤和 wait()相似。

另外，wait_for 的重载版本号的最后一个參数pred表示 wait_for的预測条件。仅仅有当 pred条件为false时调用 wait()才会堵塞当前线程，而且在收到其它线程的通知后仅仅有当 pred为 true时才会被解除堵塞，因此相当于例如以下代码：

```c++
return wait_until (lck, chrono::steady_clock::now() + rel_time, std::move(pred));
```

请看以下的样例（參考），以下的样例中，主线程等待th线程输入一个值。然后将th线程从终端接收的值打印出来。在th线程接受到值之前，主线程一直等待。每一个一秒超时一次，并打印一个 "."：　　

```c++
#include <iostream>           // std::cout
#include <thread>             // std::thread
#include <chrono>             // std::chrono::seconds
#include <mutex>              // std::mutex, std::unique_lock
#include <condition_variable> // std::condition_variable, std::cv_status
 
std::condition_variable cv;
 
int value;
 
void do_read_value()
{
    std::cin >> value;
    cv.notify_one();
}
 
int main ()
{
    std::cout << "Please, enter an integer (I'll be printing dots): \n";
    std::thread th(do_read_value);
 
    std::mutex mtx;
    std::unique_lock<std::mutex> lck(mtx);
    while (cv.wait_for(lck,std::chrono::seconds(1)) == std::cv_status::timeout) {
        std::cout << '.';
        std::cout.flush();
    }
 
    std::cout << "You entered: " << value << '\n';
 
    th.join();
    return 0;
}
```

输出结果：
```
Please, enter an integer (I'll be printing dots): 
...8...............
You entered: 8
```

### 1.2、std::condition_variable::wait_until 介绍

```c++
template <class Clock, class Duration>
  cv_status wait_until (unique_lock<mutex>& lck,
                        const chrono::time_point<Clock,Duration>& abs_time);
 
template <class Clock, class Duration, class Predicate>
       bool wait_until (unique_lock<mutex>& lck,
                        const chrono::time_point<Clock,Duration>& abs_time,
                        Predicate pred);
```

与 std::condition_variable::wait_for 相似，可是wait_until能够指定一个时间点，在当前线程收到通知或者指定的时间点 abs_time超时之前，该线程都会处于堵塞状态。而一旦超时或者收到了其它线程的通知，wait_until返回。剩下的处理步骤和 wait_until() 相似。

另外，wait_until的重载版本号的最后一个參数 pred表示 wait_until 的预測条件。仅仅有当 pred 条件为 false时调用 wait()才会堵塞当前线程，而且在收到其它线程的通知后仅仅有当pred为 true时才会被解除堵塞，因此相当于例如以下代码：

```c++
while (!pred())
  if ( wait_until(lck,abs_time) == cv_status::timeout)
    return pred();
return true;
```

### 1.3 std::condition_variable::notify_one() 介绍

唤醒某个等待(wait)线程。假设当前没有等待线程，则该函数什么也不做，假设同一时候存在多个等待线程，则唤醒某个线程是不确定的(unspecified)。

请看下例（參考）：

```c++
#include <iostream>                // std::cout
#include <thread>                // std::thread
#include <mutex>                // std::mutex, std::unique_lock
#include <condition_variable>    // std::condition_variable
 
std::mutex mtx;
std::condition_variable cv;
 
int cargo = 0; // shared value by producers and consumers
 
void consumer()
{
    std::unique_lock < std::mutex > lck(mtx);
    while (cargo == 0)
        cv.wait(lck);
    std::cout << cargo << '\n';
    cargo = 0;
}
 
void producer(int id)
{
    std::unique_lock < std::mutex > lck(mtx);
    cargo = id;
    cv.notify_one();
}
 
int main()
{
    std::thread consumers[10], producers[10];
 
    // spawn 10 consumers and 10 producers:
    for (int i = 0; i < 10; ++i) {
        consumers[i] = std::thread(consumer);
        producers[i] = std::thread(producer, i + 1);
    }
 
    // join them back:
    for (int i = 0; i < 10; ++i) {
        producers[i].join();
        consumers[i].join();
    }
 
    return 0;
}
```

输出结果：
```
3
4
5
6
7
8
9
2
10
1
```

### 1.4 std::condition_variable::notify_all() 介绍

唤醒全部的等待(wait)线程。假设当前没有等待线程，则该函数什么也不做。请看以下的样例：

```c++
#include <iostream>                // std::cout
#include <thread>                // std::thread
#include <mutex>                // std::mutex, std::unique_lock
#include <condition_variable>    // std::condition_variable
 
std::mutex mtx; // 全局相互排斥锁.
std::condition_variable cv; // 全局条件变量.
bool ready = false; // 全局标志位.
 
void do_print_id(int id)
{
    std::unique_lock <std::mutex> lck(mtx);
    while (!ready) // 假设标志位不为 true, 则等待...
        cv.wait(lck); // 当前线程被堵塞, 当全局标志位变为 true 之后,
    // 线程被唤醒, 继续往下运行打印线程编号id.
    std::cout << "thread " << id << '\n';
}
 
void go()
{
    std::unique_lock <std::mutex> lck(mtx);
    ready = true; // 设置全局标志位为 true.
    cv.notify_all(); // 唤醒全部线程.
}
 
int main()
{
    std::thread threads[10];
    // spawn 10 threads:
    for (int i = 0; i < 10; ++i)
        threads[i] = std::thread(do_print_id, i);
 
    std::cout << "10 threads ready to race...\n";
    go(); // go!
 
  for (auto & th:threads)
        th.join();
 
    return 0;
}
```

输出结果：
```
10 threads ready to race...
thread 6
thread 7
thread 8
thread 9
thread 5
thread 4
thread 3
thread 2
thread 1
thread 0
```

## 2、std::condition_variable_any 介绍

与 std::condition_variable相似。仅仅只是std::condition_variable_any的 wait 函数能够接受不论什么 lockable參数，而 std::condition_variable仅仅能接受 std::unique_lock<std::mutex>类型的參数，除此以外，和std::condition_variable差点儿全然一样。

### 2.1 std::cv_status枚举类型介绍

cv_status::no_timeout wait_for 或者wait_until没有超时，即在规定的时间段内线程收到了通知。

cv_status::timeout  wait_for 或者 wait_until 超时。
std::notify_all_at_thread_exit

函数原型为：

`void notify_all_at_thread_exit (condition_variable& cond, unique_lock<mutex> lck);`

当调用该函数的线程退出时，全部在 cond 条件变量上等待的线程都会收到通知。

请看下例（參考）：

```c++
#include <iostream>           // std::cout
#include <thread>             // std::thread
#include <mutex>              // std::mutex, std::unique_lock
#include <condition_variable> // std::condition_variable
 
std::mutex mtx;
std::condition_variable cv;
bool ready = false;
 
void print_id (int id) {
  std::unique_lock<std::mutex> lck(mtx);
  while (!ready) cv.wait(lck);
  // ...
  std::cout << "thread " << id << '\n';
}
 
void go() {
  std::unique_lock<std::mutex> lck(mtx);
  std::notify_all_at_thread_exit(cv,std::move(lck));
  ready = true;
}
 
int main ()
{
  std::thread threads[10];
  // spawn 10 threads:
  for (int i=0; i<10; ++i) {
      threads[i] = std::thread(print_id,i);
  }
    
  std::cout << "10 threads ready to race...\n";
  std::thread(go).detach();   // go!
 
  for (auto& th : threads) {
      th.join();
  } 
 
  return 0;
}
```

<condition_variable> 头文件里的两个条件变量类（std::condition_variable和std::condition_variable_any）、枚举类型（std::cv_status）、以及辅助函数（std::notify_all_at_thread_exit()）都已经介绍完　　

 