C++11-unique_ptr详解

防秃从C++练起 2021-03-08 20:17:06  74  收藏
分类专栏： C/C++
版权
在《拥抱智能指针，告别内存泄露》中说到了内存泄漏问题，也提到了C++中的智能指针基本原理，今天就来说说类模板unique_ptr。
在此之前，先回答读者的一个提问：C语言中该怎么办？有几点建议：

编写时尽量遵循函数内申请，函数内释放的原则
注意成对编写malloc和free
使用静态扫描工具，如《pclint检查》
使用内存检测工具，如valgrind
unique_ptr
一个unique_ptr独享它指向的对象。也就是说，同时只有一个unique_ptr指向同一个对象，当这个unique_ptr被销毁时，指向的对象也随即被销毁。使用它需要包含下面的头文件

#include<memory>
1
基本使用
常见方式有：

std::unique_ptr<int> up;//可以指向int的unique_ptr，不过是空的
up = std::unique_ptr<int>(new int(12));
1
2
此时它是一个空的unique_ptr，即没有指向任何对象。

//unique_ptr<T>
std::unique_ptr<string> up1(new string("bianchengzhuji"));
std::unique_ptr<int[]> up2(new int[10]);//数组需要特别注意
1
2
3
也可以指向一个new出来的对象。

你也可以结合上面两种方式，如：

std::unique_ptr<int> up;//声明空的unique_ptr
int *p= new int(1111);
up.reset(p);//令up指向新的对象，p为内置指针
1
2
3
通常来说，在销毁对象的时候，都是使用delete来销毁，但是也可以使用指定的方式进行销毁。举个简单的例子，假如你打开了一个连接，获取到了一个文件描述符，现在你想通过unique_ptr来管理，希望在不需要的时候，能够借助unique_ptr帮忙关闭它。

#include<iostream>
#include<unistd.h>
#include<memory>
void myClose(int *fd)
{
    close(*fd);
}
int main()
{
    int socketFd = 10;//just for example
    std::unique_ptr<int,decltype(myClose)*> up(&socketFd,myClose);
    /*下面是另外两种写法，后面一种是使用lambda表达式*/
    //std::unique_ptr<int,void(*)(int*)> up(&socketFd,myClose);
    //std::unique_ptr<int,void(*)(int*)> ip(&socketFd,[](int *fd){close(*fd);});
    return 0;
}
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
它的用法如下：

std::unique_ptr<T,D> up(t,d);
std::unique_ptr<T,D> up(d);//空的unique_ptr
1
2
含义分别如下：

T unique_ptr管理的对象类型
D 删除器类型
t unique_ptr管理的对象
d 删除器函数/function对象等，用于释放对象指针
这里使用了decltype(myClose)*用于获取myClose函数的类型，表明它是一个指针类型，即函数指针，它传入参数是int。你也可以使用注释中的方式。关于函数指针，可参考《》。

即便后面执行出现异常时，这个socket连接也能够正确关闭。

后面我们也可以看到，与shared_ptr不同，unique_ptr在编译时绑定删除器，避免了运行时开销。

释放指向的对象
一般来说，unique_ptr被销毁时（如离开作用域），对象也就自动释放了，也可以通过其他方式下显示释放对象。如：

up = nullptr;//置为空，释放up指向的对象
up.release();//放弃控制权，返回裸指针，并将up置为空
up.reset();//释放up指向的对象
1
2
3
可以看到release和reset的区别在于，前者会释放控制权，返回裸指针，你还可以继续使用。而后者直接释放了指向对象。

unique_ptr不支持普通的拷贝和赋值
需要特别注意的是，由于unique_ptr“独有”的特点，它不允许进行普通的拷贝或赋值，例如：

std::unique_ptr<int> up0;
std::unique_ptr<int> up1(new int(1111));
up0 = up1 //错误，不可赋值
std::unique_ptr<int> up2(up1);//错误，不支持拷贝
1
2
3
4
总之记住，既然unique_ptr是独享对象，那么任何可能被共享的操作都是不允许的，但是可以移动。

移动unique_ptr的对象
虽然unique_ptr独享对象，但是也可以移动，即转移控制权。如：

std::unique_ptr<int> up1(new int(42));
std::unique_ptr<int> up2(up1.release());
1
2
up2接受up1 release之后的指针，或者：

std::unique_ptr<int> up1(new int(42));
std::unique_ptr<int> up2;
up2.reset(up1.release());
1
2
3
或者使用move：

std::unique_ptr<int> up1(new int(42));
std::unique_ptr<int> up2(std::move(up1));
1
2
在函数中的使用
既然unique_ptr独享对象，那么就无法直接作为参数，应该怎么办呢？

作为参数
如果函数以unique_ptr作为参数呢？如果像下面这样直接把unique_ptr作为参数肯定就报错了，因为它不允许被复制：

#include<iostream>
#include<memory>
void test(std::unique_ptr<int> p)
{
    *p = 10;
}
int main()
{
    std::unique_ptr<int> up(new int(42));
    test(up);//试图传入unique_ptr，编译报错
    std::cout<<*up<<std::endl;
    return 0;
}
1
2
3
4
5
6
7
8
9
10
11
12
13
上面的代码编译将直接报错。

当然我们可以向函数中传递普通指针，使用get函数就可以获取，如：

#include<iostream>
#include<memory>
void test(int *p)
{
    *p = 10;
}
int main()
{
    std::unique_ptr<int> up(new int(42));
    test(up.get());//传入裸指针作为参数
    std::cout<<*up<<std::endl;//输出10
    return 0;
}
1
2
3
4
5
6
7
8
9
10
11
12
13
或者使用引用作为参数：

#include<iostream>
#include<memory>
void test(std::unique_ptr<int> &p)
{
    *p = 10;
}
int main()
{
    std::unique_ptr<int> up(new int(42));
    test(up);
    std::cout<<*up<<std::endl;//输出10
    return 0;
}
1
2
3
4
5
6
7
8
9
10
11
12
13
当然如果外部不再需要使用了，那么你完全可以转移，将对象交给你调用的函数管理，这里可以使用move函数：

#include<iostream>
#include<memory>
void test(std::unique_ptr<int> p)
{
    *p = 10;
}
int main()
{
    std::unique_ptr<int> up(new int(42));
    test(std::unique_ptr<int>(up.release()));
    //test(std::move(up));//这种方式也可以
    return 0;
}
1
2
3
4
5
6
7
8
9
10
11
12
13
作为返回值
unique_ptr可以作为参数返回：

#include<iostream>
#include<memory>
std::unique_ptr<int> test(int i)
{
    return std::unique_ptr<int>(new int(i));
}
int main()
{
    std::unique_ptr<int> up = test(10);
    //std::shared_ptr<int> up = test(10);
    std::cout<<*up<<std::endl;
    return 0;
}
1
2
3
4
5
6
7
8
9
10
11
12
13
你还可以把unique_ptr转换为shared_ptr使用，如注释行所示。

为什么优先选用unique_ptr
回到标题的问题，问什么优先选用unique_ptr。

避免内存泄露
避免更大开销
第一点相信很好理解，自动管理，不需要时即释放，甚至可以防止下面这样的情况：

int * p = new int(1111);
/*do something*/
delete p;
1
2
3
如果在do something的时候，出现了异常，退出了，那delete就永远没有执行的机会，就会造成内存泄露，而如果使用unique_ptr就不会有这样的困扰了。

第二点为何这么说？因为相比于shared_ptr，它的开销更小，甚至可以说和裸指针相当，它不需要维护引用计数的原子操作等等。

所以说，如果有可能，优先选用unique_ptr。

总结
本文介绍了uniqueptr的基本使用情况和使用场景，它能够有效地避免内存泄露并且效率可控，因此如果能够满足需求，则优先选择unique\ptr。