# C++11 emplace

emplace操作是C++11新特性，新引入的的三个成员emlace_front、empace 和 emplace_back，这些操作分别对应push_front、insert 和push_back，允许我们将元素放在容器头部、一个指定的位置和容器尾部，而他们之间却有一些不同，**emplace*等操作根据参数执行相应的构造函数，如果传入的参数为容器元素类型则执行拷贝构造（这点和后三则相同），后三者在执行时会调用拷贝构造或则赋值运算符。**

* 两者的区别 

当调用insert时，我们将元素类型的对象传递给insert，元素的对象被拷贝到容器中，而当我们使用emplace时，我们将参数传递元素类型的构造函，emplace使用这些参数在容器管理的内存空间中直接构造元素。

|拷贝构造函数|赋值运算符|
|:--|:--|
|emlace_front|push_front|
|empace |insert|
|emplace_back | push_back|


1. vector

`emplace <->  insert`
`emplace_back​  <-> ​push_back`

2. set

`emplcace <->  insert`

3. map

`emplace <->  insert`


如下我们有一个类定义如下：

```c++
class emTest {
public:
    emTest():data(0){
        cout << "emTest()" << endl;
    }
    emTest(int data):data(data) {
        cout << "emTest(int)" << endl;
    }
    ~emTest() {
        cout << "~emTest()" << endl;
    }
    emTest(const emTest& that) {
        cout << "emTest(&)" << endl;
        data = that.data;
    }
    friend ostream& operator<< (ostream& os, const emTest& that);
private:
    int data;
};
ostream& operator<< (ostream& os, const emTest& that) {
    os << that.data;

    return os;
}
```

通过类定义我们看到，有默认构造函数，接收一个整形参数的构造函数，还有一个拷贝构造函数，咱们可以使用emplace在vector首部放置元素，使用方法如下：

```c++
//调用默认构造函数
vec.emplace(vec.begin()); 
//调用接收int参数的构造函数
vec.emplace(vec.begin(), 4);
//调用拷贝构造函数
emTest em;
vec.emplace(vec.begin(), em);
```

main函数内的测试代码如下：

```c++
int main(){
    vector<emTest> vm;
    vm.reserve(4);
    emTest et(3);
    cout << "------------------------------------" << endl;
    vm.emplace(vm.begin(), et);
    vm.emplace(vm.begin() + 1);
    vm.emplace(vm.begin() + 2, 4);
    for(vector<emTest>::iterator i = vm.begin(); i < vm.end(); i++){
        cout<< *i;
    }
    printf("\n");
    cout << "------------------------------------" << endl;
}
```

测试输出如下：
```c++
emTest(int) ------------------------------------
emTest(&) //vm.emplace(vm.begin(), et); 拷贝构造
emTest()  //vm.emplace(vm.begin() + 1); 默认构造
emTest(int) //vm.emplace(vm.begin() + 2, 4); 接收一个int参数的构造函数
304 //vector内的元素值 ------------------------------------
~emTest()  //程序退出时释放vector内存空间
~emTest()
~emTest()
~emTest()
```


