# inet_pton()和inet_ntop()函数详解

## 1、把ip地址转化为用于网络传输的二进制数值

`int inet_aton(const char *cp, struct in_addr *inp);`

inet_aton() 转换网络主机地址ip(如192.168.1.10)为二进制数值，并存储在struct in_addr结构中，即第二个参数*inp,函数返回非0表示cp主机有地有效，返回0表示主机地址无效。（这个转换完后不能用于网络传输，还需要调用htons或htonl函数才能将主机字节顺序转化为网络字节顺序）

`in_addr_t inet_addr(const char *cp);`

inet_addr函数转换网络主机地址（如192.168.1.10)为网络字节序二进制值，如果参数char *cp无效，函数返回-1(INADDR_NONE),这个函数在处理地址为255.255.255.255时也返回－1,255.255.255.255是一个有效的地址，不过inet_addr无法处理;


## 2、将网络传输的二进制数值转化为成点分十进制的ip地址

`char *inet_ntoa(struct in_addr in);`

inet_ntoa 函数转换网络字节排序的地址为标准的ASCII以点分开的地址,该函数返回指向点分开的字符串地址（如192.168.1.10)的指针，该字符串的空间为静态分配的，这意味着在第二次调用该函数时，上一次调用将会被重写（复盖），所以如果需要保存该串最后复制出来自己管理！

 
我们如何输出一个点分十进制的IP呢？我们来看看下面的程序：

```c++
#include <stdio.h>   
#include <sys/socket.h>   
#include <netinet/in.h>   
#include <arpa/inet.h>   
#include <string.h>   
int main()   
{   
	struct in_addr addr1,addr2;   
	ulong l1,l2;   
	l1= inet_addr("192.168.0.74");   
	l2 = inet_addr("211.100.21.179");   
	memcpy(&addr1, &l1, 4);   
	memcpy(&addr2, &l2, 4);   
	printf("%s : %s\n", inet_ntoa(addr1), inet_ntoa(addr2)); //注意这一句的运行结果   
	printf("%s\n", inet_ntoa(addr1));   
	printf("%s\n", inet_ntoa(addr2));  
	return 0;   
}   
```

实际运行结果如下：　

```
192.168.0.74 : 192.168.0.74          //从这里可以看出,printf里的inet_ntoa只运行了一次。　　

192.168.0.74　　

211.100.21.179　
```

inet_ntoa返回一个char *,而这个char *的空间是在inet_ntoa里面静态分配的，所以inet_ntoa后面的调用会覆盖上一次的调用。第一句printf的结果只能说明在printf里面的可变参数的求值是从右到左的，仅此而已。


## 3、新型网路地址转化函数inet_pton和inet_ntop

这两个函数是随IPv6出现的函数，对于IPv4地址和IPv6地址都适用，函数中p和n分别代表表达（presentation)和数值（numeric)。地址的表达格式通常是ASCII字符串，数值格式则是存放到套接字地址结构的二进制值。

```c++
#include <arpe/inet.h>
int inet_pton(int family, const char *strptr, void *addrptr);     //将点分十进制的ip地址转化为用于网络传输的数值格式
        返回值：若成功则为1，若输入不是有效的表达式则为0，若出错则为-1
 
const char * inet_ntop(int family, const void *addrptr, char *strptr, size_t len);     //将数值格式转化为点分十进制的ip地址格式
        返回值：若成功则为指向结构的指针，若出错则为NULL
```

（1）这两个函数的family参数既可以是AF_INET（ipv4）也可以是AF_INET6（ipv6）。如果，以不被支持的地址族作为family参数，这两个函数都返回一个错误，并将errno置为EAFNOSUPPORT.
（2）第一个函数尝试转换由strptr指针所指向的字符串，并通过addrptr指针存放二进制结果，若成功则返回值为1，否则如果所指定的family而言输入字符串不是有效的表达式格式，那么返回值为0.

（3）inet_ntop进行相反的转换，从数值格式（addrptr）转换到表达式（strptr)。inet_ntop函数的strptr参数不可以是一个空指针。调用者必须为目标存储单元分配内存并指定其大小，调用成功时，这个指针就是该函数的返回值。len参数是目标存储单元的大小，以免该函数溢出其调用者的缓冲区。如果len太小，不足以容纳表达式结果，那么返回一个空指针，并置为errno为ENOSPC。


4.示例

inet_pton(AF_INET, ip, &foo.sin_addr);   //  代替 foo.sin_addr.addr=inet_addr(ip);
 
 
char str[INET_ADDRSTRLEN];
