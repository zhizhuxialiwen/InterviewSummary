# C++11标准库(STL)使用总结

STL定义了强大的、基于模板的、可复用的组件，实现了许多通用的数据结构及处理这些数据结构的算法。其中包含三个关键组件——容器（container，流行的模板数据结构）、迭代器（iterator）和算法（algorithm）。

|组件|	描述|
|:---|:----|
|容器|	容器是用来管理某一类对象的集合。C++ 提供了各种不同类型的容器，比如 deque、list、vector、map 等。|
|迭代器	|迭代器用于遍历对象集合的元素。这些集合可能是容器，也可能是容器的子集。|
|算法|算法作用于容器。它们提供了执行各种操作的方式，包括对容器内容执行初始化、排序、搜索和转换等操作。|

## 1、容器简介

STL容器，可将其分为四类：序列容器、有序关联容器、无序关联容器、容器适配器。

### 1.1 序列容器

|序号|标准库容器类|	描述|
|:---|:----------|:---|
|1|array|	固定大小，直接访问任意元素|
|2|vector|	从后部进行快速插入和删除操作，直接访问任意元素|
|3|forward_list|	单链表，在任意位置快速插入和删除|
|4|list|双向链表，在任意位置进行快速插入和删除操作|
|5|deque|双向队列，	从前部或后部进行快速插入和删除操作，直接访问任何元素|


### 1.2 有序关联容器（键按顺序保存）

|序号|标准库容器类|	描述|
|:---|:----------|:----|
|1|set	|快速查找，无重复元素|
|2|multiset|快速查找，可有重复元素|
|3|map|	一对一映射，无重复元素，基于键快速查找|
|4|multimap|一对一映射，可有重复元素，基于键快速查找|

### 1.3 无序关联容器

|序号|标准库容器类|	描述|
|:---|:----------|:---|
|1|unordered_set|	快速查找，无重复元素|
|2|unordered_multiset|	快速查找，可有重复元素|
|3|unordered_map	一对一映射，无重复元素，基于键快速查找|
|4|unordered_multimap|	一对一映射，可有重复元素，基于键快速查找|

### 1.4 容器适配器

|序号|标准库容器类|	描述|
|:---|:----------|:----|
|1|stack|后进先出（LIFO）|
|2|queue|先进先出（FIFO）|
|3|priority_queue|优先级最高的元素先出|

1）序列容器描述了线性的数据结构（也就是说，其中的元素在概念上” 排成一行"), 例如数组、向量和 链表。
2）关联容器描述非线性的容器，它们通常可以快速锁定其中的元素。这种容器可以存储值的集合或者键－值对。

C++11中，关联容器中的键是不可变的（不能被修改）。序列容器和关联容器一起称为首类 容器。栈和队列都是在序列容器的基础上加以约束条件得到的，因此STL把stack和queue作为容器适配 器来实现，这样就可以使程序以一种约束方式来处理线性容器。类型string支持的功能跟线性容器一样， 但是它只能存储字符数据。

除此之外，有一些其他的容器种类被称为“ 近容器" (near con即ner): C类型的基于指针的数组用于维护标志位的 bitset, 以及用千进行高速向量运算的 valarray( 这个类对运算进行了优化，也不像首类容器那么复杂）。 这些类型称为“近容器”，是因为它们展现出来的功能与首类容器类似，但是不支持所有的首类容器的功能。

## 2、迭代器简介

迭代器在很多方面与指针类似，也是用于指向首类容器中的元素（还有一些其他用途，后面将会提到）。 迭代器存有它们所指的特定容器的状态信息，即迭代器对每种类型的容器都有一个实现。 有些迭代器的操作在不同容器间是统一的。 例如，＊运算符间接引用一个迭代器，这样就可以使用它所指向的元素。++运算符使得迭代器指向容器中的下一个元素（和数组中指针递增后指向数组的下一个元素类似）。

### 2.1 STL 首类容器提供了成员函数 begin 和 end。

1. 函数 begin 返回一个指向容器中第一个元素的迭代器;
2. 函数 end 返回一个指向容器中最后一个元素的下一个元素（这个元素并不存在，常用于判断是否到达了容器的结束位仅）的迭代器。

3. 如果迭代器 i 指向一个特定的元素，那么 ++i 指向这个元素的下一个元素。* i 指代的是i指向的元素。
4. 从函数 end 中返回的迭代器只在相等或不等的比较中使用，来判断这个“移动的迭代器” （在这里指i)是否到达了容器的末端。

### 2.2 iterator与const_iterator

使用一个 iterator 对象来指向一个可以修改的容器元素，使用一个 const_iterator 对象来指向一个不能修改的容器元素。

|序号|类型|	描述|
|:--|:---|:----|
|1|随机访问迭代器(random access)|在双向迭代的基础上增加了直接访问容器中任意元素的功能，即可以向前或 向后跳转任意个元素|
|2|双向迭代器(bidirectional)|在前向迭代器基础上增加了向后移动的功能。支持多遍扫描算法|
|3|前向迭代器(forword)|	综合输入和输出迭代器的功能，并能保持它们在容器中的位置（作为状态信息），可以使用同一个迭代器两次遍历一个容器（称为多遍扫描算法）|
|4|输出迭代器(output)|用于将元素写入容器。 输出迭代中每次只能向前移动一个元索。 输出迭代器只支持一遍扫描算法，不能使用相同的输出迭代器两次遍历一个序列容器|
|5|输入迭代器(input)|用于从容器读取元素。 输入迭代器每次只能向前移动一个元素。 输入迭代器只支持一遍扫描算法，不能使用相同的输入迭代器两次遍历一个序列容器|

每种容器所支持的迭代器类型决定了这种容器是否可以在指定的 STL 算 法中使用。 支持随机访问迭代器的容器可用千所有的 STL 算法（除了那些需要改变容器大小的算法，这样的算法不能在数组和 array 对象中使用）。 指向 数组的指针可以代替迭代器用于几乎所有的 STL 算法中，包括那些要求随机访问迭代器的算法。 下表显示了每种 STL 容器所支持的迭代器类型。 注意， vector 、 deque 、 list 、 set 、 multiset 、 map 、 multimap( 首类容器）以及 string 和数组都可以使用迭代韶遍历。

|序号|容器|	支持的迭代器类型|
|:---|:---|:--------------|
|1|vector|随机访问迭代器|
|2|array|随机访问迭代器|
|3|deque|随机访问迭代器|
|4|list	|双向迭代器|
|5|forword_list|前向迭代器|
|6|set|双向迭代器|
|7|multiset|双向迭代器|
|8|unordered_set|双向迭代器|
|9|unordered_multiset|双向迭代器|
|10|map|双向迭代器|
|11|multimap|双向迭代器|
|12|unordered_map|双向迭代器|
|13|unordered_multimap|双向迭代器|
|14|stack|不支持迭代器|
|15|queue|不支持迭代器|
|16|priority_queue|不支持迭代器|

下表显示了在 STL容器的类定义中出现的数种预定义的迭代器 typedef。不是每种 typedef 都出现在每个容器中。 我们使用常量版本的迭代器来访问只读容器或不应该被更改的非只读容器，使用反向迭代器来以相反的方向访问容器。

- 为迭代器预先定义

|序号|typedef|++的方向|读写能力|
|:---|:------|:------|:-------|
|1|iterator|向前|读/写|
|2|const_iterator|向前|读|
|3|reverse_iterator|向后|读/写|
|4|const_reverse_iterator|向后|读|

### 2.3 迭代器上的操作

下表显示了可作用在每种迭代器上的操作。 除了给出的对于所有迭代器都有的运算符，迭代器还必须提供默认构造函数、拷贝构造函数和拷贝赋值操作符。 

1）前向迭代器支持＋＋ 和所有的输入和输出迭代器的功能。 2）双向迭代器支持–操作和前向迭代器的功能。 
3）随机访问迭代器支持所有在表中给出的操作。 
4）另外， 对于输入迭代器和输出迭代器，不能在保存迭代器之后再使用保存的值。

1. 适用所有迭代器的操作

|序号|迭代器操作|描述|
|:---|:--------|:---|
|1|++p|前置自增迭代器|
|2|p++|后置自增迭代器|
|3|p=p1|将一个迭代器赋值给另一个迭代器|

2. 输入迭代器	

|序号|迭代器操作|描述|
|:---|:--------|:---|
|1|*p|间接引用一个迭代器|
|2|p->m|使用迭代器读取元素m|
|3|p==p1|比较两个迭代器是否相等|
|4|p!=p1|比较两个迭代器是否不相等|

3. 输出迭代器

|序号|迭代器操作|描述|
|:---|:--------|:---|
|1|*p|间接引用一个迭代器|
|2|p=p1|把一个迭代器赋值给另一个|
|3|前向迭代器|前向迭代器提供了输入和输出迭代器的所有功能|

4. 双向迭代器

|序号|迭代器操作|描述|
|:---|:--------|:---|
|1|–p|q|
|2|p–|后置自减迭代器|

5. 随机访问迭代器

|序号|迭代器操作|描述|
|:---|:--------|:---|
|1|p+=i|迭代器p前进i个位置|
|2|p-=i|迭代器p后退i个位置|
|3|p+i|在迭代器p 的位置上前进i个位置|
|4|p-i|在迭代器p的位置上后退i个位置|
|5|p-p1|表达式的值是一个整数，它代表同一个容器中两个元素间的距离|
|6|p[i]|返回与迭代器p的位置相距i的元素|
|7|p<p1|若迭代器p小于p1(即容器中p在p1前），则返回 true； 否则返回 false。|
|8|p<=p1|若迭代器p小千或等于p1 (即容器中p 在p1前或位咒相同），则返回 true； 否则返回 false。|
|9|p>p1|若迭代器p大于p1(即容器中p在p1后），则返回true,；否则返回false。|
|10|p>=p1|若迭代器p大于或等于p1(即容楛中p在p1后或位置相同），则返回 true； 否则返回 false|

## 3、算法简介

STL提供了可以用于多种容器的算法，其中很多算法都是常用的。插入、删除、搜索、排序及其他一些对部分或全部序列容器和关联容器适用的算法。

STL包含了大约70个标准算法，表格中提供了这些算法的实例及概述。作用在容器元素上的算法只是间接地通过迭代器来实现。很多作用在序列元素上的算法通过一对迭代器定义：第一个迭代器指向这列元素的第一个，第二个迭代器指向最后一个元素之后的位置。 另外，还可以使用相似的方法创建自己的算法，这样它们就能和STL容器及迭代器一起使用了。

### 3.1 算法概述 

算法部分主要由头文件`<algorithm>`，`<numeric>`和`<functional>`组成。

1. `<algorithm>`是所有STL头文件中最大的一个，其中常用到的功能范围涉及到比较、 交换、查找、遍历操作、复制、修改、反转、排序、合并等等。
2. `<numeric>`体积很小，只包括几个在序列上面进行简单数学运算的模板函数，包括加法和乘法在序列上的一些操作。
3. `<functional>`中则定义了一些模板类，用以声明函数对象。
STL提供了大量实现算法的模版函数，只要我们熟悉了STL之后，许多代码可以被大大的化简，只需要通过调用一两个算法模板，就可以完成所需要的功能，从而大大地提升效率。 

```c++
#include <algorithm>
#include <numeric>
#include <functional>
```

### 3.2 STL中算法分类

1. 操作对象

1）直接改变容器的内容。
2）将原容器的内容复制一份,修改其副本,然后传回该副本。

2. 功能:

1）非可变序列算法 

|序号|非可变序列算法|指不直接修改其所操作的容器内容的算法|
|:--|:----|:---|
|1|计数算法|count、count_if|
|2|搜索算法|search、find、find_if、find_first_of、…|
|3|比较算法 |equal、mismatch、lexicographical_compare|

2）可变序列算法 
|可变序列算法 |指可以修改它们所操作的容器内容的算法|
|:--|:----|:---|
|1|删除算法 |remove、remove_if、remove_copy、…|
|2|修改算法 |for_each、transform |
|3|排序算法 |sort、stable_sort、partial_sort、|

3）排序算法：包括对序列进行排序和合并的算法、搜索算法以及有序序列上的集合操作。
4）数值算法：对容器内容进行数值计算。

#### 3.2.1 查找算法(13个)：判断容器中是否包含某个值

 

|序号|函数名|头文件|函数功能|函数|
|:---|:----|:-----|:------|:---|
|1|adjacent_find|`<algorithm>`|在iterator对标识元素范围内,查找一对相邻重复元素,找到则返回指向这对元素的第一个元素的ForwardIterator .否则返回last.重载版本使用输入的二元操作符代替相等的判断。|`template<class FwdIt> FwdIt adjacent_find(FwdIt first, FwdIt last); template<class FwdIt, class Pred> FwdIt adjacent_find(FwdIt first, FwdIt last, Pred pr);`|
|2|binary_search|`<algorithm>`|在有序序列中查找value,找到返回true.重载的版本实用指定的比较函数对象或函数指针来判断相等|`template<class FwdIt, class T> bool binary_search(FwdIt first, FwdIt last, const T& val);template<class FwdIt, class T, class Pred> bool binary_search(FwdIt first, FwdIt last, const T& val,Pred pr);`|
|3|count|`<algorithm>`|利用等于操作符,把标志范围内的元素与输入值比较,返回相等元素个数|`template<class InIt, class Dist> size_t count(InIt first, InIt last,const T& val, Dist& n);`|
|4|count_if|`<algorithm>`|利用输入的操作符,对标志范围内的元素进行操作,返回结果为true的个数|`template<class InIt, class Pred, class Dist> size_t count_if(InIt first, InIt last, Pred pr);`|
|5|equal_range|`<algorithm>`|功能类似equal，返回一对iterator，第一个表示lower_bound，第二个表upper_bound|`template<class FwdIt, class T> pair<FwdIt, FwdIt> equal_range(FwdIt first, FwdIt last,const T& val); template<class FwdIt, class T, class Pred> pair<FwdIt, FwdIt> equal_range(FwdIt first, FwdIt last,const T& val, Pred pr);`|
|6|find|`<algorithm>`|利用底层元素的等于操作符,对指定范围内的元素与输入值进行比较.当匹配时,结束搜索,返回该元素的一个InputIterator|`template<class InIt, class T> InIt find(InIt first, InIt last, const T& val);`|
|7|find_end|`<algorithm>`|在指定范围内查找"由输入的另外一对iterator标志的第二个序列"的最后一次出现.找到则返回最后一对的第一个ForwardIterator,否则返回输入的"另外一对"的第一个ForwardIterator.重载版本使用用户输入的操作符代替等于操作|`template<class FwdIt1, class FwdIt2> FwdIt1 find_end(FwdIt1 first1, FwdIt1 last1,FwdIt2 first2, FwdIt2 last2); template<class FwdIt1, class FwdIt2, class Pred> FwdIt1 find_end(FwdIt1 first1, FwdIt1 last1,FwdIt2 first2, FwdIt2 last2, Pred pr);`|
|8|find_first_of|`<algorithm>`|在指定范围内查找"由输入的另外一对iterator标志的第二个序列"中任意一个元素的第一次出现。重载版本中使用了用户自定义操作符|`template<class FwdIt1, class FwdIt2> FwdIt1 find_first_of(FwdIt1 first1, FwdIt1 last1,FwdIt2 first2, FwdIt2 last2); template<class FwdIt1, class FwdIt2, class Pred> FwdIt1 find_first_of(FwdIt1 first1, FwdIt1 last1,FwdIt2 first2, FwdIt2 last2, Pred pr);`|
|9|find_if|`<algorithm>`|使用输入的函数代替等于操作符执行find|`template<class InIt, class Pred> InIt find_if(InIt first, InIt last, Pred pr);`
|10|lower_bound|`<algorithm>`|返回一个ForwardIterator，指向在有序序列范围内的可以插入指定值而不破坏容器顺序的第一个位置.重载函数使用自定义比较操作|`template<class FwdIt, class T> FwdIt lower_bound(FwdIt first, FwdIt last, const T& val);template<class FwdIt, class T, class Pred> FwdIt lower_bound(FwdIt first, FwdIt last, const T& val, Pred pr);`|
|11|upper_bound|`<algorithm>`|返回一个ForwardIterator,指向在有序序列范围内插入value而不破坏容器顺序的最后一个位置，该位置标志一个大于value的值.重载函数使用自定义比较操作|`template<class FwdIt, class T> FwdIt upper_bound(FwdIt first, FwdIt last, const T& val); template<class FwdIt, class T, class Pred> FwdIt upper_bound(FwdIt first, FwdIt last, const T& val, Pred pr);`|
|12|search|`<algorithm>`|给出两个范围，返回一个ForwardIterator,查找成功指向第一个范围内第一次出现子序列(第二个范围)的位置，查找失败指向last1,重载版本使用自定义的比较操作|`template<class FwdIt1, class FwdIt2> FwdIt1 search(FwdIt1 first1, FwdIt1 last1,FwdIt2 first2, FwdIt2 last2); template<class FwdIt1, class FwdIt2, class Pred> FwdIt1 search(FwdIt1 first1, FwdIt1 last1, FwdIt2 first2, FwdIt2 last2, Pred pr);`|
|13|search_n|`<algorithm>`|在指定范围内查找val出现n次的子序列。重载版本使用自定义的比较操作|`template<class FwdIt, class Dist, class T> FwdIt search_n(FwdIt first, FwdIt last,Dist n, const T& val); template<class FwdIt, class Dist, class T, class Pred> FwdIt search_n(FwdIt first, FwdIt last,Dist n, const T& val, Pred pr);`| 

#### 3.2.2 堆算法(4个)

|序号|函数名|头文件|函数功能|函数|
|:---|:----|:-----|:------|:---|
|1|make_heap|`<algorithm>`|把指定范围内的元素生成一个堆。重载版本使用自定义比较操作|`template<class RanIt> void make_heap(RanIt first, RanIt last); template<class RanIt, class Pred> void make_heap(RanIt first, RanIt last, Pred pr);`
|2|pop_heap|`<algorithm>`|并不真正把最大元素从堆中弹出，而是重新排序堆。它把first和last-1交换，然后重新生成一个堆。可使用容器的back来访问被"弹出"的元素或者使用pop_back进行真正的删除。重载版本使用自定义的比较操作|`template<class RanIt> void pop_heap(RanIt first, RanIt last); template<class RanIt, class Pred> void pop_heap(RanIt first, RanIt last, Pred pr);`|
|3|push_heap|`<algorithm>`|假设first到last-1是一个有效堆，要被加入到堆的元素存放在位置last-1，重新生成堆。在指向该函数前，必须先把元素插入容器后。重载版本使用指定的比较操作|`template<class RanIt>void push_heap(RanIt first, RanIt last); template<class RanIt, class Pred> void push_heap(RanIt first, RanIt last, Pred pr);`|
|4|sort_heap|`<algorithm>`|对指定范围内的序列重新排序，它假设该序列是个有序堆。重载版本使用自定义比较操作|`template<class RanIt> void sort_heap(RanIt first, RanIt last); template<class RanIt, class Pred> void sort_heap(RanIt first, RanIt last, Pred pr);`|

#### 3.2.3 关系算法(8个)

|序号|函数名|头文件|函数功能|函数|
|:---|:----|:-----|:------|:---|
|1|equal|`<algorithm>`|如果两个序列在标志范围内元素都相等，返回true。重载版本使用输入的操作符代替默认的等于操作符|`template<class InIt1, class InIt2> bool equal(InIt1 first, InIt1 last, InIt2 x); template<class InIt1, class InIt2, class Pred> bool equal(InIt1 first, InIt1 last, InIt2 x, Pred pr);`|
|2|includes|`<algorithm>`|判断第一个指定范围内的所有元素是否都被第二个范围包含，使用底层元素的<操作符，成功返回true。重载版本使用用户输入的函数|`template<class InIt1, class InIt2> bool includes(InIt1 first1, InIt1 last1,InIt2 first2, InIt2 last2);template<class InIt1, class InIt2, class Pred> bool includes(InIt1 first1, InIt1 last1,InIt2 first2, InIt2 last2, Pred pr);`|
|3|lexicographical_compare|`<algorithm>`|比较两个序列。重载版本使用用户自定义比较操作|`template<class InIt1, class InIt2> bool lexicographical_compare(InIt1 first1, InIt1 last1,InIt2 first2, InIt2 last2); template<class InIt1, class InIt2, class Pred> bool lexicographical_compare(InIt1 first1, InIt1 last1,InIt2 first2, InIt2 last2, Pred pr);`
|4|max|`<algorithm>`|返回两个元素中较大一个。重载版本使用自定义比较操作|`template<class T> const T& max(const T& x, const T& y); template<class T, class Pred> const T& max(const T&  x, const T& y, Pred pr);`|
|5|max_element|`<algorithm>`|返回一个ForwardIterator，指出序列中最大的元素。重载版本使用自定义比较操作|`template<class FwdIt> FwdIt max_element(FwdIt first, FwdIt last); template<class FwdIt, class Pred> FwdIt max_element(FwdIt first, FwdIt last, Pred pr);`|
|6|min|`<algorithm>`|返回两个元素中较小一个。重载版本使用自定义比较操作|`template<class T> const T& min(const T& x, const T& y); template<class T, class Pred> const T& min(const T& x, const T& y, Pred pr);`|
|6|min_element|`<algorithm>`|返回一个ForwardIterator，指出序列中最小的元素。重载版本使用自定义比较操作|`template<class FwdIt> FwdIt min_element(FwdIt first, FwdIt last); template<class FwdIt, class Pred> FwdIt min_element(FwdIt first, FwdIt last, Pred pr);`|
|7|mismatch|`<algorithm>`|并行比较两个序列，指出第一个不匹配的位置，返回一对iterator，标志第一个不匹配元素位置。如果都匹配，返回每个容器的last。重载版本使用自定义的比较操作|`template<class InIt1, class InIt2> pair<InIt1, InIt2> mismatch(InIt1 first, InIt1 last, InIt2 x); template<class InIt1, class InIt2, class Pred> pair<InIt1, InIt2> mismatch(InIt1 first, InIt1 last, InIt2 x, Pred pr);`|

#### 3.2.4 集合算法(4个)

|序号|函数名|头文件|函数功能|函数|
|:---|:----|:-----|:------|:---|
|1|set_union|`<algorithm>`|构造一个有序序列，包含两个序列中所有的不重复元素。重载版本使用自定义的比较操作|`template<class InIt1, class InIt2, class OutIt> OutIt set_union(InIt1 first1, InIt1 last1, InIt2 first2, InIt2 last2, OutIt x); template<class InIt1, class InIt2, class OutIt, class Pred> OutIt set_union(InIt1 first1, InIt1 last1,InIt2 first2, InIt2 last2,OutIt x, Pred pr);`|
|2|set_intersection|`<algorithm>`|构造一个有序序列，其中元素在两个序列中都存在。重载版本使用自定义的比较操作|`template<class InIt1, class InIt2, class OutIt> OutIt set_intersection(InIt1 first1, InIt1 last1,InIt2 first2, InIt2 last2, OutIt x); template<class InIt1, class InIt2, class OutIt, class Pred> OutIt set_intersection(InIt1 first1, InIt1 last1,InIt2 first2,InIt2 last2, OutIt x, Pred pr);`|
|3|set_difference|`<algorithm>`|构造一个有序序列，该序列仅保留第一个序列中存在的而第二个中不存在的元素。重载版本使用自定义的比较操作|`template<class InIt1, class InIt2, class OutIt> OutIt set_difference(InIt1 first1, InIt1 last1,InIt2 first2, InIt2 last2, OutIt x); template<class InIt1, class InIt2, class OutIt, class Pred> OutIt set_difference(InIt1 first1, InIt1 last1, InIt2 first2, InIt2 last2, OutIt x, Pred pr);`|
|4|set_symmetric_difference|`<algorithm>`|构造一个有序序列，该序列取两个序列的对称差集(并集-交集)|`template<class InIt1, class InIt2, class OutIt> OutIt set_symmetric_difference(InIt1 first1, InIt1 last1, InIt2 first2, InIt2 last2, OutIt x);template<class InIt1, class InIt2, class OutIt, class Pred> OutIt set_symmetric_difference(InIt1 first1, InIt1 last1, InIt2 first2, InIt2 last2, OutIt x, Pred pr);`|

#### 3.2.5 列组合算法(2个)

提供计算给定集合按一定顺序的所有可能排列组合

|序号|函数名|头文件|函数功能|函数|
|:---|:----|:-----|:------|:---|
|1|next_permutation|`<algorithm>`|取出当前范围内的排列，并重新排序为下一个排列。重载版本使用自定义的比较操作|`template<class BidIt> bool next_permutation(BidIt first, BidIt last); template<class BidIt, class Pred> bool next_permutation(BidIt first, BidIt last, Pred pr);`|
|2|prev_permutation|`<algorithm>`|取出指定范围内的序列并将它重新排序为上一个序列。如果不存在上一个序列则返回false。重载版本使用自定义的比较操作|`template<class BidIt> bool prev_permutation(BidIt first, BidIt last); template<class BidIt, class Pred> bool prev_permutation(BidIt first, BidIt last, Pred pr);`|

#### 3.2.6 排序和通用算法(14个)：提供元素排序策略

|序号|函数名|头文件|函数功能|函数|
|:---|:----|:-----|:------|:---|
|1|inplace_merge|`<algorithm>`|合并两个有序序列，结果序列覆盖两端范围。重载版本使用输入的操作进行排序|`template<class BidIt> void inplace_merge(BidIt first, BidIt middle, BidIt last); template<class BidIt, class Pred> void inplace_merge(BidIt first, BidIt middle, BidIt last, Pred pr);`|
|2|merge|`<algorithm>`|合并两个有序序列，存放到另一个序列。重载版本使用自定义的比较|`template<class InIt1, class InIt2, class OutIt> OutIt merge(InIt1 first1, InIt1 last1,InIt2 first2, InIt2 last2, OutIt x);template<class InIt1, class InIt2, class OutIt, class Pred> OutIt merge(InIt1 first1, InIt1 last1,InIt2 first2, InIt2 last2, OutIt x, Pred pr);`|
|3|nth_element|`<algorithm>`|将范围内的序列重新排序，使所有小于第n个元素的元素都出现在它前面，而大于它的都出现在后面。重载版本使用自定义的比较操作|`template<class RanIt> void nth_element(RanIt first, RanIt nth, RanIt last); template<class RanIt, class Pred> void nth_element(RanIt first, RanIt nth, RanIt last, Pred pr);`|
|4|partial_sort|`<algorithm>`|对序列做部分排序，被排序元素个数正好可以被放到范围内。重载版本使用自定义的比较操作|`template<class RanIt> void partial_sort(RanIt first, RanIt middle, RanIt last); template<class RanIt, class Pred> void partial_sort(RanIt first, RanIt middle, RanIt last, Pred pr);`|
|5|partial_sort_copy|`<algorithm>`|与partial_sort类似，不过将经过排序的序列复制到另一个容器|`template<class InIt, class RanIt> RanIt partial_sort_copy(InIt first1, InIt last1,RanIt first2, RanIt last2); template<class InIt, class RanIt, class Pred> RanIt partial_sort_copy(InIt first1, InIt last1,RanIt first2, RanIt last2, Pred pr);`|
|6|partition|`<algorithm>`|对指定范围内元素重新排序，使用输入的函数，把结果为true的元素放在结果为false的元素之前|`template<class BidIt, class Pred> BidIt partition(BidIt first, BidIt last, Pred pr);`|
|7|random_shuffle|`<algorithm>`|对指定范围内的元素随机调整次序。重载版本输入一个随机数产生操作|`template<class RanIt> void random_shuffle(RanIt first, RanIt last); template<class RanIt, class Fun> void random_shuffle(RanIt first, RanIt last, Fun& f);`|
|8|reverse|`<algorithm>`|将指定范围内元素重新反序排序|`template<class BidIt> void reverse(BidIt first, BidIt last);`|
|9|reverse_copy|`<algorithm>`|与reverse类似，不过将结果写入另一个容器|`template<class BidIt, class OutIt> OutIt reverse_copy(BidIt first, BidIt last, OutIt x);`|
|10|rotate|`<algorithm>`|将指定范围内元素移到容器末尾，由middle指向的元素成为容器第一个元素|`template<class FwdIt> void rotate(FwdIt first, FwdIt middle, FwdIt last);`|
|12|rotate_copy|`<algorithm>`|与rotate类似，不过将结果写入另一个容器|`template<class FwdIt, class OutIt> OutIt rotate_copy(FwdIt first, FwdIt middle, FwdIt last, OutIt x);`|
|13|sort|`<algorithm>`|以升序重新排列指定范围内的元素。重载版本使用自定义的比较操作|`template<class RanIt> void sort(RanIt first, RanIt last);template<class RanIt, class Pred> void sort(RanIt first, RanIt last, Pred pr);`|
|14|stable_sort|`<algorithm>`|与sort类似，不过保留相等元素之间的顺序关系|`template<class BidIt> void stable_sort(BidIt first, BidIt last); template<class BidIt, class Pred> void stable_sort(BidIt first, BidIt last, Pred pr);`|
|15|stable_partition|`<algorithm>`|与partition类似，不过不保证保留容器中的相对顺序|`template<class FwdIt, class Pred> FwdIt stable_partition(FwdIt first, FwdIt last, Pred pr);`|

#### 3.2.7 删除和替换算法(15个)


|序号|函数名|头文件|函数功能|函数|
|:---|:----|:-----|:------|:---|
|1|copy|`<algorithm>`|复制序列|`template<class InIt, class OutIt> OutIt copy(InIt first, InIt last, OutIt x);`|
|2|copy_backward|`<algorithm>`|与copy相同，不过元素是以相反顺序被拷贝|`template<class BidIt1, class BidIt2> BidIt2 copy_backward(BidIt1 first, BidIt1 last, BidIt2 x);`|
|3|iter_swap|`<algorithm>`|交换两个ForwardIterator的值|`template<class FwdIt1, class FwdIt2> void iter_swap(FwdIt1 x, FwdIt2 y);`|
|4|remove|`<algorithm>`|删除指定范围内所有等于指定元素的元素。注意，该函数不是真正删除函数。内置函数不适合使用remove和remove_if函数|`template<class FwdIt, class T> FwdIt remove(FwdIt first, FwdIt last, const T& val);`|
|5|remove_copy|`<algorithm>`|将所有不匹配元素复制到一个制定容器，返回OutputIterator指向被拷贝的末元素的下一个位置|`template<class InIt, class OutIt, class T> OutIt remove_copy(InIt first, InIt last, OutIt x, const T& val);`|
|6|remove_if|`<algorithm>`|删除指定范围内输入操作结果为true的所有元素|`template<class FwdIt, class Pred> FwdIt remove_if(FwdIt first, FwdIt last, Pred pr);`|
|7|remove_copy_if|`<algorithm>`|将所有不匹配元素拷贝到一个指定容器|`template<class InIt, class OutIt, class Pred> OutIt remove_copy_if(InIt first, InIt last, OutIt x, Pred pr);`|
|8|replace|`<algorithm>`|将指定范围内所有等于vold的元素都用vnew代替|`template<class FwdIt, class T> void replace(FwdIt first, FwdIt last,const T& vold, const T& vnew);`|
|9|replace_copy|`<algorithm>`|与replace类似，不过将结果写入另一个容器|`template<class InIt, class OutIt, class T> OutIt replace_copy(InIt first, InIt last, OutIt x,const T& vold, const T& vnew);`|
|10|replace_if|`<algorithm>`|将指定范围内所有操作结果为true的元素用新值代替|`template<class FwdIt, class Pred, class T> void replace_if(FwdIt first, FwdIt last,Pred pr, const T& val);`|
|11|replace_copy_if|`<algorithm>`|与replace_if，不过将结果写入另一个容器|`template<class InIt, class OutIt, class Pred, class T> OutIt replace_copy_if(InIt first, InIt last, OutIt x, Pred pr, const T& val);`|
|12|swap|`<algorithm>`|交换存储在两个对象中的值|`template<class T> void swap(T& x, T& y);`|
|13|swap_range|`<algorithm>`|将指定范围内的元素与另一个序列元素值进行交换|`template<class FwdIt1, class FwdIt2> FwdIt2 swap_ranges(FwdIt1 first, FwdIt1 last, FwdIt2 x);`|
|14|unique|`<algorithm>`|清除序列中重复元素，和remove类似，它也不能真正删除元素。重载版本使用自定义比较操作|`template<class FwdIt> FwdIt unique(FwdIt first, FwdIt last); template<class FwdIt, class Pred> FwdIt unique(FwdIt first, FwdIt last, Pred pr);`|
|15|unique_copy|`<algorithm>`|与unique类似，不过把结果输出到另一个容器|`template<class InIt, class OutIt> OutIt unique_copy(InIt first, InIt last, OutIt x); template<class InIt, class OutIt, class Pred> OutIt unique_copy(InIt first, InIt last, OutIt x, Pred pr);`|

 
#### 3.2.8 生成和变异算法(6个)


|序号|函数名|头文件|函数功能|函数|
|:---|:----|:-----|:------|:---|
|1|fill|`<algorithm>`|将输入值赋给标志范围内的所有元素|`template<class FwdIt, class T> void fill(FwdIt first, FwdIt last, const T& x);`|
|2|fill_n|`<algorithm>`|将输入值赋给first到first+n范围内的所有元素|`template<class OutIt, class Size, class T> void fill_n(OutIt first, Size n, const T& x);`|
|3|for_each|`<algorithm>`|用指定函数依次对指定范围内所有元素进行迭代访问，返回所指定的函数类型。该函数不得修改序列中的元素|`template<class InIt, class Fun> Fun for_each(InIt first, InIt last, Fun f);`|
|4|generate|`<algorithm>`|连续调用输入的函数来填充指定的范围|`template<class FwdIt, class Gen> void generate(FwdIt first, FwdIt last, Gen g);`|
|5|generate_n|`<algorithm>`|与generate函数类似，填充从指定iterator开始的n个元素|`template<class OutIt, class Pred, class Gen> void generate_n(OutIt first, Dist n, Gen g);`|
|6|transform|`<algorithm>`|将输入的操作作用与指定范围内的每个元素，并产生一个新的序列。重载版本将操作作用在一对元素上，另外一个元素来自输入的另外一个序列。结果输出到指定容器|`template<class InIt, class OutIt, class Unop> OutIt transform(InIt first, InIt last, OutIt x, Unop uop); template<class InIt1, class InIt2, class OutIt, class Binop> OutIt transform(InIt1 first1, InIt1 last1, InIt2 first2,OutIt x, Binop bop);`|

#### 3.2.9 算数算法(4个)


|序号|函数名|头文件|函数功能|函数|
|:---|:----|:-----|:------|:---|
|1|accumulate|`<numeric>`|iterator对标识的序列段元素之和，加到一个由val指定的初始值上。重载版本不再做加法，而是传进来的二元操作符被应用到元素上|`template<class InIt, class T> T accumulate(InIt first, InIt last, T val); template<class InIt, class T, class Pred> T accumulate(InIt first, InIt last, T val, Pred pr);`|
|2|partial_sum|`<numeric>`|创建一个新序列，其中每个元素值代表指定范围内该位置前所有元素之和。重载版本使用自定义操作代替加法|`template<class InIt, class OutIt> OutIt partial_sum(InIt first, InIt last,OutIt result); template<class InIt, class OutIt, class Pred> OutIt partial_sum(InIt first, InIt last,OutIt result, Pred pr);`|
|3|product|`<numeric>`|对两个序列做内积(对应元素相乘，再求和)并将内积加到一个输入的初始值上。重载版本使用用户定义的操作|`template<class InIt1, class InIt2, class T> T product(InIt1 first1, InIt1 last1,Init2 first2, T val); template<class InIt1, class InIt2, class T,class Pred1, class Pred2> T product(InIt1 first1, InIt1 last1,Init2 first2, T val, Pred1 pr1, Pred2 pr2);`|
|4|adjacent_difference|`<numeric>`|创建一个新序列，新序列中每个新值代表当前元素与上一个元素的差。重载版本用指定二元操作计算相邻元素的差|`template<class InIt, class OutIt> OutIt adjacent_difference(InIt first, InIt last,OutIt result); template<class InIt, class OutIt, class Pred> OutIt adjacent_difference(InIt first, InIt last,OutIt result, Pred pr);`|

## 4、常用算法汇总

### 4.1 常用的查找算法

adjacent_find()（ adjacent 是邻近的意思）,binary_search(),count(), count_if(),equal_range(),find(),find_if()。

## 4.2 常用的排序算法

merge(),sort(),random_shuffle()（shuffle是洗牌的意思） ,reverse()。

## 4.3 常用的拷贝和替换算法

copy(), replace(), replace_if(),swap()

## 4.4 常用的算术和生成算法

accumulate()（ accumulate 是求和的意思）,fill()。

## 4.5 常用的集合算法

set_union(),set_intersection(), set_difference()。

## 4.6 常用的遍历算法

for_each(), transform()（ transform 是变换的意思）

## 5、常用头文件

C++标准库分为很多部分，每个部分都有自己的头文件。头文件包含了形成标准库哥哥部分的相关函数的原型。头文件中还包含了各种各样的类类型和函数类型的定义，以及这些函数所需的常量。头文件可以“指示”编译器怎么处理标准库和用户编写的组件的接口问题。

|序号|C++标准库头|	文件说明|
|:---|:------|:-----|
|1|< iostream >|	包含C++标准输入和输出函数的原型。|
|2|< iomanip >|	包含格式化数据流的流操纵符的函数原型。|
|3|< cmath >|	包含数学库函数原型。|
|4|< cstdlib >	|包含数转换为文本、文本转换为数、内存分配、随机数及其他各种工具函数的函数原型。|
|5|< ctime >|包含处理时间和日期的函数原型和类型。|
|6|< array >,< vector >,< list >,< forword_list >,< deque >,< queue >,< stack >,< map >,< unordered_map >,< unordered_set >,< set >,< bitset >|这些头文件包含了实现C++标准库容器的类。 在程序执行期间 ， 容器保存数据。|
|7|< ctype >|包含测试字符特定属性（例如字符是否是数字字符或者标点符号）的函数原型和 用于将小写字母转换成大写字母 、 将大写字母转换成小写字母的函数原型。|
|8|< cstring >|包含C风格字符串处理函数的函数原型。|
|9|< typeinfo >|包含运行时类型识别（在执行时确定数据类型）的类。|
|10|< exception >,< stdexcept >|这两个头文件包含用于异常处理的类。|
|11|< memory >	|包含被C++标准库用来向C++标准库容器分配内存的类和函数。|
|12|< fstream >	|包含执行由磁盘文件输入和向磁盘文件输出的函数的函数原型。|
|13|< string >|包含来自C++标准库的 string类的定义。|
|14|< sstream >	|包含执行由内存字符串输人和向内存字符串输出的函数的函数原型。|
|15|< functional >|	包含C++标准库算法所用的类和函数。|
|16|< iterator >|包含访问C++标准库容器中数据的类。|
|17|< algorithm >|包含操作C++标准库容器中数据的函数。|
|18|< cassert >|包含为辅助程序调试而添加诊断的宏。|
|19|< cfloat >|包含系统的浮点数长度限制。|
|20|< climits >	|包含系统的整数长度限制。|
|21|< cstdio >	|包含C风格标准输入和输出库函数的函数原型。|
|22|< locale >|包含流处理通常所用的类和函数，用来处理不同语言自然形式的数据（例如货币格式 、 排序字符串、字符表示 ， 等等）。|
|23|< limits >|	包含为各计算机平台定义数字数据类型限制的类。|
|24|< utility >|包含被许多C++标准库头文件所用的类和函数。|

## 参考

[STL总结](https://blog.csdn.net/weixin_41921520/article/details/100585664)