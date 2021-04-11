# Android 调试桥 (adb)

Android 调试桥 (adb) 是一种功能多样的命令行工具，可让您与设备进行通信。adb 命令可用于执行各种设备操作（例如安装和调试应用），并提供对 Unix shell（可用来在设备上运行各种命令）的访问权限。它是一种客户端-服务器程序，包括以下三个组件：

* 客户端：用于发送命令。客户端在开发计算机上运行。您可以通过发出 adb 命令从命令行终端调用客户端。
* 守护程序 (adbd)：用于在设备上运行命令。守护程序在每个设备上作为后台进程运行。
* 服务器：用于管理客户端与守护程序之间的通信。服务器在开发机器上作为后台进程运行。

adb 包含在 Android SDK 平台工具软件包中。您可以使用 SDK 管理器下载此软件包，该管理器会将其安装在 android_sdk/platform-tools/ 下。或者，如果您需要独立的 Android SDK 平台工具软件包，也可以点击此处进行下载。

如需了解如何连接设备以使用 ADB，包括如何使用 Connection Assistant 对常见问题进行排查，请参阅在硬件设备上运行应用。

## 1、adb 的工作原理

当您启动某个 adb 客户端时，该客户端会先检查是否有 adb 服务器进程正在运行。如果没有，它会启动服务器进程。服务器在启动后会与本地 TCP 端口 5037 绑定，并监听 adb 客户端发出的命令 - 所有 adb 客户端均通过端口 5037 与 adb 服务器通信。

然后，服务器会与所有正在运行的设备建立连接。它通过扫描 5555 到 5585 之间（该范围供前 16 个模拟器使用）的奇数号端口查找模拟器。服务器一旦发现 adb 守护程序 (adbd)，便会与相应的端口建立连接。请注意，每个模拟器都使用一对按顺序排列的端口 - 用于控制台连接的偶数号端口和用于 adb 连接的奇数号端口。例如：

模拟器 1，控制台：5554
模拟器 1，adb：5555
模拟器 2，控制台：5556
模拟器 2，adb：5557
依此类推

如上所示，在端口 5555 处与 adb 连接的模拟器与控制台监听端口为 5554 的模拟器是同一个。

服务器与所有设备均建立连接后，您便可以使用 adb 命令访问这些设备。由于服务器管理与设备的连接，并处理来自多个 adb 客户端的命令，因此您可以从任意客户端（或从某个脚本）控制任意设备。

## 2、在设备上启用 adb 调试

如要在通过 USB 连接的设备上使用 adb，您必须在设备的系统设置中启用 USB 调试（位于开发者选项下）。如需在通过 WLAN 连接的设备上使用 adb，请参阅通过 WLAN 连接到设备。

在搭载 Android 4.2 及更高版本的设备上，“开发者选项”屏幕默认情况下处于隐藏状态。如需将其显示出来，请依次转到设置 > 关于手机，然后点按版本号七次。返回上一屏幕，在底部可以找到开发者选项。

在某些设备上，“开发者选项”屏幕所在的位置或名称可能有所不同。

现在，您已经可以通过 USB 连接设备。您可以通过从 android_sdk/platform-tools/ 目录执行 adb devices 验证设备是否已连接。如果已连接，您将看到设备名称以“设备”形式列出。

注意：当您连接搭载 Android 4.2.2 或更高版本的设备时，系统会显示一个对话框，询问您是否接受允许通过此计算机进行调试的 RSA 密钥。这种安全机制可以保护用户设备，因为它可以确保只有在您能够解锁设备并确认对话框的情况下才能执行 USB 调试和其他 adb 命令。

要详细了解如何通过 USB 连接到设备，请参阅在硬件设备上运行应用。

## 3、通过 Wi-Fi 连接到设备（Android 11 及更高版本）

Android 11 及更高版本支持使用 Android 调试桥 (adb) 从工作站以无线方式部署和调试应用。例如，您可以将可调试应用部署到多台远程设备，而无需通过 USB 实际连接设备。这样就可以避免常见的 USB 连接问题，例如驱动程序安装方面的问题。

如需使用无线调试，您需要使用配对码将您的设备与工作站配对。您的工作站和设备必须连接到同一无线网络。如需连接到您的设备，请按以下步骤操作：

无线 adb 配对对话框
图 1. 无线 ADB 配对对话框。
在您的工作站上，更新到最新版本的 SDK 平台工具。
在设备上启用开发者选项。
启用无线调试选项。
在询问要允许在此网络上进行无线调试吗？的对话框中，点击允许。
选择使用配对码配对设备。记下设备上显示的配对码、IP 地址和端口号（参见图片）。
在工作站上，打开一个终端并导航到 android_sdk/platform-tools。
运行 adb pair ipaddr:port。 使用第 5 步中的 IP 地址和端口号。
当系统提示时，输入您在第 5 步中获得的配对码。系统会显示一条消息，表明您的设备已成功配对。

    none
    Enter pairing code: 482924
    Successfully paired to 192.168.1.130:37099 [guid=adb-235XY]
    
（仅适用于 Linux 或 Microsoft Windows）运行 adb connect ipaddr:port。使用无线调试下的 IP 地址和端口。

无线 adb IP 地址和端口号
图 2. 无线 adb IP 地址和端口号。
通过 WLAN 连接到设备（Android 10 及更低版本）
一般情况下，adb 通过 USB 与设备进行通信，但您也可以在以下情况下通过 WLAN 使用 adb：

如需连接到搭载 Android 11（及更高版本）的设备，请参阅在硬件设备上运行应用的“WLAN”部分。
如要连接到搭载早期 Android 版本的设备，您必须通过 USB 执行一些初始步骤。下文对这些步骤做了说明。
如果您开发的是 Wear OS 应用，请参阅调试 Wear OS 应用指南，其中提供了有关如何通过 WLAN 和蓝牙使用 adb 的特别说明。
将 Android 设备和 adb 主机连接到这两者都可以访问的同一 WLAN 网络。请注意，并非所有接入点都适用；您可能需要使用防火墙已正确配置为支持 adb 的接入点。
如果您要连接到 Wear OS 设备，请关闭手机上与该设备配对的蓝牙。
使用 USB 线将设备连接到主机。
设置目标设备以监听端口 5555 上的 TCP/IP 连接。

`adb tcpip 5555`
拔掉连接目标设备的 USB 线。
找到 Android 设备的 IP 地址。例如，对于 Nexus 设备，您可以在设置 > 关于平板电脑（或关于手机）> 状态 > IP 地址下找到 IP 地址。或者，对于 Wear OS 设备，您可以在设置 > WLAN 设置 > 高级 > IP 地址下找到 IP 地址。
通过 IP 地址连接到设备。

`adb connect device_ip_address`
确认主机已连接到目标设备：
```
$ adb devices
List of devices attached
device_ip_address:5555 device
```

现在，您可以开始操作了！

如果 adb 连接断开：

确保主机仍与 Android 设备连接到同一个 WLAN 网络。
通过再次执行 adb connect 步骤重新连接。
如果上述操作未解决问题，重置 adb 主机：

`adb kill-server`
然后，从头开始操作。

查询设备
在发出 adb 命令之前，了解哪些设备实例已连接到 adb 服务器会很有帮助。您可以使用 devices 命令生成已连接设备的列表。


  `adb devices -l`
  
作为回应，adb 会针对每个设备输出以下状态信息：

序列号：由 adb 创建的字符串，用于通过端口号唯一标识设备。 下面是一个序列号示例：emulator-5554
状态：设备的连接状态可以是以下几项之一：
offline：设备未连接到 adb 或没有响应。
device：设备现已连接到 adb 服务器。请注意，此状态并不表示 Android 系统已完全启动并可正常运行，因为在设备连接到 adb 时系统仍在启动。不过，在启动后，这将是设备的正常运行状态。
no device：未连接任何设备。
说明：如果您包含 -l 选项，devices 命令会告知您设备是什么。当您连接了多个设备时，此信息很有用，可帮助您将它们区分开来。
以下示例展示了 devices 命令及其输出。有三个设备正在运行。列表中的前两行表示模拟器，第三行表示连接到计算机的硬件设备。

```
$ adb devices
List of devices attached
emulator-5556 device product:sdk_google_phone_x86_64 model:Android_SDK_built_for_x86_64 device:generic_x86_64
emulator-5554 device product:sdk_google_phone_x86 model:Android_SDK_built_for_x86 device:generic_x86
0a388e93      device usb:1-1 product:razor model:Nexus_7 device:flo
```

模拟器未列出
adb devices 命令的极端命令序列会导致正在运行的模拟器不显示在 adb devices 输出中（即使在您的桌面上可以看到该模拟器）。当满足以下所有条件时，就会发生这种情况：

adb 服务器未在运行，
您在使用 emulator 命令时，将 -port 或 -ports 选项的端口值设为 5554 到 5584 之间的奇数，
您选择的奇数号端口处于空闲状态，因此可以与指定端口号的端口建立连接，或者该端口处于忙碌状态，模拟器切换到了符合第 2 条中要求的另一个端口，以及
启动模拟器后才启动 adb 服务器。
避免出现这种情况的一种方法是让模拟器自行选择端口，并且每次运行的模拟器数量不要超过 16 个。另一种方法是始终先启动 adb 服务器，然后再使用 emulator 命令，如下例所示。

示例 1：在下面的命令序列中，adb devices 命令启动了 adb 服务器，但是设备列表未显示。

停止 adb 服务器，然后按照所示顺序输入以下命令。对于 avd 名称，请提供系统中有效的 avd 名称。如需获取 avd 名称列表，请输入 emulator -list-avds。 emulator 命令位于 android_sdk/tools 目录下。

```
$ adb kill-server
$ emulator -avd Nexus_6_API_25 -port 5555
$ adb devices

List of devices attached
* daemon not running. starting it now on port 5037 *
* daemon started successfully *
```

示例 2：在下面的命令序列中，adb devices 显示了设备列表，因为先启动了 adb 服务器。

如果想在 adb devices 输出中看到模拟器，请停止 adb 服务器，然后在使用 emulator 命令之后、使用 adb devices 命令之前，重新启动该服务器，如下所示：

```
$ adb kill-server
$ emulator -avd Nexus_6_API_25 -port 5557
$ adb start-server
$ adb devices

List of devices attached
emulator-5557 device
```

如需详细了解模拟器命令行选项，请参阅使用命令行参数。

将命令发送至特定设备
如果有多个设备在运行，您在发出 adb 命令时必须指定目标设备。为此，请使用 devices 命令获取目标设备的序列号。获得序列号后，请结合使用 -s 选项与 adb 命令来指定序列号。如果您要发出很多 adb 命令，可以将 $ANDROID_SERIAL 环境变量设为包含序列号。如果您同时使用 -s 和 $ANDROID_SERIAL，-s 会替换 $ANDROID_SERIAL。

在以下示例中，先获得了已连接设备的列表，然后使用其中一个设备的序列号在该设备上安装了 helloWorld.apk。

```
$ adb devices
List of devices attached
emulator-5554 device
emulator-5555 device

$ adb -s emulator-5555 install helloWorld.apk
```

注意：如果您在多个设备可用时发出命令但未指定目标设备，adb 会生成错误。

如果有多个可用设备，但只有一个是模拟器，请使用 -e 选项将命令发送至该模拟器。同样，如果有多个设备，但只连接了一个硬件设备，请使用 -d 选项将命令发送至该硬件设备。

安装应用
您可以使用 adb 的 install 命令在模拟器或连接的设备上安装 APK：


`adb install path_to_apk`
安装测试 APK 时，必须在 install 命令中使用 -t 选项。如需了解详情，请参阅 -t。

要详细了解如何创建可安装在模拟器/设备实例上的 APK 文件，请参阅构建和运行应用。

请注意，如果您使用的是 Android Studio，则无需直接使用 adb 在模拟器/设备上安装您的应用。Android Studio 会为您执行应用的打包和安装操作。

设置端口转发
您可以使用 forward 命令设置任意端口转发，将特定主机端口上的请求转发到设备上的其他端口。以下示例设置了主机端口 6100 到设备端口 7100 的转发：


`adb forward tcp:6100 tcp:7100`
以下示例设置了主机端口 6100 到 local:logd 的转发：


`adb forward tcp:6100 local:logd`
将文件复制到设备/从设备复制文件
您可以使用 pull 和 push 命令将文件复制到设备或从设备复制文件。与 install 命令（仅将 APK 文件复制到特定位置）不同，使用 pull 和 push 命令可将任意目录和文件复制到设备中的任何位置。

如需从设备中复制某个文件或目录（及其子目录），请使用以下命令：


`adb pull remote local`
如需将某个文件或目录（及其子目录）复制到设备，请使用以下命令：


`adb push local remote`
将 local 和 remote 替换为开发机器（本地）和设备（远程）上的目标文件/目录的路径。例如：


`adb push foo.txt /sdcard/foo.txt`
停止 adb 服务器
在某些情况下，您可能需要终止 adb 服务器进程，然后重启以解决问题（例如，如果 adb 不响应命令）。

如需停止 adb 服务器，请使用 adb kill-server 命令。然后，您可以通过发出其他任何 adb 命令来重启服务器。

发出 adb 命令
您可以从开发机器上的命令行发出 adb 命令，也可以通过脚本发出。用法如下：


`adb [-d | -e | -s serial_number] command`
如果只有一个模拟器在运行或者只连接了一个设备，系统会默认将 adb 命令发送至该设备。如果有多个模拟器正在运行并且/或者连接了多个设备，您需要使用 -d、-e 或 -s 选项指定应向其发送命令的目标设备。

您可以使用以下命令来查看所有支持的 adb 命令的详细列表：


`adb --help`
发出 shell 命令
您可以使用 shell 命令通过 adb 发出设备命令，也可以启动交互式 shell。如需发出单个命令，请使用 shell 命令，如下所示：


`adb [-d |-e | -s serial_number] shell shell_command`
要在设备上启动交互式 shell，请使用 shell 命令，如下所示：


`adb [-d | -e | -s serial_number] shell`
要退出交互式 shell，请按 Ctrl + D 键或输入 exit。

注意：在 Android 平台工具 23 及更高版本中，adb 处理参数的方式与 ssh(1) 命令相同。这项变更解决了很多命令注入方面的问题，还使安全执行包含 shell 元字符的命令（如 adb install Let\'sGo.apk）成为可能。不过，这项变更还意味着，对包含 shell 元字符的所有命令的解释也发生了变化。例如，adb shell setprop foo 'a b' 命令现在会返回错误，因为单引号 (') 会被本地 shell 消去，设备看到的是 adb shell setprop foo a b。如需使该命令正常运行，请引用两次，一次用于本地 shell，另一次用于远程 shell，与处理 ssh(1) 的方法相同。例如，adb shell setprop foo "'a b'"。

Android 提供了大多数常见的 Unix 命令行工具。如需查看可用工具的列表，请使用以下命令：


`adb shell ls /system/bin`
对于大多数命令，都可通过 --help 参数获得命令帮助。许多 shell 命令都由 toybox 提供。对于所有 toybox 命令，都可通过 toybox --help 可获得命令的常规帮助。

另请参阅 Logcat 命令行工具，该工具对监控系统日志很有用。

调用 Activity 管理器 (am)
在 adb shell 中，您可以使用 Activity 管理器 (am) 工具发出命令以执行各种系统操作，如启动 Activity、强行停止进程、广播 intent、修改设备屏幕属性，等等。在 shell 中，相应的语法为：


am command
您也可以直接从 adb 发出 Activity 管理器命令，无需进入远程 shell。例如：


adb shell am start -a android.intent.action.VIEW
表 2. 可用的 Activity 管理器命令

命令	说明
start [options] intent	启动由 intent 指定的 Activity。
请参阅 intent 参数的规范。

具体选项包括：

-D：启用调试功能。
-W：等待启动完成。
--start-profiler file：启动性能剖析器并将结果发送至 file。
-P file：类似于 --start-profiler，但当应用进入空闲状态时剖析停止。
-R count：重复启动 Activity count 次。在每次重复前，将完成顶层 Activity。
-S：在启动 Activity 前，强行停止目标应用。
--opengl-trace：启用 OpenGL 函数的跟踪。
--user user_id | current：指定要作为哪个用户运行；如果未指定，则作为当前用户运行。
startservice [options] intent	启动由 intent 指定的 Service。
请参阅 intent 参数的规范。

具体选项包括：

--user user_id | current：指定要作为哪个用户运行；如果未指定，则作为当前用户运行。
force-stop package	强行停止与 package（应用的软件包名称）关联的所有进程。
kill [options] package	终止与 package（应用的软件包名称）关联的所有进程。此命令仅终止可安全终止且不会影响用户体验的进程。
具体选项包括：

--user user_id | all | current：指定要终止哪个用户的进程；如果未指定，则终止所有用户的进程。
kill-all	终止所有后台进程。
broadcast [options] intent	发出广播 intent。
请参阅 intent 参数的规范。

具体选项包括：

[--user user_id | all | current]：指定要发送给哪个用户；如果未指定，则发送给所有用户。
instrument [options] component	使用 Instrumentation 实例启动监控。通常情况下，目标 component 采用 test_package/runner_class 格式。
具体选项包括：

-r：输出原始结果（否则，对 report_key_streamresult 进行解码）。与 [-e perf true] 结合使用可生成性能测量的原始输出。
-e name value：将参数 name 设为 value。 对于测试运行程序，通用格式为 -e testrunner_flag value[,value...]。
-p file：将剖析数据写入 file。
-w：等待插桩完成后再返回。测试运行程序需要使用此选项。
--no-window-animation：运行时关闭窗口动画。
--user user_id | current：指定以哪个用户身份运行插桩；如果未指定，则以当前用户身份运行。
profile start process file	启动 process 的性能剖析器，将结果写入 file。
profile stop process	停止 process 的性能剖析器。
dumpheap [options] process file	转储 process 的堆，写入 file。
具体选项包括：

--user [user_id | current]：提供进程名称时，指定要转储的进程的用户；如果未指定，则使用当前用户。
-n：转储原生堆，而非托管堆。
set-debug-app [options] package	设置要调试的应用 package。
具体选项包括：

-w：应用启动时等待调试程序。
--persistent：保留此值。
clear-debug-app	清除之前使用 set-debug-app 设置的待调试软件包。
monitor [options]	开始监控崩溃或 ANR。
具体选项包括：

--gdb：在崩溃/ANR 时，在给定的端口上启动 gdbserv。
screen-compat {on | off} package	控制 package 的屏幕兼容性模式。
display-size [reset | widthxheight]	替换设备显示尺寸。此命令支持使用大屏设备模仿小屏幕分辨率（反之亦然），对于在不同尺寸的屏幕上测试应用非常有用。
示例：
am display-size 1280x800

display-density dpi	替换设备显示密度。此命令支持使用低密度屏幕在高密度屏幕环境上进行测试（反之亦然），对于在不同密度的屏幕上测试应用非常有用。
示例：
am display-density 480

to-uri intent	以 URI 的形式输出给定的 intent 规范。
请参阅 intent 参数的规范。

to-intent-uri intent	以 intent: URI 的形式输出给定的 intent 规范。
请参阅 intent 参数的规范。

intent 参数的规范
对于采用 intent 参数的 Activity 管理器命令，您可以使用以下选项指定 intent：

全部显示

调用软件包管理器 (pm)
在 adb shell 中，您可以使用软件包管理器 (pm) 工具发出命令，以对设备上安装的应用软件包执行操作和查询。在 shell 中，相应的语法为：


pm command
您也可以直接从 adb 发出软件包管理器命令，无需进入远程 shell。例如：


adb shell pm uninstall com.example.MyApp
表 3. 可用的软件包管理器命令。

命令	说明
list packages [options] filter	输出所有软件包，或者，仅输出软件包名称包含 filter 中的文本的软件包。
具体选项：

-f：查看它们的关联文件。
-d：进行过滤以仅显示已停用的软件包。
-e：进行过滤以仅显示已启用的软件包。
-s：进行过滤以仅显示系统软件包。
-3：进行过滤以仅显示第三方软件包。
-i：查看软件包的安装程序。
-u：也包括已卸载的软件包。
--user user_id：要查询的用户空间。
list permission-groups	输出所有已知的权限组。
list permissions [options] group	输出所有已知的权限，或者，仅输出 group 中的权限。
具体选项：

-g：按组进行整理。
-f：输出所有信息。
-s：简短摘要。
-d：仅列出危险权限。
-u：仅列出用户将看到的权限。
list instrumentation [options]	列出所有测试软件包。
具体选项：

-f：列出测试软件包的 APK 文件。
target_package：仅列出此应用的测试软件包。
list features	输出系统的所有功能。
list libraries	输出当前设备支持的所有库。
list users	输出系统中的所有用户。
path package	输出给定 package 的 APK 的路径。
install [options] path	将软件包（通过 path 指定）安装到系统。
具体选项：

-r：重新安装现有应用，并保留其数据。
-t：允许安装测试 APK。仅当您运行或调试了应用或者使用了 Android Studio 的 Build > Build APK 命令时，Gradle 才会生成测试 APK。如果是使用开发者预览版 SDK（如果 targetSdkVersion 是字母，而非数字）构建的 APK，那么安装测试 APK 时必须在 install 命令中包含 -t 选项。
-i installer_package_name：指定安装程序软件包名称。
--install-location location：使用以下某个值设置安装位置：
0：使用默认安装位置。
1：在内部设备存储上安装。
2：在外部介质上安装。
-f：在内部系统内存上安装软件包。
-d：允许版本代码降级。
-g：授予应用清单中列出的所有权限。
--fastdeploy：通过仅更新已更改的 APK 部分来快速更新安装的软件包。
--incremental：仅安装 APK 中启动应用所需的部分，同时在后台流式传输剩余数据。如要使用此功能，您必须为 APK 签名，创建一个 APK 签名方案 v4 文件，并将此文件放在 APK 所在的目录中。只有部分设备支持此功能。此选项会强制 adb 使用该功能，如果该功能不受支持，则会失败（并提供有关失败原因的详细信息）。附加 --wait 选项，可等到 APK 完全安装完毕后再授予对 APK 的访问权限。
--no-incremental 可阻止 adb 使用此功能。

uninstall [options] package	从系统中移除软件包。
具体选项：

-k：移除软件包后保留数据和缓存目录。
clear package	删除与软件包关联的所有数据。
enable package_or_component	启用给定的软件包或组件（写为“package/class”）。
disable package_or_component	停用给定的软件包或组件（写为“package/class”）。
disable-user [options] package_or_component	
具体选项：

--user user_id：要停用的用户。
grant package_name permission	向应用授予权限。在搭载 Android 6.0（API 级别 23）及更高版本的设备上，该权限可以是应用清单中声明的任何权限。在搭载 Android 5.1（API 级别 22）及更低版本的设备上，该权限必须是应用定义的可选权限。
revoke package_name permission	从应用撤消权限。在搭载 Android 6.0（API 级别 23）及更高版本的设备上，该权限可以是应用清单中声明的任何权限。在搭载 Android 5.1（API 级别 22）及更低版本的设备上，该权限必须是应用定义的可选权限。
set-install-location location	更改默认安装位置。位置值如下：
0：自动：让系统决定最合适的位置。
1：内部：在内部设备存储上安装。
2：外部：在外部介质上安装。
注意：此命令仅用于调试目的；使用此命令可能会导致应用中断和其他意外行为。

get-install-location	返回当前安装位置。返回值如下：
0 [auto]：让系统决定最合适的位置
1 [internal]：在内部设备存储上安装
2 [external]：在外部介质上安装
set-permission-enforced permission [true | false]	指定是否应强制执行指定权限。
trim-caches desired_free_space	减少缓存文件以达到给定的可用空间。
create-user user_name	创建具有给定 user_name 的新用户，从而输出该用户的新用户标识符。
remove-user user_id	移除具有给定 user_id 的用户，从而删除与该用户关联的所有数据。
get-max-users	输出设备支持的最大用户数。
调用设备政策管理器 (dpm)
为便于您开发和测试设备管理（或其他企业）应用，您可以向设备政策管理器 (dpm) 工具发出命令。使用该工具可控制活动管理应用，或更改设备上的政策状态数据。在 shell 中，语法如下：


dpm command
您也可以直接从 adb 发出设备政策管理器命令，无需进入远程 shell：


adb shell dpm command
表 4. 可用的设备政策管理器命令

命令	说明
set-active-admin [options] component	将 component 设为活动管理。
具体选项包括：

--user user_id：指定目标用户。您也可以传递 --user current 以选择当前用户。
set-profile-owner [options] component	将 component 设为活动管理，并将其软件包设为现有用户的个人资料所有者。
具体选项包括：

--user user_id：指定目标用户。您也可以传递 --user current 以选择当前用户。
--name name：指定简单易懂的组织名称。
set-device-owner [options] component	将 component 设为活动管理，并将其软件包设为设备所有者。
具体选项包括：

--user user_id：指定目标用户。您也可以传递 --user current 以选择当前用户。
--name name：指定简单易懂的组织名称。
remove-active-admin [options] component	停用活动管理。应用必须在清单中声明 android:testOnly。此命令还会移除设备所有者和个人资料所有者。
具体选项包括：

--user user_id：指定目标用户。您也可以传递 --user current 以选择当前用户。
clear-freeze-period-record	清除设备之前设置的系统 OTA 更新冻结期记录。在开发管理冻结期的应用时，这有助于避免设备存在调度方面的限制。请参阅管理系统更新。
在搭载 Android 9.0（API 级别 28）及更高版本的设备上受支持。

force-network-logs	强制系统让任何现有网络日志随时可供 DPC 检索。如果有可用的连接或 DNS 日志，DPC 会收到 onNetworkLogsAvailable() 回调。请参阅网络活动日志。
此命令有调用频率限制。在搭载 Android 9.0（API 级别 28）及更高版本的设备上受支持。

force-security-logs	强制系统向 DPC 提供任何现有安全日志。如果有可用的日志，DPC 会收到 onSecurityLogsAvailable() 回调。请参阅记录企业设备活动。
此命令有调用频率限制。在搭载 Android 9.0（API 级别 28）及更高版本的设备上受支持。

截取屏幕截图
screencap 命令是一个用于对设备显示屏截取屏幕截图的 shell 实用程序。在 shell 中，语法如下：


screencap filename
如需从命令行使用 screencap，请输入以下命令：


adb shell screencap /sdcard/screen.png
以下屏幕截图会话示例展示了如何使用 adb shell 截取屏幕截图，以及如何使用 pull 命令从设备下载屏幕截图文件：


$ adb shell
shell@ $ screencap /sdcard/screen.png
shell@ $ exit
$ adb pull /sdcard/screen.png
录制视频
screenrecord 命令是一个用于录制设备（搭载 Android 4.4（API 级别 19）及更高版本）显示屏的 shell 实用程序。该实用程序将屏幕 Activity 录制为 MPEG-4 文件。您可以使用此文件创建宣传视频或培训视频，或将其用于调试或测试。

在 shell 中，使用以下语法：


screenrecord [options] filename
如需从命令行使用 screenrecord，请输入以下命令：


adb shell screenrecord /sdcard/demo.mp4
按 Ctrl + C 键（在 Mac 上，按 Command + C 键）可停止屏幕录制；如果不手动停止，到三分钟或 --time-limit 设置的时间限制时，录制将会自动停止。

如需开始录制设备屏幕，请运行 screenrecord 命令以录制视频。然后，运行 pull 命令以将视频从设备下载到主机。下面是一个录制会话示例：


$ adb shell
shell@ $ screenrecord --verbose /sdcard/demo.mp4
(press Control + C to stop)
shell@ $ exit
$ adb pull /sdcard/demo.mp4
screenrecord 实用程序能以您要求的任何支持的分辨率和比特率进行录制，同时保持设备显示屏的宽高比。默认情况下，该实用程序以本机显示分辨率和屏幕方向进行录制，时长不超过三分钟。

screenrecord 实用程序的局限性：

音频不与视频文件一起录制。
无法在搭载 Wear OS 的设备上录制视频。
某些设备可能无法以它们的本机显示分辨率进行录制。如果在录制屏幕时出现问题，请尝试使用较低的屏幕分辨率。
不支持在录制时旋转屏幕。如果在录制期间屏幕发生了旋转，则部分屏幕内容在录制时将被切断。
表 5. screenrecord 选项

选项	说明
--help	显示命令语法和选项
--size widthxheight	设置视频大小：1280x720。默认值为设备的本机显示屏分辨率（如果支持）；如果不支持，则为 1280x720。为获得最佳效果，请使用设备的 Advanced Video Coding (AVC) 编码器支持的大小。
--bit-rate rate	设置视频的视频比特率（以 MB/秒为单位）。默认值为 4Mbps。您可以增加比特率以提升视频品质，但这样做会导致视频文件变大。下面的示例将录制比特率设为 6Mbps：


screenrecord --bit-rate 6000000 /sdcard/demo.mp4
--time-limit time	设置最大录制时长（以秒为单位）。默认值和最大值均为 180（3 分钟）。
--rotate	将输出旋转 90 度。此功能处于实验阶段。
--verbose	在命令行屏幕显示日志信息。如果您不设置此选项，则该实用程序在运行时不会显示任何信息。
读取应用的 ART 配置文件
从 Android 7.0（API 级别 24）开始，Android Runtime (ART) 会收集已安装应用的执行配置文件，这些配置文件用于优化应用性能。您可能需要检查收集的配置文件，以了解在应用启动期间，系统频繁执行了哪些方法和使用了哪些类。

要生成文本格式的配置文件信息，请使用以下命令：


adb shell cmd package dump-profiles package
要检索生成的文件，请使用：


adb pull /data/misc/profman/package.txt
重置测试设备
如果您在多个测试设备上测试应用，则在两次测试之间重置设备可能很有用，例如，可以移除用户数据并重置测试环境。您可以使用 testharness adb shell 命令对搭载 Android 10（API 级别 29）或更高版本的测试设备执行恢复出厂设置，如下所示。


adb shell cmd testharness enable
使用 testharness 恢复设备时，设备会自动将允许通过当前工作站调试设备的 RSA 密钥备份在一个持久性位置。也就是说，在重置设备后，工作站可以继续调试设备并向设备发出 adb 命令，而无需手动注册新密钥。

此外，为了帮助您更轻松且更安全地继续测试您的应用，使用 testharness 恢复设备还会更改以下设备设置：

设备会设置某些系统设置，以便不会出现初始设备设置向导。也就是说，设备会进入一种状态，供您快速安装、调试和测试您的应用。
设置：
停用锁定屏幕
停用紧急提醒
停用帐户自动同步
停用自动系统更新
其他：
停用预装的安全应用
如果您的应用需要检测并适应 testharness 命令的默认设置，您可以使用 ActivityManager.isRunningInUserTestHarness()。

sqlite
sqlite3 可启动用于检查 sqlite 数据库的 sqlite 命令行程序。它包含用于输出表格内容的 .dump 以及用于输出现有表格的 SQL CREATE 语句的 .schema 等命令。您也可以从命令行执行 SQLite 命令，如下所示。


$ adb -s emulator-5554 shell
$ sqlite3 /data/data/com.example.app/databases/rssitems.db
SQLite version 3.3.12
Enter ".help" for instructions