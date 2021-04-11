#include <cstdlib>
#include <gmock/gmock.h>
#include <gtest/gtest.h>
#include <iostream>
#include <string>

#include "MockFoo.h"

using namespace seamless;
using namespace std;

using ::testing::Return;

int main(int argc, char** argv) {
        ::testing::InitGoogleMock(&argc, argv);

        string value = "Hello World!";
        MockFoo mockFoo;
        EXPECT_CALL(mockFoo, getArbitraryString()).Times(1).
                WillOnce(Return(value));
        string returnValue = mockFoo.getArbitraryString();
        cout << "Returned Value: " << returnValue << endl;
        EXPECT_TRUE(value == returnValue);
        return EXIT_SUCCESS;
}