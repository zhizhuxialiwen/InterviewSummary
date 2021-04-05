# C++ STL 之 unordered_set 介绍

tiny丶 2016-06-23 14:10:58  66105  收藏 34
分类专栏： c++ 文章标签： STL
版权
简介
           C++ 11中出现了两种新的关联容器:unordered_set和unordered_map，其内部实现与set和map大有不同，set和map内部实现是基于RB-Tree，而unordered_set和unordered_map内部实现是基于哈希表(hashtable)，由于unordered_set和unordered_map内部实现的公共接口大致相同，所以本文以unordered_set为例。
        unordered_set是基于哈希表，因此要了解unordered_set，就必须了解哈希表的机制。哈希表是根据关键码值而进行直接访问的数据结构，通过相应的哈希函数(也称散列函数)处理关键字得到相应的关键码值，关键码值对应着一个特定位置，用该位置来存取相应的信息，这样就能以较快的速度获取关键字的信息。比如：现有公司员工的个人信息（包括年龄），需要查询某个年龄的员工个数。由于人的年龄范围大约在[0，200]，所以可以开一个200大小的数组，然后通过哈希函数得到key对应的key-value，这样就能完成统计某个年龄的员工个数。而在这个例子中，也存在这样一个问题，两个员工的年龄相同，但其他信息（如：名字、身份证）不同，通过前面说的哈希函数，会发现其都位于数组的相同位置，这里，就涉及到“冲突”。准确来说，冲突是不可避免的，而解决冲突的方法常见的有：开发地址法、再散列法、链地址法(也称拉链法)。而unordered_set内部解决冲突采用的是----链地址法，当用冲突发生时把具有同一关键码的数据组成一个链表。下图展示了链地址法的使用:

unordered_set
      模板原型:
[cpp]  view plain  copy
template < class Key,  
    class Hash = hash<Key>,  
    class Pred = equal_to<Key>,  
    class Alloc = allocator<Key>  
> class unordered_set;  

       C++ 11中对unordered_set描述大体如下：无序集合容器（unordered_set）是一个存储唯一(unique，即无重复）的关联容器（Associative container），容器中的元素无特别的秩序关系，该容器允许基于值的快速元素检索，同时也支持正向迭代。
       在一个unordered_set内部，元素不会按任何顺序排序，而是通过元素值的hash值将元素分组放置到各个槽(Bucker，也可以译为“桶”），这样就能通过元素值快速访问各个对应的元素（均摊耗时为O（1））。
       原型中的Key代表要存储的类型，而hash<Key>也就是你的hash函数，equal_to<Key>用来判断两个元素是否相等，allocator<Key>是内存的分配策略。一般情况下，我们只关心hash<Key>和equal_to<Key>参数，下面将介绍这两部分。

hash<Key>
      hash<Key>通过相应的hash函数，将传入的参数转换为一个size_t类型值，然后用该值对当前hashtable的bucket取模算得其对应的hash值。而C++标准库，为我们提供了基本数据类型的hash函数：
整型值：bool、char、unsigned char、wchar_t、char16_t、char32_t、short、int、long、long long、unsigned short、unsigned int、unsigned long、unsigned long long。上述的基本数据类型，其标准库提供的hash函数只是简单将其值转换为一个size_t类型值，具体可以参考标准库functional_hash.h头文件，如下所示：
[cpp]  view plain  copy
 /// Primary class template hash.  
  template<typename _Tp>  
    struct hash;  
  
  /// Partial specializations for pointer types.  
  template<typename _Tp>  
    struct hash<_Tp*> : public __hash_base<size_t, _Tp*>  
    {  
      size_t  
      operator()(_Tp* __p) const noexcept  
      { return reinterpret_cast<size_t>(__p); }  
    };  
  
  // Explicit specializations for integer types.  
#define _Cxx_hashtable_define_trivial_hash(_Tp)     \  
  template<>                      \  
    struct hash<_Tp> : public __hash_base<size_t, _Tp>  \  
    {                                                   \  
      size_t                                            \  
      operator()(_Tp __val) const noexcept              \  
      { return static_cast<size_t>(__val); }            \  
    };  
  
  /// Explicit specialization for bool.  
  _Cxx_hashtable_define_trivial_hash(bool)  
  
  /// Explicit specialization for char.  
  _Cxx_hashtable_define_trivial_hash(char)  
  
  /// Explicit specialization for signed char.  
  _Cxx_hashtable_define_trivial_hash(signed char)  
  
  /// Explicit specialization for unsigned char.  
  _Cxx_hashtable_define_trivial_hash(unsigned char)  
  
  /// Explicit specialization for wchar_t.  
  _Cxx_hashtable_define_trivial_hash(wchar_t)  
  
  /// Explicit specialization for char16_t.  
  _Cxx_hashtable_define_trivial_hash(char16_t)  
  
  /// Explicit specialization for char32_t.  
  _Cxx_hashtable_define_trivial_hash(char32_t)  
  
  /// Explicit specialization for short.  
  _Cxx_hashtable_define_trivial_hash(short)  
  
  /// Explicit specialization for int.  
  _Cxx_hashtable_define_trivial_hash(int)  
  
  /// Explicit specialization for long.  
  _Cxx_hashtable_define_trivial_hash(long)  
  
  /// Explicit specialization for long long.  
  _Cxx_hashtable_define_trivial_hash(long long)  
  
  /// Explicit specialization for unsigned short.  
  _Cxx_hashtable_define_trivial_hash(unsigned short)  
  
  /// Explicit specialization for unsigned int.  
  _Cxx_hashtable_define_trivial_hash(unsigned int)  
  
  /// Explicit specialization for unsigned long.  
  _Cxx_hashtable_define_trivial_hash(unsigned long)  
  
  /// Explicit specialization for unsigned long long.  
  _Cxx_hashtable_define_trivial_hash(unsigned long long)  
对于指针类型，标准库只是单一将地址转换为一个size_t值作为hash值，这里特别需要注意的是char *类型的指针，其标准库提供的hash函数只是将指针所指地址转换为一个sieze_t值，如果，你需要用char *所指的内容做hash，那么，你需要自己写hash函数或者调用系统提供的hash<string>。
标准库为string类型对象提供了一个hash函数，即：Murmur hash，。对于float、double、long double标准库也有相应的hash函数，这里，不做过多的解释，相应的可以参看functional_hash.h头文件。
      上述只是介绍了基本数据类型，而在实际应用中，有时，我们需要使用自己写的hash函数，那怎么自定义hash函数？参考标准库基本数据类型的hash函数，我们会发现这些hash函数有个共同的特点：通过定义函数对象，实现相应的hash函数，这也就意味我们可以通过自定义相应的函数对象，来实现自定义hash函数。比如：已知平面上有N，每个点的x轴、y轴范围为[0，100]，现在需要统计有多少个不同点？hash函数设计为：将每个点的x、y值看成是101进制，如下所示:
[cpp]  view plain  copy
#include<bits\stdc++.h>  
using namespace std;  
struct myHash   
{  
    size_t operator()(pair<int, int> __val) const  
    {  
        return static_cast<size_t>(__val.first * 101 + __val.second);  
    }  
};  
int main()  
{  
    unordered_set<pair<int, int>, myHash> S;  
    int x, y;  
    while (cin >> x >> y)  
        S.insert(make_pair(x, y));  
    for (auto it = S.begin(); it != S.end(); ++it)  
        cout << it->first << " " << it->second << endl;  
    return 0;  
}  

equal_to<key>
             该参数用于实现比较两个关键字是否相等，至于为什么需要这个参数？这里做点解释，前面我们说过，当不同关键字，通过hash函数，可能会得到相同的关键字值，每当我们在unordered_set里面做数据插入、删除时，由于unordered_set关键字唯一性，所以我们得确保唯一性。标准库定义了基本类型的比较函数，而对于自定义的数据类型，我们需要自定义比较函数。这里有两种方法:重载==操作符和使用函数对象，下面是STL中实现equal_to<Key>的源代码：
[cpp]  view plain  copy
template<typename _Arg, typename _Result>  
    struct unary_function  
    {  
      /// @c argument_type is the type of the argument  
      typedef _Arg  argument_type;     
  
      /// @c result_type is the return type  
      typedef _Result   result_type;    
    };  
template<typename _Tp>  
    struct equal_to : public binary_function<_Tp, _Tp, bool>  
    {  
      bool  
      operator()(const _Tp& __x, const _Tp& __y) const  
      { return __x == __y; }  
    };  
扩容与缩容
             在vector中，每当我们插入一个新元素时，如果当前的容量（capacity)已不足，需要向系统申请一个更大的空间，然后将原始数据拷贝到新空间中。这种现象在unordered_set中也存在，比如当前的表长为100，而真实存在表中的数据已经大于1000个元素，此时，每个bucker均摊有10个元素，这样就会影响到unordered_set的存取效率，而标准库通过采用某种策略来对当前空间进行扩容，以此来提高存取效率。当然，这里也存在缩容，原理和扩容类似，不过，需要注意的是，每当unordered_set内部进行一次扩容或者缩容，都需要对表中的数据重新计算，也就是说，扩容或者缩容的时间复杂度至少为。
code：
// unordered_set::find
#include <iostream>
#include <string>
#include <unordered_set>
 
int main ()
{
  std::unordered_set<std::string> myset = { "red","green","blue" };
 
  std::string input;
  std::cout << "color? ";
  getline (std::cin,input);
 
  std::unordered_set<std::string>::const_iterator got = myset.find (input);
 
  if ( got == myset.end() )
    std::cout << "not found in myset";
  else
    std::cout << *got << " is in myset";
 
  std::cout << std::endl;
 
  return 0;
}

————————————————
版权声明：本文为CSDN博主「tiny丶」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/vevenlcf/article/details/51743058