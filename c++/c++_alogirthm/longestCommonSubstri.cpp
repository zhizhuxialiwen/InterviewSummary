
/*
题目描述
给定两个字符串str1和str2,输出两个字符串的最长公共子串，如果最长公共子串为空，输出-1。
示例1
输入
复制
"1AB2345CD","12345EF"
返回值
复制
"2345"
*/

#include <iostream>
#include <string>
#include <vector>

#using namespace std;

class Solution {
public:
    /**
     * longest common substring
     * @param str1 string字符串 the string
     * @param str2 string字符串 the string
     * @return string字符串
     */
    string LCS(string str1, string str2){
        // write code here
        int len;
        string res;
        int m = str1.size();
        int n = str2.size();
        int x, y;//记住最大长度时的坐标。
        vector< vector<int> > tag(m + 1, vector<int>(n + 1, 0));
        for (int i = 1; i <= m; ++i) {
            for (int j = 1; j <= n; ++j) {
                if (str1[i - 1] == str2[j - 1]) {
                    tag[i][j] = tag[i - 1][j - 1] + 1;
                    if (len < tag[i][j]) {
                        len = tag[i][j];
                        x = i;
                        y = j;
                    }
                }
                else
                    tag[i][j] = 0;
            }
        }
        if (len == 0) res = "-1";
        else {
            for (int i = len; i >= 1; --i) {
                res += str1[x - i];
            }
        }
        return res;
        }
};