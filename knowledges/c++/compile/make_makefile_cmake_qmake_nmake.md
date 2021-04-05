# MAKE/MAKEFILE/CMAKE/QMAKE/NMAKE/区别与介绍

# 1、GCC

在正式开始介绍这些工具之前，先来看看GCC：

gcc是GNU Compiler Collection（GNU编译器集合），GCC是所谓的“ GNU工具链”的关键组件，用于开发应用程序和编写操作系统），可以简单认为是编译器，它可以编译很多种编程语言C ++（g++），Objective-C，Objective-C ++，Java（gcj），Fortran（gfortran），Ada（gnat），Go（gccgo），OpenMP，Cilk Plus和OpenAcc


母网站http://gcc.gnu.org/
gcc编译过程

![makefile1](../../../images/makefile1.PNG)
在这里插入图片描述
用gcc编译一个程序实例：

```c++
/*hello.c*/
#include <stdio.h>
int main(){
    printf("Hello world!\n");
    return 0;
}
```

编译并执行
```
$ gcc -Wall -g -o hello hello.c
$ ./hello 
# 执行
Hello world!
```

## 2、MAKE和MAKEFILE

编译源代码文件可能会很麻烦，尤其是当您必须包含多个源文件并在每次需要编译时都键入compiling命令时。Make工具（GNU Make）是简化此任务的解决方案.

make工具可以看成是一个智能的批处理工具，它本身并没有编译和链接的功能，而是用类似于批处理的方式—通过调用makefile文件中用户指定的命令来进行编译和链接的。
Makefile是特殊格式的文件，可帮助自动构建和管理项目。

推荐几个makefile学习网址：
1.https://www3.ntu.edu.sg/home/ehchua/programming/cpp/gcc_make.html
2.https://www.tutorialspoint.com/makefile/makefile_quick_guide.htm
3.https://makefiletutorial.com/

下面将用一个最简单的实例来展示makefile的使用过程

makefile的一个实例：

```c++
/*hello.c*/
#include <stdio.h>
int main(){
    printf("Hello world!\n");
    return 0;
}
```

makefile文件：
```
all: hello

hello.exe: hello.o
	 gcc -o hello hello.o

hello.o: hello.c
	 gcc -c hello.c
     
clean:
	 rm hello.o hello
```

用 make 运行“ make”实用程序，如下所示：
```
$ make
gcc -c hello.c
gcc -o hello hello.o

$ ./hello 
Hello world!

$ make clean
rm hello.o hello
```

## 3、 CMAKE和CMAKELISTS

makefile在一些简单的工程完全可以人工编写，但是当工程非常大的时候，手写makefile也非常麻烦。如果软件想跨平台，必须要保证能够在不同平台编译。而如果使用不同的 Make 工具，就得为每一种标准写一次 Makefile ，这将是一件让人抓狂的工作。

CMake是一个跨平台的安装（编译）工具，可以用简单的语句来描述所有平台的安装(编译过程)。他能够输出各种各样的makefile或者project文件，能测试编译器所支持的C++特性,类似UNIX下的automake。只是 CMake 的组态档取名为 CMakeLists.txt。Cmake 并不直接建构出最终的软件，而是产生标准的建构档（如 Unix 的 Makefile 或 Windows Visual C++ 的 projects/workspaces），然后再依一般的建构方式使用。这使得熟悉某个集成开发环境（IDE）的开发者可以用标准的方式建构他的软件，这种可以使用各平台的原生建构系统的能力是 CMake 和 SCons 等其他类似系统的区别之处。
cmake官网：https://cmake.org/
在 linux 平台下使用 CMake 生成 Makefile 并编译的流程如下：

编写 CMake 配置文件 CMakeLists.txt 。
执行命令 cmake PATH （或ccmake PATH）生成 Makefile 1，ccmake 和 cmake 的区别在于前者提供了一个交互式的界面。 PATH 指CMakeLists.txt 所在的目录。
使用 make 命令进行编译。
下面将用一个最简单的实例来展示Cmake的使用过程:

Cmake的一个实例：

```c
/*hello.c*/
#include <stdio.h>
int main(){
    printf("Hello world!\n");
    return 0;
}
```

编写CMakeLists.txt文件：
```
# CMake 最低版本号要求

cmake_minimum_required (VERSION 3.13.0)

# 项目信息

project (Hello)

# 指定生成目标hello

add_executable(hello hello.c)
```

* 执行cmake . ，生成Makefile文件：

![makefile2](../../../images/makefile2.PNG)

执行make，运行“ make”实用程序

![makefile3](../../../images/makefile3.PNG)

## 4、 其他MAKE工具

类似于GNU Make ，有其他的公司按照自己的标准构造了一套make工具，QT 的 qmake ，微软的 MS nmake，BSD Make（pmake），Makepp，等等。这些 Make 工具遵循着不同的规范和标准，所执行的 Makefile 格式也千差万别. 因为暂时只用make和cmake工具，下面其他的工具我只做引入.（其实是我不会）在这里插入图片描述

### 4.1 QMAKE

qmake是一个工具，可帮助简化跨不同平台的开发项目的构建过程（对，和cmake一样也能跨平台）。qmake自动生成Makefile，因此只需要几行信息即可创建每个Makefile。qmake无论是否用Qt编写，它都可以用于任何软件项目。
qmake根据项目文件中的信息生成Makefile。项目文件由开发人员创建，通常很简单，但是可以为复杂项目创建更复杂的项目文件。qmake包含其他支持Qt开发的功能，自动包括moc和uic的构建规则。qmake还可以为Microsoft Visual Studio生成项目，而无需开发人员更改项目文件。
在这里插入图片描述

官方网址https://doc.qt.io/archives/qt-4.8/qmake-manual.html

### 4.2 NMAKE

Microsoft 程序维护实用工具 (NMAKE。EXE) 是基于说明文件中包含的命令生成项目的 Visual Studio 附带的命令行工具。要使用 NMAKE，必须在开发人员命令提示窗口中运行它。 开发人员命令提示窗口具有为工具、库设置的环境变量，并且包括在命令行上生成所需的文件路径。

官方网址https://docs.microsoft.com/zh-cn/cpp/build/reference/running-nmake?view=vs-2019
QMAKE和CMAKE有什么区别？
cmake也是同样支持Qt程序的，cmake也能生成针对qt 程序的那种特殊makefile， 只是cmake的CMakeLists.txt 写起来相对与qmake的pro文件复杂点。 qmake 是为 Qt 量身打造的，使用起来非常方便，但是cmake功能比qmake强大。 一般的Qt工程直接使用qmake就可以了，cmake的强大功能一般人是用不到的。 当你的工程非常大的时候，又有qt部分的子工程，又有其他语言的部分子工程，据说用cmake会 方便，我也没试过。
链接：https://www.zhihu.com/question/27455963/answer/89770919