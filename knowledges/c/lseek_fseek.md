# lseek()函数和fseek()函数的使用详解

## 1、lseek()函数--local seek

C语言lseek()函数：移动文件的读写位置

1. 头文件：

```c
#include <sys/types.h>  
#include <unistd.h>
```

2. 定义函数：

`off_t lseek(int fildes, off_t offset, int whence);`
函数说明：
每一个已打开的文件都有一个读写位置, 当打开文件时通常其读写位置是指向文件开头, 若是以附加的方式打开文件(如O_APPEND), 则读写位置会指向文件尾. 当read()或write()时, 读写位置会随之增加,lseek()便是用来控制该文件的读写位置. 参数fildes 为已打开的文件描述词, 参数offset 为根据参数whence来移动读写位置的位移数.

* 参数 whence 为下列其中一种:

    * SEEK_SET 参数offset 即为新的读写位置.
    * SEEK_CUR 以目前的读写位置往后增加offset 个位移量.
    * SEEK_END 将读写位置指向文件尾后再增加offset 个位移量. 当whence 值为SEEK_CUR 或
    * SEEK_END 时, 参数offet 允许负值的出现.

* 下列是教特别的使用方式:

    * 欲将读写位置移到文件开头时:lseek(int fildes, 0, SEEK_SET);
    * 欲将读写位置移到文件尾时:lseek(int fildes, 0, SEEK_END);
    * 想要取得目前文件位置时:lseek(int fildes, 0, SEEK_CUR);

* 返回值：当调用成功时则返回目前的读写位置, 也就是距离文件开头多少个字节. 若有错误则返回-1, errno 会存放错误代码.

附加说明：Linux 系统不允许lseek()对tty 装置作用, 此项动作会令lseek()返回ESPIPE.

## 2、fseek()函数

C语言fseek()函数：移动文件流的读写位置

1. 头文件：

`#include <stdio.h>`

2. 定义函数：

`int fseek(FILE * stream, long offset, int whence);`

函数说明：fseek()用来移动文件流的读写位置。

* 参数stream 为已打开的文件指针,
* 参数offset 为根据参数whence 来移动读写位置的位移数。参数 whence 为下列其中一种:

    * SEEK_SET 从距文件开头offset 位移量为新的读写位置. 
    * SEEK_CUR 以目前的读写位置往后增加offset 个位移量.
    * SEEK_END 将读写位置指向文件尾后再增加offset 个位移量. 当whence 值为SEEK_CUR 或
    * SEEK_END 时, 参数offset 允许负值的出现.

下列是较特别的使用方式：

1) 欲将读写位置移动到文件开头时:fseek(FILE *stream, 0, SEEK_SET);
2) 欲将读写位置移动到文件尾时:fseek(FILE *stream, 0, 0SEEK_END);

* 返回值：当调用成功时则返回0, 若有错误则返回-1, errno 会存放错误代码.

**附加说明：fseek()不像lseek()会返回读写位置, 因此必须使用ftell()来取得目前读写的位置. **

范例

```c++
#include <stdio.h>
main()
{
  FILE * stream;
  long offset;
  fpos_t pos;
  stream = fopen("/etc/passwd", "r");
  fseek(stream, 5, SEEK_SET);
  printf("offset = %d\n", ftell(stream));
  rewind(stream);
  fgetpos(stream, &pos);
  printf("offset = %d\n", pos);
  pos = 10;
  fsetpos(stream, &pos);
  printf("offset = %d\n", ftell(stream));
  fclose(stream);
}
```

执行

offset = 5
offset = 0
offset = 10