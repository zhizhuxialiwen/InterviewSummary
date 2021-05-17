# poll、select与epoll的多路复用

## 1、 IO 多路转接 (复用)

IO 多路转接也称为 IO 多路复用，它是一种网络通信的手段（机制），通过这种方式可以同时监测多个文件描述符并且这个过程是阻塞的，一旦检测到有文件描述符就绪（ 可以读数据或者可以写数据），程序的阻塞就会被解除，之后就可以基于这些（一个或多个）就绪的文件描述符进行通信了。通过这种方式在单线程 / 进程的场景下也可以在服务器端实现并发。常见的 IO 多路转接方式有：select、poll、epoll。

下面先对多线程 / 多进程并发和 IO 多路转接的并发处理流程进行对比（服务器端）：

1. 多线程 / 多进程并发

* 主线程 / 父进程：调用 accept() 监测客户端连接请求

1）如果没有新的客户端的连接请求，当前线程 / 进程会阻塞
2）如果有新的客户端连接请求解除阻塞，建立连接

* 子线程 / 子进程：和建立连接的客户端通信

1）调用 read() / recv() 接收客户端发送的通信数据，如果没有通信数据，当前线程 / 进程会阻塞，数据到达之后阻塞自动解除。
2）调用 write() / send() 给客户端发送数据，如果写缓冲区已满，当前线程 / 进程会阻塞，否则将待发送数据写入写缓冲区中。

2. IO 多路转接并发

使用 IO 多路转接函数委托内核检测服务器端所有的文件描述符（通信和监听两类），这个检测过程会导致进程 / 线程的阻塞，如果检测到已就绪的文件描述符阻塞解除，并将这些已就绪的文件描述符传出。
根据类型对传出的所有已就绪文件描述符进行判断，并做出不同的处理。

监听的文件描述符：和客户端建立连接
此时调用 accept() 是不会导致程序阻塞的，因为监听的文件描述符是已就绪的（有新请求）
通信的文件描述符：调用通信函数和已建立连接的客户端通信
调用 read() / recv() 不会阻塞程序，因为通信的文件描述符是就绪的，读缓冲区内已有数据。
调用 write() / send() 不会阻塞程序，因为通信的文件描述符是就绪的，写缓冲区不满，可以往里面写数据。

对这些文件描述符继续进行下一轮的检测（循环往复。。。）
与多进程和多线程技术相比，I/O 多路复用技术的最大优势是系统开销小，系统不必创建进程 / 线程，也不必维护这些进程 / 线程，从而大大减小了系统的开销。

## 2、 select

### 2.1 函数原型

使用 select 这种 IO 多路转接方式需要调用一个同名函数 select，这个函数是跨平台的，Linux、Mac、Windows 都是支持的。程序猿通过调用这个函数可以委托内核帮助我们检测若干个文件描述符的状态，其实就是检测这些文件描述符对应的读写缓冲区的状态：

* 读缓冲区：检测里边有没有数据，如果有数据该缓冲区对应的文件描述符就绪。
* 写缓冲区：检测写缓冲区是否可以写 (有没有容量)，如果有容量可以写，缓冲区对应的文件描述符就绪。
* 读写异常：检测读写缓冲区是否有异常，如果有该缓冲区对应的文件描述符就绪。

委托检测的文件描述符被遍历检测完毕之后，已就绪的这些满足条件的文件描述符会通过 select() 的参数分 3 个集合传出，程序猿得到这几个集合之后就可以分情况依次处理了。

下面来看一下这个函数的函数原型：

```c++
#include <sys/select.h>
struct timeval {
    time_t      tv_sec;         /* seconds */
    suseconds_t tv_usec;        /* microseconds */
};

int select(int nfds, fd_set *readfds, fd_set *writefds, fd_set *exceptfds, struct timeval * timeout);
```

1. 函数参数：

* nfds：委托内核检测的这三个集合中最大的文件描述符 + 1
内核需要线性遍历这些集合中的文件描述符，这个值是循环结束的条件。
在 Window 中这个参数是无效的，指定为 - 1 即可
* readfds：文件描述符的集合，内核只检测这个集合中文件描述符对应的读缓冲区。
传入传出参数，读集合一般情况下都是需要检测的，这样才知道通过哪个文件描述符接收数据
* writefds：文件描述符的集合，内核只检测这个集合中文件描述符对应的写缓冲区。
传入传出参数，如果不需要使用这个参数可以指定为 NULL
* exceptfds：文件描述符的集合，内核检测集合中文件描述符是否有异常状态
传入传出参数，如果不需要使用这个参数可以指定为 NULL
* timeout：超时时长，用来强制解除 select () 函数的阻塞的
NULL：函数检测不到就绪的文件描述符会一直阻塞。
等待固定时长（秒）：函数检测不到就绪的文件描述符，在指定时长之后强制解除阻塞，函数返回 0
不等待：函数不会阻塞，直接将该参数对应的结构体初始化为 0 即可。

2. 函数返回值：

大于 0：成功，返回集合中已就绪的文件描述符的总个数
等于 - 1：函数调用失败
等于 0：超时，没有检测到就绪的文件描述符

* 另外初始化 fd_set 类型的参数还需要使用相关的一些列操作函数，具体如下：

```c++
// 将文件描述符fd从set集合中删除 == 将fd对应的标志位设置为0        
void FD_CLR(int fd, fd_set *set);
// 判断文件描述符fd是否在set集合中 == 读一下fd对应的标志位到底是0还是1
int  FD_ISSET(int fd, fd_set *set);
// 将文件描述符fd添加到set集合中 == 将fd对应的标志位设置为1
void FD_SET(int fd, fd_set *set);
// 将set集合中, 所有文件文件描述符对应的标志位设置为0, 集合中没有添加任何文件描述符
void FD_ZERO(fd_set *set);
```

### 2.2 细节描述

在 select() 函数中第 2、3、4 个参数都是 fd_set 类型，它表示一个文件描述符的集合，类似于信号集 sigset_t，这个类型的数据有 128 个字节，也就是 1024 个标志位，和内核中文件描述符表中的文件描述符个数是一样的。

`sizeof(fd_set) = 128 字节 * 8 = 1024 bit      // int [32]`

这并不是巧合，而是故意为之。这块内存中的每一个 bit 和 文件描述符表中的每一个文件描述符是一一对应的关系，这样就可以使用最小的存储空间将要表达的意思描述出来了。

下图中的 fd_set 中存储了要委托内核检测读缓冲区的文件描述符集合。

* 如果集合中的标志位为 0 代表不检测这个文件描述符状态
* 如果集合中的标志位为 1 代表检测这个文件描述符状态

![select6](../../../images/select6.PNG)

内核在遍历这个读集合的过程中，如果被检测的文件描述符对应的读缓冲区中没有数据，内核将修改这个文件描述符在读集合 fd_set 中对应的标志位，改为 0，如果有数据那么这个标志位的值不变，还是 1。

![select7](../../../images/select7.PNG)

当 select() 函数解除阻塞之后，被内核修改过的读集合通过参数传出，此时集合中只要标志位的值为 1，那么它对应的文件描述符肯定是就绪的，我们就可以基于这个文件描述符和客户端建立新连接或者通信了。

### 2.3. 并发处理

#### 2.3.1 处理流程

如果在服务器基于 select 实现并发，其处理流程如下：

第一步：创建监听的套接字 lfd = socket ();
第二步：将监听的套接字和本地的 IP 和端口绑定 bind ()
第三步：给监听的套接字设置监听 listen ()
第四步：创建一个文件描述符集合 fd_set，用于存储需要检测读事件的所有的文件描述符
* 通过 FD_ZERO () 初始化
* 通过 FD_SET () 将监听的文件描述符放入检测的读集合中
第五步：循环调用 select ()，周期性的对所有的文件描述符进行检测
第六步：select () 解除阻塞返回，得到内核传出的满足条件的就绪的文件描述符集合
* 通过 FD_ISSET () 判断集合中的标志位是否为 1
1）如果这个文件描述符是监听的文件描述符，调用 accept () 和客户端建立连接。
1.1）将得到的新的通信的文件描述符，通过 FD_SET () 放入到检测集合中。
2）如果这个文件描述符是通信的文件描述符，调用通信函数和客户端通信。
2.1）如果客户端和服务器断开了连接，使用 FD_CLR () 将这个文件描述符从检测集合中删除
2.2）如果没有断开连接，正常通信即可
重复第 6 步

![select8](../../../images/select8.PNG)

#### 2.3.2 通信代码

1. 服务器端代码如下：

```c++
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <string.h>
#include <arpa/inet.h>

int main()
{
    // 1. 创建监听的fd
    int lfd = socket(AF_INET, SOCK_STREAM, 0);

    // 2. 绑定
    struct sockaddr_in addr;
    addr.sin_family = AF_INET;
    addr.sin_port = htons(9999); //htonl()--"Host to Network Long"
    addr.sin_addr.s_addr = INADDR_ANY; //转换过来就是0.0.0.0，泛指本机的意思，也就是表示本机的所有IP
    bind(lfd, (struct sockaddr*)&addr, sizeof(addr));

    // 3. 设置监听
    listen(lfd, 128);

    // 4. 将监听的fd的状态检测委托给内核检测
    int maxfd = lfd;
    // 初始化检测的读集合
    fd_set rdset;
    fd_set rdtemp;
    // 4.1清零
    FD_ZERO(&rdset);
    // 4.2 将监听的lfd设置到检测的读集合中
    FD_SET(lfd, &rdset);
    // 通过select委托内核检测读集合中的文件描述符状态, 检测read缓冲区有没有数据
    // 如果有数据, select解除阻塞返回
    // 应该让内核持续检测
    while(1)
    {
        // 默认阻塞
        // rdset 中是委托内核检测的所有的文件描述符
        rdtemp = rdset;
        //5 num文件描述符的总个数
        int num = select(maxfd+1, &rdtemp, NULL, NULL, NULL);
        if(num <= 0) {
            cout<<"error"<<endl;
            return -1;
        }
        // rdset中的数据被内核改写了, 只保留了发生变化的文件描述的标志位上的1, 没变化的改为0
        // 只要rdset中的fd对应的标志位为1 -> 缓冲区有数据了
        // 判断
        // 5.1 有没有新连接
        if(FD_ISSET(lfd, &rdtemp))
        {
            // 6.接受连接请求, 这个调用不阻塞
            struct sockaddr_in cliaddr;
            socklen_t cliLen = sizeof(cliaddr);
            int cfd = accept(lfd, (struct sockaddr*)&cliaddr, &cliLen);

            // 6.1 得到了有效的文件描述符
            // 通信的文件描述符添加到读集合
            // 在下一轮select检测的时候, 就能得到缓冲区的状态
            FD_SET(cfd, &rdset);
            // 重置最大的文件描述符
            maxfd = cfd > maxfd ? cfd : maxfd;
        }

        // 没有新连接, 通信
        for(int i=0; i<maxfd+1; ++i)
        {
			// 判断从监听的文件描述符之后到maxfd这个范围内的文件描述符是否读缓冲区有数据
            if(i != lfd && FD_ISSET(i, &rdtemp))
            {
                // 接收数据
                char buf[10] = {0};
                // 一次只能接收10个字节, 客户端一次发送100个字节
                // 一次是接收不完的, 文件描述符对应的读缓冲区中还有数据
                // 下一轮select检测的时候, 内核还会标记这个文件描述符缓冲区有数据 -> 再读一次
                // 	循环会一直持续, 知道缓冲区数据被读完位置
                int len = read(i, buf, sizeof(buf));
                if(len == 0)
                {
                    printf("客户端关闭了连接...\n");
                    // 将检测的文件描述符从读集合中删除
                    FD_CLR(i, &rdset);
                    close(i);
                }
                else if(len > 0)
                {
                    // 收到了数据
                    // 发送数据
                    write(i, buf, strlen(buf)+1);
                    printf("recv:%s\n",buf);
                }
                else
                {
                    // 异常
                    perror("read");
                }
            }
        }
    }

    return 0;
}
```

在上面的代码中，创建了两个 fd_set 变量，用于保存要检测的读集合：

// 初始化检测的读集合
fd_set rdset;
fd_set rdtemp;

rdset 用于保存要检测的原始数据，这个变量不能作为参数传递给 select 函数，因为在函数内部这个变量中的值会被内核修改，函数调用完毕返回之后，里边就不是原始数据了，大部分情况下是值为 1 的标志位变少了，不可能每一轮检测，所有的文件描述符都是就行的状态。因此需要通过 rdtemp 变量将原始数据传递给内核，select () 调用完毕之后再将内核数据传出，这两个变量的功能是不一样的。

* 客户端代码:

```c++
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <string.h>
#include <arpa/inet.h>

int main()
{
    // 1. 创建用于通信的套接字
    int fd = socket(AF_INET, SOCK_STREAM, 0);
    if(fd == -1)
    {
        perror("socket");
        exit(0);
    }

    // 2. 连接服务器
    struct sockaddr_in addr;
    addr.sin_family = AF_INET;     // ipv4
    addr.sin_port = htons(9999);   // 服务器监听的端口, 字节序应该是网络字节序
    inet_pton(AF_INET, "127.0.0.1", &addr.sin_addr.s_addr);
    int ret = connect(fd, (struct sockaddr*)&addr, sizeof(addr));
    if(ret == -1)
    {
        perror("connect");
        exit(0);
    }

    // 通信
    while(1)
    {
        // 读数据
        char recvBuf[1024];
        printf("The input: ");
        // 写数据
        // sprintf(recvBuf, "data: %d\n", i++);
        fgets(recvBuf, sizeof(recvBuf), stdin);
        write(fd, recvBuf, strlen(recvBuf)+1);
        // 如果客户端没有发送数据, 默认阻塞
        read(fd, recvBuf, sizeof(recvBuf));
        printf("recv buf: %s\n", recvBuf);
        sleep(1);
    }

    // 释放资源
    close(fd); 

    return 0;
}
```

编译：
```c++
select服务端:
wen@wen-virtual-machine:~/gopath/src/IO/select_poll_epoll2$ g++ -o select_server select_server.cpp 
wen@wen-virtual-machine:~/gopath/src/IO/select_poll_epoll2$ ./select_server 

select客户端：
wen@wen-virtual-machine:~/gopath/src/IO/select_poll_epoll2$ g++ -o select_client select_client.cpp 
wen@wen-virtual-machine:~/gopath/src/IO/select_poll_epoll2$ ./select_client 
```

客户端不需要使用 IO 多路转接进行处理，因为客户端和服务器的对应关系是 1：N，也就是说客户端是比较专一的，只能和一个连接成功的服务器通信。

虽然使用 select 这种 IO 多路转接技术可以降低系统开销，提高程序效率，但是它也有局限性(缺点)：

`int select(int nfds, fd_set *readfds, fd_set *writefds, fd_set *exceptfds, struct timeval * timeout);`

1. 待检测集合（第 2、3、4 个参数）需要频繁的在用户区和内核区之间进行数据的拷贝，效率低。
2. 内核对于 select 传递进来的待检测集合的检测方式是线性的。

* 如果集合内待检测的文件描述符很多，检测效率会比较低
* 如果集合内待检测的文件描述符相对较少，检测效率会比较高

3. 使用select能够检测的最大文件描述符个数有上限，默认是1024，这是在内核中被写死了的。

4. select应用于Linux、Mac、Windows 平台，即支持的平台。

## 2、 poll 函数

poll 的机制与 select 类似，与 select 在本质上没有多大差别，使用方法也类似，下面的是对于二者的对比：

内核对应文件描述符的检测也是以线性的方式进行轮询，根据描述符的状态进行处理：

1）poll 和 select 检测的文件描述符集合会在检测过程中频繁的进行用户区和内核区的拷贝，它的开销随着文件描述符数量的增加而线性增大，从而效率也会越来越低。
2）select检测的文件描述符个数上限是1024，poll没有最大文件描述符数量的限制
3）select可以跨平台使用，poll只能在Linux平台使用

poll 函数的函数原型如下：

```c++
#include <poll.h>
// 每个委托poll检测的fd都对应这样一个结构体
struct pollfd {
    int   fd;         /* 委托内核检测的文件描述符 */
    short events;     /* 委托内核检测文件描述符的什么事件 */
    short revents;    /* 文件描述符实际发生的事件 -> 传出 */
};

struct pollfd myfd[100];
int poll(struct pollfd *fds, nfds_t nfds, int timeout);
```

函数参数：

* fds: 这是一个 struct pollfd 类型的数组，里边存储了待检测的文件描述符的信息，这个数组中有三个成员：

1) fd：委托内核检测的文件描述符
2) events：委托内核检测的 fd 事件（输入、输出、错误），每一个事件有多个取值
3) revents：这是一个传出参数，数据由内核写入，存储内核检测之后的结果

* nfds: 这是第一个参数数组中最后一个有效元素的下标 + 1（也可以指定参数 1 数组的元素总个数）

* timeout: 指定 poll 函数的阻塞时长

1) -1：一直阻塞，直到检测的集合中有就绪的文件描述符（有事件产生）解除阻塞
2) 0：不阻塞，不管检测集合中有没有已就绪的文件描述符，函数马上返回
3) 大于 0：阻塞指定的毫秒（ms）数之后，解除阻塞

* 函数返回值：

失败： 返回 - 1
成功：返回一个大于 0 的整数，表示检测的集合中已就绪的文件描述符的总个数

### 1.2 测试代码

1. 服务器端

```c++
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <string.h>
#include <arpa/inet.h> //definitions for internet operations
#include <sys/select.h> //系统提供select函数来实现多路复用输入/输出模型
#include <poll.h>

int main()
{
    // 1.创建套接字
    int lfd = socket(AF_INET, SOCK_STREAM, 0);
    if(lfd == -1)
    {
        perror("socket");
        exit(0);
    }
    // 2. 绑定 ip, port
    struct sockaddr_in addr;
    addr.sin_port = htons(9999);
    addr.sin_family = AF_INET;
    addr.sin_addr.s_addr = INADDR_ANY;
    //inet_addr("127.0.0.1")
    int ret = bind(lfd, (struct sockaddr*)&addr, sizeof(addr));
    if(ret == -1)
    {
        perror("bind");
        exit(0);
    }
    // 3. 监听
    ret = listen(lfd, 100);
    if(ret == -1)
    {
        perror("listen");
        exit(0);
    }
    
    // 4. 等待连接 -> 循环
    // 检测 -> 读缓冲区, 委托内核去处理
    // 数据初始化, 创建自定义的文件描述符集
    struct pollfd fds[1024];
    // 初始化
    for(int i=0; i<1024; ++i)
    {
        fds[i].fd = -1;
        fds[i].events = POLLIN;
    }
    fds[0].fd = lfd;

    int maxfd = 0;
    while(1)
    {
        //4.1 委托内核检测
        ret = poll(fds, maxfd+1, -1);
        if(ret == -1)
        {
            perror("select");
            exit(0);
        }

        // 4.2 检测的度缓冲区有变化
        // 有新连接
        if(fds[0].revents & POLLIN)
        {
            // 接收连接请求
            struct sockaddr_in sockcli;
            socklen_t len = sizeof(sockcli);
            //5 这个accept是不会阻塞的
            int connfd = accept(lfd, (struct sockaddr*)&sockcli, &len);
            // 委托内核检测connfd的读缓冲区
            int i;
            for(i=0; i<1024; ++i)
            {
                if(fds[i].fd == -1)
                {
                    fds[i].fd = connfd;
                    break;
                }
            }
            maxfd = i > maxfd ? i : maxfd;
        }
        // 通信, 有客户端发送数据过来
        for(int i=1; i<=maxfd; ++i)
        {
            // 如果在集合中, 说明读缓冲区有数据
            if(fds[i].revents & POLLIN)
            {
                char buf[128];
                int ret = read(fds[i].fd, buf, sizeof(buf));
                if(ret == -1)
                {
                    perror("read");
                    exit(0);
                }
                else if(ret == 0)
                {
                    printf("对方已经关闭了连接...\n");
                    close(fds[i].fd);
                    fds[i].fd = -1;
                }
                else
                {
                    printf("客户端say: %s\n", buf);
                    write(fds[i].fd, buf, strlen(buf)+1);
                }
            }
        }
    }
    close(lfd);
    return 0;
}
```

从上面的测试代码可以得知，使用 poll 和 select 进行 IO 多路转接的处理思路是完全相同的，但是使用 poll 编写的代码看起来会更直观一些，select 使用的位图的方式来标记要委托内核检测的文件描述符（每个比特位对应一个唯一的文件描述符），并且对这个 fd_set 类型的位图变量进行读写还需要借助一系列的宏函数，操作比较麻烦。而 poll 直接将要检测的文件描述符的相关信息封装到了一个结构体 struct pollfd 中，我们可以直接读写这个结构体变量。

另外 poll 的第二个参数有两种赋值方式，但是都和第一个参数的数组有关系：

使用参数 1 数组的元素个数
使用参数 1 数组中存储的最后一个有效元素对应的下标值 + 1
内核会根据第二个参数传递的值对参数 1 数组中的文件描述符进行线性遍历，这一点和 select 也是类似的。

2. 客户端

```c++
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <string.h>
#include <arpa/inet.h>

int main()
{
    // 1. 创建用于通信的套接字
    /*
    第三个参数：0 是指定协议类型，系统自动根据情况指定,你不必显式制定这个参数，使用0则根据tcp或udp两个参数使用默认的协议。
    */
    int fd = socket(AF_INET, SOCK_STREAM, 0);
    if(fd == -1)
    {
        perror("socket");
        exit(0);
    }

    // 2. 连接服务器
    struct sockaddr_in addr;
    addr.sin_family = AF_INET;  // ipv4
    addr.sin_port = htons(9999);   // 服务器监听的端口, 字节序应该是网络字节序
    inet_pton(AF_INET, "127.0.0.1", &addr.sin_addr.s_addr);
    int ret = connect(fd, (struct sockaddr*)&addr, sizeof(addr));
    if(ret == -1)
    {
        perror("connect");
        exit(0);
    }

    // 通信
    while(1)
    {
        // 读数据
        char recvBuf[1024];
        printf("Input: ");
        // 写数据
        // sprintf(recvBuf, "data: %d\n", i++);
        fgets(recvBuf, sizeof(recvBuf), stdin);
        write(fd, recvBuf, strlen(recvBuf)+1);
        // 如果客户端没有发送数据, 默认阻塞
        read(fd, recvBuf, sizeof(recvBuf));
        printf("recv buf: %s\n", recvBuf);
        sleep(1);
    }
    // 释放资源
    close(fd); 
    return 0;
}
```

客户端不需要使用 IO 多路转接进行处理，因为客户端和服务器的对应关系是 1：N，也就是说客户端是比较专一的，只能和一个连接成功的服务器通信。

poll多路复用的编译：
![poll3](../../../images/poll3.PNG)
![poll4](../../../images/poll4.PNG)
![poll5](../../../images/poll5.PNG)
![poll6](../../../images/poll6.PNG)

## 3、epoll

### 3.1 概述

epoll 全称 eventpoll，是 linux 内核实现 IO 多路转接 / 复用（IO multiplexing）的一个实现。IO 多路转接的意思是在一个操作里同时监听多个输入输出源，在其中一个或多个输入输出源可用的时候返回，然后对其的进行读写操作。epoll 是 select 和 poll 的升级版，相较于这两个前辈，epoll 改进了工作方式，因此它更加高效。

* epoll与select和poll的区别：

1. 对于待检测集合select和poll是基于线性方式处理的，epoll是基于红黑树来管理待检测集合的。
2. select和poll每次都会线性扫描整个待检测集合，集合越大速度越慢，epoll使用的是回调机制，效率高，处理效率也不会随着检测集合的变大而下降。
3. select和poll工作过程中存在内核/用户空间数据的频繁拷贝问题(2此拷贝，用户和内核都拷贝一次)，在epoll中内核和用户区使用的是共享内存（基于mmap内存映射区实现，memory map，一次拷贝），省去了不必要的内存拷贝。
4. 程序猿需要对select和poll返回的集合进行判断才能知道哪些文件描述符是就绪的，通过epoll可以直接得到已就绪的文件描述符集合，无需再次检测
5. select的最大文件描述符为1024；poll没有限制说明，尽量不要太大；使用 epoll 没有最大文件描述符的限制，仅受系统中进程能打开的最大文件数目限制。

当多路复用的文件数量庞大、IO 流量频繁的时候，一般不太适合使用 select () 和 poll ()，这种情况下 select () 和 poll () 表现较差，推荐使用 epoll ()。

### 3.2 操作函数

在 epoll 中一共提供是三个 API 函数，分别处理不同的操作，函数原型如下：

```c++
#include <sys/epoll.h>
// 创建epoll实例，通过一棵红黑树管理待检测集合
int epoll_create(int size);
// 管理红黑树上的文件描述符(添加、修改、删除)
int epoll_ctl(int epfd, int op, int fd, struct epoll_event *event);
// 检测epoll树中是否有就绪的文件描述符
int epoll_wait(int epfd, struct epoll_event * events, int maxevents, int timeout);
```

select/poll 低效的原因之一是将 “添加 / 维护待检测任务” 和 “阻塞进程 / 线程” 两个步骤合二为一。每次调用 select 都需要这两步操作，然而大多数应用场景中，需要监视的 socket 个数相对固定，并不需要每次都修改。epoll 将这两个操作分开，先用 epoll_ctl() 维护等待队列，再调用 epoll_wait() 阻塞进程（解耦）。通过下图的对比显而易见，epoll 的效率得到了提升。

![select_poll_epoll1](../../../images/select_poll_epoll1.PNG)

1. epoll_create() 函数的作用是创建一个红黑树模型的实例，用于管理待检测的文件描述符的集合。

`int epoll_create(int size);`

函数参数 size：在 Linux 内核 2.6.8 版本以后，这个参数是被忽略的，只需要指定一个大于 0 的数值就可以了。
创建一个epoll的句柄，size用来告诉内核这个监听的数目一共有多大。这个参数不同于select()中的第一个参数，给出最大监听的fd+1的值。需要注意的是，当创建好epoll句柄后，它就是会占用一个fd值，在linux下如果查看/proc/进程id/fd/，是能够看到这个fd的，所以在使用完epoll后，必须调用close()关闭，否则可能导致fd被耗尽。

* 函数返回值：
失败：返回 - 1
成功：返回一个有效的文件描述符，通过这个文件描述符就可以访问创建的 epoll 实例了。

2. epoll_ctl() 函数的作用是管理红黑树实例上的节点，可以进行添加、删除、修改操作。

`int epoll_ctl(int epfd, int op, int fd, struct epoll_event *event);`

```c++
// 联合体, 多个变量共用同一块内存        
typedef union epoll_data {
 	void        *ptr;
	int          fd;	// 通常情况下使用这个成员, 和epoll_ctl的第三个参数相同即可
	uint32_t     u32;
	uint64_t     u64;
} epoll_data_t;

struct epoll_event {
	uint32_t     events;      /* Epoll events */
	epoll_data_t data;        /* User data variable */
};
int epoll_ctl(int epfd, int op, int fd, struct epoll_event *event);
```

* 函数参数：

1) epfd：epoll_create () 函数的返回值，通过这个参数找到 epoll 实例
2) op：这是一个枚举值，控制通过该函数执行什么操作
EPOLL_CTL_ADD：往 epoll 模型中添加新的节点
EPOLL_CTL_MOD：修改 epoll 模型中已经存在的节点
EPOLL_CTL_DEL：删除 epoll 模型中的指定的节点
3) fd：文件描述符，即要添加 / 修改 / 删除的文件描述符
4) event：epoll 事件，用来修饰第三个参数对应的文件描述符的，指定检测这个文件描述符的什么事件

4.1) events：委托 epoll 检测的事件
EPOLLIN：读事件，接收数据，检测读缓冲区，如果有数据该文件描述符就绪
EPOLLOUT：写事件，发送数据，检测写缓冲区，如果可写该文件描述符就绪
EPOLLERR：异常事件

4.2) data：用户数据变量，这是一个联合体类型，通常情况下使用里边的 fd 成员，用于存储待检测的文件描述符的值，在调用 epoll_wait() 函数的时候这个值会被传出。
函数返回值：
失败：返回 - 1
成功：返回 0

3. epoll_wait() 函数的作用是检测创建的 epoll 实例中有没有就绪的文件描述符。

`int epoll_wait(int epfd, struct epoll_event * events, int maxevents, int timeout);`

* 函数参数：

1) epfd：epoll_create () 函数的返回值，通过这个参数找到 epoll 实例。
2）events：传出参数，这是一个结构体数组的地址，里边存储了已就绪的文件描述符的信息
3）maxevents：修饰第二个参数，结构体数组的容量（元素个数）
4）timeout：如果检测的 epoll 实例中没有已就绪的文件描述符，该函数阻塞的时长，单位 ms 毫秒
0：函数不阻塞，不管 epoll 实例中有没有就绪的文件描述符，函数被调用后都直接返回
大于 0：如果 epoll 实例中没有已就绪的文件描述符，函数阻塞对应的毫秒数再返回
-1：函数一直阻塞，直到 epoll 实例中有已就绪的文件描述符之后才解除阻塞

* 函数返回值：
成功：
等于 0：函数是阻塞被强制解除了，没有检测到满足条件的文件描述符
大于 0：检测到的已就绪的文件描述符的总个数
失败：返回 - 1

###  3.3 epoll 的使用

#### 3.3.1 操作步骤

* 在服务器端使用 epoll 进行 IO 多路转接的操作步骤如下：

1. 创建监听的套接字

`int lfd = socket(AF_INET, SOCK_STREAM, 0);`

2. 设置端口复用（可选）

```c++
int opt = 1;
setsockopt(lfd, SOL_SOCKET, SO_REUSEADDR, &opt, sizeof(opt));
```

3. 使用本地的IP与端口和监听的套接字进行绑定

```c++
int ret = bind(lfd, (struct sockaddr*)&serv_addr, sizeof(serv_addr));
```

4. 给监听的套接字设置监听

`listen(lfd, 128);` //128表示client链接服务器最大值

5. 创建epoll实例对象

`int epfd = epoll_create(100);`

6. 将用于监听的套接字添加到epoll实例中

```c++
struct epoll_event ev;
ev.events = EPOLLIN;    // 检测lfd读读缓冲区是否有数据
ev.data.fd = lfd;
int ret = epoll_ctl(epfd, EPOLL_CTL_ADD, lfd, &ev);
```

7. 检测添加到epoll实例中的文件描述符是否已就绪，并将这些已就绪的文件描述符进行处理

`int num = epoll_wait(epfd, evs, size, -1);`

7.1) 如果是监听的文件描述符，和新客户端建立连接，将得到的文件描述符添加到epoll实例中

```c++
int cfd = accept(curfd, NULL, NULL);
ev.events = EPOLLIN;
ev.data.fd = cfd;
// 新得到的文件描述符添加到epoll模型中, 下一轮循环的时候就可以被检测了
epoll_ctl(epfd, EPOLL_CTL_ADD, cfd, &ev);
```

7.2) 如果是通信的文件描述符，和对应的客户端通信，如果连接已断开，将该文件描述符从epoll实例中删除

```c++
int len = recv(curfd, buf, sizeof(buf), 0);
if(len == 0)
{
    // 将这个文件描述符从epoll模型中删除
    epoll_ctl(epfd, EPOLL_CTL_DEL, curfd, NULL);
    close(curfd);
}
else if(len > 0)
{
    send(curfd, buf, len, 0);
}
```

重复第 7 步的操作

#### 3.3.2 示例代码

1. epoll_server

```c++
#include <stdio.h>
#include <ctype.h>
#include <unistd.h>
#include <stdlib.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <string.h>
#include <arpa/inet.h>
#include <sys/socket.h>
#include <sys/epoll.h>

// server
int main(int argc, const char* argv[])
{
    // 1.创建监听的套接字
    int lfd = socket(AF_INET, SOCK_STREAM, 0);
    if(lfd == -1)
    {
        perror("socket error");
        exit(1);
    }

   
    // 2.设置端口复用
    /*SO_REUSEADDR:打开或关闭地址复用功能。
    当option_value不等于0时，打开，否则，关闭。它实际所做的工作是置sock->sk->sk_reuse为1或0。
    */
    int opt = 1; 
    setsockopt(lfd, SOL_SOCKET, SO_REUSEADDR, &opt, sizeof(opt));

    // 3. 绑定地址与端口
    struct sockaddr_in serv_addr;
    memset(&serv_addr, 0, sizeof(serv_addr));
    serv_addr.sin_family = AF_INET;
    serv_addr.sin_port = htons(9999);
    serv_addr.sin_addr.s_addr = htonl(INADDR_ANY);  // 转换过来就是0.0.0.0，泛指本机的意思，也就是表示本机的所有IP
    
    int ret = bind(lfd, (struct sockaddr*)&serv_addr, sizeof(serv_addr));
    if(ret == -1)
    {
        perror("bind error");
        exit(1);
    }

    // 4. 监听
    ret = listen(lfd, 64);
    if(ret == -1)
    {
        perror("listen error");
        exit(1);
    }

    // 现在只有监听的文件描述符
    // 所有的文件描述符对应读写缓冲区状态都是委托内核进行检测的epoll
    // 5.创建一个epoll模型
    int epfd = epoll_create(100); //size用来告诉内核这个监听的数目一共有多大
    if(epfd == -1)
    {
        perror("epoll_create");
        exit(0);
    }

    //6. 往epoll实例中添加需要检测的节点, 现在只有监听的文件描述符
    struct epoll_event ev;
    ev.events = EPOLLIN;    // 检测lfd读读缓冲区是否有数据
    ev.data.fd = lfd;
    ret = epoll_ctl(epfd, EPOLL_CTL_ADD, lfd, &ev);
    if(ret == -1)
    {
        perror("epoll_ctl");
        exit(0);
    }

    struct epoll_event evs[1024];
    int size = sizeof(evs) / sizeof(struct epoll_event);
    // 7. 持续检测
    while(1)
    {
        // 7.1调用一次, 检测一次
        int num = epoll_wait(epfd, evs, size, -1);
        for(int i=0; i<num; ++i)
        {
            // 取出当前的文件描述符
            int curfd = evs[i].data.fd;
            // 判断这个文件描述符是不是用于监听的
            if(curfd == lfd)
            {
                // 建立新的连接
                int cfd = accept(curfd, NULL, NULL);
                // 新得到的文件描述符添加到epoll模型中, 下一轮循环的时候就可以被检测了
                ev.events = EPOLLIN;    // 读缓冲区是否有数据
                ev.data.fd = cfd;
                ret = epoll_ctl(epfd, EPOLL_CTL_ADD, cfd, &ev);
                if(ret == -1)
                {
                    perror("epoll_ctl-accept");
                    exit(0);
                }
            }
            else
            {
                // 处理通信的文件描述符
                // 接收数据
                char buf[1024];
                memset(buf, 0, sizeof(buf));
                int len = recv(curfd, buf, sizeof(buf), 0);
                if(len == 0)
                {
                    printf("客户端已经断开了连接\n");
                    // 将这个文件描述符从epoll模型中删除
                    epoll_ctl(epfd, EPOLL_CTL_DEL, curfd, NULL);
                    close(curfd);
                }
                else if(len > 0)
                {
                    printf("客户端say: %s\n", buf);
                    send(curfd, buf, len, 0);
                }
                else
                {
                    perror("recv");
                    exit(0);
                } 
            }
        }
    }

    return 0;
}
```

当在服务器端循环调用 epoll_wait() 的时候，就会得到一个就绪列表，并通过该函数的第二个参数传出：

```c++
struct epoll_event evs[1024];
int num = epoll_wait(epfd, evs, size, -1);
```

每当 epoll_wait() 函数返回一次，在 evs 中最多可以存储 size 个已就绪的文件描述符信息，但是在这个数组中实际存储的有效元素个数为 num 个，如果在这个 epoll 实例的红黑树中已就绪的文件描述符很多，并且 evs 数组无法将这些信息全部传出，那么这些信息会在下一次 epoll_wait() 函数返回的时候被传出。

通过 evs 数组被传递出的每一个有效元素里边都包含了已就绪的文件描述符的相关信息，这些信息并不是凭空得来的，这取决于我们在往 epoll 实例中添加节点的时候，往节点中初始化了哪些数据：

```c++
struct epoll_event ev;
// 节点初始化
ev.events = EPOLLIN;    
ev.data.fd = lfd;	// 使用了联合体中 fd 成员
// 添加待检测节点到epoll实例中
int ret = epoll_ctl(epfd, EPOLL_CTL_ADD, lfd, &ev);
```

在添加节点的时候，需要对这个 struct epoll_event 类型的节点进行初始化，当这个节点对应的文件描述符变为已就绪状态，这些被传入的初始化信息就会被原样传出，这个对应关系必须要搞清楚。

### 3.3.3 recv函数与send函数

1. recv函数：
功能：在已建立连接的套接字上接收数据。
格式：int recv(SOCKET s, char *buf, int len, int flags)。
参数：s-已建立连接的套接字；buf-存放接收到的数据的缓冲区指针；len-buf的长度；flags-调用方式：
（1）0：接收的是正常数据，无特殊行为。
（2）MSG_PEEK：系统缓冲区数据复制到提供的接收缓冲区，但是系统缓冲区内容并没有删除。
（3）MSG_OOB：表示处理带外数据。
返回值：接收成功时返回接收到的数据长度，连接结束时返回0，连接失败时返回SOCKET_ERROR。

2. send函数：
功能：在已建立连接的套接字上发送数据.
格式：int send(SOCKET s, const char *buf, int len, int flags)。
参数：参数：s-已建立连接的套接字；buf-存放将要发送的数据的缓冲区指针；len-发送缓冲区中的字符数；flags-控制数据传输方式：
（1）0：接收的是正常数据，无特殊行为。
（2）MSG_DONTROUTE：表示目标主机就在本地网络中，无需路由选择。
（3）MSG_OOB：表示处理带外数据。
返回值：发送成功时返回发送的数据长度，连接结束时返回0，连接失败时返回SOCKET_ERROR。

## 3.4 epoll 的工作模式

#### 3.4.1 水平模式

水平模式可以简称为 LT 模式，LT（level triggered）是缺省的工作方式，并且同时支持block和no-block socket。在这种做法中，内核通知使用者哪些文件描述符已经就绪，之后就可以对这些已就绪的文件描述符进行 IO 操作了。如果我们不作任何操作，内核还是会继续通知使用者。

水平模式的特点：

* 读事件：如果文件描述符对应的读缓冲区还有数据，读事件就会被触发，epoll_wait () 解除阻塞。

当读事件被触发，epoll_wait () 解除阻塞，之后就可以接收数据了
如果接收数据的 buf 很小，不能全部将缓冲区数据读出，那么读事件会继续被触发，直到数据被全部读出，如果接收数据的内存相对较大，读数据的效率也会相对较高（减少了读数据的次数）
因为读数据是被动的，必须要通过读事件才能指定有数据达到了，因此对于读事件的检测是必须的

* 写事件：如果文件描述符对应的写缓冲区可写，写事件就会被触发，epoll_wait () 解除阻塞。

当写事件被触发，epoll_wait () 解除阻塞，之后就可以将数据写入到写缓冲区了
写事件的触发发生在写数据之前而不是之后，被写入到写缓冲区中的数据是由内核自动发送出去的
如果写缓冲区没有被写满，写事件会一直被触发
因为写数据是主动的，并且写缓冲区一般情况下都是可写的（缓冲区不满），因此对于写事件的检测不是必须的
epoll 水平模式示例代码

#### 3.4.2 边沿模式

边沿模式可以简称为 ET 模式，ET（edge-triggered）是高速工作方式，只支持no-block socket。在这种模式下，当文件描述符从未就绪变为就绪时，内核会通过epoll通知使用者。然后它会假设使用者知道文件描述符已经就绪，并且不会再为那个文件描述符发送更多的就绪通知（only once）。如果我们对这个文件描述符做 IO 操作，从而导致它再次变成未就绪，当这个未就绪的文件描述符再次变成就绪状态，内核会再次进行通知，并且还是**只通知一次。** ET模式在很大程度上减少了epoll事件被重复触发的次数，因此效率要比LT模式高。

边沿模式的特点:

* 读事件：当读缓冲区有新的数据进入，读事件被触发一次，没有新数据不会触发该事件。

如果有新数据进入到读缓冲区，读事件被触发，epoll_wait () 解除阻塞
读事件被触发，可以通过调用 read ()/recv () 函数将缓冲区数据读出
如果数据没有被全部读走，并且没有新数据进入，读事件不会再次触发，只通知一次
如果数据被全部读走或者只读走一部分，此时有新数据进入，读事件被触发，并且只通知一次

* 写事件：当写缓冲区状态可写，写事件只会触发一次

如果写缓冲区被检测到可写，写事件被触发，epoll_wait () 解除阻塞
写事件被触发，就可以通过调用 write ()/send () 函数，将数据写入到写缓冲区中
写缓冲区从不满到被写满，期间写事件只会被触发一次
写缓冲区从满到不满，状态变为可写，写事件只会被触发一次
综上所述：epoll 的边沿模式下 epoll_wait () 检测到文件描述符有新事件才会通知，如果不是新的事件就不通知，通知的次数比水平模式少，效率比水平模式要高。

1. ET 模式的设置

**边沿模式不是默认的 epoll 模式，需要额外进行设置,即LT模式是默认模式。**
epoll 设置边沿模式是非常简单的，epoll 管理的红黑树示例中每个节点都是 struct epoll_event 类型，只需要将 EPOLLET 添加到结构体的 events 成员中即可：

```c++
struct epoll_event ev;
ev.events = EPOLLIN | EPOLLET;	// 设置边沿模式
```

示例代码如下：

```c++
int num = epoll_wait(epfd, evs, size, -1);
for(int i=0; i<num; ++i)
{
    // 取出当前的文件描述符
    int curfd = evs[i].data.fd;
    // 判断这个文件描述符是不是用于监听的
    if(curfd == lfd)
    {
        // 建立新的连接
        int cfd = accept(curfd, NULL, NULL);
        // 新得到的文件描述符添加到epoll模型中, 下一轮循环的时候就可以被检测了
        // 读缓冲区是否有数据, 并且将文件描述符设置为边沿模式
        struct epoll_event ev;
        ev.events = EPOLLIN | EPOLLET; //ET模式  
        ev.data.fd = cfd;
        ret = epoll_ctl(epfd, EPOLL_CTL_ADD, cfd, &ev);
        if(ret == -1)
        {
            perror("epoll_ctl-accept");
            exit(0);
        }
    }
}
```

2. 设置非阻塞

* 对于写事件的触发一般情况下是不需要进行检测的，因为写缓冲区大部分情况下都是有足够的空间可以进行数据的写入。

* 对于读事件的触发就必须要检测了，因为服务器也不知道客户端什么时候发送数据，如果使用 epoll 的边沿模式进行读事件的检测，有新数据达到只会通知一次，那么必须要保证得到通知后将数据全部从读缓冲区中读出。那么，应该如何读这些数据呢？

方式 1：准备一块特别大的内存，用于存储从读缓冲区中读出的数据，但是这种方式有很大的弊端：

内存的大小没有办法界定，太大浪费内存，太小又不够用
系统能够分配的最大堆内存也是有上限的，栈内存就更不必多言了
方式 2：循环接收数据

```c++
//弊端
int len = 0;
while((len = recv(curfd, buf, sizeof(buf), 0)) > 0)
{
    // 数据处理...
}
```

这样做也是有弊端的，因为套接字操作默认是阻塞的，当读缓冲区数据被读完之后，读操作就阻塞了也就是调用的 read()/recv() 函数被阻塞了，当前进程 / 线程被阻塞之后就无法处理其他操作了。

* 要解决阻塞问题，就需要将套接字默认的阻塞行为修改为非阻塞，需要使用 fcntl() 函数进行处理：

```c++
// 设置完成之后, 读写都变成了非阻塞模式
int flag = fcntl(cfd, F_GETFL);
flag |=  O_NONBLOCK;                                                        
fcntl(cfd, F_SETFL, flag);
```

* fcntl 函数的使用详解

通过上述分析就可以得出一个结论：epoll 在边沿模式下，必须要将套接字设置为非阻塞模式，但是，这样就会引发另外的一个 bug，在非阻塞模式下，循环地将读缓冲区数据读到本地内存中，当缓冲区数据被读完了，调用的 read()/recv() 函数还会继续从缓冲区中读数据，此时函数调用就失败了，返回 - 1，对应的全局变量 errno 值为 EAGAIN 或者 EWOULDBLOCK 如果打印错误信息会得到如下的信息：Resource temporarily unavailable

```c++
// 非阻塞模式下recv() / read()函数返回值 len == -1
int len = recv(curfd, buf, sizeof(buf), 0);
if(len == -1)
{
    if(errno == EAGAIN)
    {
        printf("数据读完了...\n");
    }
    else
    {
        perror("recv");
        exit(0);
    }
}
```

3. ET模式epoll server的示例代码

```c++
#include <stdio.h>
#include <ctype.h>
#include <unistd.h>
#include <stdlib.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <string.h>
#include <arpa/inet.h>
#include <sys/socket.h>
#include <sys/epoll.h>
#include <fcntl.h>
#include <errno.h>

// server
int main(int argc, const char* argv[])
{
    //2. 创建监听的套接字
    int lfd = socket(AF_INET, SOCK_STREAM, 0);
    if(lfd == -1)
    {
        perror("socket error");
        exit(1);
    }

    //2. 绑定
    struct sockaddr_in serv_addr;
    memset(&serv_addr, 0, sizeof(serv_addr));
    serv_addr.sin_family = AF_INET;
    serv_addr.sin_port = htons(9999);
    serv_addr.sin_addr.s_addr = htonl(INADDR_ANY);  // 本地多有的ＩＰ
    // 127.0.0.1
    // inet_pton(AF_INET, "127.0.0.1", &serv_addr.sin_addr.s_addr);
    
    // 3.设置端口复用
    int opt = 1;
    setsockopt(lfd, SOL_SOCKET, SO_REUSEADDR, &opt, sizeof(opt));

    // 4.绑定端口
    int ret = bind(lfd, (struct sockaddr*)&serv_addr, sizeof(serv_addr));
    if(ret == -1)
    {
        perror("bind error");
        exit(1);
    }

    // 5.监听
    ret = listen(lfd, 64);
    if(ret == -1)
    {
        perror("listen error");
        exit(1);
    }

    // 6.现在只有监听的文件描述符
    // 所有的文件描述符对应读写缓冲区状态都是委托内核进行检测的epoll
    // 创建一个epoll模型
    int epfd = epoll_create(100);
    if(epfd == -1)
    {
        perror("epoll_create");
        exit(0);
    }

    // 往epoll实例中添加需要检测的节点, 现在只有监听的文件描述符
    struct epoll_event ev;
    ev.events = EPOLLIN;    // 检测lfd读读缓冲区是否有数据
    ev.data.fd = lfd;
    ret = epoll_ctl(epfd, EPOLL_CTL_ADD, lfd, &ev);
    if(ret == -1)
    {
        perror("epoll_ctl");
        exit(0);
    }


    struct epoll_event evs[1024];
    int size = sizeof(evs) / sizeof(struct epoll_event);
    // 持续检测
    while(1)
    {
        // 调用一次, 检测一次
        int num = epoll_wait(epfd, evs, size, -1);
        printf("==== num: %d\n", num);
        //num表示文件描述符的个数
        for(int i=0; i<num; ++i)
        {
            // 取出当前的文件描述符
            int curfd = evs[i].data.fd;
            // 判断这个文件描述符是不是用于监听的
            if(curfd == lfd)
            {
                // 建立新的连接
                int cfd = accept(curfd, NULL, NULL);
                // 将文件描述符设置为非阻塞
                // 得到文件描述符的属性
                int flag = fcntl(cfd, F_GETFL);
                flag |= O_NONBLOCK;
                fcntl(cfd, F_SETFL, flag);
                // 新得到的文件描述符添加到epoll模型中, 下一轮循环的时候就可以被检测了
                // 通信的文件描述符检测读缓冲区数据的时候设置为边沿模式
                ev.events = EPOLLIN | EPOLLET;    // 读缓冲区是否有数据
                ev.data.fd = cfd;
                ret = epoll_ctl(epfd, EPOLL_CTL_ADD, cfd, &ev);
                if(ret == -1)
                {
                    perror("epoll_ctl-accept");
                    exit(0);
                }
            }
            else
            {
                // 处理通信的文件描述符
                // 接收数据
                char buf[5];
                memset(buf, 0, sizeof(buf));
                // 循环读数据
                while(1)
                {
                    int len = recv(curfd, buf, sizeof(buf), 0);
                    if(len == 0)
                    {
                        // 非阻塞模式下和阻塞模式是一样的 => 判断对方是否断开连接
                        printf("客户端断开了连接...\n");
                        // 将这个文件描述符从epoll模型中删除
                        epoll_ctl(epfd, EPOLL_CTL_DEL, curfd, NULL);
                        close(curfd);
                        break;
                    }
                    else if(len > 0)
                    {
                        // 通信
                        // 接收的数据打印到终端
                        write(STDOUT_FILENO, buf, len);
                        // 发送数据
                        send(curfd, buf, len, 0);
                    }
                    else
                    {
                        // len == -1
                        if(errno == EAGAIN)
                        {
                            printf("数据读完了...\n");
                            break;
                        }
                        else
                        {
                            perror("recv");
                            exit(0);
                        }
                    }
                }
            }
        }
    }

    return 0;
}
```

epoll client的代码与select client代码一样。

编译：

![epoll8](../../../images/epoll8.PNG)
![epoll9](../../../images/epoll9.PNG)
![epoll10](../../../images/epoll10.PNG)
![epoll11](../../../images/epoll11.PNG)