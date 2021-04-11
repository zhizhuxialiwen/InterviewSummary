gn、ninja的安装-Ubuntu18.04

邱国禄 2020-01-05 14:02:22  9646  收藏 19
分类专栏： 环境搭建
版权
版权声明：原创文章，欢迎转载，但请注明出处，谢谢。https://blog.csdn.net/qiuguolu1108/article/details/103842556


如果你不想编译gn、ninja，想直接使用gn、ninja的二进制程序，可以直接到博客的最后，通过链接直接下载，省去自己编译。分享的链接中有测试用例，可以直接测试gn、ninja是否可用。


文章目录
ninja的安装
一、安装依赖
二、下载ninja
三、编译ninja
四、安装ninja
gn的安装
一、先安装clang
二、下载gn
三、编译 gn
四、安装gn
gn和ninja的二进制程序
一、二进制的gn和ninja
二、测试用例

gn的安装需要使用ninja，所以首先安装ninja。
ninja的安装需要依赖re2c，gn的安装需要使用clang编译器，并且gn中使用了C++17，在Ubuntu16安装的clang-6.0是不支持C++17的，为了方便安装转战到Ubuntu18。

每次最头疼的都是搭建环境，特别浪费时间。目标是学习gn和ninja，所以怎么方便怎么来。在Ubuntu18搭建环境要比Ubuntu16方便很多，所以选用Ubuntu18。

ninja的安装
一、安装依赖
在安装ninja之前，需要安装其依赖re2c。

root@ubuntu:~# apt-get install re2c
root@ubuntu:~# re2c --version
re2c 1.0.1
1
2
3
我安装的是 1.0.1 版本

二、下载ninja
在github中下载ninja，ninja github地址https://github.com/ninja-build/ninja。

git clone https://github.com/ninja-build/ninja.git
1
三、编译ninja
进入刚才下载的ninja目录中，执行编译脚本。

./configure.py --bootstrap   #在ninja目录中执行
1
四、安装ninja
编译结束后，会在ninja目录中生成ninja的可执行程序ninja。可以直接将ninja程序拷贝到/usr/bin中，方便又省事。

cp ./ninja  /usr/bin  #在ninja目录中执行
1
现在就可以在任意位置使用ninja了。

效果如下：

root@ubuntu:~# ninja --version
1.9.0.git
1
2
gn的安装
下载最新版的gn貌似需要翻墙，直接在github中找了一个，虽然不是最新版的，但可以用。

gn的官方源：https://chromium.googlesource.com/chromium/src/tools/gn

我自己使用的github链接：https://github.com/timniederhausen/gn

一、先安装clang
gn的编译需要使用clang编译器，并且gn使用了C++17的语法，所以需要使用较高版本的clang。

我自己安装的是clang 7.0，是可以使用的。

sudo apt-get install clang-7
1
安装clang以后需要做一点修改，用上述方法安装的clang，直接在命令行中输入clang是无法使用的，在/usr/bin目录下看到是clang-7、clang++-7、clang-cpp-7。但在编译gn的时候，需要使用clang++命令，所以给这个三个可执行程序做一个软连接，修改一下它们的名字。

进入/usr/bin/目录，修改如下：

ln -s clang-7 clang
ln -s clang++-7 clang++
ln -s clang-cpp-7 clang-cpp
1
2
3
在命令行中输入clang --version，显示如下内容，说明clang安装成功。

root@ubuntu:~# clang++ --version
clang version 7.0.0-3~ubuntu0.18.04.1 (tags/RELEASE_700/final)
Target: x86_64-pc-linux-gnu
Thread model: posix
InstalledDir: /usr/bin
1
2
3
4
5
二、下载gn
git clone https://github.com/timniederhausen/gn.git
1
三、编译 gn
进入刚才下载的gn目录中，先执行gn的配置脚本。

./build/gen.py
1
然后在gn目录中执行：

ninja -C out
1
编译结束后，gn程序就在gn/out目录中。

四、安装gn
将gn/out目录下的gn复制到/usr/bin目录就可以在任意位置使用gn了。

cp ./out/gn /usr/bin     #在gn目录下执行
1
效果如下：

root@ubuntu:~# gn --version
1641 (0a06cb92a)
1
2
gn和ninja的二进制程序
一、二进制的gn和ninja
如果你嫌麻烦，可以直接使用我编译好的可执行文件。这两个二进制文件，我测试了一下，可以在Ubuntu-16.04和Ubuntu-18.04上运行。
链接：https://pan.baidu.com/s/1_l8JMfuhLJgD7RKs-IDOnQ 提取码：1e0r
如果链接失效了，可以提醒我，我会及时更新链接。
将gn、ninja下载后，直接放在/usr/bin目录下，这样在任意位置可以直接使用这个两个程序了。

二、测试用例
在分享的文件夹中，我提供了一份测试用例，将上述两个可执行文件放到/usr/bin目录后，进入gn-demo目录。

给gn和ninja添加执行权限
下载拷贝过去后，可能gn和ninja没有了执行权限，如果出现下面情况，可以给其添加执行权限。
root@learner:~/gn-demo# gn gen ./out
-bash: /usr/bin/gn: Permission denied
1
2
使用以下命令添加执行权限：

root@learner:~/gn-demo# chmod +x /usr/bin/gn 
root@learner:~/gn-demo# chmod +x /usr/bin/ninja 
1
2
先使用gn生成工程。
root@learner:~/gn-demo# gn gen ./out                #使用gn生成ninja工程
Done. Made 2 targets from 4 files in 4ms
1
2
再使用ninja生成可执行文件。
root@learner:~/gn-demo# ninja -C ./out             #使用ninja生成可执行文件
ninja: Entering directory `./out'
[3/3] STAMP obj/default.stamp
1
2
3
现在可执行文件就在./out目录中。
root@learner:~/gn-demo# cd out/
root@learner:~/gn-demo/out# ls
args.gn  build.ninja  build.ninja.d  hello  obj  toolchain.ninja
1
2
3
执行可执行文件hello
root@learner:~/gn-demo/out# ./hello               #运行可执行程序hello
hello world
————————————————
版权声明：本文为CSDN博主「邱国禄」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/qiuguolu1108/article/details/103842556