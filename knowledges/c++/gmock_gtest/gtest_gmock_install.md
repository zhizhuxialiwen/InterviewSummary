# gtest/gmock


闲来无事，想尝试一下gtest/gmock，根据下载的源码包里有README，并根据自己安装过程补充记录如下，以便以后查询
## 1、linux安装gtest/gmock

### 1.1 方式一：

1. 将googletest-release-1.10.0.tar.gz解压，并进入解压后的目录


网页：https://github.com/google/googletest
下载googletest
git clone https://github.com/google/googletest.git

```
tar -xzvf googletest-release-1.10.0.tar.gz
cd googletest-release-1.10.0
```

2. 编译

```
mkdir build
cd build
cmake ../CMakeLists.txt
cd ..
make
```

3. 拷贝库文件和包含文件

```
sudo cp lib/* /usr/lib
sudo cp -r googletest/include/gtest /usr/include
sudo cp -r googlemock/include/gmock/  /usr/include/
```

4. 编译

```
gtest: g++ gmock_eq.cpp -o out -lgtest -lpthread -std=c++11
gmock: g++ gmock_eq.cpp -o out -lgtest -lgmock -lpthread -std=c++11
```

### 1.2 方式三: Linux下CMake工程中gtest&gmock的安装与使用


1. 源码

gtest
《玩转Google开源C++单元测试框架Google Test系列(gtest)(总)》
http://www.cnblogs.com/coderzh/archive/2009/04/06/1426755.html

gmock
《Google Mock 入门概述》
http://www.cnblogs.com/welkinwalker/archive/2011/11/29/2267225.html

这篇文章主要想讲的就是如何在Linux下编译gtest&gmock的代码，并且在cmake工程中配置它。

2. 编译gtest&gmock

首先，下载代码，地址如下：
https://github.com/google/googletest
第二步，下载完成后用unzip命令解压代码；
第三步，解压完成后，进入目录，利用g++来编译代码，命令如下：
* gtest

```
g++ -isystem ${GTEST_DIR}/include -I${GTEST_DIR} \
    -pthread -c ${GTEST_DIR}/src/gtest-all.cc
ar -rv libgtest.a gtest-all.o
```

Note that (We need `-pthread` as Google Test uses threads.)
* gmock

```
g++ -isystem ${GTEST_DIR}/include -I${GTEST_DIR} \
    -isystem ${GMOCK_DIR}/include -I${GMOCK_DIR} \
    -pthread -c ${GTEST_DIR}/src/gtest-all.cc
g++ -isystem ${GTEST_DIR}/include -I${GTEST_DIR} \
    -isystem ${GMOCK_DIR}/include -I${GMOCK_DIR} \
    -pthread -c ${GMOCK_DIR}/src/gmock-all.cc
ar -rv libgmock.a gtest-all.o gmock-all.o
```

其中，GTEST_DIR、GMOCK_DIR就是代码的位置。

3. 配置Cmake工程

第一步，在工程目录下创建lib文件夹和include文件夹；

第二步，把GTEST_DIR和GMOCK_DIR目录下的include文件夹复制到工程的include中，以及把之前编译的libgmock.a和libgtest.a复制到lib下；

第三步，在CMakeLists.txt中添加相应代码，例如：
```
cmake_minimum_required(VERSION 3.2)
project(gtest_test)
LINK_DIRECTORIES( ${PROJECT_SOURCE_DIR}/lib )
INCLUDE_DIRECTORIES(  ${PROJECT_SOURCE_DIR}/include )
add_executable(gtest_test Foomain.cpp)
#下面这条语句中，链接了gmock、gtest以及pthread
#pthread是必要的，因为前两者会用到
TARGET_LINK_LIBRARIES(gtest_test gmock gtest pthread)
install(TARGETS gtest_test RUNTIME DESTINATION bin)
```

## 2、gtest与gmock案例

### 2.1 gtest测试案例

```c++
#include <gtest/gtest.h> 

int fun1() {
  return 10;
}

class test : public ::testing::Test{
public:
  int fun2() {
    return 1;
  };
};

TEST(fun1, test_fun) {
  EXPECT_EQ(10, fun1());        //单个函数的测试
}

TEST_F(test, test_class) {
  EXPECT_EQ(10, fun2());       //类中函数的测试
}

int main(int argc, char **argv) {
  ::testing::InitGoogleTest(&argc, argv);
  return RUN_ALL_TESTS();
}
```

编译: `g++ gtest_test.cpp -o gtest_test -lgtest -lpthread`

### 2.2 gmock测试案例

```c++
#include <gtest/gtest.h>  
#include <gmock/gmock.h>  
using namespace testing;  
class A {
public:
    int set(int num) {
        value = num;
        return num;
    }
    int get() {
    return value;
    }
    int value;
};

class MockA : public A {
public:
    MOCK_METHOD1(set, int(int num));
    MOCK_METHOD0(get, int());

};

TEST(Atest, getnum)  
{  
    MockA m_A;  
    int a = 10;
    EXPECT_CALL(m_A, set(_)).WillRepeatedly(Return(a));
    int k = m_A.set(200);
    EXPECT_EQ(10, k);  
}

int main(int argc, char *argv[]) {
    ::testing::InitGoogleTest(&argc, argv);
    return RUN_ALL_TESTS();
}
```
编译
`g++ gmock_test.cpp -o gmock_test -lgtest -lgmock -lpthread`


最后附上GTest/GMock学习文档：

GTest学习文章：

http://www.cnblogs.com/coderzh/archive/2009/04/06/1426755.html

GMock三篇学习文章：

1，http://code.google.com/p/googlemock/wiki/ForDummies
2，http://code.google.com/p/googlemock/wiki/CheatSheet

3，http://code.google.com/p/googlemock/wiki/CookBook
