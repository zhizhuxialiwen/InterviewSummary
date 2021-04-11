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

/*g++ gmock_eq.cpp -o out -lgtest -lgmock -lpthread -std=c++11*/