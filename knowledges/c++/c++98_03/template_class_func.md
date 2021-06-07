# C++函数模板与类模板实例解析

本文针对C++函数模板与类模板进行了较为详尽的实例解析，有助于帮助读者加深对C++函数模板与类模板的理解。具体内容如下：

泛型编程（Generic Programming）是一种编程范式，通过将类型参数化来实现在同一份代码上操作多种数据类型，泛型是一般化并可重复使用的意思。泛型编程最初诞生于C++中，目的是为了实现C++的STL（标准模板库）。

模板（template）是泛型编程的基础，一个模板就是一个创建类或函数的蓝图或公式。例如，当使用一个vector这样的泛型类型或者find这样的泛型函数时，我们提供足够的信息，将蓝图转换为特定的类或函数。

## 1、函数模板

一个通用的函数模板（function template）就是一个公式，可用来生成针对特定类型或特定值的函数版本。模板定义以关键字template开始，后面跟一个模板参数列表，列表中的多个模板参数（template parameter）以逗号分隔。模板参数表示在类或函数定义中用到的类型或值。

### 1.1、类型参数

一个模板类型参数（type parameter）表示的是一种类型。我们可以将类型参数看作类型说明符，就像内置类型或类类型说明符一样使用。类型参数前必须使用关键字`class 或typename`：

```c++
template <typename T> // typename和class一样的 
T function(T* p) 
{ 
  T tmp = *p;  // 临时变量类型为T 
  //... 
  return tmp;  // 返回值类型为T 
} 
```

关键字typename和class是一样的作用，但显然typename比class更为直观，它更清楚地指出随后的名字是一个类型名。

编译器用模板类型实参为我们实例化（instantiate）特定版本的函数，一个版本称做模板的一个实例（instantiation）。当我们调用一个函数模板时，编译器通常用函数实参来为我们推断模板实参。当然如果函数没有模板类型的参数，则我们需要特别指出来：

```c++
int a = 10; 
cout << function(&a) << endl;   // 编译器根据函数实参推断模板实参 
  
cout << function<int>(&a) << endl;  // <int>指出模板参数为int 
```

### 1.2、非类型参数

在模板中还可以定义非类型参数（nontype parameter），一个非类型参数表示一个值而非一个类型。我们通过一个特定的类型名而非关键字class或typename来指定非类型参数：

```c++
// 整形模板 
template<unsigned M, unsigned N> 
void add() 
{ 
  cout<< M+N << endl; 
} 
  
// 指针 
template<const char* C> 
void func1(const char* str) 
{ 
  cout << C << " " << str << endl; 
} 
  
// 引用 
template<char (&R)[9]> 
void func2(const char* str) 
{ 
  cout << R << " " << str << endl; 
} 
  
// 函数指针 
template<void (*f)(const char*)> 
void func3(const char* c) 
{ 
  f(c); 
} 
  
void print(const char* c) { cout << c << endl;} 
  
char arr[9] = "template";  // 全局变量，具有静态生存期 
  
int main() 
{ 
  add<10, 20>(); 
  func1<arr>("pointer"); 
  func2<arr>("reference"); 
  func3<print>("template function pointer"); 
  return 0; 
}
```

当实例化时，非类型参数被一个用户提供的或编译器推断出的值所替代。一个非类型参数可以是一个整型，或者是一个指向对象或函数的指针或引用：绑定到整形（非类型参数）的实参必须是一个常量表达式，绑定到指针或引用（非类型参数）的实参必须具有静态的生存期（比如全局变量），不能把普通局部变量 或动态对象绑定到指针或引用的非类型形参。

## 2、类模板

相应的，类模板（class template）是用来生成类的蓝图。与函数模板的不同之处是，编译器不能为类模板推断模板参数类型，所以我们必须显式的提供模板实参。与函数模板一样，类模板参数可以是类型参数，也可以是非类型参数，这里就不再赘述了。

```c++
template<typename T> 
class Array { 
public: 
  Array(T arr[], int s); 
  void print(); 
private: 
  T *ptr; 
  int size; 
}; 
  
// 类模板外部定义成员函数 
template<typename T> 
Array<T>::Array(T arr[], int s) 
{ 
  ptr = new T[s]; 
  size = s; 
  for(int i=0; i<size; ++i) 
    ptr[i]=arr[i]; 
} 
  
template<typename T> 
void Array<T>::print() 
{ 
  for(int i=0; i<size; ++i) 
    cout << " " << *(ptr+i); 
  cout << endl; 
} 
  
int main() 
{ 
  char a[5] = {'J','a','m','e','s'}; 
  Array<char> charArr(a, 5); 
  charArr.print(); 
  
  int b[5] = { 1, 2, 3, 4, 5}; 
  Array<int> intArr(b, 5); 
  intArr.print(); 
  
  return 0; 
}
```

### 2.1 类模板的成员函数

与其他类一样，我们既可以在类模板内部，也可以在类模板外部定义其成员函数。定义在类模板之外的成员函数必须以关键字template开始，后接类模板参数列表。

```c++
template <typename T> 
return_type class_name<T>::member_name(parm-list) { } 
```

默认情况下，对于一个实例化了的类模板，其成员函数只有在使用时才被实例化。如果一个成员函数没有被使用，则它不会被实例化。

### 2.2 类模板和友元

当一个类包含一个友元声明时，类与友元各自是否是模板是相互无关的。如果一个类模板包含一个非模板的友元，则友元被授权可以访问所有模板的实例。如果友元自身是模板，类可以授权给所有友元模板的实例，也可以只授权给特定实例。

// 前置声明，在将模板的一个特定实例声明为友元时要用到 

```c++
template<typename T> class Pal; 
  
// 普通类 
class C { 
  friend class Pal<C>; // 用类C实例化的Pal是C的一个友元 
  template<typename T> friend class Pal2; //Pal2所有实例都是C的友元;无须前置声明 
}; 
  
// 模板类 
template<typename T> class C2 { 
  // C2的每个实例将用相同类型实例化的Pal声明为友元,一对一关系 
  friend class Pal<T>; 
  // Pal2的所有实例都是C2的每个实例的友元，不需要前置声明 
  template<typename X> friend class Pal2;  
  // Pal3是普通非模板类，它是C2所有实例的友元 
  friend class Pal3; 
}; 
```

### 2.3 类模板的static成员

类模板可以声明static成员。类模板的每一个实例都有其自己独有的static成员对象，对于给定的类型X，所有class_name<X>类型的对象共享相同的一份static成员实例。
```c++
template<typename T> 
class Foo { 
public: 
  void print(); 
  //...其他操作 
private: 
  static int i; 
}; 
  
template<typename T> 
void Foo<T>::print() 
{ 
  cout << ++i << endl; 
} 
  
template<typename T> 
int Foo<T>::i = 10; // 初始化为10 
  
int main() 
{ 
  Foo<int> f1; 
  Foo<int> f2; 
  Foo<float> f3; 
  f1.print();  // 输出11 
  f2.print();  // 输出12 
  f3.print();  // 输出11 
  return 0; 
}
```

我们可以通过类类型对象来访问一个类模板的static对象，也可以使用作用域运算符（::）直接访问静态成员。类似模板类的其他成员函数，一个static成员函数也只有在使用时才会实例化。