# Android之adb的使用

## Android的adb

一、adb定义
二、adb的作用及细节
    2.1 adb作用
    2.2 adb服务原理
        通过WLAN连接设备：
三、ADB基本用法
    3.1 命令语法
    3.2 为命令指定目标设备
        查询设备
        网络命令
        安装应用：
        卸载应用
        设置端口转发
        将文件复制到设备/从设备复制文件
        停止 adb 服务器
四、Logcat的使用
    4.1 命令行语法

## 一、adb定义

adb的全称为Android Debug Bridge，就是起到调试桥的作用。通过adb我们可以在Eclipse中方便通过DDMS来调试Android程序，说白了就是debug工具。adb的工作方式比较特殊，采用监听Socket TCP 5554等端口的方式让IDE和Qemu通讯，默认情况下adb会daemon相关的网络端口，所以当我们运行Eclipse时adb进程就会自动运行。

官方文档：

1. Android Debug Bridge (adb) is a versatile command line tool that lets you communicate with an emulator instance or connected Android-powered device. It is a client-server program that includes three components:

2. A client, which runs on your development machine. You can invoke a client from a shell by issuing an adb command. Other Android tools such as the ADT plugin and DDMS also create adb clients.
3. A server, which runs as a background process on your development machine. The server manages communication between the client and the adb daemon running on an emulator or device.
4. A daemon, which runs as a background process on each emulator or device instance.
.......

Android Debug Bridge（adb）是一个多功能命令行工具，可让您与虚拟机或连接的Android设备进行通信。 它是一个客户端 - 服务器程序，包括三个组件：

* 客户端。 您可以通过发出adb命令从shell调用客户端。 其他Android工具（如ADT插件和DDMS）也可以创建adb客户端。
* 服务器，在开发计算机上作为后台进程运行。 服务器管理客户端与在仿真器或设备上运行的adb守护程序之间的通信。
* 后台程序，在每个模拟器或设备实例上作为后台进程运行。

## 二、adb的作用及细节

### 2.1 adb作用

ADB 是 Android SDK 里的一个工具, 用这个工具可以直接操作管理Android模拟器或者真实的Android设备。它的主要功能有:

在设备上运行Shell命令； 将本地APK软件安装至模拟器或Android设备；
管理设备或手机模拟器上的预定端口；
在设备或手机模拟器上复制或粘贴文件。

### 2.2 adb服务原理

adb是Android SDK中的一个工具，可以在<sdk>/platform-rools/找到
当开启一个adb的客户端时，客户端会在第一次检测是否已经存在adb服务。如果没有，那它将会开启一个服务。当adb服务启动时，它会绑定TCP的5037端口，用来监听来自客户端的命令。
注意：所有adb的客户端都共用5037这个端口号。
之后，adb服务器会与所有正在运行的虚拟机/连接设备建立连接，它通过扫描从5555到5585的奇数（odd）端口号来定位虚拟机/设备。
当服务器发现一个adb的守护进程（daemon）时，它就会与之建立连接。

注意：每一个虚拟机/连接设备都需要一对连续的端口号 
-- 偶数端口号用来连接控制台
-- 奇数端口号用来连接adb

例如：
假设用adb devices输出的是：

Emulator 1, console: 5554
Emulator 1, adb: 5555
Emulator 2, console: 5556
Emulator 2, adb: 5557 ...

你可以看到连接5554端口的虚拟机接口和5555端口是一样的

一旦服务器设置了与所有仿真器实例的连接，您就可以使用adb命令来控制和访问这些实例。 由于服务器管理与虚拟机/设备实例的连接并处理来自多个adb客户端的命令，因此您可以从任何客户端（或从脚本）控制任何虚拟机/设备实例。

以下部分描述了可用于访问adb功能和管理虚拟机/设备状态的命令。 请注意，如果您在Eclipse中开发Android应用程序并安装了ADT插件，则无需从命令行访问adb。 ADT插件提供了adb到Eclipse IDE的透明集成（transparent integration）。 但是，您仍然可以根据需要直接使用adb，例如用于调试。

* 通过WLAN连接设备：

1. 设置目标设备以监听端口 5555 上的 TCP/IP 连接。
	`adb tcpip 5555`

2. 通过 IP 地址连接到设备
	`adb connect device_ip_address`

3. 确认主机已连接到目标设备：
	`$ adb devices`
	
    List of devices attached
    device_ip_address:5555 device

## 三、ADB基本用法

### 3.1 命令语法

adb 命令的基本语法如下：

`adb [-d|-e|-s <serialNumber>] <command>`

如果只有一个设备/模拟器连接时，可以省略掉 [-d|-e|-s <serialNumber>] 这一部分，直接使用 adb <command>。

### 3.2 为命令指定目标设备

如果有多个设备/模拟器连接，则需要为命令指定目标设备。

|参数|	含义|
|:--|:--|
|-d	|指定当前唯一通过 USB 连接的 Android 设备为命令目标|
|-e	|指定当前唯一运行的模拟器为命令目标|
|-s	|<serialNumber>指定相应 serialNumber 号的设备/模拟器为命令目标|

* 在多个设备/模拟器连接的情况下较常用的是 -s 
<serialNumber> 参数，serialNumber 可以通过 adb devices 命令获取。如：

`$ adb devices`

List of devices attached
cf264b8f	device
emulator-5554	device
10.129.164.6:5555	device

输出里的 cf264b8f、emulator-5554 和 10.129.164.6:5555 即为 serialNumber。

比如这时想指定 cf264b8f 这个设备来运行 adb 命令获取屏幕分辨率：

`adb -s cf264b8f shell wm size`

又如想给 emulator:5554 这个设备安装应用（这种形式的 serialNumber 格式为 <IP>:<Port>，一般为无线连接的设备或 Genymotion 等第三方 Android 模拟器）：

`adb -s 10.129.164.6:5555 install test.apk`

#### 3.2.1 查询设备

在发出 adb 命令之前，了解哪些设备实例已连接到 adb 服务器会很有帮助。您可以使用 devices 命令生成已连接设备的列表。

  `adb devices -l`

作为回应，adb 会针对每个设备输出此状态信息：

1. 序列号：由 adb 创建的字符串，用于通过端口号唯一标识设备。下面是一个序列号示例：emulator-5554
2. 状态：设备的连接状态可为下列状态之一：

* offline：设备未连接到 adb 或没有响应。
* device：设备已连接到 adb 服务器。请注意，此状态并不表示 Android 系统已完全启动并可正常运行，因为在设备连接到 adb 时系统仍在启动。不过，在启动后，这将是设备的正常运行状态。
* no device：未连接到设备。

说明：如果包含 -l 选项，devices 命令会告知您设备是什么。如果您连接了多个设备，此信息可帮助您区分这些设备。

#### 3.2.2 网络命令

|命令|	说明|
|:--|:--|
|connect host[:port]|	通过 TCP/IP 连接到设备。如果您未指定端口，则使用默认端口 5555。|
|disconnect [host | host:port]|	断开与在指定端口上运行的指定 TCP/IP 设备的连接。如果未指定主机或端口，则所有设备都将与所有 TCP/IP 端口断开连接。如果指定了主机，但未指定端口，则使用默认端口 5555。|
|forward --list	|列出所有转发的套接字连接。|
|forward [–no-rebind] local remote|	将套接字连接从指定的本地端口转发到设备上指定的远程端口。您可以通过以下方式指定本地和远程端口tcp:port。要选择任何开放端口，请将 local 值设置为 tcp:0.|
|forward --remove local	|移除指定的转发套接字连接。|
|reverse --list	|列出设备的所有反向套接字连接。|
|reverse [–no-rebind] remote local|	反向连接套接字。||–no-rebind |选项表示如果指定的套接字已通过之前的 reverse 命令完成绑定，则反向连接失败。|
|reverse --remove remote|	从设备中移除指定的反向套接字连接。|
|reverse --remove-all|	从设备中移除所有反向套接字连接|

#### 3.2.3 安装应用

安装应用的基本命令格式是：

`adb install [-l] [-r] [-t] [-s] [-d] [-g] <apk-file>`

adb install 后面可以跟一些可选参数来控制安装 APK 的行为，可用参数及含义如下：

|参数|	含义|
|:--|:--|
|-l	|将应用安装到保护目录 /mnt/asec|
|-r	|允许覆盖安装|
|-t	|允许安装 AndroidManifest.xml 里 application 指定 android:testOnly=“true” 的应用|
|-s	|将应用安装到 sdcard|
|-d	|允许降级覆盖安装|
|-g	|授予所有运行时权限|

运行命令后可以看到输出内容，包含安装进度和状态，安装状态如下：

Success：代表安装成功。
Failure：代表安装失败。APK 安装失败的情况有很多，Failure状态之后有安装失败输出代码。

1. adb install 实际是分三步完成：

1) push apk 文件到 /data/local/tmp。
2) 调用 pm install 安装。
3) 删除 /data/local/tmp 下的对应 apk 文件。

#### 3.2.4 卸载应用

卸载应用的基本命令格式是：

`adb uninstall [-k] <package-name>`

<package-name> 表示应用的包名，-k 参数可选，表示卸载应用但保留数据和缓存目录。

#### 3.2.5 设置端口转发

您可以使用 forward 命令设置任意端口转发，将对特定主机端口上的请求转发到设备上的其他端口。以下示例介绍了如何设置主机端口 6100 到设备端口 7100 的转发：

    `adb forward tcp:6100 tcp:7100`

以下示例介绍了如何设置主机端口 6100 到 local:logd 的转发：

    `adb forward tcp:6100 local:logd`

#### 3.2.6 将文件复制到设备/从设备复制文件

您可以使用 pull 和 push 命令将文件复制到某个设备或从中复制文件。与 install 命令（仅将 APK 文件复制到特定位置）不同，pull 和 push 命令使您能够将任意目录和文件复制到设备中的任何位置。

1. 要从设备中复制某个文件或目录（及其子目录），请使用以下命令：

`adb pull remote local`

2. 要将某个文件或目录（及其子目录）复制到某个设备，请使用以下命令：

`adb push local remote`

将 local 和 remote 替换为开发计算机（本地）和设备（远程）上的目标文件/目录的路径。例如：

`adb push foo.txt /sdcard/foo.txt`

#### 3.2.7 停止 adb 服务器

在某些情况下，您可能需要终止 adb 服务器进程，然后重新启动以解决问题（例如，如果 adb 不响应命令）。

要停止 adb 服务器，请使用 `adb kill-server` 命令。然后，您可以通过发出任意其他 adb 命令重启服务器。

## 四、Logcat的使用

### 4.1 命令行语法

Logcat
要通过 adb shell 运行 Logcat，一般用法是：

`[adb] logcat [<option>] ... [<filter-spec>] ...`

您可以将 logcat 作为 adb 命令运行，也可以直接在模拟器或关联设备的 Shell 提示中运行它。要使用 adb 查看日志输出，请转到您的 SDK platform-tools/ 目录并执行：

`adb logcat`

要获取 logcat 在线帮助，请启动设备，然后执行：

`adb logcat --help`

您可以建立与设备的 shell 连接并执行：

    `$ adb shell`
    `# logcat`

借鉴的博客和文档：
https://blog.csdn.net/lb245557472/article/details/84068519
https://github.com/mzlogin/awesome-adb/blob/master/README.md
https://www.android-doc.com/tools/help/adb.html
https://developer.android.google.cn/studio/command-line/adb.html

