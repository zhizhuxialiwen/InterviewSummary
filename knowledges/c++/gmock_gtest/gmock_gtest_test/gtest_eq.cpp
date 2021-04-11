#include <gtest/gtest.h> 

int fun1() {
  return 10;
}

class test : public ::testing::Test{
public:
  int fun2() {
    return 10;
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


/*g++ gtest_eq.cpp -o out -lgtest -lpthread  -std=c++11*/