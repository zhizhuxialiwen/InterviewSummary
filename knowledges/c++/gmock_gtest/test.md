# googletest-gmock使用示例


gmock是一个非常好用的单元测试工具。它可以模拟接口，对指定的类进行测试。

官方说明： https://github.com/google/googletest/blob/v1.8.x/googlemock/docs/Documentation.md

注意：googletest的版本为v1.8.0。不同版本的使用方法不同。

以下代码是可以正常使用的。

1. demo.h

```c++
#include <string>
using namespace std;

class Parent {
  public:
    virtual ~Parent() {}

    virtual int getNum() const = 0;
    virtual void setResult(int value) = 0;
    virtual void print(const string &str) = 0;
    virtual int calc(int a, double b) = 0;
};

class Target {
  public:
    Target(Parent *parent) :
        parent_(parent)
    { }

    int doThis()
    {
        int v = parent_->getNum();
        parent_->setResult(v);
        while (v -- > 0) {
            parent_->print(to_string(v));
        }
        return parent_->getNum();
    }

    int doThat()
    {
        return parent_->calc(1, 2.2);
    }

  private:
    Parent *parent_;
};
```

2. demo.cpp

```c++
#include "demo.h"
#include <gmock/gmock.h>

class MockParent : public Parent {
  public:
//! MOCK_[CONST_]METHODx(方法名, 返回类型(参数类型列表));
    MOCK_CONST_METHOD0(getNum, int());    //! 由于 getNum() 是 const 成员函数，所以要使用 MOCK_CONST_METHODx
    MOCK_METHOD1(setResult, void(int));
    MOCK_METHOD1(print, void(const string &));
    MOCK_METHOD2(calc, int(int, double));
};

using ::testing::Return;
using ::testing::_;

TEST(demo, 1) {
    MockParent p;
    Target t(&p);

    //! 设置 p.getNum() 方法的形为
    EXPECT_CALL(p, getNum())
        .Times(2)    //! 期望被调两次
        .WillOnce(Return(2))   //! 第一次返回值为2
        .WillOnce(Return(10)); //! 第二次返回值为10

    //! 设置 p.setResult(), 参数为2的调用形为
    EXPECT_CALL(p, setResult(2))
        .Times(1);

    EXPECT_CALL(p, print(_))  //! 表示任意的参数，其中 _ 就是 ::testing::_ ，如果冲突要注意
        .Times(2);

    EXPECT_EQ(t.doThis(), 10);
}

TEST(demo, 2) {
    MockParent p;
    Target t(&p);

    EXPECT_CALL(p, calc(1, 2.2))
        .Times(1)
        .WillOnce(Return(3));

    EXPECT_EQ(t.doThat(), 3);
}
```

3. 编译


Makefile文件

```s
OBJS=demo.o

TARGET=test

$(TARGET) : $(OBJS)
    g++ -o $@ $^ $(LDFLAGS) -lgmock -lgmock_main -lpthread

clean:
    rm *.o
```