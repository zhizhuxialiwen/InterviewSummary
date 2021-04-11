# google mock模拟C++对象

单元测试 googletest testing 测试
版权
google mock是用来配合google test对C++项目做单元测试的。它依赖于googletest（参见我上篇文章《如何用googletest写单元测试》： http://blog.csdn.net/russell_tao/article/details/7333226），下面我来说说linux上怎么用它来做单元测试。

本文包括：1、如何获取、编译google mock；2、如何使用gmock（下面用gmock来代称google mock）配合gtest做单元测试。

## 1、如何获取、编译google mock

gmock的当前版本与gtest一样，是1.6.0。可以从这个网址获取：http://code.google.com/p/googlemock/downloads/list。

下载到压缩包解压后，下面我们开始编译出静态库文件（必须得自己编译出），以在我们自己的单元测试工程中使用。

与gtest相同，我们执行完./configure; make后，不能执行make install，理由与上篇相同。

验证这个包有没有问题，依然可以执行如下命令：
```
cd make
make
./gmock_test
```

如果你看到类似下文的输出屏幕，证明你的机器运行gmock没有问题。
```
[----------] Global test environment tear-down
[==========] 13 tests from 3 test cases ran. (2 ms total)
[  PASSED  ] 13 tests.
```

这时还没有编译出我们要的libgmock.a呢。继续在gmock解包目录下执行：
```
g++ -I gtest/include/ -I gtest/ -I include/ -I ./ -c gtest/src/gtest-all.cc 
g++ -I gtest/include/ -I gtest/ -I include/ -I ./ -c src/gmock-all.cc       
ar -rv libgmock.a gtest-all.o gmock-all.o 
```

如此，当前目录下会链接出我们需要的libgmock.a。注意，这个gmock.a静态库里，把gtest需要的gtest-all.cc都编译进来了，所以我们的单元测试工程只需要链接libgmock，不再需要链接上文说的libgtest了。

## 2、如何使用gmock

首先，编译我们自己的单元测试工程时，需要在makefile里加入以下编译选项：-I${GTEST_DIR}/include -I${GMOCK_DIR}/include，这两个目录我们自己从上面的包里拷贝出来即可。链接时，需要加上libgmock.a。

还是以一个例子来说明怎么在mock对象的情况下写单元测试。
我现在有一个生产者消费者网络模型，消费者（例如client）会先发TCP请求到我的SERVER去订阅某个对象。生产者（另一台SERVER）产生关于某个对象的事件后发给我的SERVER后，我的SERVER再把事件发给消费者。

就是这么简单。


我现在想写一个单元测试，主要测试代码逻辑，不想去管网络包的收发这些事情。

我现在有两个类，一个叫CSubscriber，它封装为一个订阅的消费者，功能主要是操作网络，包括网络收发包，协议解析等。另一个叫CSubEventHandler，它主要做逻辑处理，去操作CSubscriber对象，例如epoll返回读事件后，会构造一个CSubscriber对象，然后CSubEventHandler::handleRead方法就来处理这个CSubscriber对象。


我单元测试的目的是，测试CSubEventHandler::handleRead的业务逻辑，我同时也想测试CSubscriber方法里的协议解析逻辑，但是对于CSubscriber封装的读写包部分，我希望可以mock成我想要的网络包。

怎么做呢？

### 2.1、先mock一个CSubscriber类如下：

```c++
class MockCSubscriber : public CSubscriber
{
public:
	MockCSubscriber(int fd):CSubscriber(fd){}
	MOCK_METHOD1(readBuf, int(int len));
	MOCK_METHOD1(writeBuf, int(int len));
	MOCK_METHOD0(closeSock, void());
};
```

其中，CSubscriber的构造方法必须有一个int型的fd，而readBuf和writeBuf都只接收一个int型的参数，而closeSock方法 没有参数传递。于是我使用了MOCK_METHOD0和MOCK_METHOD1这两个宏来声明想MOCK的方法。这两个宏的使用很简单，解释下：
`MOCK_METHOD#1(#2, #3(#4) )`

`#2`是你要mock的方法名称！`#1`表示你要mock的方法共有几个参数，`#4`是这个方法具体的参数，`#3`表示这个方法的返回值类型。

很简单不是？！


### 2.2、如果只关心mock方法的返回值。

这里用到一个宏ON_CALL。看例子：

`ON_CALL(subObj, readBuf(1000)).WillByDefault(Return(blen));`

什么意思呢？再用刚才的解释方法：
`ON_CALL(#1, #2(#3)).WillByDefault(Return(#4));`

`#1`表示mock对象。就像我上面所说，对CSubscriber我定义了一个Mock类，那么就必须生成相应的mock对象，例如：
`MockCSubscriber subObj(5);`

`#2`表示想定义的那个方法名称。上例中我想定义readBuf这个方法的返回值。
`#3`表示readBuf方法的参数。这里的1000表示，只有调用CSubscriber::readBuf同时传递参数为1000时，才会用到ON_CALL的定义。

`#4`表示调用CSubscriber::readBuf同时传递参数为1000时，返回blen这个变量的值。


### 2.3、如果还希望mock方法有固定的被调用方式

这里用到宏EXPECT_CALL，看个例子：

`EXPECT_CALL(subObj, readBuf(1000)).Times(1);`

很相似吧?最后的Times表示，只希望readBuf在传递参数为1000时，被调用且仅被调用一次。

其实这些宏有很复杂的用法的，例如：

```c++
EXPECT_CALL(subObj, readBuf(1000))
    .Times(5)
    .WillOnce(Return(100))
    .WillOnce(Return(150))
    .WillRepeatedly(Return(200));
```

表示，readBuf希望被调用五次，第一次返回100，第二次返回150，后三次返回200。如果不满足，会报错。

### 2.4、实际的调用测试

其实调用跟上篇googletest文章里的测试是一致的，我这里只列下上文的完整用例代码（不包括被测试类的实现代码）：

1. mian.cpp

```c++
#include "gtest/gtest.h"
#include "gmock/gmock.h"
#include "CSubscriber.h"
#include "CPublisher.h"
#include "CSubEventHandler.h"
#include "CPubEventHandler.h"
 
using ::testing::AtLeast;
using testing::Return;
 
 
class MockCSubscriber : public CSubscriber
{
public:
	MockCSubscriber(int fd):CSubscriber(fd){}
	MOCK_METHOD1(readBuf, int(int len));
	MOCK_METHOD1(writeBuf, int(int len));
	MOCK_METHOD0(closeSock, void());
};
 
class MockCPublisher : public CPublisher
{
public:
	MockCPublisher(int fd):CPublisher(fd){}
	MOCK_METHOD1(readBuf, int(int len));
	MOCK_METHOD1(writeBuf, int(int len));
	MOCK_METHOD0(closeSock, void());
};
 
 
TEST(subpubHandler, sub1pub1) {
	MockCSubscriber subObj(5);
	MockCPublisher pubObj(5);
 
	subObj.m_iRecvBufLen = 1000;
	pubObj.m_iRecvBufLen = 1000;
 
	char* pSubscribeBuf = "GET / HTTP/1.1\r\nobject: /tt/aa\r\ntime: 112\r\n\r\n";
	char* pMessageBuf = "GET / HTTP/1.1\r\nobject: /tt/aa\r\ntime: 112\r\nmessage: tttt\r\n\r\n";
	subObj.m_pRecvBuf = pSubscribeBuf;
	int blen = strlen(pSubscribeBuf);
	subObj.m_iRecvPos = blen;
 
	pubObj.m_pRecvBuf = pMessageBuf;
	int mlen = strlen(pMessageBuf);
	pubObj.m_iRecvPos = mlen;
 
 
	ON_CALL(subObj, readBuf(1000)).WillByDefault(Return(blen));
	ON_CALL(subObj, writeBuf(CEventHandler::InternalError.size())).WillByDefault(Return(0));
 
	CSubEventHandler subHandler(NULL);
	CPubEventHandler pubHandler(NULL);
 
	CHashTable ht1(100);
	CHashTable ht2(100);
	subHandler.initial(100, &ht1, &ht2);
	pubHandler.initial(100, &ht1, &ht2);
 
	EXPECT_CALL(subObj, readBuf(1000)).Times(1);
	//EXPECT_CALL(subObj, closeSock()).Times(1);
	EXPECT_CALL(subObj, writeBuf(4)).Times(1);
 
	EXPECT_TRUE(subHandler.handleRead(&subObj));
 
	ON_CALL(pubObj, readBuf(1000)).WillByDefault(Return(mlen));
	ON_CALL(pubObj, writeBuf(4)).WillByDefault(Return(0));
 
	EXPECT_CALL(pubObj, readBuf(1000)).Times(1);
	EXPECT_CALL(pubObj, closeSock()).Times(1);
	EXPECT_CALL(pubObj, writeBuf(CEventHandler::Success.size())).Times(1);
 
	EXPECT_TRUE(pubHandler.handleRead(&pubObj));
}
```

2. CSubscriber的头文件：

```c++
class CSubscriber : public CBaseConnection, public CHashElement
{
public:
	CSubscriber(int fd);
	
	virtual ~CSubscriber();
 
	bool initial();
 
	bool reset();
 
	//function return:
	//0: means complete read, all elements parsed OK
	//1: means it need recv more buf, not it's not complete
	//-1: means the packet is not valid.
	//-2: means connection wrong.
	int readPacket();
 
	//max send buf length
	static int m_iSendBufLen;
 
	//max recv buf length
	static int m_iRecvBufLen;
 
private:
	/*request format:
	 * GET /objectname?ts=xxx HTTP/1.x\r\n\r\n*/
	bool parsePacket();
};
```

e)、main函数的写法

与gtest相同，唯一的区别是初始化参数，如下：

```c++
#include <gmock/gmock.h>
 
int main(int argc, char** argv) {
	testing::InitGoogleMock(&argc, argv);
	//testing::InitGoogleTest(&argc, argv);
 
	// Runs all tests using Google Test.
	return RUN_ALL_TESTS();
}
```