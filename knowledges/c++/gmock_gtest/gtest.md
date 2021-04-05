# gtest1.10.0安装过程及简单使用

Linux环境：Ubuntu 16.04
Windows环境：Windows10 x64，Visual Stuido 2017

## 1、Linux环境下安装及使用

### 1.1 准备

下载googletest-release-1.10.0.tar.gz
https://github.com/google/googletest/releases/tag/release-1.10.0 

### 1.2 需要安装g++和cmake

```
sudo apt-get install g++
sudo apt-get install cmake
```

### 1.3 安装过程

1. 将googletest-release-1.10.0.tar.gz解压，并进入解压后的目录
tar -xzvf googletest-release-1.10.0.tar.gz
cd googletest-release-1.10.0

2. 编译
mkdir build
cd build
cmake ../CMakeLists.txt
cd ..
make

3. 拷贝库文件和包含文件
sudo cp lib/* /usr/lib
sudo cp -r googletest/include/gtest /usr/include

## 2、简单使用

编写测试文件test_add.cpp，内容如下：

```c++
#include <gtest/gtest.h>
 
int add(int a, int b)
{
    return a + b;
}
 
TEST(testCase, should_return_sum_correctly)
{
    EXPECT_EQ(10, add(4, 6));
}
 
int main(int argc,char **argv)
{
  testing::InitGoogleTest(&argc, argv);
  return RUN_ALL_TESTS();
}
```

执行编译命令： 

`g++ test_add.cpp -o test_add -lgtest -lpthread -std=c++11`

运行

`./test_add`

![gtest](../../../images/gtest.PNG)

## 2、GTest使用教程（二）-- 断言和宏测试


[windows-gtest安装与编译(一)](https://blog.csdn.net/W_Y2010/article/details/84674115?utm_medium=distribute.pc_relevant.none-task-blog-2%7Edefault%7EBlogCommendFromBaidu%7Edefault-5.control&dist_request_id=1328741.41223.16170260946925081&depth_1-utm_source=distribute.pc_relevant.none-task-blog-2%7Edefault%7EBlogCommendFromBaidu%7Edefault-5.control)

上一讲介绍了GTest的安装和如何在项目中使用GTest，这一讲主要介绍GTest的断言机制和宏测试

### 2.1、断言

一般的，要测试一个方法（函数）是否是正常执行的，可以提供一些输入数据，在调用这个方法（函数）后，得到输出数据，然后检查输出的数据是否与我们期望的结果是一致的，若一致，则说明这个方法的逻辑是正确的，否则，就有问题。 在对输出结果进行检查（check）时，GTest为我提供了一系列的断言（assertion）来进行代码测试，这些宏有点类似于函数调用。
当断言失败时GTest将会打印出assertion时的源文件和出错行的位置，以及附加的失败信息。这些输出的附加信息用户可以直接通过“<<”在这些断言宏后面。 如：

`EXPRCT_TRUE(bFlag) << " bFlag  is false";`

测试宏可以分为两大类：
```
ASSERT_*
EXPECT_*
```
这些成对的断言功能相同，但效果不同。

* ASSERT_*将会在失败时产生致命错误并中止当前调用它的函数执行（注意不是当前测试用例）。
* EXPECT_会生成非致命错误，不会中止当前函数，而是继续执行当前函数。

通常情况应该首选使用EXPECT_，因为ASSERT_*在报告完错误后不会进行清理工作，有可能导致内容泄露问题。

#### 2.1.1 基本断言（真值比较）

```
Fatal assertion	Nonfatal assertion	Verifies
ASSERT_TRUE(condition);	   EXPECT_TRUE(condition);	condition is true
ASSERT_FALSE(condition);	EXPECT_FALSE(condition);	condition is false
```

#### 2.1.2 二值比较

```
Fatal assertion	Nonfatal assertion	Verifies
ASSERT_EQ(val1,val2);	EXPECT_EQ(val1,val2);	val1 == val2
ASSERT_NE(val1,val2);	EXPECT_NE(val1,val2);	val1 != val2
ASSERT_LT(val1,val2);	EXPECT_LT(val1,val2);	val1 < val2
ASSERT_LE(val1,val2);	EXPECT_LE(val1,val2);	val1 <= val2
ASSERT_GT(val1,val2);	EXPECT_GT(val1,val2);	val1 > val2
ASSERT_GE(val1,val2);	EXPECT_GE(val1,val2);	val1 >= val2
```

一般来说二进制比较，都是对比其结构体所在内存的内容。C++大部分原生类型都是可以使用二进制对比的。但是对于自定义类型，我们就要定义一些操作符的行为，比如=、<等。

#### 2.1.3 字符串比较

```
Fatal assertion	Nonfatal assertion	Verifies
ASSERT_STREQ(str1,str2);	EXPECT_STREQ(str1,str2);	the two C strings have the same content
ASSERT_STRNE(str1,str2);	EXPECT_STRNE(str1,str2);	the two C strings have different content
ASSERT_STRCASEEQ(str1,str2);	EXPECT_STRCASEEQ(str1,str2);	the two C strings have the same content, ignoring case
ASSERT_STRCASENE(str1,str2);	EXPECT_STRCASENE(str1,str2);	the two C strings have different content, ignoring case
```

#### 2.1.4 浮点对比断言

在对比数据方面，我们往往会讨论到浮点数的对比。因为在一些情况下，浮点数的计算精度将影响对比结果，所以这块都会单独拿出来说。GTest对于浮点数的对比也是单独的

```
Fatal assertion	Nonfatal assertion	Verifies
ASSERT_FLOAT_EQ(val1, val2);	EXPECT_FLOAT_EQ(val1, val2);	the two float values are almost equal
ASSERT_DOUBLE_EQ(val1, val2);	EXPECT_DOUBLE_EQ(val1, val2);	the two double values are almost equal
```

almost euqal表示两个数只是近似相似，默认的是是指两者的差值在4ULP之内（Units in the Last Place）。我们还可以自己制定精度

```
Fatal assertion	Nonfatal assertion	Verifies
ASSERT_NEAR(val1, val2, abs_error);	EXPECT_NEAR(val1, val2, abs_error);	the difference between val1 and val2 doesn’t exceed the given absolute error
```

* 使用方法是

```
  ASSERT_NEAR(-1.0f, -1.1f, 0.2f);
  ASSERT_NEAR(2.0f, 3.0f, 1.0f);
```

#### 2.1.5 成功失败断言

该类断言用于直接标记是否成功或者失败。可以使用SUCCEED()宏标记成功，使用FAIL()宏标记致命错误（同ASSERT_)，ADD_FAILURE()宏标记非致命错误（同EXPECT_）。举个例子

```
if (Check) {
  SUCCEED();
}
else {
  FAIL();
}
```

我们直接在自己的判断下设置断言。这儿有个地方需要说一下，SUCCEED()宏会调用GTEST_MESSAGE_AT_宏，从而会影响TestResult的test_part_results结构体，这也是唯一的成功情况下影响该结构体的地方。

#### 2.1.6 异常断言

异常断言是在断言中接收一定类型的异常，并转换成断言形式。它有如下几种

```
Fatal assertion	Nonfatal assertion	Verifies
ASSERT_THROW(statement, exception_type);	EXPECT_THROW(statement, exception_type);	statement throws an exception of the given type
ASSERT_ANY_THROW(statement);	EXPECT_ANY_THROW(statement);	statement throws an exception of any type
ASSERT_NO_THROW(statement);	EXPECT_NO_THROW(statement);	statement doesn’t throw any exception
```

举一个例子

```
void ThrowException(int n) {
    switch (n) {
    case 0:
        throw 0;
    case 1:
        throw "const char*";
    case 2:
        throw 1.1f;
    case 3:
        return;
    }
}
 
TEST(ThrowException, Check) {
    EXPECT_THROW(ThrowException(0), int);
    EXPECT_THROW(ThrowException(1), const char*);
    ASSERT_ANY_THROW(ThrowException(2)); 
    ASSERT_NO_THROW(ThrowException(3));  
}
```

这组测试特例中，我们预期ThrowException在传入0时，会返回int型异常；传入1时，会返回const char*异常。传入2时，会返回异常，但是异常类型我们并不关心。传入3时，不返回任何异常。当然ThrowExeception的实现也是按以上预期设计的。

#### 2.1.7 参数名输出断言

在之前的介绍的断言中，如果在出错的情况下，我们会对局部测试相关信息进行输出，但是并不涉及其可能传入的参数。参数名输出断言，可以把参数名和对应的值给输出出来。

```
Fatal assertion	Nonfatal assertion	Verifies
ASSERT_PRED1(pred1, val1);	EXPECT_PRED1(pred1, val1);	pred1(val1) returns true
ASSERT_PRED2(pred2, val1, val2);	EXPECT_PRED2(pred2, val1, val2);	pred2(val1, val2) returns true
```

目前版本的GTest支持5个参数的版本ASSERT/EXPECT_PRED5宏。其使用方法是:

```
template <typename T1, typename T2>
bool GreaterThan(T1 x1, T2 x2) {
  return x1 > x2;
}
TEST(PredicateAssertionTest, AcceptsTemplateFunction) {
  int a = 5;
  int b = 6;
  ASSERT_PRED2((GreaterThan<int, int>), a, b);
}
```

其输出是
```
error: (GreaterThan<int, int>)(a, b) evaluates to false, where
a evaluates to 5
b evaluates to 6
```

#### 2.1.8 子过程中使用断言

经过之前的分析，我们可以想到，如果子过程中使用了断言，则结果输出只会指向子过程，而不会指向父过程中的某个调用，如果在父过程中多次调用这个子过程，那么就无法分析是哪一次调用失败。为了便于阅读我们可以使用SCOPED_TRACE宏去标记下位置

```
void Sub(int n) {
    ASSERT_EQ(1, n);
}
 
TEST(SubTest, Test1) {
    {
        SCOPED_TRACE("A");
        Sub(2);
    }
    Sub(3);
}
```

其结果输出时标记了下A这行位置，可见如果没有这个标记，是很难区分出是哪个Sub失败的。

```
..\test\gtest_unittest.cc(87): error:       Expected: 1
To be equal to: n
      Which is: 2
Google Test trace:
..\test\gtest_unittest.cc(92): A
..\test\gtest_unittest.cc(87): error:       Expected: 1
To be equal to: n
      Which is: 3
```

我们再注意下Sub的实现，其使用了ASSERT_EQ断言，该断言并不会影响Test1测试特例的运行，其原因我们在之前做过分析了。为了消除这种可能存在的误解，GTest推荐使用在子过程中使用

`ASSERT/EXPECT_NO_FATAL_FAILURE(statement);`

如果父过程一定要在子过程发生错误时退出怎么办？我们可以使用`::testing::Test::HasFatalFailure()`去判断当前线程中是否产生过错误。

```
TEST(SubTest, Test1) {
    {
        SCOPED_TRACE("A");
        Sub(2);
    }
    if (::testing::Test::HasFatalFailure())
        return;
    Sub(3);
}
```

### 2.2 宏测试

#### 2.2.1 TEST 宏

为了更好的介绍这些宏，以其自带的Sample1为例：

```
// Tests factorial of negative numbers.
TEST(FactorialTest, Negative) {
  EXPECT_EQ(1, Factorial(-5));
  EXPECT_EQ(1, Factorial(-1));
  EXPECT_GT(Factorial(-10), 0);
}
 
// Tests factorial of 0.
TEST(FactorialTest, Zero) {
  EXPECT_EQ(1, Factorial(0));
}
 
// Tests factorial of positive numbers.
TEST(FactorialTest, Positive) {
  EXPECT_EQ(1, Factorial(1));
  EXPECT_EQ(2, Factorial(2));
  EXPECT_EQ(6, Factorial(3));
  EXPECT_EQ(40320, Factorial(8));
  }
```

TEST宏是一个很重要的宏，它构成一个测试特例，它的原型是：

```
#if !GTEST_DONT_DEFINE_TEST
# define TEST(test_case_name, test_name) GTEST_TEST(test_case_name, test_name)
#endif
```

TEST宏的**第一个参数是test_case_name（测试用例名），第二个参数是test_name（测试特例名）。**

这里简单介绍一下测试用例名和测试特例名（也叫测试名）的区别和联系:

1. 测试用例（Test Case）是为某个特殊目标而编制的一组测试输入、执行条件以及预期结果，以便测试某个程序路径或核实是否满足某个特定需求，测试特例是测试用例下的一个（组）测试。

以上述Sample1代码为例，三段TEST宏构成的是一个测试用例——测试用例名是FactorialTest（阶乘方法检测，测试Factorial函数），该用例覆盖了三种测试特例——Negative、Zero和Positive——即检测输入参数是负数、零和正数这三种特例情况。

我们再看一组检测素数的测试用例

```
TEST(IsPrimeTest, Negative) {
 // This test belongs to the IsPrimeTest test case.
 EXPECT_FALSE(IsPrime(-1));
 EXPECT_FALSE(IsPrime(-2));
 EXPECT_FALSE(IsPrime(INT_MIN));
}

// Tests some trivial cases.
TEST(IsPrimeTest, Trivial) {
 EXPECT_FALSE(IsPrime(0));
 EXPECT_FALSE(IsPrime(1));
 EXPECT_TRUE(IsPrime(2));
 EXPECT_TRUE(IsPrime(3));
}

// Tests positive input.
TEST(IsPrimeTest, Positive) {
 EXPECT_FALSE(IsPrime(4));
 EXPECT_TRUE(IsPrime(5));
 EXPECT_FALSE(IsPrime(6));
 EXPECT_TRUE(IsPrime(23));
}
```

这组测试用例名是IsPrimeTest（测试IsPrime函数），三个测试特例是Negative（错误结果场景）、Trivial（有对有错的场景）和Positive（正确结果场景）。

对于测试用例名和测试特例名，不能有下划线（_）。因为GTest源码中需要使用下划线把它们连接成一个独立的类名

```
// Expands to the name of the class that implements the given test.
#define GTEST_TEST_CLASS_NAME_(test_case_name, test_name) \
  test_case_name##_##test_name##_Test
```

这样也就要求，我们不能有相同的“测试用例名和特例名”的组合——否则类名重合。

测试用例名和测试特例名的分开，使得我们编写的测试代码有着更加清晰的结构——即有相关性也有独立性。相关性是通过相同的测试用例名联系的，而独立性通过不同的测试特例名体现的。

#### 2.2.2 TEST_F 宏

场景：我们要测试向数据库插入（id,name,location）这样的三个数据，那要先构建一个基础数据（0,Fang,Beijing)。我们第一个测试特例可能需要关注于id这个字段，于是它要在基础数据上做出修改，将（1,Fang,Beijing)插入数据库。第二个测试特例可能需要关注于name字段，于是它要在基础数据上做出修改，将(0,Wang,Beijing)插入数据库。第三个测试特例可能需要关注于location字段，于是它要修改基础数据，将（0,Fang,Nanjing)插入数据库。如果使用GTEST宏来测试的话，那么每个测试特例前，我们需要把所有的数据填充好，再去操作。真实场景中一条记录往往不止三个数据，这样做会显得非常繁琐和不直观。

Google工程师早就考虑到这样的场景，可以将上述的场景提炼一下，其实我们只要在每个特例执行前，获取一份基础数据（原始数据），然后修改其中本次测试特例关心的一项就可以了。同时这份基础数据不可以在每个测试特例中被修改——即本次测试特例获取的基础数据不会受之前测试特例对基础数据修改而影响——获取的是一个恒定的数据。
这个时候我们就需要使用GTEST_F宏了，**GTEST_F叫作测试套件**。

Test Fixtures类继承于::testing::Test类。
在类内部使用public或者protected描述其成员，为了保证实际执行的测试子类可以使用其成员变量（这个我们后面会分析下）
在构造函数或者继承于::testing::Test类中的SetUp方法中，可以实现我们需要构造的数据。
在析构函数或者继承于::testing::Test类中的TearDown方法中，可以实现一些资源释放的代码（在3中申请的资源）。
其代码：

```
#define TEST_F(test_fixture, test_name)\
  GTEST_TEST_(test_fixture, test_name, test_fixture, \
              ::testing::internal::GetTypeId<test_fixture>())
```

**第一个参数要求是1中定义的类名；第二个参数是测试特例名。**

其中4这步并不是必须的，因为我们的数据可能不是申请来的数据，不需要释放。还有就是“构造函数/析构函数”和“SetUp/TearDown”的选择，对于什么时候选择哪对，没有统一的标准。一般来说就是构造/析构函数里忌讳做什么就不要在里面做，比如抛出异常等。
与TEST宏不同的是测试用例名必须是一个已定义的类名

1. 测试特例级别预处理

我们以一个例子来讲解

```c++
class TestFixtures : public ::testing::Test {
public:
    TestFixtures() {
        printf("\nTestFixtures\n");
    };
    ~TestFixtures() {
        printf("\n~TestFixtures\n");
    }
protected:
    void SetUp() {
        printf("\nSetUp\n");
        data = 0;
    };
    void TearDown() {
        printf("\nTearDown\n");
    }
protected:
    int data;
};
 
TEST_F(TestFixtures, First) {
    EXPECT_EQ(data, 0);
    data =  1;
    EXPECT_EQ(data, 1);
}
 
TEST_F(TestFixtures, Second) {
    EXPECT_EQ(data, 0);
    data =  1;
    EXPECT_EQ(data, 1);
}
```

First测试特例中，我们修改了data的数据（23行），第24行验证了修改的有效性和正确性。在second的测试特例中，一开始就检测了data数据（第28行），如果First特例中修改data（23行）影响了基础数据，则本次检测将失败。我们将First和Second测试特例的实现定义成一样的逻辑，可以避免编译器造成的执行顺序不确定从而影响测试结果。我们看下测试输出

```
[----------] 2 tests from TestFixtures
[ RUN      ] TestFixtures.First
TestFixtures
SetUp
TearDown
~TestFixtures
[       OK ] TestFixtures.First (9877 ms)
[ RUN      ] TestFixtures.Second
TestFixtures
SetUp
TearDown
~TestFixtures
[       OK ] TestFixtures.Second (21848 ms)
[----------] 2 tests from TestFixtures (37632 ms total)
```

可以见得，所有局部测试都是正确的，验证了Test Fixtures类中数据的恒定性。我们从输出应该可以看出来，每个测试特例都是要新建一个新的Test Fixtures对象，并在该测试特例结束时销毁它。这样可以保证数据的干净。

2. 测试用例级别预处理

这种预处理方式也是要使用Test Fixtures。不同的是，我们需要定义几个静态成员：

静态成员变量，用于指向数据。
静态方法SetUpTestCase()
静态方法TearDownTestCase()
举个例子，我们需要自定义测试用例开始和结束时的行为

测试开始时输出Start Test Case
测试结束时统计结果

```c++
class TestFixturesS : public ::testing::Test {
public:
    TestFixturesS() {
        printf("\nTestFixturesS\n");
    };
    ~TestFixturesS() {
        printf("\n~TestFixturesS\n");
    }
protected:
    void SetUp() {
    };
    void TearDown() {
    };
 
    static void SetUpTestCase() {
        UnitTest& unit_test = *UnitTest::GetInstance();
        const TestCase& test_case = *unit_test.current_test_case();
        printf("Start Test Case %s \n", test_case.name());
    };
 
    static void TearDownTestCase() {
        UnitTest& unit_test = *UnitTest::GetInstance();
        const TestCase& test_case = *unit_test.current_test_case();
        int failed_tests = 0;
        int suc_tests = 0;
        for (int j = 0; j < test_case.total_test_count(); ++j) {
            const TestInfo& test_info = *test_case.GetTestInfo(j);
            if (test_info.result()->Failed()) {
                failed_tests++;
            }
            else {
                suc_tests++;
            }
        }
        printf("End Test Case %s. Suc : %d, Failed: %d\n", test_case.name(), suc_tests, failed_tests);
    };
 
};
 
TEST_F(TestFixturesS, SUC) {
    EXPECT_EQ(1,1);
}
 
TEST_F(TestFixturesS, FAI) {
    EXPECT_EQ(1,2);
}
```

测试用例中，我们分别测试一个成功结果和一个错误的结果。然后输出如下

```
[----------] 2 tests from TestFixturesS
Start Test Case TestFixturesS
[ RUN      ] TestFixturesS.SUC
TestFixturesS
~TestFixturesS
[       OK ] TestFixturesS.SUC (2 ms)
[ RUN      ] TestFixturesS.FAI
TestFixturesS
..\test\gtest_unittest.cc(126): error:       Expected: 1
To be equal to: 2
~TestFixturesS
[  FAILED  ] TestFixturesS.FAI (5 ms)
End Test Case TestFixturesS. Suc : 1, Failed: 1
[----------] 2 tests from TestFixturesS (12 ms total)
```

从输出上看，SetUpTestCase在测试用例一开始时就被执行了，TearDownTestCase在测试用例结束前被执行了。

3. 全局级别预处理

顾名思义，它是在测试用例之上的一层初始化逻辑。如果我们要使用该特性，则要声明一个继承于::testing::Environment的类，并实现其SetUp/TearDown方法。这两个方法的关系和之前介绍Test Fixtures类是一样的。

我们看一个例子，我们例子中的预处理

测试开始时输出Start Test
测试结束时统计结果

```c++
namespace testing {
namespace internal {
class EnvironmentTest : public ::testing::Environment {
public:
    EnvironmentTest() {
        printf("\nEnvironmentTest\n");
    };
    ~EnvironmentTest() {
        printf("\n~EnvironmentTest\n");
    }
public:
    void SetUp() {
        printf("\n~Start Test\n");
    };
    void TearDown() {
        UnitTest& unit_test = *UnitTest::GetInstance();
        for (int i = 0; i < unit_test.total_test_case_count(); ++i) {
            int failed_tests = 0;
            int suc_tests = 0;
            const TestCase& test_case = *unit_test.GetTestCase(i);
            for (int j = 0; j < test_case.total_test_count(); ++j) {
                const TestInfo& test_info = *test_case.GetTestInfo(j);
                // Counts failed tests that were not meant to fail (those without
                // 'Fails' in the name).
                if (test_info.result()->Failed()) {
                    failed_tests++;
                }
                else {
                    suc_tests++;
                }
            }
            printf("End Test Case %s. Suc : %d, Failed: %d\n", test_case.name(), suc_tests, failed_tests);
        }
    };
};
}
}
 
GTEST_API_ int main(int argc, char **argv) {
  printf("Running main() from gtest_main.cc\n");
  ::testing::AddGlobalTestEnvironment(new testing::internal::EnvironmentTest);
  testing::InitGoogleTest(&argc, argv);
  return RUN_ALL_TESTS();
}
```

我们可以关注下::testing::AddGlobalTestEnvironment(new testing::internal::EnvironmentTest);这句，我们要在调用RUN_ALL_TESTS之前，使用该函数将全局初始化对象加入到框架中。通过这种方式，可以猜测出，我们可以加入多个对象到框架中。

4. 总结

GTEST对于每个测试间共享数据提供了三种不同级别的方式：

测试特例间共享数据：SetUp() 和TearDown()，
测试用例间共享数据：SetUpTestCase()和TearDownTestCase()
全局共享数据：SetUpEnvironment()和TearDownEnvironment()

#### 2.2.3 GTEST_P宏

在我们设计测试用例时，我们需要考虑很多场景。每个场景都可能要细致地考虑到到各个参数的选择。比如我们要测试一个函数，它有两个参数，两个参数输入是整型，第一个参数我们需要测试[-5,-1,0,1,9]五种场景，第二个参数我们需要测试[0,10,100,1000]四种场景，那么我们组合在一起就是20种场景，我们需要这样写代码：

```
EXPECT_TRUE(Func(-5,0));
EXPECT_TRUE(Func(-1,0));
...
EXPECT_TRUE(Func(-5,10));
EXPECT_TRUE(Func(-1,10));
...
EXPECT_TRUE(Func(9,1000));
```

这种写法明显是不合理的。GTest框架当然也会考虑到这点，它设计了一套自动生成上述检测的机制，让我们用很少的代码就可以解决这个问题，那就是TEST_P宏。

```
# define TEST_P(test_case_name, test_name) \
  class GTEST_TEST_CLASS_NAME_(test_case_name, test_name) \
      : public test_case_name { \
   public: \
    GTEST_TEST_CLASS_NAME_(test_case_name, test_name)() {} \
    virtual void TestBody(); \
   private:
    GTEST_DISALLOW_COPY_AND_ASSIGN_(\
        GTEST_TEST_CLASS_NAME_(test_case_name, test_name));   \
......
```

与TEST_F宏一样，**第一个参数是测试用例名，第二个参数是测试特例名。**

我们先从应用的角度讲解其使用，首先我们设计一个需要被测试的类。

```c++
class Bis 
{
public:
    bool Even(int n) 
    {
        if (n % 2 == 0) 
        {
            return true;
        }
        else 
        {
            return false;
        }
    };
 
    bool Suc(bool bSuc) 
    {
        return bSuc;
    }
};
```

该类暴露了两个返回bool类型的方法：Even用于判断是否是偶数；Suc只是返回传入的参数。

1. bool型入参

Suc函数的入参类型是bool，于是我们可以新建一个测试用例类，让它继承于template < typename T> class WithParamInterface模板类，并把模板指定为bool，由于GTest要求提供测试的类要继承于::testing::Test，所以这个UT还要继承::testing::Test。

```
class CheckBisSucUT :
    public ::testing::Test,
    public ::testing::WithParamInterface<bool>
{
};
```

我们再设置一个测试特例，在特例中使用GetParam()方法获取框架指定的参数
```
TEST_P(CheckBisSucUT , Test) 
{
	CheckBisSuc objBis;
    EXPECT_TRUE(objBis.Suc(GetParam()));
}
```

最后，我们使用INSTANTIATE_TEST_CASE_P宏向框架注册“定制化测试”

`INSTANTIATE_TEST_CASE_P(TestBisBool, CheckBisSucUT, Bool());`

该宏的第一个参数是测试前缀，第二个参数是测试类名，第三个参数是参数生成规则。如此我们就相当于执行了
```
    EXPECT_TRUE(objBis.Suc(true));
    EXPECT_TRUE(objBis.Suc(false));
```

2. 可选择入参

我们再看下针对Even函数的测试。我们要定义一个继承于template < typename T> class WithParamInterface模板类的类CheckBisEvenUT，用于指定Even的入参类型为int

```
class CheckBisEvenUT:
    public ::testing::Test,
    public ::testing::WithParamInterface<int>
{
};
```

然后我们建立一个针对该类的测试特例

```
TEST_P(CheckBisEvenUT, Test) 
{
	CheckBisSuc objBis;
    EXPECT_TRUE(objBis.Even(GetParam()));
}
```

最后我们可以使用Range、Values或者ValuesIn的方式指定Even的参数值

```
INSTANTIATE_TEST_CASE_P(TestBisValuesRange, CheckBisEvenUT, Range(0, 9, 2));
 
INSTANTIATE_TEST_CASE_P(TestBisValues, CheckBisEvenUT, Values(11, 12, 13, 14));
 
int values[] = {0, 1};
INSTANTIATE_TEST_CASE_P(TestBisValuesIn, CheckBisEvenUT, ValuesIn(values));
 
int moreValues[] = {0,1,2,3,4,5,6,7,8,9,10};
vector<int> IntVecValues(moreValues, moreValues + sizeof(moreValues));
INSTANTIATE_TEST_CASE_P(TestBisValuesInVector, CheckBisEvenUT, ValuesIn(IntVecValues));
```

Range的第一个参数是起始参数值，第二个值是结束参数值，第三个参数是递增值。于是Range这组测试测试的是0、2、4、6、8这些入参。如果第三个参数没有， 则默认是递增1。

Values中罗列的是将被选择作为参数的值。

ValuesIn的参数是个容器或者容器的起始迭代器和结束迭代器。

3. 参数组合

参数组合要求编译器支持tr/tuple，所以一些不支持tr库的编译器将无法使用该功能。

什么是参数组合？顾名思义，就是将不同参数集组合在一起衍生出多维的数据。比如（true,false）和（1,2）可以组合成（true，1）、（true，2）、（false，1）和（false，2）等四种参数组合，然后我们使用这四组数据进行测试。

我们看个例子，首先我们要定义一个待测类。需要注意的是，它继承了模板类TestWithParam，且模板参数是组合的类型::testing::tuple<bool, int>。这个类并没有继承Bis，而是让Bis成为其成员变量，在checkData函数中检测Bis的各个函数

```c++
class CombineTest : 
    public TestWithParam< ::testing::tuple<bool, int> > {
protected:
	bool checkData() {
		bool suc = ::testing::get<0>(GetParam());
        int n = ::testing::get<1>(GetParam());
		return bis.Suc(suc) &&  bis.Even(n);
	}
private:
	Bis bis;
};
        然后我们定义一个（true，false）和（1，2，3，4）组合测试

TEST_P(CombineTest, Test) {
	EXPECT_TRUE(checkData());
}
 
INSTANTIATE_TEST_CASE_P(TestBisValuesCombine, CombineTest, Combine(Bool(), Values(0, 1, 2, 3, 4)));
```

如何我们便可以衍生出8组测试。我们看下部分测试结果输出

```
[----------] 8 tests from TestBisValuesCombine/CombineTest
......
[ RUN      ] TestBisValuesCombine/CombineTest.Test/6
[       OK ] TestBisValuesCombine/CombineTest.Test/6 (0 ms)
[ RUN      ] TestBisValuesCombine/CombineTest.Test/7
../samples/sample11_unittest.cc:175: Failure
Value of: checkData()
  Actual: false
Expected: true
[  FAILED  ] TestBisValuesCombine/CombineTest.Test/7, where GetParam() = (true, 3) (1 ms)
[----------] 8 tests from TestBisValuesCombine/CombineTest (2 ms total)
```

上例中TestBisValuesCombine/CombineTest是最终的测试用例名，Test/6和Test/7是其下两个测试特例名。

我们最后把参数生成函数罗列下

```
Range(begin, end[, step])	Yields values {begin, begin+step, begin+step+step, …}. The values do not include end. step defaults to 1.
Values(v1, v2, …, vN)	Yields values {v1, v2, …, vN}.
ValuesIn(container) and ValuesIn(begin, end)	Yields values from a C-style array, an STL-style container, or an iterator range [begin, end). container, begin, and end can be expressions whose values are determined at run time.
Bool()	Yields sequence {false, true}.
Combine(g1, g2, …, gN)	Yields all combinations (the Cartesian product for the math savvy) of the values generated by the N generators. This is only available if your system provides the <tr1/tuple> header. If you are sure your system does, and Google Test disagrees, you can override it by defining GTEST_HAS_TR1_TUPLE=1. See comments in include/gtest/internal/gtest-port.h for more information.
```

4. 总结

TEST_P 大致与TEST_F相同，都是第一个参数是一个已定义类名，第二个参数是测试特例名。不同的是，TEST_P测试用例类需要继承::testing::WithParamInterface< T> ,并且可以用GetPara方法取得参数。
————————————————
版权声明：本文为CSDN博主「会会会飞的鱼」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/W_Y2010/article/details/92405343
