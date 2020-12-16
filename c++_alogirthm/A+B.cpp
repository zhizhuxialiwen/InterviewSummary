//大数相加，例如A、B超过32位整数长度，  A+B?
//关键点：1、使用字符串进行表达；2、采用ISCII原理：'5' - '0' = 5; '2' + '7' - 2*'0' = 9; 3、相加需要进位

#include <iostream>
#include <string>
#include <algorithm>

using namespace std;

string sum(string str1, string str2) {
    if("" == str1) {
        return str2;
    }
    if("" == str2) {
        return str1;
    }

    string strRes = "";
    int str1Len = str1.length();
    int str2Len = str2.length();
    int minLen = str1Len < str2len ? str1Len : str2Len;

    //反转
    reverse(str1.begin(), str1.end());
    reverse(str2.begin(), str2.end());

    //进位
    int carry = 0;
    //当前值
    int currentNum = 0;
    int i = 0;
    for(; i < currentNum; i++) {
        currentNum = str1.at(i) + str2.at(i) - 2*'0' + carry;
        carry = currentNum / 10;
        currentNum %= 10;
        strRes.append(to_string(currentNum));
    }

    //求较大的那一部分
    string strMax;
    if( str1Len < str2Len) {
        strMax = str2;
    } else {
        strMax = str1;
    }

    for(; i < strMax.length(); i++) {
        currentNum = strMax.at(i) - '0' + carry;
        carry = currentNum / 10;
        currentNum %= 10;
        strRes.append(to_string(currentNum));
    }

    //处理最后一个进位
    if (carry > 0) {
        strRes.append(to_string(carry));
    }

    reverse(strRes.begin(), strRes.end());
    return strRes;
}

int main() {
    string str[20][2];
    int T;
    cin >> T;
    for(int i = 0; i < T) {
        cin >> str[i][0] >> str[i][1];
        string str1 = str[i][0];
        string str2 = str[i][1];
        string strRes = sum(str1, str2);
        cout << "Case "<< i + 1 <<":"<<endl;
        if(i == T-1) {
            cout << str1 << " + " << " = " << strRes;
        } else {
            cout << str1 << " + " << " = " << strRes << endl;
            cout << endl;
        }
    }

    return 0;
}