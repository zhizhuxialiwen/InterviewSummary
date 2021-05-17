# c++11 for_each()

## 1、for循环多了新的语法

```c++
#include<iostream>
#include<algorithm>
#include<vector>
using namespace std;

int main()
{
    vector<int> vec;
    for(int i=0;i<10;++i)
    {
        vec.push_back(i);
    }

    for(int &it: vec)
    {
        cout<<it<<" ";
    }
    cout<<endl;
    return 0;
}
```

输出结果：
0 1 2 3 4 5 6 7 8 9

## 2、for_each

c++11 还增加了for_each，目前很多地方使用for_each语法来使用STL遍历容器。
需要包含头文件 #include”algorithm”

```c++
#include<iostream>
#include<algorithm>
#include<vector>
using namespace std;

int main()
{
    vector<int> vec;
    for(int i=0;i<10;++i)
    {
        vec.push_back(i);
    }
    for_each(vec.begin(),vec.end(), [](int i)->void{ cout << i <<" "; }); 


    cout<<endl;
    return 0;
}
```

输出结果：
0 1 2 3 4 5 6 7 8 9