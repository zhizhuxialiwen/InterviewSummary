# C++11新特性（10）- 容器的cbegin和cend函数

## 1、const迭代器

先看下面的程序：

```c++
sum = 0;
vector<int> v{1, 2, 3, 4, 5, 6};
vector<int>::iterator it = v.begin();
while(it != v.end()){
       sum += *it;
       it++;
 }
```

代码先是取得了vector的迭代器，然后遍历vector求和。再看下面的代码：

```c++
//错误
sum = 0;
const vector<int> cv{1, 2, 3, 4, 5, 6};
vector<int>::iterator cit = cv.begin();
while(cit != v.end()){
       sum += *cit;
       cit++;
 }
```

这段代码是不能通过编译的，原因是定义的vector是const类型，所以迭代器必须也是const类型。代码需要做如下修改：


```c++
sum = 0;
const vector<int> cv{1, 2, 3, 4, 5, 6};
vector<int>::const_iterator cit = cv.begin();
while(cit != v.end()){
    sum += *cit;
    cit++;
}
```

另一个办法是使用auto类型修饰符：

```c++
sum = 0;
auto ait = cv.begin();
while(ait != cv.end()){ 
       sum += *ait;
       ait++;
 }
```

省去了人工区分迭代器类型的麻烦，又不会妨碍const类型迭代器的功能。

更进一步

vector本身是const类型，生成的迭代器就必须是const类型。这样，在编译层次就避免了可能发生的对vector数据的修改。

还有另外一种情况，数据本身不是const类型，但是从设计的角度来讲有些处理不应该修改该数据。这时也应该要求const类型的迭代器，以避免数据被意外修改。

C++11为此提供了cbegin和cend方法。

```c++
vector<int> v{1, 2, 3, 4, 5, 6};、
auto ait = v.cbegin();
while(ait != v.cend()){
           sum += *ait;
           *ait = sum;  //编译错误
           ait++;
}
```


cbegin()/cend()决定了返回的迭代器类型为const。这时即使vector的类型不是const，也可以防止对该数据的误操作。

作者观点：

为了安全，不该给的不给，不该拿的不拿。做人如此，编程亦然。