# linux的常见命令

Linux常用命令英文全称与中文解释 （pwd、su、df、du等）
转自:https://www.cnblogs.com/shttke/p/11495080.html

|序号|命令|英文|中文|
|:--|:---|:---|:---|
|1|apt|Advanced Packaging Tool|高级包装工具|
|2|cp|Copy file  |复制文件|
|3|cat| Concatenate|  串联|
|4|cd|Change directory |切换目录|
|5|chmod|Change mode |改变模式|
|6|chown| Change owner |改变所有者|
|7|chgrp| Change group |改变用户组|
|8|df| Disk free | 空余硬盘|
|9|du| Disk usage| 硬盘使用率|
|10|dd|本来应根据其功能描述“Convert an copy”命名为“cc”，但“cc”已经被用以代表“C Complier”，所以命名为“dd”|用于读取、转换并输出数据|
|11|fg| Foreground| 前景|
|12|fsck|File system check|  文件系统检查|
|13|insmod|Install module| 安装模块|
|14|ln| Link files | 链接文件|
|15|lsmod| List module|  列表模块|
|16|ls|List files | 列出目录下的文件|
|17|ldd|List dynamic dependencies |列出动态相依|
|18|mv| Move file|  移动文件|
|19|mkdir|Make directory |建立目录|
|20|mkfs| Make file system| 建立文件系统|
|21|man| Manual  | 意思是手册，可以用这个命令查询其他命令的用法。|
|22|ps|Process Status | 进程状态|
|23|pwd|Print working directory |显示当前工作路径。|
|24|rm| Remove file|  删除文件|
|25|rmmod|Remove module |删除模块|
|26|rmdir|Remove directory | 移动目录|
|27|su|Swith user | 切换用户，切换到root用户|
|28tar|Tape archive |解压文件|
|29|umount| Unmount |卸载|
|30|uname| Unix name  |系统名称|
|31|/bin = |BINaries ||
|32|/dev = |DEVices ||
|33|/etc = |editable text configurationChest | 存放配置文件的地方,配置文件的目录|
|34|/opt = |Optional application software packages |可选应用软件包|
|35|/lib = |library |库|
|36|/proc = |processes  |进程|
|37|/sbin = |Superuser binaries|超级用户的二进制文件|
|38|/tmp = |temporary   |临时文件存放|
|39|/usr = |Unix shared resources  |Unix共享资源|
|40|/var = |Variable |是储存各种变化的文件，比如log等等|

## 1、绝对路径用什么符号表示？当前目录、上层目录用什么表示？主目录用什么表示? 切换目录用什么命令？

答案：
绝对路径： 如/etc/init.d
当前目录和上层目录： ./  ../
主目录： ~/
切换目录： cd

## 2、怎么查看当前进程？怎么执行退出？怎么查看当前路径？

答案：
查看当前进程： ps
执行退出： exit
查看当前路径： pwd

## 3、怎么清屏？怎么退出当前命令？怎么执行睡眠？怎么查看当前用户 id？查看指定帮助用什么命令？

答案：
清屏： clear
退出当前命令： ctrl+c 彻底退出
执行睡眠 ： ctrl+z 挂起当前进程fg 恢复后台
查看当前用户 id： ”id“：查看显示目前登陆账户的 uid 和 gid 及所属分组及用户名
查看指定帮助： 如 man adduser 这个很全 而且有例子； adduser --help 这个告诉你一些常用参数； info adduesr；

## 4、Ls 命令执行什么功能？ 可以带哪些参数，有什么区别？

答案：
ls 执行的功能： 列出指定目录中的目录，以及文件
哪些参数以及区别： a 所有文件l 详细信息，包括大小字节数，可读可写可执行的权限等

## 5、建立软链接(快捷方式)，以及硬链接的命令。

答案：
 ln [参数][源文件或目录][目标文件或目录]
软链接： ln -s slink source
硬链接： ln link source

- 软链接：

1.软链接，以路径的形式存在。类似于Windows操作系统中的快捷方式
2.软链接可以 跨文件系统 ，硬链接不可以
3.软链接可以对一个不存在的文件名进行链接
4.软链接可以对目录进行链接

- 硬链接：

1.硬链接，以文件副本的形式存在。但不占用实际空间。
2.不允许给目录创建硬链接
3.硬链接只有在同一个文件系统中才能创建

## 6、目录创建用什么命令？创建文件用什么命令？复制文件用什么命令？

答案：
创建目录： mkdir
创建文件：典型的如 touch，vi 也可以创建文件，其实只要向一个不存在的文件输出，都会创建文件
复制文件： cp 7. 文件权限修改用什么命令？格式是怎么样的？

## 7、给文件分配权限

文件权限修改： chmod
格式如下：

chmod u+x file给file的属主增加执行权限 chmod 751 file 给 file 的属主用户分配读、写、执行(7)的权限，给 file 的所在组分配读、执行(5)的权限，给其他用户分配执行(1)的权限
chmod u=rwx,g=rx,o=xfile上例的另一种形式 chmod =r file 为所有用户分配读权限
chmod 444 file同上例 chmod a-wx,a+r file同上例
$ chmod -R u+r directory 递归地给 directory 目录下所有文件和子目录的属主分配读的权限
r:4; w:2; x:1
 

## 8、查看文件内容有哪些命令可以使用？

答案：
vi 文件名 #编辑方式查看，可修改
cat 文件名 #显示全部文件内容
more 文件名 #分页显示文件内容
less 文件名 #与 more 相似，更好的是可以往前翻页
tail 文件名 #仅查看尾部，还可以指定行数
head 文件名 #仅查看头部,还可以指定行数

## 9、随意写文件命令？怎么向屏幕输出带空格的字符串，比如”hello world”? 

答案：
写文件命令：vi
向屏幕输出带空格的字符串:echo hello world

## 10、终端是哪个文件夹下的哪个文件？黑洞文件是哪个文件夹下的哪个命令？

答案：
终端  /dev/tty
黑洞文件  /dev/null

## 11、移动文件用哪个命令？改名用哪个命令？

答案：
移动文件：mv c/sourcefile d/sourcefile
修改文件名：mv sourcefile desfile

## 12、复制文件用哪个命令？如果需要连同文件夹一块复制呢？如果需要有提示功能呢？

答案：
cp sourcefile desfile
cp -r  ？？？？

## 13、删除文件用哪个命令？如果需要连目录及目录下文件一块删除呢？删除空文件夹用什么命令？

答案：
rm 
rm -r 
rmdir

## 14、Linux 下命令有哪几种可使用的通配符？分别代表什么含义?

答案：
“？”可替代单个字符。

“*”可替代任意多个字符。

方括号“[charset]”可替代 charset 集中的任何单个字符，如[a-z]，[abABC]

## 15、用什么命令对一个文件的内容进行统计？(行号、单词数、字节数)

答案：
wc 命令 - c 统计字节数 - l 统计行数 - w 统计字数。

## 16、Grep 命令有什么用？ 如何忽略大小写？ 如何查找不含该串的行?

答案：
是一种强大的文本搜索工具，它能使用正则表达式搜索文本，并把匹 配的行打印出来。
grep [stringSTRING] filename 
grep [^string] filename

## 17、Linux 中进程有哪几种状态？在 ps 显示出来的信息中，分别用什么符号表示的？

答案：
（1）、不可中断状态：进程处于睡眠状态，但是此刻进程是不可中断的。不可中断， 指进程不响应异步信号。
（2）、暂停状态/跟踪状态：向进程发送一个 SIGSTOP 信号，它就会因响应该信号 而进入 TASK_STOPPED 状态;当进程正在被跟踪时，它处于 TASK_TRACED 这个特殊的状态。
“正在被跟踪”指的是进程暂停下来，等待跟踪它的进程对它进行操作。

（3）、就绪状态：在 run_queue 队列里的状态

（4）、运行状态：在 run_queue 队列里的状态
（5）、可中断睡眠状态：处于这个状态的进程因为等待某某事件的发生（比如等待 socket 连接、等待信号量），而被挂起
（6）、zombie 状态（僵尸）：父亲没有通过 wait 系列的系统调用会顺便将子进程的尸体（task_struct）也释放掉
（7）、退出状态

D 不可中断 Uninterruptible（usually IO）
R 正在运行，或在队列中的进程
S 处于休眠状态
T 停止或被追踪
Z 僵尸进程
W 进入内存交换（从内核 2.6 开始无效）
X 死掉的进程

## 18、怎么使一个命令在后台运行?

答案：
一般都是使用 & 在命令结尾来让程序自动运行。(命令后可以不追加空格)

可以在命令的末尾加上一个 & 号，将这个任务放到后台去执行：

`cp -R original/dir/ backup/dir/ &`

任务被放到后台执行之后，就可以立即继续在同一个终端上工作了，甚至关闭终端也不影响这个任务的正常执行。需要注意的是，如果要求这个任务输出内容到标准输出中（例如 echo 或 ls），即使使用了 &，也会等待这些输出任务在前台运行完毕。


## 19、利用 ps 怎么显示所有的进程? 怎么利用 ps 查看指定进程的信息？

答案：
1)ps -ef (system v 输出) 
2)ps -aux (bsd 格式输出)
3)top
以全屏交互式的界面显示进程排名，及时跟踪包括CPU、内存等系统资源占用情况，默认情况下每三秒刷新一次，其作用基本类似于Windows系统中的任务管理器。
4)pstree -aup
 可以带上|grep  查询特定进程。例如 pstree -aup | grep php

ps -ef | grep pid

## 20、哪个命令专门用来查看后台任务? 

答案：
job -l

## 21、把后台任务调到前台执行使用什么命令?把停下的后台任务在后台执行起来用什么命令?

答案：
把后台任务调到前台执行 fg
把停下的后台任务在后台执行起来 bg

- 下列命令可以用来操纵进程任务：
1、ps 列出系统中正在运行的进程；
2、kill 发送信号给一个或多个进程（经常用来杀死一个进程）；
3、jobs 列出当前shell环境中已启动的任务状态，若未指定jobsid，则显示所有活动的任务状态信息；如果报告了一个任务的终止(即任务的状态被标记为Terminated)，shell
从当前的shell环境已知的列表中删除任务的进程标识；
4、bg 将进程搬到后台运行（Background）；
5、fg 将进程搬到前台运行（Foreground）；
 
## 22、终止进程用什么命令? 带什么参数? 

答案：

kill [-s <信息名称或编号>][程序] 或 kill [-l <信息编号>] 

kill -9 pid

最常用的信号是：
1 (HUP)：重新加载进程。
9 (KILL)：杀死一个进程。
15 (TERM)：正常停止一个进程。
 
Linux pkill 用于杀死一个进程，与 kill 不同的是它会杀死指定名字的所有进程，类似于 killall 命令。

kill 命令杀死指定进程 PID，需要配合 ps 使用，而 pkill 直接对进程对名字进行操作，更加方便。

- pkill

语法
  pkill [选项]  name
参数说明：

name ： 进程名
选项包含如下几个参数：
-o 仅向找到的最小（起始）进程号发送信号 -n 仅向找到的最大（结束）进程号发送信号
-P 指定父进程号发送信号
-g 指定进程组
-t 指定开启进程的终端
实例
`pkill -9  php-fpm`   //结束所有的 php-fpm 进程

## 23、怎么查看系统支持的所有信号？

答案：
kill -l

## 24、搜索文件用什么命令? 格式是怎么样的? 

答案：

1. find <指定目录> <指定条件> <指定动作>
2. whereis 加参数与文件名;
该指令会在特定目录中查找符合条件的文件。这些文件应属于原始代码、二进制文件，或是帮助文件。
3. locate 只加文件名;
用于查找符合条件的文档，他会去保存文档和目录名称的数据库内，查找合乎范本样式条件的文档或目录。

一般情况我们只需要输入 locate your_file_name 即可查找指定文件。

find 直接搜索磁盘，较慢。
find / -name "string*"

## 25、查看当前谁在使用该主机用什么命令? 查找自己所在的终端信息用什么命令?

答案：
查找自己所在的终端信息：who am i
查看当前谁在使用该主机：who

## 26、使用什么命令查看用过的命令列表?

答案：history

## 27、使用什么命令查看磁盘使用空间？ 空闲空间呢?

答案：
df -hl
文件系统 容量 已用 可用 已用% 挂载点
Filesystem Size Used Avail Use% Mounted on /dev/hda2 45G 19G 24G 44% /
/dev/hda1 494M 19M 450M 4% /boot

## 28、使用什么命令查看网络是否连通?

答案：
netstat

## 29、使用什么命令查看 ip 地址及接口信息？

答案：
ifconfig

## 30、查看各类环境变量用什么命令?

答案：
查看所有 env
查看某个，如 home： env $HOME

## 31、通过什么命令指定命令提示符?

答案：
\u：显示当前用户账号
\h：显示当前主机名
\W：只显示当前路径最后一个目录
\w：显示当前绝对路径（当前用户目录会以~代替）
$PWD：显示当前全路径
$：显示命令行’$'或者’#'符号
\#：下达的第几个命令
\d：代表日期，格式为week day month date，例如："MonAug1"
\t：显示时间为24小时格式，如：HH：MM：SS
\T：显示时间为12小时格式
\A：显示时间为24小时格式：HH：MM
\v：BASH的版本信息 如export PS1=’[\u@\h\w\#]$‘

## 32、查找命令的可执行文件是去哪查找的? 怎么对其进行设置及添加? 

答案：
`whereis [-bfmsu][-B <目录>...][-M <目录>...][-S <目录>...][文件...]`

补充说明：

1. whereis 指令会在特定目录中查找符合条件的文件。这些文件的烈性应属于原始代码，二进制文件，或是帮助文件。

-b   只查找二进制文件。
-B<目录> 只在设置的目录下查找二进制文件。 -f 不显示文件名前的路径名称。
-m   只查找说明文件。
-M<目录> 只在设置的目录下查找说明文件。 -s 只查找原始代码文件。
-S<目录> 只在设置的目录下查找原始代码文件。 -u 查找不包含指定类型的文件。

2. which 指令会在 PATH 变量指定的路径中，搜索某个系统命令的位置，并且返回第一个搜索结果。

-n 指定文件名长度，指定的长度必须大于或等于所有文件中最长的文件名。
-p 与-n 参数相同，但此处的包括了文件的路径。 -w 指定输出时栏位的宽度。
-V   显示版本信息

## 33、通过什么命令查找执行命令?

答案：
which 只能查可执行文件
whereis 只能查二进制文件、说明文档，源文件等


## 34、怎么对命令进行取别名？

答案：
alias la='ls -a'

## 35、du 和 df 的定义，以及区别？

答案：
du 显示目录或文件的大小
df 显示每个<文件>所在的文件系统的信息，默认是显示所有文件系统。
（文件系统分配其中的一些磁盘块用来记录它自身的一些数据，如 i 节点，磁盘分布图，间接块，超级块等。这些数据对大多数用户级的程序来说是不可见的，通常称为 Meta Data。） du 命令是用户级的程序，它不考虑 Meta Data，而 df 命令则查看文件系统的磁盘分配图并考虑 Meta Data。
df 命令获得真正的文件系统数据，而 du 命令只查看文件系统的部分情况。

## 36、处理文本文件的语言

AWK 是一种处理文本文件的语言，是一个强大的文本分析工具。

答案：

```linux
awk '{pattern + action}' {filenames}
#cat /etc/passwd |awk -F ':' '{print 1"\t"7}' //-F 的意思是以':'分隔 root /bin/bash
daemon /bin/sh 搜索/etc/passwd 有 root 关键字的所有行

#awk -F: '/root/' /etc/passwd root:x:0:0:root:/root:/bin/bash
```
 

## 37、当你需要给命令绑定一个宏或者按键的时候，应该怎么做呢？

答案：
可以使用bind命令，bind可以很方便地在shell中实现宏或按键的绑定。
在进行按键绑定的时候，我们需要先获取到绑定按键对应的字符序列。

比如获取F12的字符序列获取方法如下：先按下Ctrl+V,然后按下F12 .我们就可以得到F12的字符序列 ^[[24~。

接着使用bind进行绑定。

[root@localhost ~]# bind ‘”\e[24~":"date"'

注意：相同的按键在不同的终端或终端模拟器下可能会产生不同的字符序列。

【附】也可以使用showkey -a命令查看按键对应的字符序列。

## 38、如果一个linux新手想要知道当前系统支持的所有命令的列表，他需要怎么做？

答案：

使用命令compgen ­-c，可以打印出所有支持的命令列表。
[root@localhost ~]$ compgen -c

l.
ll
ls
which
if
then
else
elif
fi
case
esac
for
select
while
until
do
done
…

## 39、如果你的助手想要打印出当前的目录栈，你会建议他怎么做？

答案：

使用Linux 命令dirs可以将当前的目录栈打印出来。
[root@localhost ~]# dirs
/usr/share/X11
【附】：目录栈通过pushd popd 来操作。

## 40、你的系统目前有许多正在运行的任务，在不重启机器的条件下，有什么方法可以把所有正在运行的进程移除呢？

答案：
使用linux命令 ’disown -r ’可以将所有正在运行的进程移除。

disown命令可以从当前shell的作业列表中移除全部作业，也可移除指定的一到多个作业；正在运行的作业也可以移除；也可以标记作业，使得它们在当前shell退出后也不会结束。
 
## 41、bash shell 中的hash 命令有什么作用？

答案：

linux命令’hash’管理着一个内置的哈希表，记录了已执行过的命令的完整路径, 用该命令可以打印出你所使用过的命令以及执行的次数。

```linux
[root@localhost ~]# hash
hits command
2 /bin/ls
2 /bin/su
```

## 42、哪一个bash内置命令能够进行数学运算。

答案：
bash shell 的内置命令let 可以进行整型数的数学运算。

```linux
#! /bin/bash
…
…
let c=a+b
…
…
```
 

## 43、怎样一页一页地查看一个大文件的内容呢？

答案：

通过管道将命令”cat file_name.txt” 和 ’more’ 连接在一起可以实现这个需要.

[root@localhost ~]# cat file_name.txt | more

## 44、数据字典属于哪一个用户的？

答案：

数据字典是属于’SYS’用户的，用户‘SYS’ 和 ’SYSEM’是由系统默认自动创建的

## 45、怎样查看一个linux命令的概要与用法？假设你在/bin目录中偶然看到一个你从没见过的的命令，怎样才能知道它的作用和用法呢？

答案：

使用命令whatis 可以先出显示出这个命令的用法简要，比如，你可以使用whatis zcat 去查看‘zcat’的介绍以及使用简要。

```linux
[root@localhost ~]# whatis zcat
zcat [gzip] (1) – compress or expand files
```

## 46、使用哪一个命令可以查看自己文件系统的磁盘空间配额呢？

答案：

使用命令repquota 能够显示出一个文件系统的配额信息

【附】只有root用户才能够查看其它用户的配额。

## 47、查看cpu占有率？

1. vmstat

vmstat是Virtual Meomory Statistics（虚拟内存统计）的缩写，可对操作系统的虚拟内存、进程、CPU活动进行监控。是对系统的整体情况进行统计，不足之处是无法对某个进程进行深入分析。

2. top

Linux top命令用于实时显示 process 的动态。