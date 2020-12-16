//大数相加，例如A、B超过32位整数长度，  A-B?
//关键：1、总是大值减小值；2、若A>B,则为正；如A<B，则进行交换位置，添加'-'；3、'9' - '4' - '0' = 5; 4、借位：10

#include <iostream>
#include <string>
#include <algorithm>


string sub1(string str1, int str1Len, string str2, int str2Len) {
    string strRes = "";
    int str1Num, str2Num;
    int over = 0;
    int flag = 0;
    for(int i = 0; i < str1Len; i++) {
        //从右往左取值
        str1Num = i > (str1Len - 1) ? 0 :str1[str1Len - 1 - i] - '0';
        str2Num = i > (str2Len - 1) ? 0 :str2[str2Len - 1 - i] - '0';

        int subNum = str1Num - str2Num - over;
        //借位
        if(subNum < 0) {
            subNum += 10;
            over = 1;
        } else {
            over = 0;
        }

        if(subNum == 0) {
            flag++;
        } else {
            while(flag > 0) {
                strRes = '0' + strRes;
                flag--;
            }
            strRes = to_string(subNum) + strRes;
        }
    }

    return strRes;
}

string  sub(string str1, string str2) {
    string strRes = "";
    if("" == str1) {
        strRes = "-" + str2;
        return strRes;
    }

    if("" == str2) {
        strRes = str1;
        return strRes;
    }

    int str1Len = str1.length();
    int str2Len = str2.length();

    if(str1Len > str2Len) {
        strRes = sub1(str1, str1Len, str2, str2Len);
    } else if (str1Len == str2Len) {
        if(str1 > str2) {
            strRes = sub1(str1, str1Len, str2, str2Len);
        } else if(str1 == str2) {
            strRes = '0';
        } else {
            strRes = '-' + sub1(str2, str2Len, str1, str1Len);
        }
    } else {
        strRes = '-' + sub1(str2, str2Len, str1, str1Len);
    }

    return strRes;
}

int main() {
    int T;
    string str[20][2];
    cin >> T;
    for(int i = 0; i < T) {
        cin >> str[i][0] >> str[i][1];
        string str1 = str[i][0];
        string str2 = str[i][1];
        string strRes = sub(str1, str2);
        cout << "Case "<< i + 1 <<":"<<endl;
        if(i == T-1) {
            cout << str1 << " - " << " = " << strRes;
        } else {
            cout << str1 << " - " << " = " << strRes << endl;
            cout << endl;
        }
    }
    return 0;
}