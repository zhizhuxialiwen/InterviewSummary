# C++11常用新特性汇总
 
感谢博主的分享，转载自：http://www.cnblogs.com/feng-sc/p/5710724.html


1、关键字及新语法
    1.1、auto关键字及用法
    1.2、nullptr关键字及用法
    1.3、for循环语法
2、STL容器
    2.1、std::array
    2.2、std::forward_list
    2.3、std::unordered_map
    2.4、std::unordered_set
3、多线程
    3.1、std::thread
    3.2、st::atomic
    3.3、std::condition_variable
4、智能指针内存管理
    4.1、std::shared_ptr
    4.2、std::weak_ptr
5、其他
    5.1、std::function、std::bind封装可执行对象
    5.2、lamda表达式

## 1、关键字及新语法

　　C++11相比C++98增加了许多关键字及新的语法特性，很多人觉得这些语法可有可无，没有新特性也可以用传统C++去实现。

　　也许吧，但个人对待新技术总是抱着渴望而热衷的态度对待，也许正如很多人所想，用传统语法也可以实现，但新技术可以让你的设计更完美。这就如同在原来的维度里，你要完成一件事情，需要很多个复杂的步骤，但在新语法特性里，你可以从另外的维度，很干脆，直接就把原来很复杂的问题用很简单的方法解决了，我想着就是新的技术带来的一些编程体验上非常好的感觉。大家也不要觉得代码写得越复杂就先显得越牛B，有时候在理解的基础上，尽量选择“站在巨人的肩膀上”，可能你会站得更高，也看得更远。

　　本章重点总结一些常用c++11新语法特点。后续会在本人理解的基础上，会继续在本博客内更新或增加新的小章节。

### 1.1、auto关键字及用法

A、auto关键字能做什么？

　　auto并没有让C++成为弱类型语言，也没有弱化变量什么，只是使用auto的时候，编译器根据上下文情况，确定auto变量的真正类型。

//示例代码1.0 http://www.cnblogs.com/feng-sc/p/5710724.html

```c++
auto AddTest(int a, int b) 
{
    return a + b;
}

int main()
{
    auto index = 10;
    auto str = "abc";
    auto ret = AddTest(1,2);
    std::cout << "index:" << index << std::endl;
    std::cout << "str:" << str << std::endl;
    std::cout << "res:" << ret << std::endl;
}
```

　　是的，你没看错，代码也没错，auto在C++14中可以作为函数的返回值，因此auto AddTest(int a, int b)的定义是没问题的。

　　运行结果：
![c++11_1](../../../images/c++11_1.PNG)
　　  

B、auto不能做什么？

　　auto作为函数返回值时，只能用于定义函数，不能用于声明函数。

```c++
#pragma once
class Test
{
public:
    auto TestWork(int a ,int b);
};
```

　　如下函数中，在引用头文件的调用TestWork函数是，编译无法通过。

　　但如果把实现写在头文件中，可以编译通过，因为编译器可以根据函数实现的返回值确定auto的真实类型。如果读者用过inline类成员函数，这个应该很容易明白，此特性与inline类成员函数类似。

```c++
#pragma once
class Test
{
public:
    auto TestWork(int a, int b)
    {
        return a + b;
    }
};
```

### 1.2、nullptr关键字及用法

　　为什么需要nullptr? NULL有什么毛病？

　　我们通过下面一个小小的例子来发现NULL的一点问题：

```c++
class Test
{
public:
    void TestWork(int index)
    {
        std::cout << "TestWork 1" << std::endl;
    }
    void TestWork(int * index)
    {
        std::cout << "TestWork 2" << std::endl;
    }
};

int main()
{
    Test test;
    test.TestWork(NULL);
    test.TestWork(nullptr);
}
```

　　运行结果：
![c++11_2](../../../images/c++11_2.PNG)
　  　 

　　NULL在c++里表示空指针，看到问题了吧，我们调用test.TestWork(NULL)，其实期望是调用的是void TestWork(int * index)，但结果调用了void TestWork(int index)。但使用nullptr的时候，我们能调用到正确的函数。

### 1.3、for循环语法

　　习惯C#或java的同事之前使用C++的时候曾吐槽C++ for循环没有想C#那样foreach的用法，是的，在C++11之前，标准C++是无法做到的。熟悉boost库读者可能知道boost里面有foreach的宏定义BOOST_FOREACH，但个人觉得看起并不是那么美观。

　　OK，我们直接以简单示例看看用法吧。

```c++
int main()
{
    int numbers[] = { 1,2,3,4,5 };
    std::cout << "numbers:" << std::endl;
    for (auto number : numbers)
    {
        std::cout << number << std::endl;
    }
}
```

## 2、STL容器

　　C++11在STL容器方面也有所增加，给人的感觉就是越来越完整，越来越丰富的感觉，可以让我们在不同场景下能选择跟具合适的容器，提高我们的效率。

　　本章节总结C++11新增的一些容器，以及对其实现做一些简单的解释。

### 2.1、std::array

　　个人觉得std::array跟数组并没有太大区别，对于多维数据使用std::array，个人反而有些不是很习惯吧。

　　std::array相对于数组，增加了迭代器等函数（接口定义可参考C++官方文档）。

```c++
#include <array>
int main()
{
    std::array<int, 4> arrayDemo = { 1,2,3,4 };
    std::cout << "arrayDemo:" << std::endl;
    for (auto itor : arrayDemo)
    {
        std::cout << itor << std::endl;
    }
    int arrayDemoSize = sizeof(arrayDemo);
    std::cout << "arrayDemo size:" << arrayDemoSize << std::endl;
    return 0;
}
```

运行结果：
![c++11_3](../../../images/c++11_3.PNG)
 　  　

　　打印出来的size和直接使用数组定义结果是一样的。

### 2.2、std::forward_list

　　std::forward_list为从++新增的线性表，与list区别在于它是单向链表。我们在学习数据结构的时候都知道，链表在对数据进行插入和删除是比顺序存储的线性表有优势，因此在插入和删除操作频繁的应用场景中，使用list和forward_list比使用array、vector和deque效率要高很多。

```c++
#include <forward_list>
int main()
{
    std::forward_list<int> numbers = {1,2,3,4,5,4,4};
    std::cout << "numbers:" << std::endl;
    for (auto number : numbers)
    {
        std::cout << number << std::endl;
    }
    numbers.remove(4);
    std::cout << "numbers after remove:" << std::endl;
    for (auto number : numbers)
    {
        std::cout << number << std::endl;
    }
    return 0;
}
```

运行结果：
![c++11_4](../../../images/c++11_4.PNG)

　  　

### 2.3、std::unordered_map

　　std::unordered_map与std::map用法基本差不多，但STL在内部实现上有很大不同，std::map使用的数据结构为二叉树，而std::unordered_map内部是哈希表的实现方式，哈希map理论上查找效率为O(1)。但在存储效率上，哈希map需要增加哈希表的内存开销。

　　下面代码为C++官网实例源码实例：

```c++
#include <iostream>
#include <string>
#include <unordered_map>
int main()
{
    std::unordered_map<std::string, std::string> mymap =
    {
        { "house","maison" },
        { "apple","pomme" },
        { "tree","arbre" },
        { "book","livre" },
        { "door","porte" },
        { "grapefruit","pamplemousse" }
    };
    unsigned n = mymap.bucket_count();
    std::cout << "mymap has " << n << " buckets.\n";
    for (unsigned i = 0; i<n; ++i) 
    {
        std::cout << "bucket #" << i << " contains: ";
        for (auto it = mymap.begin(i); it != mymap.end(i); ++it)
            std::cout << "[" << it->first << ":" << it->second << "] ";
        std::cout << "\n";
    }
    return 0;
}
```
　
运行结果：
![c++11_5](../../../images/c++11_5.PNG)
 　 　

　　运行结果与官网给出的结果不一样。实验证明，不同编译器编译出来的结果不一样，如下为linux下gcc 4.6.3版本编译出来的结果。或许是因为使用的哈希算法不一样，个人没有深究此问题。

![c++11_6](../../../images/c++11_6.PNG)　

### 2.4、std::unordered_set

　　std::unordered_set的数据存储结构也是哈希表的方式结构，除此之外，std::unordered_set在插入时不会自动排序，这都是std::set表现不同的地方。

　　我们来测试一下下面的代码：　　

```c++
#include <iostream>
#include <string>
#include <unordered_set>
#include <set>
int main()
{
    std::unordered_set<int> unorder_set;
    unorder_set.insert(7);
    unorder_set.insert(5);
    unorder_set.insert(3);
    unorder_set.insert(4);
    unorder_set.insert(6);
    std::cout << "unorder_set:" << std::endl;
    for (auto itor : unorder_set)
    {
        std::cout << itor << std::endl;
    }

    std::set<int> set;
    set.insert(7);
    set.insert(5);
    set.insert(3);
    set.insert(4);
    set.insert(6);
    std::cout << "set:" << std::endl;
    for (auto itor : set)
    {
        std::cout << itor << std::endl;
    }
}
```

运行结果：
![c++11_7](../../../images/c++11_7.PNG)
　 　 

## 3、多线程

　　在C++11以前，C++的多线程编程均需依赖系统或第三方接口实现，一定程度上影响了代码的移植性。C++11中，引入了boost库中的多线程部分内容，形成C++标准，形成标准后的boost多线程编程部分接口基本没有变化，这样方便了以前使用boost接口开发的使用者切换使用C++标准接口，把容易把boost接口升级为C++接口。

　　我们通过如下几部分介绍C++11多线程方面的接口及使用方法。

### 3.1、std::thread

　　std::thread为C++11的线程类，使用方法和boost接口一样，非常方便，同时，C++11的std::thread解决了boost::thread中构成参数限制的问题，我想着都是得益于C++11的可变参数的设计风格。

　　我们通过如下代码熟悉下std::thread使用风格。

```c++
#include <thread>
void threadfun1()
{
    std::cout << "threadfun1 - 1\r\n" << std::endl;
    std::this_thread::sleep_for(std::chrono::seconds(1));
    std::cout << "threadfun1 - 2" << std::endl;
}

void threadfun2(int iParam, std::string sParam)
{
    std::cout << "threadfun2 - 1" << std::endl;
    std::this_thread::sleep_for(std::chrono::seconds(5));
    std::cout << "threadfun2 - 2" << std::endl;
}

int main()
{
    std::thread t1(threadfun1);
    std::thread t2(threadfun2, 10, "abc");
    t1.join();
    std::cout << "join" << std::endl;
    t2.detach();
    std::cout << "detach" << std::endl;
}
```

　　运行结果：
![c++11_8](../../../images/c++11_8.PNG)
 　　 

　　有以上输出结果可以得知，t1.join()会等待t1线程退出后才继续往下执行，t2.detach()并不会并不会把，detach字符输出后，主函数退出，threadfun2还未执行完成，但是在主线程退出后，t2的线程也被已经被强退出。

### 3.2、std::atomic

std::atomic为C++11分装的原子数据类型。

什么是原子数据类型？

　　从功能上看，简单地说，原子数据类型不会发生数据竞争，能直接用在多线程中而不必我们用户对其进行添加互斥资源锁的类型。从实现上，大家可以理解为这些原子类型内部自己加了锁。

　　我们下面通过一个测试例子说明原子类型std::atomic_int的特点。

　　下面例子中，我们使用10个线程，把std::atomic_int类型的变量iCount从100减到1。


```c++
#include <thread>
#include <atomic>
#include <stdio.h>
std::atomic_bool bIsReady = false;
std::atomic_int iCount = 100;
void threadfun1()
{
    if (!bIsReady) {
        std::this_thread::yield();
    }
    while (iCount > 0)
    {
        printf("iCount:%d\r\n", iCount--);
    }
}

int main()
{
    std::atomic_bool b;
    std::list<std::thread> lstThread;
    for (int i = 0; i < 10; ++i)
    {
        lstThread.push_back(std::thread(threadfun1));
    }
    for (auto& th : lstThread)
    {
        th.join();
    }
}
```

运行结果：
![c++11_9](../../../images/c++11_9.PNG)
　　

　　 注：屏幕太短的原因，上面结果没有截完屏

　　从上面的结果可以看到，iCount的最小结果都是1，单可能不是最后一次打印，没有小于等于0的情况，大家可以代码复制下来多运行几遍对比看看。

 

### 3.3、std::condition_variable

　　C++11中的std::condition_variable就像Linux下使用pthread_cond_wait和pthread_cond_signal一样，可以让线程休眠，直到别唤醒，现在在从新执行。线程等待在多线程编程中使用非常频繁，经常需要等待一些异步执行的条件的返回结果。

　　OK，在此不多解释，下面我们通过C++11官网的例子看看。

```c++
 1 // webset address: http://www.cplusplus.com/reference/condition_variable/condition_variable/%20condition_variable
 2 // condition_variable example
 3 #include <iostream>           // std::cout
 4 #include <thread>             // std::thread
 5 #include <mutex>              // std::mutex, std::unique_lock
 6 #include <condition_variable> // std::condition_variable
 7 
 8 std::mutex mtx;
 9 std::condition_variable cv;
10 bool ready = false;
11 
12 void print_id(int id) {
13     std::unique_lock<std::mutex> lck(mtx);
14     while (!ready)  cv.wait(lck);
15     // ...
16     std::cout << "thread " << id << '\n';
17 }
18 
19 void go() {
20     std::unique_lock<std::mutex> lck(mtx);
21     ready = true;
22     cv.notify_all();
23 }
24 
25 int main()
26 {
27     std::thread threads[10];
28     // spawn 10 threads:
29     for (int i = 0; i<10; ++i)
30         threads[i] = std::thread(print_id, i);
31 
32     std::cout << "10 threads ready to race...\n";
33     go();                       // go!
34 
35     for (auto& th : threads) th.join();
36 
37     return 0;
38 }
```

　　运行结果：
![c++11_10](../../../images/c++11_10.PNG)
　  

　　上面的代码，在14行中调用cv.wait(lck)的时候，线程将进入休眠，在调用33行的go函数之前，10个线程都处于休眠状态，当22行的cv.notify_all()运行后，14行的休眠将结束，继续往下运行，最终输出如上结果。

## 4、智能指针内存管理

　　在内存管理方面，C++11的std::auto_ptr基础上，移植了boost库中的智能指针的部分实现，如std::shared_ptr、std::weak_ptr等，当然，想boost::thread一样，C++11也修复了boost::make_shared中构造参数的限制问题。把智能指针添加为标准，个人觉得真的非常方便，毕竟在C++中，智能指针在编程设计中使用的还是非常广泛。

　　什么是智能指针？网上已经有很多解释，个人觉得“智能指针”这个名词似乎起得过于“霸气”，很多初学者看到这个名词就觉得似乎很难。

　　简单地说，智能指针只是用对象去管理一个资源指针，同时用一个计数器计算当前指针引用对象的个数，当管理指针的对象增加或减少时，计数器也相应加1或减1，当最后一个指针管理对象销毁时，计数器为1，此时在销毁指针管理对象的同时，也把指针管理对象所管理的指针进行delete操作。

　　如下图所示，简单话了一下指针、智能指针对象和计数器之间的关系：

![c++11_11](../../../images/c++11_11.PNG)

　　下面的小章节中，我们分别介绍常用的两个智能指针std::shared_ptr、std::weak_ptr的用法。

### 4.1、std::shared_ptr

　　std::shared_ptr包装了new操作符动态分别的内存，可以自由拷贝复制，基本上是使用最多的一个智能指针类型。

　　我们通过下面例子来了解下std::shared_ptr的用法:

复制代码

```c++
#include <memory>
class Test
{
public:
    Test()
    {
        std::cout << "Test()" << std::endl;
    }
    ~Test()
    {
        std::cout << "~Test()" << std::endl;
    }
};
int main()
{
    std::shared_ptr<Test> p1 = std::make_shared<Test>();
    std::cout << "1 ref:" << p1.use_count() << std::endl;
    {
        std::shared_ptr<Test> p2 = p1;
        std::cout << "2 ref:" << p1.use_count() << std::endl;
    }
    std::cout << "3 ref:" << p1.use_count() << std::endl;
    return 0;
}
```

　　运行结果：
![c++11_12](../../../images/c++11_12.PNG)
　　

　　从上面代码的运行结果，需要读者了解的是：

　　1、std::make_shared封装了new方法，boost::make_shared之前的原则是既然释放资源delete由智能指针负责，那么应该把new封装起来，否则会让人觉得自己调用了new，但没有调用delete，似乎与谁申请，谁释放的原则不符。C++也沿用了这一做法。

　　2、随着引用对象的增加std::shared_ptr<Test> p2 = p1，指针的引用计数有1变为2，当p2退出作用域后，p1的引用计数变回1，当main函数退出后，p1离开main函数的作用域，此时p1被销毁，当p1销毁时，检测到引用计数已经为1，就会在p1的析构函数中调用delete之前std::make_shared创建的指针。

### 4.2、std::weak_ptr

　　std::weak_ptr网上很多人说其实是为了解决std::shared_ptr在相互引用的情况下出现的问题而存在的，C++官网对这个只能指针的解释也不多，那就先甭管那么多了，让我们暂时完全接受这个观点。

　　std::weak_ptr有什么特点呢？与std::shared_ptr最大的差别是在赋值是，不会引起智能指针计数增加。

　　我们下面将继续如下两点：
1、std::shared_ptr相互引用会有什么后果；
2、std::weak_ptr如何解决第一点的问题。

A、std::shared_ptr相互引用的问题示例：

```c++
#include <memory>
class TestB;
class TestA
{
public:
    TestA()
    {
        std::cout << "TestA()" << std::endl;
    }
    void ReferTestB(std::shared_ptr<TestB> test_ptr)
    {
        m_TestB_Ptr = test_ptr;
    }
    ~TestA()
    {
        std::cout << "~TestA()" << std::endl;
    }
private:
    std::shared_ptr<TestB> m_TestB_Ptr; //TestB的智能指针
}; 

class TestB
{
public:
    TestB()
    {
        std::cout << "TestB()" << std::endl;
    }

    void ReferTestB(std::shared_ptr<TestA> test_ptr)
    {
        m_TestA_Ptr = test_ptr;
    }
    ~TestB()
    {
        std::cout << "~TestB()" << std::endl;
    }
    std::shared_ptr<TestA> m_TestA_Ptr; //TestA的智能指针
};


int main()
{
    std::shared_ptr<TestA> ptr_a = std::make_shared<TestA>();
    std::shared_ptr<TestB> ptr_b = std::make_shared<TestB>();
    ptr_a->ReferTestB(ptr_b);
    ptr_b->ReferTestB(ptr_a);
    return 0;
}
```

　　运行结果：
![c++11_13](../../../images/c++11_13.PNG)
　  　

　　大家可以看到，上面代码中，我们创建了一个TestA和一个TestB的对象，但在整个main函数都运行完后，都没看到两个对象被析构，这是什么问题呢？

　　原来，智能指针ptr_a中引用了ptr_b，同样ptr_b中也引用了ptr_a，在main函数退出前，ptr_a和ptr_b的引用计数均为2，退出main函数后，引用计数均变为1，也就是相互引用。

　　这等效于说：

　　　　ptr_a对ptr_b说，哎，我说ptr_b，我现在的条件是，你先释放我，我才能释放你，这是天生的，造物者决定的，改不了。

　　　　ptr_b也对ptr_a说，我的条件也是一样，你先释放我，我才能释放你，怎么办？

　　是吧，大家都没错，相互引用导致的问题就是释放条件的冲突，最终也可能导致内存泄漏。

 B、std::weak_ptr如何解决相互引用的问题

　　我们在上面的代码基础上使用std::weak_ptr进行修改：

```c++
#include <memory>
class TestB;
class TestA
{
public:
    TestA()
    {
        std::cout << "TestA()" << std::endl;
    }
    void ReferTestB(std::shared_ptr<TestB> test_ptr)
    {
        m_TestB_Ptr = test_ptr;
    }
    void TestWork()
    {
        std::cout << "~TestA::TestWork()" << std::endl;
    }
    ~TestA()
    {
        std::cout << "~TestA()" << std::endl;
    }
private:
    std::weak_ptr<TestB> m_TestB_Ptr;
};

class TestB
{
public:
    TestB()
    {
        std::cout << "TestB()" << std::endl;
    }

    void ReferTestB(std::shared_ptr<TestA> test_ptr)
    {
        m_TestA_Ptr = test_ptr;
    }
    void TestWork()
    {
        std::cout << "~TestB::TestWork()" << std::endl;
    }
    ~TestB()
    {
        std::shared_ptr<TestA> tmp = m_TestA_Ptr.lock();
        tmp->TestWork();
        std::cout << "2 ref a:" << tmp.use_count() << std::endl;
        std::cout << "~TestB()" << std::endl;
    }
    std::weak_ptr<TestA> m_TestA_Ptr;
};


int main()
{
    std::shared_ptr<TestA> ptr_a = std::make_shared<TestA>();
    std::shared_ptr<TestB> ptr_b = std::make_shared<TestB>();
    ptr_a->ReferTestB(ptr_b);
    ptr_b->ReferTestB(ptr_a);
    std::cout << "1 ref a:" << ptr_a.use_count() << std::endl;
    std::cout << "1 ref b:" << ptr_a.use_count() << std::endl;
    return 0;
}
```

　　运行结果：
![c++11_14](../../../images/c++11_14.PNG)
　     

　　由以上代码运行结果我们可以看到：

　　1、所有的对象最后都能正常释放，不会存在上一个例子中的内存没有释放的问题。

　　2、ptr_a 和ptr_b在main函数中退出前，引用计数均为1，也就是说，在TestA和TestB中对std::weak_ptr的相互引用，不会导致计数的增加。在TestB析构函数中，调用std::shared_ptr<TestA> tmp = m_TestA_Ptr.lock()，把std::weak_ptr类型转换成std::shared_ptr类型，然后对TestA对象进行调用。

## 5、其他

　　本章节介绍的内容如果按照分类来看，也属于以上语法类别，但感觉还是单独拿出来总结好些。

　　下面小节主要介绍std::function、std::bind和lamda表达式的一些特点和用法，希望对读者能有所帮助。

### 5.1、std::function、std::bind封装可执行对象
　　std::bind和std::function也是从boost中移植进来的C++新标准，这两个语法使得封装可执行对象变得简单而易用。此外，std::bind和std::function也可以结合我们一下所说的lamda表达式一起使用，使得可执行对象的写法更加“花俏”。

　　我们下面通过实例一步步了解std::function和std::bind的用法：

1. Test.h文件

```c++
class Test
{
public:
    void Add()
    {
        
    }
};
```

2. main.cpp文件

```c++
#include <functional>
#include <iostream>
#include "Test.h"
int add(int a,int b)
{
    return a + b;
}

int main()
{
    Test test;
    test.Add();
    return 0;
}
```

解释：
上面代码中，我们实现了一个add函数和一个Test类，Test类里面有一个Test函数也有一个函数Add。

OK，我们现在来考虑一下这个问题，假如我们的需求是让Test里面的Add由外部实现，如main.cpp里面的add函数，有什么方法呢？

　　没错，我们可以用函数指针。

　　
* 我们修改Test.h

```c++
class Test
{
public:
    typedef int(*FunType)(int, int);
    void Add(FunType fun,int a,int b)
    {
        int sum = fun(a, b);
        std::cout << "sum:" << sum << std::endl;
    }
};
```

* 修改main.cpp的调用

```c++
#include <functional>
#include <iostream>
#include "Test.h"
int add(int a,int b)
{
    return a + b;
}

int main()
{
    Test test;
    test.Add(add, 1, 2);
    return 0;
}
```

运行结果：

`sum:3`
　　  
　　到现在为止，完美了吗？如果你是Test.h的提供者，你觉得有什么问题？

　　我们把问题升级，假如add实现是在另外一个类内部，如下代码：

```c++
class TestAdd
{
public:
    int Add(int a,int b)
    {
        return a + b;
    }
};

int main()
{
    Test test;
    //test.Add(add, 1, 2);
    return 0;
}
```

　　假如add方法在TestAdd类内部，那你的Test类没辙了，因为Test里的Test函数只接受函数指针。你可能说，这个不是我的问题啊，我是接口的定义者，使用者应该遵循我的规则。但如果现在我是客户，我们谈一笔生意，就是我要购买使用你的Test类，前提是需要支持我传入函数指针，也能传入对象函数，你做不做这笔生意？

　　是的，你可以选择不做这笔生意。我们现在再假设你已经好几个月没吃肉了（别跟我说你是素食主义者），身边的苍蝇肉、蚊子肉啊都不被你吃光了，好不容易等到有机会吃肉，那有什么办法呢？

　　这个时候std::function和std::bind就帮上忙了。

*我们继续修改代码：

* Test.h

```c++
class Test
{
public:
    void Add(std::function<int(int, int)> fun, int a, int b)
    {
        int sum = fun(a, b);
        std::cout << "sum:" << sum << std::endl;
    }
};
```

解释：
Test类中std::function<int(int,int)>表示std::function封装的可执行对象返回值和两个参数均为int类型。

* main.cpp

```c++
int add(int a,int b)
{
    std::cout << "add" << std::endl;
    return a + b;
}

class TestAdd
{
public:
    int Add(int a,int b)
    {
        std::cout << "TestAdd::Add" << std::endl;
        return a + b;
    }
};

int main()
{
    Test test;
    test.Add(add, 1, 2);

    TestAdd testAdd;
    test.Add(std::bind(&TestAdd::Add, testAdd, std::placeholders::_1, std::placeholders::_2), 1, 2);
    return 0;
}
```

解释：
std::bind第一个参数为对象函数指针，表示函数相对于类的首地址的偏移量；
testAdd为对象指针；
std::placeholders::_1和std::placeholders::_2为参数占位符，表示std::bind封装的可执行对象可以接受两个参数。

　　运行结果：
![c++11_15](../../../images/c++11_15.PNG)
　　

　　是的，得出这个结果，你就可以等着吃肉了，我们的Test函数在函数指针和类对象函数中都两种情况下都完美运行。

　　
### 5.2、lamda表达式(匿名函数)

　　在众多的C++11新特性中，个人觉得lamda表达式不仅仅是一个语法新特性，对于没有用过java或C#lamda表达式读者，C++11的lamda表达式在一定程度上还冲击着你对传统C++编程的思维和想法。

　　我们先从一个简单的例子来看看lamda表达式：

复制代码

```c++
1 int main()
2 {
3     auto add = [](int a, int b)->int{
4         return a + b;
5     };
6     int ret = add(1,2);
7     std::cout << "ret:" << ret << std::endl;
8     return 0;
9 }
```

　　解释：

第3至5行为lamda表达式的定义部分

* []：中括号用于控制main函数与内，lamda表达式之前的变量在lamda表达式中的访问形式；
* int a,int b）：为函数的形参
* ->int：lamda表达式函数的返回值定义
* {}:大括号内为lamda表达式的函数体。

运行结果：
`ret: 3`
　　
我使用lamda表达式修改5.1中的例子看看：

* main.cpp

```c++
int main()
{
    Test test;
    test.Add(add, 1, 2);

    TestAdd testAdd;
    test.Add(std::bind(&TestAdd::Add, testAdd, std::placeholders::_1, std::placeholders::_2), 1, 2);

    test.Add([](int a, int b)->int {
        std::cout << "lamda add fun" << std::endl;
        return a + b;
    },1,2);
    return 0;
}
```

　　运行结果：
![c++11_16](../../../images/c++11_16.PNG)
 　 　