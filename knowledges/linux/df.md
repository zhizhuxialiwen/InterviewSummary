# Linux df 命令

Linux df（英文全拼：disk free） 命令用于显示目前在 Linux 系统上的文件系统磁盘使用情况统计。

* 语法

`df [选项]... [FILE]...`

文件-a, --all 包含所有的具有 0 Blocks 的文件系统
文件--block-size={SIZE} 使用 {SIZE} 大小的 Blocks
文件-h, --human-readable 使用人类可读的格式(预设值是不加这个选项的...)
文件-H, --si 很像 -h, 但是用 1000 为单位而不是用 1024
文件-i, --inodes 列出 inode 资讯，不列出已使用 block
文件-k, --kilobytes 就像是 --block-size=1024
文件-l, --local 限制列出的文件结构
文件-m, --megabytes 就像 --block-size=1048576
文件--no-sync 取得资讯前不 sync (预设值)
文件-P, --portability 使用 POSIX 输出格式
文件--sync 在取得资讯前 sync
文件-t, --type=TYPE 限制列出文件系统的 TYPE
文件-T, --print-type 显示文件系统的形式
文件-x, --exclude-type=TYPE 限制列出文件系统不要显示 TYPE
文件-v (忽略)
文件--help 显示这个帮手并且离开
文件--version 输出版本资讯并且离开
实例
显示文件系统的磁盘使用情况统计：

```
# df 
Filesystem     1K-blocks    Used     Available Use% Mounted on 
/dev/sda6       29640780 4320704     23814388  16%     / 
udev             1536756       4     1536752    1%     /dev 
tmpfs             617620     888     616732     1%     /run 
none                5120       0     5120       0%     /run/lock 
none             1544044     156     1543888    1%     /run/shm 
```

第一列指定文件系统的名称，第二列指定一个特定的文件系统1K-块1K是1024字节为单位的总内存。用和可用列正在使用中，分别指定的内存量。

使用列指定使用的内存的百分比，而最后一栏"安装在"指定的文件系统的挂载点。

df也可以显示磁盘使用的文件系统信息：

```
# df test 
Filesystem     1K-blocks    Used      Available Use% Mounted on 
/dev/sda6       29640780    4320600   23814492  16%       / 
```

用一个-i选项的df命令的输出显示inode信息而非块使用量。

```
wen@ubuntu:~/gopath/src$ df -i
Filesystem      Inodes  IUsed   IFree IUse% Mounted on
udev            244732    445  244287    1% /dev
tmpfs           252160    679  251481    1% /run
/dev/sda1      3874816 198791 3676025    6% /
tmpfs           252160      9  252151    1% /dev/shm
tmpfs           252160      5  252155    1% /run/lock
tmpfs           252160     17  252143    1% /sys/fs/cgroup
tmpfs           252160     27  252133    1% /run/user/1000
```

显示所有的信息:

```
# df --total 
Filesystem     1K-blocks    Used    Available Use% Mounted on 
/dev/sda6       29640780 4320720    23814372  16%     / 
udev             1536756       4    1536752   1%      /dev 
tmpfs             617620     892    616728    1%      /run 
none                5120       0    5120      0%      /run/lock 
none             1544044     156    1543888   1%      /run/shm 
total           33344320 4321772    27516860  14% 
```

我们看到输出的末尾，包含一个额外的行，显示总的每一列。

-h选项，通过它可以产生可读的格式df命令的输出：

```
wen@ubuntu:~/gopath/src$ df -h
Filesystem      Size  Used Avail Use% Mounted on
udev            956M     0  956M   0% /dev
tmpfs           197M  6.3M  191M   4% /run
/dev/sda1        58G  4.8G   51G   9% /
tmpfs           985M  212K  985M   1% /dev/shm
tmpfs           5.0M  4.0K  5.0M   1% /run/lock
tmpfs           985M     0  985M   0% /sys/fs/cgroup
tmpfs           197M   44K  197M   1% /run/user/1000
```
我们可以看到输出显示的数字形式的'G'（千兆字节），"M"（兆字节）和"K"（千字节）。

这使输出容易阅读和理解，从而使显示可读的。请注意，第二列的名称也发生了变化，为了使显示可读的"大小"。