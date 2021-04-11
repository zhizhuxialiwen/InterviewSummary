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

/* g++ gtest_add.cpp -o out -lgtest -lpthread -std=c++11 */