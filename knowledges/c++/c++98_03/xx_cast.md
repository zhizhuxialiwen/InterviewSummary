# C++强制类型转换：static_cast、dynamic_cast、const_cast、reinterpret_cast

## 1、 c强制转换与c++强制转换

 c语言强制类型转换主要用于基础的数据类型间的转换，语法为：

(type-id)expression//转换格式1

type-id(expression)//转换格式2

c++除了能使用c语言的强制类型转换外，还新增了四种强制类型转换：static_cast、dynamic_cast、const_cast、reinterpret_cast，主要运用于继承关系类间的强制转化，语法为：

```c++
static_cast<new_type>      (expression)
dynamic_cast<new_type>     (expression) 
const_cast<new_type>       (expression) 
reinterpret_cast<new_type> (expression)
```

备注：new_type为目标数据类型，expression为原始数据类型变量或者表达式。

《Effective C++》中将c语言强制类型转换称为旧式转型，c++强制类型转换称为新式转型。

## 2、 static_cast、dynamic_cast、const_cast、reinterpret_cast

### 2.1 static_cast

static_cast相当于传统的C语言里的强制转换，该运算符把expression转换为new_type类型，用来强迫隐式转换，例如non-const对象转为const对象，编译时检查，用于非多态的转换，可以转换指针及其他，但没有运行时类型检查来保证转换的安全性。它主要有如下几种用法：

1. 用于类层次结构中基类（父类）和派生类（子类）之间指针或引用的转换。
进行上行转换（把派生类的指针或引用转换成基类表示）是安全的；
进行下行转换（把基类指针或引用转换成派生类表示）时，由于没有动态类型检查，所以是不安全的。
2. 用于基本数据类型之间的转换，如把int转换成char，把int转换成enum。这种转换的安全性也要开发人员来保证。
3. 把空指针转换成目标类型的空指针。
4. 把任何类型的表达式转换成void类型。

注意：static_cast不能转换掉expression的const、volatile、或者__unaligned属性。
 

* 基本类型数据转换举例如下：

```c++
char a = 'a';
int b = static_cast<char>(a);//正确，将char型数据转换成int型数据

double *c = new double;
void *d = static_cast<void*>(c);//正确，将double指针转换成void指针

int e = 10;
const int f = static_cast<const int>(e);//正确，将int型数据转换成const int型数据

const int g = 20;
int *h = static_cast<int*>(&g);//编译错误，static_cast不能转换掉g的const属性
```
 
* 类上行和下行转换：

```c++
if(Derived *dp = static_cast<Derived *>(bp)){//下行转换是不安全的
  //使用dp指向的Derived对象  
}
else{
  //使用bp指向的Base对象  
}

if(Base*bp = static_cast<Derived *>(dp)){//上行转换是安全的
  //使用bp指向的Derived对象  
}
else{
  //使用dp指向的Base对象  
}
```

### 2.2 dynamic_cast

```c++
dynamic_cast<type*>(e)
dynamic_cast<type&>(e)
dynamic_cast<type&&>(e)
```

　　type必须是一个类类型，在第一种形式中，type必须是一个有效的指针，在第二种形式中，type必须是一个左值，在第三种形式中，type必须是一个右值。在上面所有形式中，e的类型必须符合以下三个条件中的任何一个：e的类型是是目标类型type的公有派生类、e的类型是目标type的共有基类或者e的类型就是目标type的的类型。如果一条dynamic_cast语句的转换目标是指针类型并且失败了，则结果为0。如果转换目标是引用类型并且失败了，则dynamic_cast运算符将抛出一个std::bad_cast异常(该异常定义在typeinfo标准库头文件中)。e也可以是一个空指针，结果是所需类型的空指针。

dynamic_cast主要用于类层次间的上行转换和下行转换，还可以用于类之间的交叉转换（cross cast）。
在类层次间进行上行转换时，dynamic_cast和static_cast的效果是一样的；
在进行下行转换时，dynamic_cast具有类型检查的功能，比static_cast更安全。dynamic_cast是唯一无法由旧式语法执行的动作，也是唯一可能耗费重大运行成本的转型动作。

（1）指针类型
举例，Base为包含至少一个虚函数的基类，Derived是Base的共有派生类，如果有一个指向Base的指针bp，我们可以在运行时将它转换成指向Derived的指针，代码如下：

```c++
if(Derived *dp = dynamic_cast<Derived *>(bp)){
  //使用dp指向的Derived对象  
}
else{
  //使用bp指向的Base对象  
}
```

值得注意的是，在上述代码中，if语句中定义了dp，这样做的好处是可以在一个操作中同时完成类型转换和条件检查两项任务。

（2）引用类型

因为不存在所谓空引用，所以引用类型的dynamic_cast转换与指针类型不同，在引用转换失败时，会抛出std::bad_cast异常，该异常定义在头文件typeinfo中。

```c++
void f(const Base &b){
 try{
   const Derived &d = dynamic_cast<const Base &>(b);  
   //使用b引用的Derived对象
 }
 catch(std::bad_cast){
   //处理类型转换失败的情况
 }
}
```
 
### 2.3 const_cast

const_cast，用于修改类型的const或volatile属性。 

该运算符用来修改类型的const(唯一有此能力的C++-style转型操作符)或volatile属性。除了const 或volatile修饰之外， new_type和expression的类型是一样的。
1. 常量指针被转化成非常量的指针，并且仍然指向原来的对象；
2. 常量引用被转换成非常量的引用，并且仍然指向原来的对象；
3. const_cast一般用于修改底指针。如const char *p形式。
 
举例转换如下：

```c++
const int g = 20;
int *h = const_cast<int*>(&g);//去掉const常量const属性

const int g = 20;
int &h = const_cast<int &>(g);//去掉const引用const属性

 const char *g = "hello";
char *h = const_cast<char *>(g);//去掉const指针const属性
```
 
### 2.4 reinterpret_cast

new_type必须是一个指针、引用、算术类型、函数指针或者成员指针。它可以把一个指针转换成一个整数，也可以把一个整数转换成一个指针（先把一个指针转换成一个整数，再把该整数转换成原类型的指针，还可以得到原先的指针值）。

reinterpret_cast意图执行低级转型，实际动作（及结果）可能取决于编辑器，这也就表示它不可移植。

　　举一个错误使用reintepret_cast例子，将整数类型转换成函数指针后，vc++在执行过程中会报"...中的 0xxxxxxxxx 处有未经处理的异常: 0xC0000005: Access violation"错误：

```c++
#include <iostream>
using namespace std;
int output(int p){
    cout << p <<endl;
　　return 0;
}

typedef int (*test_func)(int );//定义函数指针test_func
int main(){
    int p = 10;
    test_func fun1 = output;
    fun1(p);//正确
    test_func fun2 = reinterpret_cast<test_func>(&p);
    fun2(p);//...处有未经处理的异常: 0xC0000005: Access violation
    return 0;
}
```

IBM的C++指南、C++之父Bjarne Stroustrup的FAQ网页和MSDN的Visual C++也都指出：错误的使用reinterpret_cast很容易导致程序的不安全，只有将转换后的类型值转换回到其原始类型，这样才是正确使用reinterpret_cast方式。

　　MSDN中也提到了，实际中可将reinterpret_cast应用到哈希函数中，如下（64位系统中需将unsigned int修改为unsigned long）：
```c++
// expre_reinterpret_cast_Operator.cpp
// compile with: /EHsc
#include <iostream>

// Returns a hash code based on an address
unsigned short Hash( void *p ) {
   unsigned int val = reinterpret_cast<unsigned int>( p );
   return ( unsigned short )( val ^ (val >> 16));
}

using namespace std;
int main() {
   int a[20];
   for ( int i = 0; i < 20; i++ )
      cout << Hash( a + i ) << endl;
}
```
 

　　另外，static_cast和reinterpret_cast的区别主要在于多重继承，比如
 
```c++
class A {
    public:
    int m_a;
};
 
class B {
    public:
    int m_b;
};
 
class C : public A, public B {};
```

　　那么对于以下代码：

```　
C c;
printf("%p, %p, %p", &c, reinterpret_cast<B*>(&c), static_cast <B*>(&c));
```

 　前两个的输出值是相同的，最后一个则会在原基础上偏移4个字节，这是因为static_cast计算了父子类指针转换的偏移量，并将之转换到正确的地址（c里面有m_a,m_b，转换为B*指针后指到m_b处），而reinterpret_cast却不会做这一层转换。
　因此, 你需要谨慎使用 reinterpret_cast。
 
## 3、 c++强制转换注意事项

新式转换较旧式转换更受欢迎。原因有二，一是新式转型较易辨别，能简化“找出类型系统在哪个地方被破坏”的过程；二是各转型动作的目标愈窄化，编译器愈能诊断出错误的运用。
尽量少使用转型操作，尤其是dynamic_cast，耗时较高，会导致性能的下降，尽量使用其他方法替代。
 

## 参考资料：

a):http://en.cppreference.com/w/cpp/language/static_cast

b):http://en.cppreference.com/w/cpp/language/dynamic_cast

c):http://en.cppreference.com/w/cpp/language/const_cast

d):http://en.cppreference.com/w/cpp/language/reinterpret_cast

e):《Effective C++》条款27:尽量少做转型动作

f): 百度百科

g) 《C++ Primer》