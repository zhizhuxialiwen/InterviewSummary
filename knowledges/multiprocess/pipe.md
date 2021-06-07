# 管道

## 1、 管道

管道的是进程间通信（IPC - InterProcess Communication）的一种方式，管道的本质其实就是内核中的一块内存 (或者叫内核缓冲区)，这块缓冲区中的数据存储在一个**环形队列**中，因为管道在内核里边，因此我们不能直接对其进行任何操作。

![pipeline1](../../images/pipeline1.PNG)

因为管道数据是通过队列来维护的，我们先来分析一个管道中数据的特点：

1. 管道对应的内核缓冲区大小是固定的，默认为 4k（也就是队列最大能存储 4k 数据）
2. 管道分为两部分：读端和写端（队列的两端），数据从写端进入管道，从读端流出管道。
3. 管道中的数据只能读一次，做一次读操作之后数据也就没有了（读数据相当于出队列）。
4. 管道是单工的：数据只能单向流动，数据从写端流向读端。
5. 对管道的操作（读、写）默认是阻塞的
    * 读管道：管道中没有数据，读操作被阻塞，当管道中有数据之后阻塞才能解除
    * 写管道：管道被写满了，写数据的操作被阻塞，当管道变为不满的状态，写阻塞解除

管道在内核中，不能直接对其进行操作，我们通过什么方式去读写管道呢？其实管道操作就是文件 IO 操作，内核中管道的两端分别对应两个文件描述符，通过写端的文件描述符把数据写入到管道中，通过读端的文件描述符将数据从管道中读出来。读写管道的函数就是 Linux 中的文件 IO 函数 [read/write 详解](./readWritePipe.md)

// 读管道
`ssize_t read(int fd, void *buf, size_t count);`
// 写管道的函数
`ssize_t write(int fd, const void *buf, size_t count);`

最后分析一下为什么可以使用管道进行进程间通信，先看一下下面的图片：

![pipeline2](../../images/pipeline2.PNG)

在上图中假设父进通过一系列操作可以通过文件描述符表中的文件描述符 fd3 写管道，通过 fd4 读管道，然后再通过 fork() 创建出子进程，那么在父进程中被分配的文件描述符 fd3， fd4也就被拷贝到子进程中，子进程通过 fd3可以将数据写入到内核的管道中，通过fd4将数据从管道中读出来。

也就是说管道是独立于任何进程的，并且充当了两个进程用于数据通信的载体，只要两个进程能够得到同一个管道的入口和出口（读端和写端的文件描述符），那么他们之间就可以通过管道进行数据的交互。

## 2、 匿名（无名）管道--血缘的进程

### 2.1 创建匿名管道

匿名管道是管道的一种，既然是匿名也就是说这个管道没有名字，但其本质是不变的，就是位于内核中的一块内存，匿名管道拥有上面介绍的管道的所有特性，额外的我们需要知道，**匿名管道只能实现有血缘关系的进程间通信，什么叫有血缘的进程关系呢，比如：父子进程，兄弟进程，爷孙进程，叔侄进程。** 最后说一下创建匿名管道的函数，函数原型如下：

`#include <unistd.h>`
// 创建一个匿名的管道, 得到两个可用的文件描述符
`int pipe(int pipefd[2]);`
参数：传出参数，需要传递一个整形数组的地址，数组大小为 2，也就是说最终会传出两个元素
* pipefd[0]: 对应管道读端的文件描述符，通过它可以将数据从管道中读出
* pipefd[1]: 对应管道写端的文件描述符，通过它可以将数据写入到管道中
* 返回值：成功返回 0，失败返回 -1

### 2.2 进程间通信

使用匿名管道只能够实现有血缘关系的进程间通信，要求写一段程序完成下边的功能：

需求描述:
   在父进程中创建一个子进程, 父子进程分别执行不同的操作:
     - 子进程: 执行一个shell命令 "ps aux", 将命令的结果传递给父进程
     - 父进程: 将子进程命令的结果输出到终端
需求分析:

子进程中执行 shell 命令相当于启动一个磁盘程序，因此需要使用 `execl ()/execlp () 函数`
`execlp(“ps”, “ps”, “aux”, NULL)`

子进程中执行完 shell 命令直接就可以在终端输出结果，如果将这些信息传递给父进程呢？
数据传递需要使用管道，子进程需要将数据写入到管道中。
将默认输出到终端的数据写入到管道就需要进行输出的重定向，需要使用 dup2() 做这件事情。
`dup2(fd[1], STDOUT_FILENO);`

父进程需要读管道，将从管道中读出的数据打印到终端。
父进程最后需要释放子进程资源，防止出现僵尸进程。

在使用管道进行进程间通信的注意事项：必须要保证数据在管道中的单向流动。这句话怎么理解呢，通过下面的图来分析一下：

第一步：在父进程中创建了匿名管道，得到了两个分配的文件描述符，fd3 操作管道的读端，fd4 操作管道的写端。

![pipeline3](../../images/pipeline3.PNG)

第二步：父进程创建子进程，父进程的文件描述符被拷贝，在子进程的文件描述符表中也得到了两个被分配的可以使用的文件描述符，通过 fd3 读管道，通过 fd4 写管道。通过下图可以看到管道中数据的流动不是单向的，有以下这么几种情况：

* 父进程通过 fd4 将数据写入管道，然后父进程再通过 fd3 将数据从管道中读出
* 父进程通过 fd4 将数据写入管道，然后子进程再通过 fd3 将数据从管道中读出
* 子进程通过 fd4 将数据写入管道，然后子进程再通过 fd3 将数据从管道中读出
* 子进程通过 fd4 将数据写入管道，然后父进程再通过 fd3 将数据从管道中读出

前边说到过，管道行为默认是阻塞的，假设子进程通过写端将数据写入管道，父进程的读端将数据读出，这样子进程的读端就读不到数据，导致子进程阻塞在读管道的操作上，这样就会给程序的执行造成一些不必要的影响。如果我们本来也没有打算让进程读或者写管道，那么就可以将进程操作的读端或者写端关闭。

![pipeline4](../../images/pipeline4.PNG)

第三步：为了避免两个进程都读管道，但是可能其中某个进程由于读不到数据而阻塞的情况，我们可以关闭进程中用不到的那一端的文件描述符，这样数据就只能单向的从一端流向另外一端了，如下图，我们关闭了父进程的写端，关闭了子进程的读端：

![pipeline5](../../images/pipeline5.PNG)

根据上面的分析，最终可以写出下面的代码：

* 案例1

// 管道的数据是单向流动的:
// 操作管道的是两个进程, 进程A读管道, 需要关闭管道的写端, 进程B写管道, 需要关闭管道的读端
// 如果不做上述的操作, 会对程序的结果造成一些影响, 对管道的操作无法结束

```c++
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <string.h>
#include <fcntl.h>
#include <sys/wait.h>

int main()
{
    // 1. 创建匿名管道, 得到两个文件描述符
    int fd[2];
    int ret = pipe(fd);
    if(ret == -1)
    {
        perror("pipe");
        exit(0);
    }
    // 2. 创建子进程 -> 能够操作管道的文件描述符被复制到子进程中
    pid_t pid = fork();
    if(pid == 0)
    {
        // 关闭读端
        close(fd[0]);
        // 3. 在子进程中执行 execlp("ps", "ps", "aux", NULL);
        // 在子进程中完成输出的重定向, 原来输出到终端现在要写管道
        // 进程打印数据默认输出到终端, 终端对应的文件描述符: stdout_fileno
        // 标准输出 重定向到 管道的写端
        dup2(fd[1], STDOUT_FILENO);
        execlp("ps", "ps", "aux", NULL); //启动进程
        perror("execlp");
    }

    // 4. 父进程读管道
    else if(pid > 0)
    {
        // 关闭管道的写端
        close(fd[1]);
        // 5. 父进程打印读到的数据信息
        char buf[4096]; //4k = 4 * 1024
        // 读管道
        // 如果管道中没有数据, read会阻塞
        // 有数据之后, read解除阻塞, 直接读数据
        // 需要循环读数据, 管道是有容量的, 写满之后就不写了
        // 数据被读走之后, 继续写管道, 那么就需要再继续读数据
        while(1)
        {
            memset(buf, 0, sizeof(buf));
            int len = read(fd[0], buf, sizeof(buf));
            if(len == 0)
            {
                // 管道的写端关闭了, 如果管道中没有数据, 管道读端不会阻塞
                // 没数据直接返回0, 如果有数据, 将数据读出, 数据读完之后返回0
                break;
            }
            printf("%s, len = %d\n", buf, len);
        }
        close(fd[0]);

        // 回收子进程资源
        wait(NULL);
    }
    return 0;
}
```

编译：

![pipeline6](../../images/pipeline6.PNG)

* 案例2

```c++
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <string.h>
#include <fcntl.h>
#include <sys/wait.h>

int main()
{
    // 1. 创建匿名管道, 得到两个文件描述符
    int fd[2];
    int ret = pipe(fd);
    if(ret == -1)
    {
        perror("pipe");
        exit(0);
    }
    // 2. 创建子进程 -> 能够操作管道的文件描述符被复制到子进程中
    pid_t pid = fork();
    if(pid == 0)
    {
        // 关闭读端
        close(fd[0]);
        // 3. 在子进程中执行 execlp("ps", "ps", "aux", NULL);
        // 在子进程中完成输出的重定向, 原来输出到终端现在要写管道
        // 进程打印数据默认输出到终端, 终端对应的文件描述符: stdout_fileno
        // 标准输出 重定向到 管道的写端
        //dup2(fd[1], STDOUT_FILENO);
        char buf[4096]; //4k = 4 * 1024
        memset(buf, 0, sizeof(buf));
        const char * str1 = "hello world!";
        memcpy(buf, str1, strlen(str1));
        int len = write(fd[1], buf, sizeof(buf));
        if(len == 0)
        {
            // 管道的写端关闭了, 如果管道中没有数据, 管道读端不会阻塞
            // 没数据直接返回0, 如果有数据, 将数据读出, 数据读完之后返回0
            return -1;
        }
        //execlp("ps", "ps", "aux", NULL); //启动进程
        //perror("execlp");
        close(fd[1]);
    }

    // 4. 父进程读管道
    else if(pid > 0)
    {
        // 关闭管道的写端
        close(fd[1]);
        // 5. 父进程打印读到的数据信息
        char buf[4096]; //4k = 4 * 1024
        // 读管道
        // 如果管道中没有数据, read会阻塞
        // 有数据之后, read解除阻塞, 直接读数据
        // 需要循环读数据, 管道是有容量的, 写满之后就不写了
        // 数据被读走之后, 继续写管道, 那么就需要再继续读数据
        while(1)
        {
            memset(buf, 0, sizeof(buf));
            int len = read(fd[0], buf, sizeof(buf));
            if(len == 0)
            {
                // 管道的写端关闭了, 如果管道中没有数据, 管道读端不会阻塞
                // 没数据直接返回0, 如果有数据, 将数据读出, 数据读完之后返回0
                break;
            }
            printf("%s, len = %d\n", buf, len);
        }
        close(fd[0]);

        // 回收子进程资源
        wait(NULL);
    }
    return 0;
}
```

编译：

![pipeline9](../../images/pipeline9.PNG)

## 3、 有名管道
### 3.1 创建有名管道

有名管道拥有管道的所有特性，之所以称之为有名是因为管道在磁盘上有实体文件，文件类型为 p ，有名管道文件大小永远为 0，因为有名管道也是将数据存储到内存的缓冲区中，打开这个磁盘上的管道文件就可以得到操作有名管道的文件描述符，通过文件描述符读写管道存储在内核中的数据。

**有名管道也可以称为 fifo (first in first out)，使用有名管道既可以进行有血缘关系的进程间通信，也可以进行没有血缘关系的进程间通信。** 创建有名管道的方式有两种，一种是通过命令，一种是通过函数。

1. 通过命令

`$ mkfifo 有名管道的名字`

2. 通过函数

```c++
#include <sys/types.h>
#include <sys/stat.h>
// int open(const char *pathname, int flags, mode_t mode);
int mkfifo(const char *pathname, mode_t mode);
```

参数:
* pathname: 要创建的有名管道的名字
* mode: 文件的操作权限，和 open () 的第三个参数一个作用，最终权限: (mode & ~umask)
* 返回值：创建成功返回 0，失败返回 -1

### 3.2 进程间通信

不管是有血缘关系还是没有血缘关系，使用有名管道实现进程间通信的方式是相同的，就是在两个进程中分别以读、写的方式打开磁盘上的管道文件，得到用于读管道、写管道的文件描述符，就可以调用对应的 read ()、write () 函数进行读写操作了。

小贴士：

有名管道操作需要通过 open () 操作得到读写管道的文件描述符，如果只是读端打开了或者只是写端打开了，进程会阻塞在这里不会向下执行，直到在另一个进程中将管道的对端打开，当前进程的阻塞也就解除了。所以当发现进程阻塞在了 open () 函数上不要感到惊讶。

```c
#include <stdlib.h>
#include <unistd.h>
#include <fcntl.h>
#include <sys/stat.h>
```
* 　stdlib 头文件即standard library标准库头文件,stdlib 头文件里包含了C、C++语言的最常用的系统函数,该文件包含了的C语言标准库函数的定义。
　　stdlib.h里面定义了五种类型、一些宏和通用工具函数。 类型例如size_t、wchar_t、div_t、ldiv_t和lldiv_t；宏例如EXIT_FAILURE、EXIT_SUCCESS、RAND_MAX和MB_CUR_MAX等等；常用的函数如malloc()、calloc()、realloc()、free()、system()、atoi()、atol()、rand()、srand()、exit()等等。具体的内容你自己可以打开编译器的include目录里面的stdlib.h头文件看看。
* unistd.h是unix std的意思，是POSIX标准定义的unix类系统定义符号常量的头文件，包含了许多UNIX系统服务的函数原型，例如read函数、write函数和getpid函数。 
* fcntl.h定义了很多宏和open,fcntl函数原型
* sys/stat.h: 文件状态，是unix/linux系统定义文件状态所在的伪标准头文件。

* 写管道的进程

1. 创建有名管道文件 
    mkfifo()
2. 打开有名管道文件, 打开方式是 O_WRONLY
    int wfd = open("xx", O_WRONLY);
3. 调用write函数写文件 ==> 数据被写入管道中
    write(wfd, data, strlen(data));
4. 写完之后关闭文件描述符
    close(wfd);

```c++
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <string.h>
#include <fcntl.h>
#include <sys/stat.h>

int main()
{
    // 1. 创建有名管道文件
    int ret = mkfifo("./testfifo", 0664);
    if(ret == -1)
    {
        perror("mkfifo");
        exit(0);
    }
    printf("管道文件创建成功...\n");

    // 2. 打开管道文件
    // 因为要写管道, 所有打开方式, 应该指定为 O_WRONLY
    // 如果先打开写端, 读端还没有打开, open函数会阻塞, 当读端也打开之后, open解除阻塞
    int wfd = open("./testfifo", O_WRONLY);
    if(wfd == -1)
    {
        perror("open");
        exit(0);
    }
    printf("以只写的方式打开文件成功...\n");

    // 3. 循环写管道
    int i = 0;
    while(i<100)
    {
        char buf[1024];
        sprintf(buf, "hello, fifo, 我在写管道...%d\n", i);
        write(wfd, buf, strlen(buf));
        i++;
        sleep(1);
    }
    close(wfd);

    return 0;
}
```

编译：

![pipeline7](../../images/pipeline7.PNG)

* 读管道的进程

1. 这两个进程需要操作相同的管道文件
2. 打开有名管道文件, 打开方式是 o_rdonly
    int rfd = open("xx", O_RDONLY);
3. 调用read函数读文件 ==> 读管道中的数据
    char buf[4096];
    read(rfd, buf, sizeof(buf));
4. 读完之后关闭文件描述符
    close(rfd);

```c++
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <string.h>
#include <fcntl.h>
#include <sys/stat.h>

int main()
{
    // 1. 打开管道文件
    // 因为要read管道, so打开方式, 应该指定为 O_RDONLY
    // 如果只打开了读端, 写端还没有打开, open阻塞, 当写端被打开, 阻塞就解除了
    int rfd = open("./testfifo", O_RDONLY);
    if(rfd == -1)
    {
        perror("open"); // 把一个描述性错误消息输出到标准错误 stder
        exit(0);
    }
    printf("以只读的方式打开文件成功...\n");

    // 2. 循环读管道
    while(1)
    {
        char buf[1024];
        memset(buf, 0, sizeof(buf));
        // 读是阻塞的, 如果管道中没有数据, read自动阻塞
        // 有数据解除阻塞, 继续读数据
        int len = read(rfd, buf, sizeof(buf));
        printf("读出的数据: %s\n", buf);
        if(len == 0)
        {
            // 写端关闭了, read解除阻塞返回0
            printf("管道的写端已经关闭, 拜拜...\n");
            break;
        }

    }
    close(rfd);

    return 0;
}
```

编译：

![pipeline8](../../images/pipeline8.PNG)

## 4、 管道的读写行为

关于管道不管是有名的还是匿名的，在进行读写的时候，它们表现出的行为是一致的，下面是对其读写行为的总结:

* 读管道，需要根据写端的状态进行分析：
    * 写端没有关闭 (操作管道写端的文件描述符没有被关闭)
        * 如果管道中没有数据 ==> 读阻塞 , 如果管道中被写入了数据，阻塞解除
        * 如果管道中有数据 ==> 不阻塞，管道中的数据被读完了，再继续读管道还会阻塞
    * 写端已经关闭了 (没有可用的文件描述符可以写管道了)
        * 管道中没有数据 ==> 读端解除阻塞，read 函数返回 0
        * 管道中有数据 ==> read 先将数据读出，数据读完之后返回 0, 不会阻塞了

* 写管道，需要根据读端的状态进行分析：
    * 读端没有关闭
        * 如果管道有存储的空间，一直写数据
        * 如果管道写满了，写操作就阻塞，当读端将管道数据读走了，解除阻塞继续写
    * 读端关闭了，管道破裂 (异常), 进程直接退出

管道的两端默认是阻塞的，如何将管道设置为非阻塞呢？管道的读写两端的非阻塞操作是相同的，下面的代码中将匿名的读端设置为了非阻塞：

```c++
// 通过fcntl 修改就可以, 一般情况下不建议修改
// 管道操作对应两个文件描述符, 分别是管道的读端 和 写端

// 1. 获取读端的文件描述符的flag属性
int flag = fcntl(fd[0], F_GETFL);
// 2. 添加非阻塞属性到 flag中
flag |= O_NONBLOCK;
// 3. 将新的flag属性设置给读端的文件描述符
fcntl(fd[0], F_SETFL, flag);
// 4. 非阻塞读管道
char buf[4096];
read(fd[0], buf, sizeof(buf));
```


