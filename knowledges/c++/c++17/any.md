# C++17 any类（万能容器）详解


any是一种很特殊的容器，它只能容纳一个元素，但这个元素可以是任意的类型，可以是基本数据类型（int，double，char，float...）也可以是复合数据类型（类、结构体），那它究竟有什么用？目前来说我没发现什么太大的作用，如果看官发现什么有用的作用，分享一下=.=

操作演示：

```c++
#include<iostream>
#include<any>
#include<vector>
#include<set>
using namespace std;
int main() {
	any Int = 69,		//整型
		Double = 8.8,	//浮点型
		CStr = "hello",			//字符数组
		Str = string("world!");		//string类
	vector<any> anys = { Int,Double,CStr,Str };	//可直接放入序列容器中
	//set < any > any_set{ Int,Double,CStr,Str };	//Error:不能直接放入关联容器内,需要提供operator<，但 any 很难比较
	cout << *Int._Cast<int>() << endl;	//T* _Cast<typenampe T>():返回any对象的"值"，但需要一个模板参数确定对象值的类型
	cout << *Double._Cast<double>() << endl;
	cout << *CStr._Cast<char const *>() << endl;
	cout << *Str._Cast<string>() << endl;
	return 0;
}
```

输出结果：
//OutPut:
//69
//8.8
//hello
//world!
 
首先，VS默认是没有支持C++17的，需要自己修改设置，如果不能使用any，请修改标准
 VS修改C++标准（支持C++17）

* any成员：

void reset() ： 重置（清空）any对象。
bool has_value() const ： 值判断，有值返回true，无值返回false。
const type_info& type() const ： 获取类型，返回一个type_info&类的常引用，其成员有类型的哈希值及类型名称（const char*）
void swap(any & other) ：与另一个any对象交换“值”。

* 模板函数（简化版）：

`template<typename Type>`    

Type& emplace(Type&& Args) ：修改any对象的值。

`template<typename Type>`

Type* _Cast() ：返回any对象值的地址

```c++
#include<iostream>
#include<any>
using namespace std;
int main() {
	any a = 1,b=string("hello");
	cout << a.has_value() << endl;	//true
	a.reset();
	cout << a.has_value() << endl;	//false
	a.emplace<double>(1.1);	//修改值
	a.swap(b);
	cout << a.type().name() << "：" << *a._Cast<string>() << endl;		//string ：hello
	cout << b.type().name() << "：" << *b._Cast<double>() << endl;		//double ：1.1
	return 0;
}
```

* any实现原理：

通过使用模板构造函数擦除模板类的参数类型。

存储：定义一个基类Base，再派生一个模板类Data，对二者再进行一次封装，构造一个Any类，使用Any类的模板构造函数来构造一个Data对象，这样就能存储任何数据类型。

取值：只能存，但还无法把元素取出，所以Any必须有一个基类Base指针的成员变量，存储构造好的Data对象，使用模板函数_Cast()，利用其模板参数Type，进行一个再将Base类强制转换为Data<Type> 对象。

基本上是这个原理吧。

附超简版代码：

```c++
#include<iostream>
#include<any>
using namespace std;
class Any {
public:
	template<typename T>
	Any(T t) :base(new Data<T>(t)) {}		//模板构造函数
	template<typename T>
	T _Cast() {
		return dynamic_cast<Data<T>*>(base.get())->value;		//强制转换
	}
private:
	class Base {
	public:
		virtual ~Base() {}		//确定Base为多态类型
	};
	template <typename T>
	class Data :public Base {
	public:
		Data(T t) :value(t) {}
		T value;
	};
	unique_ptr<Base> base;					//基类指针
};
int main() {
	Any a(string("s123")), b = 1, c = 12.0;
	cout << a._Cast<string>() << endl;			/*调用get函数必须要填写模板参数，我感觉这样大大折扣了any的作用*/
	cout << b._Cast<int>() << endl;				/*我想能否在构造的时候就确认参数的类型，保证_Cast调用的时候不需要使用模板参数*/
	cout << c._Cast<double>();
	return 0;
}
```

