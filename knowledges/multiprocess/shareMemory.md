# 共享内存

共享内存不同于内存映射区，它不属于任何进程，并且不受进程生命周期的影响。通过调用 Linux 提供的系统函数就可得到这块共享内存。使用之前需要让进程和共享内存进行关联，得到共享内存的起始地址之后就可以直接进行读写操作了，进程也可以和这块共享内存解除关联，解除关联之后就不能操作这块共享内存了。在所有进程间通信的方式中共享内存的效率是最高的。

共享内存操作默认不阻塞，如果多个进程同时读写共享内存，可能出现数据混乱，共享内存需要借助其他机制来保证进程间的数据同步，比如：信号量，共享内存内部没有提供这种机制。

## 1. 创建 / 打开共享内存

### 1.1 shmget

在使用共享内存之前必须要先做一些准备工作，如果共享内存不存在就需要先创建出来，如果已经存在了就需要先打开这块共享内存。不管是创建还是打开共享内存使用的函数是同一个，函数原型如下:

```c++
#include <sys/ipc.h>
#include <sys/shm.h>
int shmget(key_t key, size_t size, int shmflg);
```

参数:

* key: 类型 key_t 是个整形数，通过这个key可以创建或者打开一块共享内存，该参数的值一定要大于0
* size: 创建共享内存的时候，指定共享内存的大小，如果是打开一块存在的共享内存，size 是没有意义的
* shmflg：创建共享内存的时候指定的属性

1. IPC_CREAT: 创建新的共享内存，如果创建共享内存，需要指定对共享内存的操作权限，比如：IPC_CREAT | 0664
2. IPC_EXCL: 检测共享内存是否已经存在了，必须和 IPC_CREAT 一起使用

* 返回值：共享内存创建或者打开成功返回标识共享内存的唯一的 ID，失败返回 - 1

函数使用举例:

场景 1：创建一块大小为 4k 的共享内存

`shmget(100, 4096, IPC_CREAT|0664);`

场景 2：创建一块大小为 4k 的共享内存，并且检测是否存在

// 	如果共享内存已经存在, 共享内存创建失败, 返回-1, 可以perror() 打印错误信息

`shmget(100, 4096, IPC_CREAT|0664|IPC_EXCL);`

场景 3：打开一块已经存在的共享内存

// 函数参数虽然指定了大小和IPC_CREAT, 但是都不起作用, 因为共享内存已经存在, 只能打开, 参数4096也没有意义
`shmget(100, 4096, IPC_CREAT|0664);`
`shmget(100, 0, 0);`

场景 4：打开一块共享内存，如果不存在就创建

`shmget(100, 4096, IPC_CREAT|0664);`

### 1.2 ftok

shmget () 函数的第一个参数是一个大于 0 的正整数，如果不想自己指定可以通过 ftok () 函数直接生成这个 key 值。该函数的函数原型如下：

```c++
// ftok函数原型
#include <sys/types.h>
#include <sys/ipc.h>

// 将两个参数作为种子, 生成一个 key_t 类型的数值
key_t ftok(const char *pathname, int proj_id);
```

参数:

* pathname: 当前操作系统中一个存在的路径
* proj_id: 这个参数只用到了 int 中的一个字节，传参的时候要将其作为 char 进行操作，取值范围: 1-255

* 返回值：函数调用成功返回一个可用于创建、打开共享内存的 key 值，调用失败返回 - 1

使用举例：

```c++
// 根据路径生成一个key_t
key_t key = ftok("/home/robin", 'a');
// 创建或打开共享内存
shmget(key, 4096, IPC_CREATE|0664);
```

## 2、 关联和解除关联

### 2.1 shmat关联 -- share memory at

创建 / 打开共享内存之后还必须和共享内存进行关联，这样才能得到共享内存的起始地址，通过得到的内存地址进行数据的读写操作，关联函数的原型如下：

```c++
void *shmat(int shmid, const void *shmaddr, int shmflg);
```

参数:

* shmid: 要操作的共享内存的 ID, 是 shmget () 函数的返回值
* shmaddr: 共享内存的起始地址，用户不知道，需要让内核指定，写 NULL
* shmflg: 和共享内存关联的对共享内存的操作权限

1. SHM_RDONLY: 读权限，只能读共享内存中的数据
2. 0: 读写权限，可以读写共享内存数据

* 返回值：关联成功，返回值共享内存的起始地址，关联失败返回 (void *) -1

### 2.2 shmdt共享内存删除关联--share memory delete

当进程不需要再操作共享内存，可以让进程和共享内存解除关联，另外如果没有执行该操作，进程退出之后，结束的进程和共享内存的关联也就自动解除了。

`int shmdt(const void *shmaddr);`

参数：shmat () 函数的返回值，共享内存的起始地址

返回值：关联解除成功返回 0，失败返回 - 1

## 3、 删除共享内存

### 3.1 shmctl共享内存控制-- share memory control

shmctl () 函数是一个多功能函数，可以设置、获取共享内存的状态也可以将共享内存标记为删除状态。当共享内存被标记为删除状态之后，并不会马上被删除，直到所有的进程全部和共享内存解除关联，共享内存才会被删除。因为通过 shmctl () 函数只是能够标记删除共享内存，所以在程序中多次调用该操作是没有关系的。

```c++
// 共享内存控制函数
int shmctl(int shmid, int cmd, struct shmid_ds *buf);

// 参数 struct shmid_ds 结构体原型          
struct shmid_ds {
	struct ipc_perm shm_perm;    /* Ownership and permissions */
	size_t          shm_segsz;   /* Size of segment (bytes) */
	time_t          shm_atime;   /* Last attach time */
	time_t          shm_dtime;   /* Last detach time */
	time_t          shm_ctime;   /* Last change time */
	pid_t           shm_cpid;    /* PID of creator */
	pid_t           shm_lpid;    /* PID of last shmat(2)/shmdt(2) */
    // 引用计数, 多少个进程和共享内存进行了关联
	shmatt_t        shm_nattch;  /* 记录了有多少个进程和当前共享内存进行了管联 */
	...
};
```

参数:

* shmid: 要操作的共享内存的 ID, 是 shmget () 函数的返回值
* cmd: 要做的操作

1. IPC_STAT: 得到当前共享内存的状态
2. IPC_SET: 设置共享内存的状态
3. IPC_RMID: 标记共享内存要被删除了
buf:
cmd==IPC_STAT, 作为传出参数，会得到共享内存的相关属性信息
cmd==IPC_SET, 作为传入参，将用户的自定义属性设置到共享内存中
cmd==IPC_RMID, buf 就没意义了，这时候 buf 指定为 NULL 即可
返回值：函数调用成功返回值大于等于 0，调用失败返回 - 1

### 3.2 相关 shell 命令

使用 ipcs 添加参数 -m 可以查看系统中共享内存的详细信息

`$ ipcs -m`

------------ 共享内存段 --------------
键        shmid      拥有者  权限     字节     nattch     状态      
0x00000000 425984     oracle     600        524288     2          目标       
0x00000000 327681     oracle     600        524288     2          目标       
0x00000000 458754     oracle     600        524288     2          目标 	
使用 ipcrm 命令可以标记删除某块共享内存

![shareMemory1](../../images/shareMemory1.PNG)

```c++
# key == shmget的第一个参数
$ ipcrm -M shmkey  

# id == shmget的返回值
$ ipcrm -m shmid	

```

### 3.3 共享内存状态

```c++
// 参数 struct shmid_ds 结构体原型          
struct shmid_ds {
	struct ipc_perm shm_perm;    /* Ownership and permissions */
	size_t          shm_segsz;   /* Size of segment (bytes) */
	time_t          shm_atime;   /* Last attach time */
	time_t          shm_dtime;   /* Last detach time */
	time_t          shm_ctime;   /* Last change time */
	pid_t           shm_cpid;    /* PID of creator */
	pid_t           shm_lpid;    /* PID of last shmat(2)/shmdt(2) */
    // 引用计数, 多少个进程和共享内存进行了关联
	shmatt_t        shm_nattch;  /* 记录了有多少个进程和当前共享内存进行了管联 */
	...
};
```

通过 shmctl() 我们可以得知，共享内存的信息是存储到一个叫做 struct shmid_ds 的结构体中，其中有一个非常重要的成员叫做 shm_nattch，在这个成员变量里边记录着当前共享内存关联的进程的个数，一般将其称之为引用计数。当共享内存被标记为删除状态，并且这个引用计数变为 0 之后共享内存才会被真正的被删除掉。

当共享内存被标记为删除状态之后，共享内存的状态也会发生变化，共享内存内部维护的 key 从一个正整数变为 0，其属性从公共的变为私有的。这里的私有是指只有已经关联成功的进程才允许继续访问共享内存，不再允许新的进程和这块共享内存进行关联了。下图演示了共享内存的状态变化：

![shareMemory2](../../images/shareMemory2.PNG)

## 4、 进程间通信

使用共享内存实现进程间通信的操作流程如下：

1. 调用linux的系统API创建一块共享内存
    - 这块内存不属于任何进程, 默认进程不能对其进行操作
    
2. 准备好进程A, 和进程B, 这两个进程需要和创建的共享内存进行关联
    - 关联操作: 调用linux的 api
    - 关联成功之后, 得到了这块共享内存的起始地址
        
3. 在进程A或者进程B中对共享内存进行读写操作
    - 读内存: printf() 等;
	- 写内存: memcpy() 等;

4. 通信完成, 可以让进程A和B和共享内存解除关联
    - 解除成功, 进程A和B不能再操作共享内存了
    - 共享内存不受进程生命周期的影响的
    
5. 共享内存不在使用之后, 将其删除
    - 调用linux的api函数, 删除之后这块内存被内核回收了

* 写共享内存的进程代码:

```c++
#include <stdio.h>
#include <sys/shm.h>
#include <string.h>

int main()
{
    // 1. 创建共享内存, 大小为4k
    int shmid = shmget(1000, 4096, IPC_CREAT|0664);
    if(shmid == -1)
    {
        perror("shmget error");
        return -1;
    }

    // 2. 当前进程和共享内存关联
    //用户不知道，需要让内核指定，写 NULL
    //0表示读写权限
    void* ptr = shmat(shmid, NULL, 0);
    if(ptr == (void *) -1)
    {
        perror("shmat error");
        return -1;
    }

    // 3. 写共享内存
    const char* p = "hello, world, 共享内存真香...";
    memcpy(ptr, p, strlen(p)+1);
    printf("写共享内存：wptr=%s\n", ptr);

    // 阻塞程序
    printf("按任意键继续, 删除共享内存\n");
    //从标准输入 stdin 获取一个字符（一个无符号字符）
    getchar(); 
    //解除共享内存的关联
    shmdt(ptr);

    // 删除共享内存
    shmctl(shmid, IPC_RMID, NULL);
    printf("共享内存已经被删除...\n");

    return 0;
}
```

`gcc -o writeShareMemory writeShareMemory.c`
![shareMemory3](../../images/shareMemory3.PNG)

读共享内存的进程代码:

```c++
#include <stdio.h>
#include <sys/shm.h>
#include <string.h>

int main()
{
    // 1. 创建共享内存, 大小为4k
    int shmid = shmget(1000, 0, 0);
    if(shmid == -1)
    {
        perror("shmget error");
        return -1;
    }

    // 2. 当前进程和共享内存关联
    //用户不知道，需要让内核指定，写 NULL
    //0表示读写权限
    void* ptr = shmat(shmid, NULL, 0);
    if(ptr == (void *) -1)
    {
        perror("shmat error");
        return -1;
    }

    // 3. 读共享内存
    printf("共享内存数据: %s\n", (char*)ptr);

    // 阻塞程序
    printf("按任意键继续, 删除共享内存\n");
    getchar();

    shmdt(ptr);

    // 删除共享内存
    shmctl(shmid, IPC_RMID, NULL);
    printf("共享内存已经被删除...\n");

    return 0;
}
```

`gcc -o readShareMemory readShareMemory.c`
![shareMemory4](../../images/shareMemory4.PNG)


## 5、 shm 和 mmap 的区别

共享内存和内存映射区都可以实现进程间通信，下面来分析一下二者的区别：

1. 实现进程间通信的方式

    * shm: 多个进程只需要一块共享内存就够了，共享内存不属于进程，需要和进程关联才能使用  
    * 内存映射区：位于每个进程的虚拟地址空间中，并且需要关联同一个磁盘文件才能实现进程间数据通信

2. 效率:

    * shm: 直接对内存操作，效率高  
    * 内存映射区：需要内存和文件之间的数据同步，效率低

3. 生命周期

    * shm：进程退出对共享内存没有影响，调用相关函数 / 命令 / 关机才能删除共享内存。
    * 内存映射区：进程退出，内存映射区也就没有了

4. 数据的完整性 -> 突发状态下数据能不能被保存下来（比如：突然断电）

    * shm：数据存储在物理内存中，断电之后系统关闭，内存数据也就丢失了
    * 内存映射区：可以完整的保存数据，内存映射区数据会同步到磁盘文件

