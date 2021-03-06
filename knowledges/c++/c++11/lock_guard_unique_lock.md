# c++11中的lock_guard和unique_lock使用浅析

锁
锁用来在多线程访问同一个资源时防止数据竞险，保证数据的一致性访问。

多线程本来就是为了提高效率和响应速度，但锁的使用又限制了多线程的并行执行，这会降低效率，但为了保证数据正确，不得不使用锁，它们就是这样纠缠。

作为效率优先的c++开发人员，很多人谈锁色变。

虽然有很多的无锁技术应用到项目中来，但是还是很有必要对锁的技术有一个基础的理解。本文主要讨论c++11中的两种锁：lock_guard 和 unique_lock。

结合锁进行线程间同步的条件变量使用，请参考条件变量condition variable 。

## 1 lock_guard

lock_guard是一个互斥量包装程序，它提供了一种方便的RAII（Resource acquisition is initialization ）风格的机制来在作用域块的持续时间内拥有一个互斥量。

创建lock_guard对象时，它将尝试获取提供给它的互斥锁的所有权。当控制流离开lock_guard对象的作用域时，lock_guard析构并释放互斥量。

* 它的特点如下：
    * 创建即加锁，作用域结束自动析构并解锁，无需手工解锁
    * 不能中途解锁，必须等作用域结束才解锁
    * 不能复制

示例代码如下：

```c++
#include <thread>
#include <mutex>
#include <iostream>
 
int g_i = 0;
std::mutex g_i_mutex;  // protects g_i
 
void safe_increment()
{
    const std::lock_guard<std::mutex> lock(g_i_mutex);
    ++g_i;
    std::cout << std::this_thread::get_id() << ": " << g_i << '\n';
    // g_i_mutex is automatically released when lock
    // goes out of scope
}
 
int main()
{
    std::cout << "main: " << g_i << '\n';
    std::thread t1(safe_increment);
    std::thread t2(safe_increment);
    t1.join();
    t2.join();
    std::cout << "main: " << g_i << '\n';
}
```

输出：
```c++
main: 0
140641306900224: 1
140641298507520: 2
main: 2
```

## 2、unique_lock

unique_lock是一个通用的互斥量锁定包装器，它允许延迟锁定，限时深度锁定，递归锁定，锁定所有权的转移以及与条件变量一起使用。

简单地讲，unique_lock 是 lock_guard 的升级加强版，它具有 lock_guard 的所有功能，同时又具有其他很多方法，使用起来更强灵活方便，能够应对更复杂的锁定需要。

* 特点如下：
    * 创建时可以不锁定（通过指定第二个参数为std::defer_lock），而在需要时再锁定
    * 可以随时加锁解锁
    * 作用域规则同 lock_grard，析构时自动释放锁
    * 不可复制，可移动
    * 条件变量需要该类型的锁作为参数（此时必须使用unique_lock）
示例代码：

```c++
#include <mutex>
#include <thread>
#include <chrono>
 
struct Box {
    explicit Box(int num) : num_things{num} {}
 
    int num_things;
    std::mutex m;
};
 
void transfer(Box &from, Box &to, int num)
{
    // don't actually take the locks yet
    std::unique_lock<std::mutex> lock1(from.m, std::defer_lock);
    std::unique_lock<std::mutex> lock2(to.m, std::defer_lock);
 
    // lock both unique_locks without deadlock
    std::lock(lock1, lock2);
 
    from.num_things -= num;
    to.num_things += num;
 
    // 'from.m' and 'to.m' mutexes unlocked in 'unique_lock' dtors
}
 
int main()
{
    Box acc1(100);
    Box acc2(50);
    std::thread t1(transfer, std::ref(acc1), std::ref(acc2), 10);
    std::thread t2(transfer, std::ref(acc2), std::ref(acc1), 5);
    t1.join();
    t2.join();
}
```

总结
所有 lock_guard 能够做到的事情，都可以使用 unique_lock 做到，反之则不然。

那么何时使用lock_guard呢？很简单，

需要使用锁的时候，首先考虑使用 lock_guard
它简单、明了、易读。如果用它完全ok，就不要考虑其他了。

如果现实不允许，就让实力派 unique_lock 出马吧！