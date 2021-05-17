# Golang语言

## 1、背景

### 1.1 起源及演进

最近十年来，C/C++在计算领域没有很好得到发展，并没有新的系统编程语言出现。对开发程度和系统效率在很多情况下不能兼得。要么执行效率高，但低效的开发和编译，如C++；要么执行低效，但拥有有效的编译，如.NET、Java；所以需要一种拥有较高效的执行速度、编译速度和开发速度的编程语言，Go就横空出世了。

|语言|优点|缺点|
|:--|:---|:---|
|c++|c++执行效率高于golang|c++开发和编译效率低于golang|
|golang|golang语言开发和编译效率高于c++|golang执行效率没有c++，但是执行效率也比较高|

go是一个Google公司推出的一个开源项目（系统开发语言），它是基于编译、垃圾收集和并发的编程语言。并将其开源并在BSD许可证下发行。

Go最初的设计由Robert Griesemer，Rob Pike 和Ken Thompson在2007年9月开始的，官方的发布是在2009年11月。2010年5月由Rob Pike公开的将其运用于google内部的一个后台系统。目前在google App Engine也支持go语言（目前仅支持三种：Java、Python和Go）

Go可以运行在Linux, Mac OS X, FreeBSD, OpenBSD, Plan 9 和 Microsof windows系统上，同时也支持多种处理器架构：I386， AMD64和ARM

（注：官方网站： http://golang.org）

Robert Griesemer：曾协助制作Java的HotSpot编译器和Chrome浏览器的JavaScript引擎V8
Rob Pike：曾是贝尔实验室的Unix团队和Plan9操作系统计划成员，与Thompson工同创造了UTF-8字符编码
Ken Thompson:是C语言和Unix的创造者。（1983年图灵奖和1988国家技术奖）
他们对系统编程语言、操作系统和并发有很深的理解。

### 1.2 主要特点

- Go被设计为21世纪的C语言，它属于C语言家族，比如：C/C++、Java和C#，同时它吸收了很多现在编程语言的优点。

    > 1）对Go的并发机制是源于CSP（Communication Sequential Processes），这同样的机制也被于Erlang。
    > 2）对C、C++相比，其语法得到了很大程序上的简化，使代码更简明、清楚，同时拥有动态语言的一些特点

- 基于BSD完全开源，所以能免费的被任何人用于适合商业目的。
- 语言层面对并发的支持（goroutine：独立于OS的线程，所以多个goroutine可以运行在一个OS的线程里，也可以分布到多个OS线程里。goroutine是从OS线程上抽象出来的一个轻量级的基于CSP的协程）
    1）在语言层面加入对并发的支持，而不是以库的形式提供。
    2）更高层次的并发抽象，而不是直接暴露OS的并发机制.
    3）多个goroutine间是并行的。
    4）底层混合使用非阻塞IO和线程。

- 主要目的
    > 1）融合效率、速度和安全的强类型的静态编译语言，同时能够容易的进行编程，让编程变得更有乐趣。
    > 2）较少的关键字和简洁的语法
    > 3）类型安全和内存安全：在指针类型，但不允许对指针进行操作。
    > 4）支持网络通信、并发控制、并行计算和分布式计算。
    > 5）在语言层面实现对多处理器（或多核）进行编程

- 内嵌运行时反射机制。
- 可以集成C语言实现的库
- 它不是传统意义上的面向对象语言（没有类的概念），但它有接口（interface），由此实现多态特性。
- 函数（Function）是它的基本构成单元（也可以叫着面向函数的程序设计语言）
- 是一种静态类型和安全的语言，将其编译、连接成本地代码（拥有高效的执行效率）
- 支持交叉编译，并采用编译的编码：UTF-8

### 1.3 应用领域

它最初的构想是作为一个系统编程语言，但目前也被用于像Web Server，存储架构等这类分布式、高并发系统中。当然也可以用于一般的文字处理和作为脚本程序。

Go的编译器作为Native Client被内嵌到Chrome浏览器中，可以被Web应用程序用来执行本地代码；同时Go也可以运行在Intel和ARM的处理器上。

目前已被Google集成到Google APP Engine中，在基于Google App Engine基础设施的Web应用中也得到了很好的应用。目前GAE中仅支持三种应用程序开发语言：Java、Python和Go。（注：GAE的链接）

但不适合应用到对实时性要求很高的系统中，因为Go的内存模型是基于垃圾回收机制和原子内存分配。

### 1.4 目前缺少的一些特性

目前Go对OO中涉及到的一些特点还没有很好的支持，但可能会在以后进一步完善。
没有函数和操作符的重载
不支持隐式类型转换， 避免产生Bug和迷惑。
不支持类和继承。
不支持动态代码加载
不支持动态库
不支持泛型


### 1.5 Go语言的主要特点

1. 强调简单、易学：Go语言的关键字更少，同时砍掉很多不需要的功能，例如c++的构造函数、析构函数、类的继承、指针的++/--。
2. 内存管理：Go语言和java/pathon都需要程序员管理内存，而c++需要程序员管理内存(new/delete)，目前c++的智能指针可以管理内存。
3. 快速编译：Go语言开发和编译效率高于c++，而执行效率低于c++，且执行效率也比较高。
4. 并发支持：Go支持并发，通过go关键字创建N个goroutine实现并发，而c++需要第三方框架实现并发功能，例如线程池和异步调用（一般由第三方RPC框架提供）。
5. 静态类型：将运行时、依赖库直接打包到可执行文件内部，简化了部署和发布操作，无须事先安装运行环境和下载诸多第三方库。
6. 部署简单（go install）
7. 自身就是文档（通过godoc将代码中的注释信息构造成文档）
8. 开源免费（BSD licensed）
9. 语法规范：Go语言实现规范语法统一，例如Go语言大写字母表示public公开的，小写字母表示私有的。
10. 标准工具：Go语言提供了编译、测试、调试、性能分析等系列标准工具，编译效率高，支持原生单元测试、调试、代码检查、性能分析，而c++需要第三方开源工具才能实现。
11. defer: 延迟调用功能，例如defer file.close()。
12. 函数支持多个返回值：Go语言的函数支持多个返回值，而c++只支持单个返回值，若需要支持多个返回值，可以存放在结构体。
13. 接口设计：Go语言专注结构体的方法的实现，而c++是类的继承。
14. 反射：Go通过反射实现对象的序列化（serialization，有时候也叫Marshal & Unmarshal），例如Json、XML。
15. 错误处理：Go使用panic/defer/recover，而c++使用try/catch/finnaly。
16. 匿名函数和闭包。
17. 语言的交互性：这里的交互性主要是和C的交互性，之所以这样是因为Go语言的开发者是最初贝尔实验室创建Unix系统以及C语言的一般人。

## 2、安装与使用

### 2.1 安装golang语言

golang软件官方（需要翻墙）：https://golang.org/dl/
golaang学习中文网镜像下载：https://studygolang.com/dl/

#### 2.1.1 windows10安装

第一步：安装go.1.13.windows-amd64.msi文件，点击安装。
第二步：配置golang环境
安装目录：GOROOT: C:\GO
存储目录：GOPATH: D:\wen\gopath\
配置path: %GOROOT%\bin;  %GOPATH%\bin
第三步：查看配置：go env
第四步：查看版本：go version
第五步：创建存储目录的文件夹
在GOPATH目录创建三个子目录：
src: 存储源代码（比如: .go、.c）
pkg: 编译生成的文件（比如： .a）
bin: 编译后生成的执行文件。

第六步：创建子目录

例如下载 https://github.com/apache/arrow.git 时，则需要创建子目录：D:\wen\gopath\src\github.com\apache\

第七步：编译命令

方法一：直接运行
`go run hello.go`
方法二：生成执行文件，再运行执行文件
`go build hello.go`
`./hello.exe`

#### 2.1.2 linux安装

查看系统版本
`cat /etc/issue`

第一步：获取安装包

不建议使用命令安装，原因：版本比较老

建议手动安装
`sudo tar -xzvf go1.14.4.linux-amd64.tar.gz -C /usr/local`

第二步： 配置环境变量

`$vim /etc/profile`
在文件末尾添加如下内容：
`export PATH=$PATH:/usr/local/go/bin`

让profile立即生效
`$source /etc/profile`

第三步：查看环境变量

`go env`

第四步：查看版本

`go version`

第五步：创建工作目录

`mkdir /home/wen/gopath/src /home/wen/gopath/pkg /home/wen/gopath/bin`

src: 存储源代码（比如: .go、.c）
pkg: 编译生成的文件（比如： .a）
bin: 编译后生成的执行文件。


第六步：创建子目录

例如下载 https://github.com/apache/arrow.git 时，则需要创建子目录：D: /home/wen/gopath/src/github.com/apache/

第七步：编译命令

方法一：直接运行
`go run hello.go`
方法二：生成执行文件，再运行执行文件
`go build hello.go`
`./hello.exe`

### 2.2 windows安装go软件

#### 2.2.1 Goland安装

第一步： 下载goland（30天免费）

https://www.jetbrain.com/zh-cn/go/promo/

第二步：安装
第三步：配置goland软件GOROOT和GOPATH
`File->Setting->Go->GOROOT和GOPATH`
第四步：测试软件是否安装成功

快捷键：shift+ctrl+F10运行hello.go

输出“hello world!”

#### 2.2.2 Goland安装

第一步：下载与安装VSCode

官网： https://code.visualstudio.com/Download
选择64bit

第二步：安装中文插件

在左侧菜单栏最后一项管理扩展，输入chinese,点击安装。安装完毕后重启，VSCode就会显示中文。

第三步：安装golang插件

在左侧菜单栏最后一项管理扩展，输入Go,点击安装。安装完毕后重启，VSCode就会显示go。

第四步：安装go语言开发工具包

作用：代码提示、代码自动补全等功能
windows平台按下ctrl+shif+p，则在输入框输入go:install,选择go:install/update tools

直接安装是无法安装成功的，原因是此软件为国外的，需要翻墙。

解决方法：查找别人提供go-tools工具

第五步：调试debug

在gopath/src目录下添加launch.json和setting.json文件

- settings.json 文件内容如下：主要是goroot和gopath

```json
{
    "files.autoSave": "onFocusChange",
    "go.buildOnSave": true,
    "go.lintOnSave": true,
    "go.vetOnSave": true,
    "go.buildTags": "",
    "go.buildFlags": [],
    "go.lintFlags": [],
    "go.vetFlags": [],
    "go.coverOnSave": false,
    "go.useCodeSnippetsOnFunctionSuggest": false,
    "go.formatOnSave": true,
    "go.formatTool": "goreturns",
    "go.goroot": "D:\\go\\Golang",    
    "go.gopath": "D:\\goath", 
    "go.gocodeAutoBuild": true
}
```

launch.json文件内容如下：主要是host和port

```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "igoodful",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            "remotePath": "",
            "port": 2345,
            "host": "127.0.0.1",
            "program": "${workspaceRoot}\\helloworld",
            "env": {},
            "args": []
        }
    ]
}
```

设置 launch.json 配置文件
`ctrl+shift+p` 输入 `Debug: Open launch.json` 打开 `launch.json` 文件，如果第一次打开,会新建一个配置文件，默认配置内容如下

```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${fileDirname}",
            "env": {},
            "args": []
        }
    ]
}
```

常见属性如下

|序号|属性|介绍|
|:---|:--|:----|
|1|name  |  调试界面下拉选择项的名称|
|2|type  |  设置为go无需改动，是 vs code 用于计算调试代码需要用哪个扩展|
|3|mode  |可以设置为 auto, debug, remote, test, exec 中的一个|
|4|program  |  调试程序的路径（绝对路径）|
|5|env   | 调试时使用的环境变量。例如:{ "ENVNAME": "ENVVALUE" }|
|6|envFile  |  包含环境变量文件的绝对路径，在 env 中设置的属性会覆盖 envFile 中的配置|
|7|args  |  传给正在调试程序命令行参数数组|
|8|showLog |   布尔值，是否将调试信息输出|
|9|logOutput | 配置调试输出的组件（debugger, gdbwire, lldbout, debuglineerr, rpc）,使用,分隔， showLog 设置为 true 时，此项配置生效|
|10|buildFlags   | 构建 go 程序时传给 go 编译器的标志|
|11|remotePath   | 远程调试程序的绝对路径，当 mode 设置为 remote 时有效在 debug 配置中使用 VS Code 变量|
|12|${workspaceFolder} |调试 VS Code 打开工作空间的根目录下的所有文件|
|13|${file}| 调试当前文件|
|14|${fileDirname} |调试当前文件所在目录下的所有文件|

第八步：打开SFTP的sftp.json

功能：实现远程文件安全传输，例如本地文件传输的虚拟机，保存就可以实现本地与虚拟机文件同步。

ctrl+shift+p,输入SFTP：Config

sftp.json文件

```json
{
    "host": "192.168.211.167",
    "port": 22,
    "username": "wen",
    "password": "123",
    "protocol": "sftp",
    "agent": null,
    "privateKeyPath": null,
    "passphrase": null,
    "passive": false,
    "interactiveAuth": false,
    "remotePath": "/home/wen/c++",
    "uploadOnSave": true,
    "syncMode": "update",
    "watcher": {
    "files": false,
    "autoUpload": true,
    "autoDelete": false
    },
    "ignore": [
    "**/.vscode/**",
    "**/.git/**",
    "**/.DS_Store",
    "**/.svn",
    "**/.tx",
    "**/.github"
    ]
}
```

## 3、 Go语言的依赖及如何使用

安装golang，在GOROOT\pkg下有默认安装包

常见的第三方包地址： https://pkg.go.dev

Go语言的依赖管理随着版本的更迭正逐渐完善起来。

### 3.1 依赖管理

- 为什么需要依赖管理？

最早的时候，Go所依赖的所有的第三方库都放在GOPATH这个目录下面。这就导致了同一个库只能保存一个版本的代码。如果不同的项目依赖同一个第三方的库的不同版本，应该怎么解决？

### 3.2 godep工具

Go语言从v1.5开始开始引入vendor模式，如果项目目录下有vendor目录，那么go工具链会优先使用vendor内的包进行编译、测试等。

godep是一个通过vender模式实现的Go语言的第三方依赖管理工具，类似的还有由社区维护准官方包管理工具dep。

- 安装

执行以下命令安装godep工具。
`go get github.com/tools/godep`

- 基本命令

安装好godep之后，在终端输入godep查看支持的所有命令。

```godep
godep save     将依赖项输出并复制到Godeps.json文件中
godep go       使用保存的依赖项运行go工具
godep get      下载并安装具有指定依赖项的包
godep path     打印依赖的GOPATH路径
godep restore  在GOPATH中拉取依赖的版本
godep update   更新选定的包或go版本
godep diff     显示当前和以前保存的依赖项集之间的差异
godep version  查看版本信息
```

使用`godep help [command]`可以看看具体命令的帮助信息。

- 使用godep

在项目目录下执行godep save命令，会在当前项目中创建Godeps和vender两个文件夹。

其中Godeps文件夹下有一个Godeps.json的文件，里面记录了项目所依赖的包信息。 vender文件夹下是项目依赖的包的源代码文件。

- vender机制

Go1.5版本之后开始支持，能够控制Go语言程序编译时依赖包搜索路径的优先级。

例如查找项目的某个依赖包，首先会在项目根目录下的vender文件夹中查找，如果没有找到就会去$GOAPTH/src目录下查找。

- godep开发流程

1. 保证程序能够正常编译
2. 执行godep save保存当前项目的所有第三方依赖的版本信息和代码
3. 提交Godeps目录和vender目录到代码库。
4. 如果要更新依赖的版本，可以直接修改Godeps.json文件中的对应项

### 3.3 go module

go module是Go1.11版本之后官方推出的版本管理工具，并且从Go1.13版本开始，go module将是Go语言默认的依赖管理工具。

- GO111MODULE

要启用go module支持首先要设置环境变量GO111MODULE，通过它可以开启或关闭模块支持，它有三个可选值：off、on、auto，默认值是auto。

1. `GO111MODULE=off`禁用模块支持，编译时会从GOPATH和vendor文件夹中查找包。
2. `GO111MODULE=on`启用模块支持，编译时会忽略GOPATH和vendor文件夹，只根据 go.mod下载依赖。
3. `GO111MODULE=auto`，当项目在$GOPATH/src外且项目根目录有go.mod文件时，开启模块支持。

简单来说，设置`GO111MODULE=on`之后就可以使用`go module`了，以后就没有必要在GOPATH中创建项目了，并且还能够很好的管理项目依赖的第三方包信息。

使用 `go module` 管理依赖后会在项目根目录下生成两个文件`go.mod`和`go.sum`。

### 3.4 GOPROXY

Go1.11之后设置GOPROXY命令为：

`export GOPROXY=https://goproxy.cn`

Go1.13之后GOPROXY默认值为`https://proxy.golang.org`，在国内是无法访问的，所以十分建议大家设置GOPROXY，这里我推荐使用goproxy.cn。

`go env -w GOPROXY=https://goproxy.cn,direct`

### 3.5 go mod命令

常用的go mod命令如下：

```go
go mod download    下载依赖的module到本地cache（默认为$GOPATH/pkg/mod目录）
go mod edit        编辑go.mod文件
go mod graph       打印模块依赖图
go mod init        初始化当前文件夹, 创建go.mod文件
go mod tidy        增加缺少的module，删除无用的module
go mod vendor      将依赖复制到vendor下
go mod verify      校验依赖
go mod why         解释为什么需要依赖
```

- `go.mod`

`go.mod`文件记录了项目所有的依赖信息，其结构大致如下：

```go
module github.com/Q1mi/studygo/blogger
go 1.12
require (
	github.com/DeanThompson/ginpprof v0.0.0-20190408063150-3be636683586
	github.com/gin-gonic/gin v1.4.0
	github.com/go-sql-driver/mysql v1.4.1
	github.com/jmoiron/sqlx v1.2.0
	github.com/satori/go.uuid v1.2.0
	google.golang.org/appengine v1.6.1 // indirect
)
```

其中，

1. module用来定义包名
2. require用来定义依赖包及版本
3. indirect表示间接引用

- 依赖的版本

go mod支持语义化版本号，比如`go get foo@v1.2.3`，也可以跟git的分支或tag，比如`go get foo@master`，当然也可以跟git提交哈希，比如`go get foo@e3702bed2`。关于依赖的版本支持以下几种格式：

```go
gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7
gopkg.in/vmihailenco/msgpack.v2 v2.9.1
gopkg.in/yaml.v2 <=v2.2.1
github.com/tatsushid/go-fastping v0.0.0-20160109021039-d7bb493dee3e
latest
```

- replace

在国内访问golang.org/x的各个包都需要翻墙，你可以在go.mod中使用replace替换成github上对应的库。

replace (
	golang.org/x/crypto v0.0.0-20180820150726-614d502a4dac => github.com/golang/crypto v0.0.0-20180820150726-614d502a4dac
	golang.org/x/net v0.0.0-20180821023952-922f4815f713 => github.com/golang/net v0.0.0-20180826012351-8a410e7b638d
	golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0
)

#### 3.5.1 go get

在项目中执行go get命令可以下载依赖包，并且还可以指定下载的版本。

运行`go get -u`将会升级到最新的次要版本或者修订版本(x.y.z, z是修订版本号， y是次要版本号)
运行go get -u=patch将会升级到最新的修订版本
运行go get package@version将会升级到指定的版本号version
如果下载所有依赖可以使用go mod download命令。

整理依赖
我们在代码中删除依赖代码后，相关的依赖库并不会在go.mod文件中自动移除。这种情况下我们可以使用go mod tidy命令更新go.mod中的依赖关系。

#### 3.5.2 go mod edit

- 格式化

因为我们可以手动修改go.mod文件，所以有些时候需要格式化该文件。Go提供了一下命令：

`go mod edit -fmt`

- 添加依赖项

`go mod edit -require=golang.org/x/text`

- 移除依赖项

如果只是想修改go.mod文件中的内容，那么可以运行`go mod edit -drop require=package path`，比如要在`go.mod`中移除`golang.org/x/text`包，可以使用如下命令：

`go mod edit -drop require=golang.org/x/text`

关于go mod edit的更多用法可以通过`go help mod edit`查看。

### 3.6 在项目中使用go module

- 既有项目

如果需要对一个已经存在的项目启用go module，可以按照以下步骤操作：

在项目目录下执行`go mod init`，生成一个`go.mod`文件。
执行`go get`，查找并记录当前项目的依赖，同时生成一个`go.sum`记录每个依赖库的版本和哈希值。

- 新项目

对于一个新创建的项目，我们可以在项目文件夹下按照以下步骤操作：

执行`go mod init` 项目名命令，在当前项目文件夹下创建一个`go.mod`文件。
手动编辑go.mod中的require依赖项或执行go get自动发现、维护依赖。
go module是Go1.11版本之后官方推出的版本管理工具，并且从Go1.13版本开始，go module将是Go语言默认的依赖管理工具。到今天Go1.14版本推出之后Go modules 功能已经被正式推荐在生产环境下使用了。

这几天已经有很多教程讲解如何使用go module，以及如何使用go module导入gitlab私有仓库，我这里就不再啰嗦了。但是最近我发现很多小伙伴在群里问如何使用go module导入本地包，作为初学者大家刚开始接触package的时候肯定都是先在本地创建一个包，然后本地调用一下，然后就被卡住了。。。

这里就详细介绍下如何使用go module导入本地包。

- 前提

假设我们现在有`moduledemo`和`mypackage`两个包，其中`moduledemo`包中会导入`mypackage`包并使用它的New方法。

mypackage/mypackage.go内容如下：


```go
package mypackage

import "fmt"

func New(){
	fmt.Println("mypackage.New")
}
```

我们现在分两种情况讨论：

#### 3.6.1 在同一个项目下

注意：在一个项目（project）下我们是可以定义多个包（package）的。

目录结构
现在的情况是，我们在moduledemo/main.go中调用了mypackage这个包。

```go
moduledemo
├── go.mod
├── main.go
└── mypackage
    └── mypackage.go
```

- 导入包

这个时候，我们需要在moduledemo/go.mod中按如下定义：

```go
module moduledemo

go 1.14
```

然后在moduledemo/main.go中按如下方式导入mypackage

```go
package main

import (
	"fmt"
	"moduledemo/mypackage"  // 导入同一项目下的mypackage包
)
func main() {
	mypackage.New()
	fmt.Println("main")
}
```

- 举个例子

举一反三，假设我们现在有文件目录结构如下：

```go
└── bubble
    ├── dao
    │   └── mysql.go
    ├── go.mod
    └── main.go
```

其中`bubble/go.mod`内容如下：

```go
module github.com/q1mi/bubble

go 1.14
```

`bubble/dao/mysql.go`内容如下：

```
package dao

import "fmt"

func New(){
	fmt.Println("mypackage.New")
}
```

`bubble/main.go`内容如下：

```go
package main

import (
	"fmt"
	"github.com/q1mi/bubble/dao"
)
func main() {
	dao.New()
	fmt.Println("main")
}
```


#### 3.6.2 不在同一个项目下

目录结构

```go
├── moduledemo
│   ├── go.mod
│   └── main.go
└── mypackage
    ├── go.mod
    └── mypackage.go
```

- 导入包

这个时候，mypackage也需要进行module初始化，即拥有一个属于自己的go.mod文件，内容如下：

```go
module mypackage

go 1.14
```

然后我们在`moduledemo/main.go`中按如下方式导入：

```go
import (
	"fmt"
	"mypackage"
)
func main() {
	mypackage.New()
	fmt.Println("main")
}
```

因为这两个包不在同一个项目路径下，你想要导入本地包，并且这些包也没有发布到远程的github或其他代码仓库地址。这个时候我们就需要在go.mod文件中使用replace指令。

在调用方也就是`packagedemo/go.mod`中按如下方式指定使用相对路径来寻找mypackage这个包。

```go
module moduledemo

go 1.14

require "mypackage" v0.0.0
replace "mypackage" => "../mypackage"
```

- 举个例子

最后我们再举个例子巩固下上面的内容。

我们现在有文件目录结构如下：

```go
├── p1
│   ├── go.mod
│   └── main.go
└── p2
    ├── go.mod
    └── p2.go
```

`p1/main.go`中想要导入p2.go中定义的函数。

`p2/go.mod`内容如下：

```go
module liwenzhou.com/q1mi/p2

go 1.14
```

`p1/main.go`中按如下方式导入

```go
import (
	"fmt"
	"liwenzhou.com/q1mi/p2"
)
func main() {
	p2.New()
	fmt.Println("main")
}
```

因为我并没有把liwenzhou.com/q1mi/p2这个包上传到liwenzhou.com这个网站，我们只是想导入本地的包，这个时候就需要用到replace这个指令了。

p1/go.mod内容如下：

```go
module github.com/q1mi/p1

go 1.14


require "liwenzhou.com/q1mi/p2" v0.0.0
replace "liwenzhou.com/q1mi/p2" => "../p2"
```

此时，我们就可以正常编译p1这个项目了。

## 4、目录

4.1、go test工具
4.2、测试函数
4.2.1、测试函数的格式
4.2.2、测试函数示例
4.5、测试组
4.6、子测试
4.7、测试覆盖率
4.8、基准测试
4.9、基准测试函数格式
4.10、基准测试示例
4.11、性能比较函数
4.12、重置时间
4.13、并行测试
4.14、Setup与TearDown
4.15、TestMain
4.16、子测试的Setup与Teardown
4.17、示例函数
4.18、示例函数的格式
4.19、示例函数示例

这篇文章主要介绍下在Go语言中如何做单元测试和基准测试。

### 4.1、go test工具

Go语言中的测试依赖`go test`命令。编写测试代码和编写普通的Go代码过程是类似的，并不需要学习新的语法、规则或工具。

`go test`命令是一个按照一定约定和组织的测试代码的驱动程序。在包目录内，所有以`_test.go`为后缀名的源代码文件都是`go test`测试的一部分，不会被go build编译到最终的可执行文件中。

在`*_test.go`文件中有三种类型的函数: **单元测试函数、基准测试函数和示例函数。**

|类型	|格式	|作用|
|:------|:------|:---|
|测试函数	|函数名前缀为Test|测试程序的一些逻辑行为是否正确|
|基准函数	|函数名前缀为Benchmark	|测试函数的性能|
|示例函数	|函数名前缀为Example	|为文档提供示例文档|

`go test`命令会遍历所有的`*_test.go`文件中符合上述命名规则的函数，然后生成一个临时的main包用于调用相应的测试函数，然后构建并运行、报告测试结果，最后清理测试中生成的临时文件。

### 4.2、测试函数

#### 4.2.1、测试函数的格式

每个测试函数必须导入testing包，测试函数的基本格式（签名）如下：

```go
func TestName(t *testing.T){
    // ...
}
```

测试函数的名字必须以Test开头，可选的后缀名必须以大写字母开头，举几个例子：

```go
func TestAdd(t *testing.T){ ... }
func TestSum(t *testing.T){ ... }
func TestLog(t *testing.T){ ... }
```

其中参数t用于报告测试失败和附加的日志信息。 testing.T的拥有的方法如下：

```go
func (c *T) Error(args ...interface{})
func (c *T) Errorf(format string, args ...interface{})
func (c *T) Fail()
func (c *T) FailNow()
func (c *T) Failed() bool
func (c *T) Fatal(args ...interface{})
func (c *T) Fatalf(format string, args ...interface{})
func (c *T) Log(args ...interface{})
func (c *T) Logf(format string, args ...interface{})
func (c *T) Name() string
func (t *T) Parallel()
func (t *T) Run(name string, f func(t *T)) bool
func (c *T) Skip(args ...interface{})
func (c *T) SkipNow()
func (c *T) Skipf(format string, args ...interface{})
func (c *T) Skipped() bool
```

#### 4.2.2、测试函数示例

就像细胞是构成我们身体的基本单位，一个软件程序也是由很多单元组件构成的。单元组件可以是函数、结构体、方法和最终用户可能依赖的任意东西。总之我们需要确保这些组件是能够正常运行的。单元测试是一些利用各种方法测试单元组件的程序，它会将结果与预期输出进行比较。

接下来，我们定义一个split的包，包中定义了一个Split函数，具体实现如下：

```split/split.go```

```go
package split

import &quot;strings&quot;

// split package with a single split function.
// Split slices s into all substrings separated by sep and
// returns a slice of the substrings between those separators.
func Split(s, sep string) (result []string) {
	i := strings.Index(s, sep)

	for i &gt; -1 {
		result = append(result, s[:i])
		s = s[i+1:]
		i = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}
```

在当前目录下，我们创建一个split_test.go的测试文件，并定义一个测试函数如下：

```split/split_test.go```

```go
package split

import (
	&quot;reflect&quot;
	&quot;testing&quot;
)

func TestSplit(t *testing.T) { // 测试函数名必须以Test开头，必须接收一个*testing.T类型参数
	got := Split(&quot;a:b:c&quot;, &quot;:&quot;)         // 程序输出的结果
	want := []string{&quot;a&quot;, &quot;b&quot;, &quot;c&quot;}    // 期望的结果
	if !reflect.DeepEqual(want, got) { // 因为slice不能比较直接，借助反射包中的方法比较
		t.Errorf(&quot;excepted:%v, got:%v&quot;, want, got) // 测试失败输出错误提示
	}
}
```

此时split这个包中的文件如下：

```go
split $ ls -l
total 16
-rw-r--r--  1 nickchen121  staff  408  4 29 15:50 split.go
-rw-r--r--  1 nickchen121  staff  466  4 29 16:04 split_test.go
```

在split包路径下，执行go test命令，可以看到输出结果如下：

```go
split $ go test
PASS
ok      github.com/Q1mi/studygo/code_demo/test_demo/split       0.005s
```

一个测试用例有点单薄，我们再编写一个测试使用多个字符切割字符串的例子，在`split_test.go`中添加如下测试函数：

```go
func TestMoreSplit(t *testing.T) {
	got := Split(&quot;abcd&quot;, &quot;bc&quot;)
	want := []string{&quot;a&quot;, &quot;d&quot;}
	if !reflect.DeepEqual(want, got) {
		t.Errorf(&quot;excepted:%v, got:%v&quot;, want, got)
	}
}
```

再次运行`go test`命令，输出结果如下：

```go
split $ go test
--- FAIL: TestMultiSplit (0.00s)
    split_test.go:20: excepted:[a d], got:[a cd]
FAIL
exit status 1
FAIL    github.com/Q1mi/studygo/code_demo/test_demo/split       0.006s
```

这一次，我们的测试失败了。我们可以为go test命令添加-v参数，查看测试函数名称和运行时间：

```go
split $ go test -v
=== RUN   TestSplit
--- PASS: TestSplit (0.00s)
=== RUN   TestMoreSplit
--- FAIL: TestMoreSplit (0.00s)
    split_test.go:21: excepted:[a d], got:[a cd]
FAIL
exit status 1
FAIL    github.com/Q1mi/studygo/code_demo/test_demo/split       0.005s
```

这一次我们能清楚的看到是TestMoreSplit这个测试没有成功。 还可以在go test命令后添加-run参数，它对应一个正则表达式，只有函数名匹配上的测试函数才会被go test命令执行。

```go
split $ go test -v -run=&quot;More&quot;
=== RUN   TestMoreSplit
--- FAIL: TestMoreSplit (0.00s)
    split_test.go:21: excepted:[a d], got:[a cd]
FAIL
exit status 1
FAIL    github.com/Q1mi/studygo/code_demo/test_demo/split       0.006s
```

现在我们回过头来解决我们程序中的问题。很显然我们最初的split函数并没有考虑到sep为多个字符的情况，我们来修复下这个Bug：

```go
package split

import &quot;strings&quot;

// split package with a single split function.

// Split slices s into all substrings separated by sep and
// returns a slice of the substrings between those separators.
func Split(s, sep string) (result []string) {
	i := strings.Index(s, sep)

	for i &gt; -1 {
		result = append(result, s[:i])
		s = s[i+len(sep):] // 这里使用len(sep)获取sep的长度
		i = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}
```

这一次我们再来测试一下，我们的程序。注意，当我们修改了我们的代码之后不要仅仅执行那些失败的测试函数，我们应该完整的运行所有的测试，保证不会因为修改代码而引入了新的问题。

```go
split $ go test -v
=== RUN   TestSplit
--- PASS: TestSplit (0.00s)
=== RUN   TestMoreSplit
--- PASS: TestMoreSplit (0.00s)
PASS
ok      github.com/Q1mi/studygo/code_demo/test_demo/split       0.006s
```

这一次我们的测试都通过了。

#### 4.2.3、测试组

我们现在还想要测试一下split函数对中文字符串的支持，这个时候我们可以再编写一个TestChineseSplit测试函数，但是我们也可以使用如下更友好的一种方式来添加更多的测试用例。

```go
func TestSplit(t *testing.T) {
   // 定义一个测试用例类型
	type test struct {
		input string
		sep   string
		want  []string
	}
	// 定义一个存储测试用例的切片
	tests := []test{
		{input: &quot;a:b:c&quot;, sep: &quot;:&quot;, want: []string{&quot;a&quot;, &quot;b&quot;, &quot;c&quot;}},
		{input: &quot;a:b:c&quot;, sep: &quot;,&quot;, want: []string{&quot;a:b:c&quot;}},
		{input: &quot;abcd&quot;, sep: &quot;bc&quot;, want: []string{&quot;a&quot;, &quot;d&quot;}},
		{input: &quot;沙河有沙又有河&quot;, sep: &quot;沙&quot;, want: []string{&quot;河有&quot;, &quot;又有河&quot;}},
	}
	// 遍历切片，逐一执行测试用例
	for _, tc := range tests {
		got := Split(tc.input, tc.sep)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf(&quot;excepted:%v, got:%v&quot;, tc.want, got)
		}
	}
}
```

我们通过上面的代码把多个测试用例合到一起，再次执行go test命令。

```go
split $ go test -v
=== RUN   TestSplit
--- FAIL: TestSplit (0.00s)
    split_test.go:42: excepted:[河有 又有河], got:[ 河有 又有河]
FAIL
exit status 1
FAIL    github.com/Q1mi/studygo/code_demo/test_demo/split       0.006s
```

我们的测试出现了问题，仔细看打印的测试失败提示信息：excepted:[河有 又有河], got:[ 河有 又有河]，你会发现[ 河有 又有河]中有个不明显的空串，这种情况下十分推荐使用%#v的格式化方式。

我们修改下测试用例的格式化输出错误提示部分：

```go
func TestSplit(t *testing.T) {
   ...
   
	for _, tc := range tests {
		got := Split(tc.input, tc.sep)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf(&quot;excepted:%#v, got:%#v&quot;, tc.want, got)
		}
	}
}
```

此时运行go test命令后就能看到比较明显的提示信息了：

```
split $ go test -v
=== RUN   TestSplit
--- FAIL: TestSplit (0.00s)
    split_test.go:42: excepted:[]string{&quot;河有&quot;, &quot;又有河&quot;}, got:[]string{&quot;&quot;, &quot;河有&quot;, &quot;又有河&quot;}
FAIL
exit status 1
FAIL    github.com/Q1mi/studygo/code_demo/test_demo/split       0.006s
```

#### 4.2.4、子测试

看起来都挺不错的，但是如果测试用例比较多的时候，我们是没办法一眼看出来具体是哪个测试用例失败了。我们可能会想到下面的解决办法：

- 使用t.Run()执行子测试

```go
func TestSplit(t *testing.T) {
	type test struct { // 定义test结构体
		input string
		sep   string
		want  []string
	}
	tests := map[string]test{ // 测试用例使用map存储
		&quot;simple&quot;:      {input: &quot;a:b:c&quot;, sep: &quot;:&quot;, want: []string{&quot;a&quot;, &quot;b&quot;, &quot;c&quot;}},
		&quot;wrong sep&quot;:   {input: &quot;a:b:c&quot;, sep: &quot;,&quot;, want: []string{&quot;a:b:c&quot;}},
		&quot;more sep&quot;:    {input: &quot;abcd&quot;, sep: &quot;bc&quot;, want: []string{&quot;a&quot;, &quot;d&quot;}},
		&quot;leading sep&quot;: {input: &quot;沙河有沙又有河&quot;, sep: &quot;沙&quot;, want: []string{&quot;河有&quot;, &quot;又有河&quot;}},
	}
	for name, tc := range tests {
		got := Split(tc.input, tc.sep)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf(&quot;name:%s excepted:%#v, got:%#v&quot;, name, tc.want, got) // 将测试用例的name格式化输出
		}
	}
}
```

上面的做法是能够解决问题的。同时Go1.7+中新增了子测试，我们可以按照如下方式使用t.Run执行子测试：

```go
func TestSplit(t *testing.T) {
	type test struct { // 定义test结构体
		input string
		sep   string
		want  []string
	}
	tests := map[string]test{ // 测试用例使用map存储
		&quot;simple&quot;:      {input: &quot;a:b:c&quot;, sep: &quot;:&quot;, want: []string{&quot;a&quot;, &quot;b&quot;, &quot;c&quot;}},
		&quot;wrong sep&quot;:   {input: &quot;a:b:c&quot;, sep: &quot;,&quot;, want: []string{&quot;a:b:c&quot;}},
		&quot;more sep&quot;:    {input: &quot;abcd&quot;, sep: &quot;bc&quot;, want: []string{&quot;a&quot;, &quot;d&quot;}},
		&quot;leading sep&quot;: {input: &quot;沙河有沙又有河&quot;, sep: &quot;沙&quot;, want: []string{&quot;河有&quot;, &quot;又有河&quot;}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf(&quot;excepted:%#v, got:%#v&quot;, tc.want, got)
			}
		})
	}
}
```

此时我们再执行go test命令就能够看到更清晰的输出内容了：

```go
split $ go test -v
=== RUN   TestSplit
=== RUN   TestSplit/leading_sep
=== RUN   TestSplit/simple
=== RUN   TestSplit/wrong_sep
=== RUN   TestSplit/more_sep
--- FAIL: TestSplit (0.00s)
    --- FAIL: TestSplit/leading_sep (0.00s)
        split_test.go:83: excepted:[]string{&quot;河有&quot;, &quot;又有河&quot;}, got:[]string{&quot;&quot;, &quot;河有&quot;, &quot;又有河&quot;}
    --- PASS: TestSplit/simple (0.00s)
    --- PASS: TestSplit/wrong_sep (0.00s)
    --- PASS: TestSplit/more_sep (0.00s)
FAIL
exit status 1
FAIL    github.com/Q1mi/studygo/code_demo/test_demo/split       0.006s
```

这个时候我们要把测试用例中的错误修改回来：

```go
func TestSplit(t *testing.T) {
	...
	tests := map[string]test{ // 测试用例使用map存储
		&quot;simple&quot;:      {input: &quot;a:b:c&quot;, sep: &quot;:&quot;, want: []string{&quot;a&quot;, &quot;b&quot;, &quot;c&quot;}},
		&quot;wrong sep&quot;:   {input: &quot;a:b:c&quot;, sep: &quot;,&quot;, want: []string{&quot;a:b:c&quot;}},
		&quot;more sep&quot;:    {input: &quot;abcd&quot;, sep: &quot;bc&quot;, want: []string{&quot;a&quot;, &quot;d&quot;}},
		&quot;leading sep&quot;: {input: &quot;沙河有沙又有河&quot;, sep: &quot;沙&quot;, want: []string{&quot;&quot;, &quot;河有&quot;, &quot;又有河&quot;}},
	}
	...
}
```

我们都知道可以通过`-run=RegExp`来指定运行的测试用例，还可以通过/来指定要运行的子测试用例，例如：`go test -v -run=Split/simple`只会运行simple对应的子测试用例。

#### 4.2.5、测试覆盖率

测试覆盖率是你的代码被测试套件覆盖的百分比。通常我们使用的都是语句的覆盖率，也就是在测试中至少被运行一次的代码占总代码的比例。

Go提供内置功能来检查你的代码覆盖率。我们可以使用`go test -cover`来查看测试覆盖率。例如：

```go
split $ go test -cover
PASS
coverage: 100.0% of statements
ok      github.com/Q1mi/studygo/code_demo/test_demo/split       0.005s
```

从上面的结果可以看到我们的测试用例覆盖了100%的代码。

Go还提供了一个额外的`-coverprofile`参数，用来将覆盖率相关的记录信息输出到一个文件。例如：

```go
split $ go test -cover -coverprofile=c.out
PASS
coverage: 100.0% of statements
ok      github.com/Q1mi/studygo/code_demo/test_demo/split       0.005s
```

上面的命令会将覆盖率相关的信息输出到当前文件夹下面的c.out文件中，然后我们执行go tool cover -html=c.out，使用cover工具来处理生成的记录信息，该命令会打开本地的浏览器窗口生成一个HTML报告。 cover.png 上图中每个用绿色标记的语句块表示被覆盖了，而红色的表示没有被覆盖。

### 4.3、基准测试

#### 4.3.1、基准测试函数格式

**基准测试就是在一定的工作负载之下检测程序性能的一种方法。** 基准测试的基本格式如下：

```go
func BenchmarkName(b *testing.B){
    // ...
}
```

基准测试以Benchmark为前缀，需要一个`*testing.B`类型的参数b，基准测试必须要执行`b.N`次，这样的测试才有对照性，`b.N`的值是系统根据实际情况去调整的，从而保证测试的稳定性。 `testing.B`拥有的方法如下：

```go
func (c *B) Error(args ...interface{})
func (c *B) Errorf(format string, args ...interface{})
func (c *B) Fail()
func (c *B) FailNow()
func (c *B) Failed() bool
func (c *B) Fatal(args ...interface{})
func (c *B) Fatalf(format string, args ...interface{})
func (c *B) Log(args ...interface{})
func (c *B) Logf(format string, args ...interface{})
func (c *B) Name() string
func (b *B) ReportAllocs()
func (b *B) ResetTimer()
func (b *B) Run(name string, f func(b *B)) bool
func (b *B) RunParallel(body func(*PB))
func (b *B) SetBytes(n int64)
func (b *B) SetParallelism(p int)
func (c *B) Skip(args ...interface{})
func (c *B) SkipNow()
func (c *B) Skipf(format string, args ...interface{})
func (c *B) Skipped() bool
func (b *B) StartTimer()
func (b *B) StopTimer()
```

#### 4.3.2、基准测试示例

我们为split包中的`Split函数`编写基准测试如下：

```go
func BenchmarkSplit(b *testing.B) {
	for i := 0; i &lt; b.N; i++ {
		Split(&quot;沙河有沙又有河&quot;, &quot;沙&quot;)
	}
}
```

基准测试并不会默认执行，需要增加-bench参数，所以我们通过执行`go test -bench=Split`命令执行基准测试，输出结果如下：

```go
split $ go test -bench=Split
goos: darwin
goarch: amd64
pkg: github.com/Q1mi/studygo/code_demo/test_demo/split
BenchmarkSplit-8        10000000               203 ns/op
PASS
ok      github.com/Q1mi/studygo/code_demo/test_demo/split       2.255s
```

其中`BenchmarkSplit-8`表示对`Split`函数进行基准测试，数字`8`表示`GOMAXPROCS`的值，这个对于并发基准测试很重要。``10000000`和`203ns/op`表示每次调用Split函数耗时`203ns`，这个结果是`10000000次调用的平均值。

我们还可以为基准测试添加`-benchmem`参数，来获得内存分配的统计数据。

```go
split $ go test -bench=Split -benchmem
goos: darwin
goarch: amd64
pkg: github.com/Q1mi/studygo/code_demo/test_demo/split
BenchmarkSplit-8        10000000               215 ns/op             112 B/op          3 allocs/op
PASS
ok      github.com/Q1mi/studygo/code_demo/test_demo/split       2.394s
```

其中，`112 B/op`表示每次操作内存分配了112字节，`3 allocs/op`则表示每次操作进行了3次内存分配。 我们将我们的Split函数优化如下：

```go
func Split(s, sep string) (result []string) {
	result = make([]string, 0, strings.Count(s, sep)+1)
	i := strings.Index(s, sep)
	for i &gt; -1 {
		result = append(result, s[:i])
		s = s[i+len(sep):] // 这里使用len(sep)获取sep的长度
		i = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}
```

这一次我们提前使用make函数将result初始化为一个容量足够大的切片，而不再像之前一样通过调用append函数来追加。我们来看一下这个改进会带来多大的性能提升：

```go
split $ go test -bench=Split -benchmem
goos: darwin
goarch: amd64
pkg: github.com/Q1mi/studygo/code_demo/test_demo/split
BenchmarkSplit-8        10000000               127 ns/op              48 B/op          1 allocs/op
PASS
ok      github.com/Q1mi/studygo/code_demo/test_demo/split       1.423s
```

这个使用make函数提前分配内存的改动，减少了2/3的内存分配次数，并且减少了一半的内存分配。

#### 4.3.2、性能比较函数

上面的基准测试只能得到给定操作的绝对耗时，但是在很多性能问题是发生在两个不同操作之间的相对耗时，比如同一个函数处理1000个元素的耗时与处理1万甚至100万个元素的耗时的差别是多少？再或者对于同一个任务究竟使用哪种算法性能最佳？我们通常需要对两个不同算法的实现使用相同的输入来进行基准比较测试。

性能比较函数通常是一个带有参数的函数，被多个不同的Benchmark函数传入不同的值来调用。举个例子如下：

```go
func benchmark(b *testing.B, size int){/* ... */}
func Benchmark10(b *testing.B){ benchmark(b, 10) }
func Benchmark100(b *testing.B){ benchmark(b, 100) }
func Benchmark1000(b *testing.B){ benchmark(b, 1000) }
```

例如我们编写了一个计算斐波那契数列的函数如下：

// fib.go
// Fib 是一个计算第n个斐波那契数的函数

```go
func Fib(n int) int {
	if n &lt; 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}
```

我们编写的性能比较函数如下：

// fib_test.go

```go
func benchmarkFib(b *testing.B, n int) {
	for i := 0; i &lt; b.N; i++ {
		Fib(n)
	}
}

func BenchmarkFib1(b *testing.B)  { benchmarkFib(b, 1) }
func BenchmarkFib2(b *testing.B)  { benchmarkFib(b, 2) }
func BenchmarkFib3(b *testing.B)  { benchmarkFib(b, 3) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(b, 10) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(b, 20) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(b, 40) }
```

运行基准测试：

```go
split $ go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/Q1mi/studygo/code_demo/test_demo/fib
BenchmarkFib1-8         1000000000               2.03 ns/op
BenchmarkFib2-8         300000000                5.39 ns/op
BenchmarkFib3-8         200000000                9.71 ns/op
BenchmarkFib10-8         5000000               325 ns/op
BenchmarkFib20-8           30000             42460 ns/op
BenchmarkFib40-8               2         638524980 ns/op
PASS
ok      github.com/Q1mi/studygo/code_demo/test_demo/fib 12.944s
```

这里需要注意的是，默认情况下，每个基准测试至少运行1秒。如果在Benchmark函数返回时没有到1秒，则b.N的值会按1,2,5,10,20,50，…增加，并且函数再次运行。

最终的BenchmarkFib40只运行了两次，每次运行的平均值只有不到一秒。像这种情况下我们应该可以使用`-benchtime`标志增加最小基准时间，以产生更准确的结果。例如：

```go
split $ go test -bench=Fib40 -benchtime=20s
goos: darwin
goarch: amd64
pkg: github.com/Q1mi/studygo/code_demo/test_demo/fib
BenchmarkFib40-8              50         663205114 ns/op
PASS
ok      github.com/Q1mi/studygo/code_demo/test_demo/fib 33.849s
```

这一次BenchmarkFib40函数运行了50次，结果就会更准确一些了。

使用性能比较函数做测试的时候一个容易犯的错误就是把b.N作为输入的大小，例如以下两个例子都是错误的示范：

```go
// 错误示范1
func BenchmarkFibWrong(b *testing.B) {
	for n := 0; n &lt; b.N; n++ {
		Fib(n)
	}
}

// 错误示范2
func BenchmarkFibWrong2(b *testing.B) {
	Fib(b.N)
}
```

#### 4.3.3、重置时间

b.ResetTimer之前的处理不会放到执行时间里，也不会输出到报告中，所以可以在之前做一些不计划作为测试报告的操作。例如：

```go
func BenchmarkSplit(b *testing.B) {
	time.Sleep(5 * time.Second) // 假设需要做一些耗时的无关操作
	b.ResetTimer()              // 重置计时器
	for i := 0; i &lt; b.N; i++ {
		Split(&quot;沙河有沙又有河&quot;, &quot;沙&quot;)
	}
}
```

#### 4.3.4、并行测试

`func (b *B) RunParallel(body func(*PB))`会以并行的方式执行给定的基准测试。

RunParallel会创建出多个goroutine，并将b.N分配给这些goroutine执行， 其中goroutine数量的默认值为GOMAXPROCS。用户如果想要增加非CPU受限（non-CPU-bound）基准测试的并行性， 那么可以在RunParallel之前调用SetParallelism 。RunParallel通常会与-cpu标志一同使用。

```go
func BenchmarkSplitParallel(b *testing.B) {
	// b.SetParallelism(1) // 设置使用的CPU数
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Split(&quot;沙河有沙又有河&quot;, &quot;沙&quot;)
		}
	})
}
```

执行一下基准测试：

```go
split $ go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/Q1mi/studygo/code_demo/test_demo/split
BenchmarkSplit-8                10000000               131 ns/op
BenchmarkSplitParallel-8        50000000                36.1 ns/op
PASS
ok      github.com/Q1mi/studygo/code_demo/test_demo/split       3.308s
```

还可以通过在测试命令后添加-cpu参数如go test -bench=. -cpu 1来指定使用的CPU数量。

#### 4.3.5、Setup与TearDown

测试程序有时需要在测试之前进行额外的设置（setup）或在测试之后进行拆卸（teardown）。

#### 4.3.6、TestMain

通过在`*_test.go`文件中定义TestMain函数来可以在测试之前进行额外的设置（setup）或在测试之后进行拆卸（teardown）操作。

如果测试文件包含函数:`func TestMain(m *testing.M)`，那么生成的测试会先调用 TestMain(m)，然后再运行具体测试。TestMain运行在主goroutine中, 可以在调用 m.Run前后做任何设置（setup）和拆卸（teardown）。退出测试的时候应该使用m.Run的返回值作为参数调用os.Exit。

一个使用TestMain来设置Setup和TearDown的示例如下：

```go
func TestMain(m *testing.M) {
	fmt.Println(&quot;write setup code here...&quot;) // 测试之前的做一些设置
	// 如果 TestMain 使用了 flags，这里应该加上flag.Parse()
	retCode := m.Run()                         // 执行测试
	fmt.Println(&quot;write teardown code here...&quot;) // 测试之后做一些拆卸工作
	os.Exit(retCode)                           // 退出测试
}
```

需要注意的是：在调用TestMain时, flag.Parse并没有被调用。所以如果TestMain 依赖于command-line标志 (包括 testing 包的标记), 则应该显示的调用flag.Parse。

#### 4.3.7、子测试的Setup与Teardown

有时候我们可能需要为每个测试集设置`Setup`与`Teardown`，也有可能需要为每个子测试设置Setup与Teardown。下面我们定义两个函数工具函数如下：

```go
// 测试集的Setup与Teardown
func setupTestCase(t *testing.T) func(t *testing.T) {
	t.Log(&quot;如有需要在此执行:测试之前的setup&quot;)
	return func(t *testing.T) {
		t.Log(&quot;如有需要在此执行:测试之后的teardown&quot;)
	}
}

// 子测试的Setup与Teardown
func setupSubTest(t *testing.T) func(t *testing.T) {
	t.Log(&quot;如有需要在此执行:子测试之前的setup&quot;)
	return func(t *testing.T) {
		t.Log(&quot;如有需要在此执行:子测试之后的teardown&quot;)
	}
}
```

使用方式如下：

```go
func TestSplit(t *testing.T) {
	type test struct { // 定义test结构体
		input string
		sep   string
		want  []string
	}
	tests := map[string]test{ // 测试用例使用map存储
		&quot;simple&quot;:      {input: &quot;a:b:c&quot;, sep: &quot;:&quot;, want: []string{&quot;a&quot;, &quot;b&quot;, &quot;c&quot;}},
		&quot;wrong sep&quot;:   {input: &quot;a:b:c&quot;, sep: &quot;,&quot;, want: []string{&quot;a:b:c&quot;}},
		&quot;more sep&quot;:    {input: &quot;abcd&quot;, sep: &quot;bc&quot;, want: []string{&quot;a&quot;, &quot;d&quot;}},
		&quot;leading sep&quot;: {input: &quot;沙河有沙又有河&quot;, sep: &quot;沙&quot;, want: []string{&quot;&quot;, &quot;河有&quot;, &quot;又有河&quot;}},
	}
	teardownTestCase := setupTestCase(t) // 测试之前执行setup操作
	defer teardownTestCase(t)            // 测试之后执行testdoen操作

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			teardownSubTest := setupSubTest(t) // 子测试之前执行setup操作
			defer teardownSubTest(t)           // 测试之后执行testdoen操作
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf(&quot;excepted:%#v, got:%#v&quot;, tc.want, got)
			}
		})
	}
}
```

测试结果如下：

```go
split $ go test -v
=== RUN   TestSplit
=== RUN   TestSplit/simple
=== RUN   TestSplit/wrong_sep
=== RUN   TestSplit/more_sep
=== RUN   TestSplit/leading_sep
--- PASS: TestSplit (0.00s)
    split_test.go:71: 如有需要在此执行:测试之前的setup
    --- PASS: TestSplit/simple (0.00s)
        split_test.go:79: 如有需要在此执行:子测试之前的setup
        split_test.go:81: 如有需要在此执行:子测试之后的teardown
    --- PASS: TestSplit/wrong_sep (0.00s)
        split_test.go:79: 如有需要在此执行:子测试之前的setup
        split_test.go:81: 如有需要在此执行:子测试之后的teardown
    --- PASS: TestSplit/more_sep (0.00s)
        split_test.go:79: 如有需要在此执行:子测试之前的setup
        split_test.go:81: 如有需要在此执行:子测试之后的teardown
    --- PASS: TestSplit/leading_sep (0.00s)
        split_test.go:79: 如有需要在此执行:子测试之前的setup
        split_test.go:81: 如有需要在此执行:子测试之后的teardown
    split_test.go:73: 如有需要在此执行:测试之后的teardown
=== RUN   ExampleSplit
--- PASS: ExampleSplit (0.00s)
PASS
ok      github.com/Q1mi/studygo/code_demo/test_demo/split       0.006s
```

### 4.4、示例函数

#### 4.4.1、示例函数的格式

被go test特殊对待的第三种函数就是示例函数，它们的函数名以Example为前缀。它们既没有参数也没有返回值。标准格式如下：

```go
func ExampleName() {
    // ...
}
```

#### 4.4.2、示例函数示例

下面的代码是我们为Split函数编写的一个示例函数：

```go
func ExampleSplit() {
	fmt.Println(split.Split(&quot;a:b:c&quot;, &quot;:&quot;))
	fmt.Println(split.Split(&quot;沙河有沙又有河&quot;, &quot;沙&quot;))
	// Output:
	// [a b c]
	// [ 河有 又有河]
}
```

为你的代码编写示例代码有如下三个用处：

示例函数能够作为文档直接使用，例如基于web的godoc中能把示例函数与对应的函数或包相关联。

示例函数只要包含了// Output:也是可以通过go test运行的可执行测试。

```go
split $ go test -run Example
PASS
ok      github.com/Q1mi/studygo/code_demo/test_demo/split       0.006s
```