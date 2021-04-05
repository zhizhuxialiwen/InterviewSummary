# c++17 constexpr关键字 tcy

tcy23456 2020-03-27 09:17:00  389  收藏
分类专栏： C/C++
版权
 
## 1、constexpr

### 1.1.定义

声明一个对象（表达式或函数返回值）编译期为常量(即编译时即可计算结果)

### 1.2. const与constexpr区别

const常量--    运行时常量（必须初始化）
constexpr常量--编译期常量（必须初始化）
 
### 1.3. 用途

1）常量声明,表达式,函数(参数,返回值),成员函数，类构造函数,实现编译期循环或递归
2）if constexpr(bool)根据模板参数类型做不同处理，让代码变得简洁让模版具备选择功能
3）消除enable_if，替换#ifdef宏

### 1.4.constexpr函数限制

 1）必须有返回值且仅有一条return语句（禁用非常量数据，函数，全局变量） 
 2）使用前必须定义，在该编译单元内被定义前不能调用  
 3）常量构造函数限制：函数体必为空,初始化列表只能由常量表达式来赋值

## 2、实例： 

1. 实例1：
 
```c++
#include <iostream>
#include<string>
#include<cassert>
using namespace std;

//实例1.1：constexpr函数
constexpr double get_pi() { return 3.1415926; }

template <typename T>
constexpr auto foo(T x) { return [x](auto y) { return x + y; }; }

//实例1.2：
constexpr  int increase_x(int& x) { return  ++x; }
constexpr  int foo(int x) {
    //错误：increase_x(x)不是核心常数表达式，因为x的生命周期从表达式increase_x(x)之外开始
    //constexpr  int y = increase_x(x); 
    constexpr  int y = 0;
    return y;
}

//实例1.3：
constexpr  int bar(int x)  //ok：x不需要用核心初始化/常量表达式
{
    int y = increase_x(x); 
    return y; 
} 

//用值2初始化y bar(1)是一个核心常量表达式，因为x的生命周期始于表达式bar(1)内
constexpr  int y = bar(1);  

void test(){
    static  const  int a = 10;
    constexpr  const  int& ra = a;  
    constexpr  int ia = a;  
    
    const  int b = 20;
    constexpr  const  int& rb = b;  
    constexpr  int ib = b;  
}
//实例1.4：
struct Boy {
    constexpr Boy(const char* n, int a) : name(n), age(a) {}
    constexpr const char* get_name() const { return name; }//不能为string
    constexpr int get_age() const { return age; }

private:
    const char* name{ "Tom" }; 
    int age{ 0 };
};

void test_Boy() {
    constexpr Boy boy{ "Bob",10 };// 必须是常量表达式
    constexpr const char* name = boy.get_name();//Bob
    constexpr int age = boy.get_age();//10
}
//实例1.5：
void test1() {
    //1.常量:
    constexpr const double pi = 3.14;//常量必须赋初值
    constexpr double e = pi + 1;     //常量表达式中都必须为常量

    //2.函数：
    constexpr auto f1 = [](int x, int y) {
        auto func = [=] {return x; }; 
        return [=] { return func() + y; }; 
    };

    constexpr auto f2 = [](int x, int y) {return x + y; };
    
    auto f3 = [](int x, int y) {return x + y; };

    int y1 = f1(1, 2)() + f2(3, 4);    //编译期计算=10
    int y2 = f2(y1, 1);               //运行期计算=4
    //constexpr int y3 = f3(2, 3);    //错误
}

int main() {
    test();
    test1();
    test_Boy();
}	
```

实例2：

```c++
#include <iostream>
#include<string>
#include<cassert>
using namespace std;

//2.1.C++17之前编译期整数加法：
template<int N>
constexpr int sum() { return N; }

template <int N, int N2, int... Ns>
constexpr int sum() { return N + sum<N2, Ns...>(); }

//2.2.C++17更简洁代码:
template <int N=-1, int... Ns>//若去掉=-1必须至少一个参数
constexpr auto sum17() {
    if constexpr (sizeof...(Ns) == 0)return N;
    else return N + sum17<Ns...>();
}

//2.3.C++17的fold expression：
template<typename ...Args>
constexpr int sum(Args... args) { return (0 + ... + args); }

int main() {
    assert((sum<>()) == 0);
    assert((sum<1>()) == 1);
    assert((sum<1, 2>()) == 3);

    assert((sum17<>()) == -1);
    assert((sum17<1>()) == 1);
    assert((sum17<1, 2>()) == 3);

    assert(sum() == 0);
    assert(sum(1) == 1);
    assert(sum(1,2) == 3);	
}	
实例3：if constexpr消除enable_if
#include <iostream>
#include <type_traits>
#include <string>
#include<cassert>

using namespace std;

//1.笨拙方法：如根据类型来选择函数不得不分开几个函数来写
template<typename T>
std::enable_if_t<std::is_integral<T>::value, std::string> to_str(T t) { return std::to_string(t); }

template<typename T>
std::enable_if_t<!std::is_integral<T>::value, T> to_str(T t) { return t; }

//2.通过if constexpr可以消除enable_if_t:
template<typename T>
auto to_str_Cpp17(T t) {
    if constexpr (std::is_integral<T>::value) return std::to_string(t);
    else return t;
}

//3.下面写法有问题:
template<typename T>
auto to_str_Cpp17_1(T t) {
    if constexpr (std::is_integral<T>::value) return std::to_string(t);
    //else       //当输入非int程序正常，前两行代码将失效只有最后一句有效
    return t;    //当输入int时错误：所有返回表达式必须推导为相同类型，有两个return
}

int main() {
    assert(to_str("Tom") == string("Tom") && to_str(3.14) == 3.14);
    assert(to_str_Cpp17("Tom") == string("Tom") && to_str_Cpp17(3.14) == 3.14);

    //cout << to_str_Cpp17_1(3) << endl;                           //错误
    assert(to_str_Cpp17_1("Tom") == string("Tom") && to_str_Cpp17_1(3.14) == 3.14);
}
实例4：constexpr if用来替换#ifdef宏

#include <type_traits>
enum class OS { Linux, Mac, Windows }; 

#ifdef __linux__
constexpr OS the_os = OS::Linux;
#elif __APPLE__
constexpr OS the_os = OS::Mac;
#elif __WIN32
constexpr OS the_os = OS::Windows;
#endif 

/*
#if defined( __linux) || defined(__linux__) || defined(linux)
constexpr OS the_os = OS::Linux;
#elif defined(__APPLE__)
constexpr OS the_os = OS::Mac;
#elif defined(_WIN32) || defined(__WIN32__) || defined(WIN32)
constexpr OS the_os = OS::Windows;
#endif 
*/

void do_something() {
        if constexpr (the_os == OS::Linux) {}        //Linux操作
        else if constexpr (the_os == OS::Mac) {}     //Mac操作
        else if constexpr (the_os == OS::Windows) {} //windows操作
} 
```
 
