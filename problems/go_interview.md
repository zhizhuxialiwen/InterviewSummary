# golang 语言面试题

## 1、golang面试基础题

https://www.jishuchi.com/read/go-interview/3435

### 1.1、在go语言中，new和make的区别？

- new

1. new 的作用是初始化一个指向类型的指针(*T)
2. new函数是内建函数，函数定义：func new(Type) *Type
3. 使用new函数来分配空间。传递给new 函数的是一个类型，不是一个值。返回值是指向这个新分配的0值的指针。

- make

1. make 的作用是为 slice，map 或 chan 初始化并返回引用(T)。
2. make函数是内建函数，函数定义：func make(Type, size IntegerType) Type

    > 1）第一个参数是一个类型，第二个参数是长度
    > 2）返回值是一个类型

make(T, args)函数的目的与new(T)不同。它仅仅用于创建 Slice, Map 和 Channel，并且返回类型是 T（不是T*）的一个初始化的（不是零值）的实例。

### 1.2、 在go语言中，Printf()、Sprintf()、Fprintf()函数的区别用法是什么？

都是把格式好的字符串输出，只是输出的目标不一样：

1. Printf()，是把格式字符串输出到标准输出（一般是屏幕，可以重定向）。Printf() 是和标准输出文件(stdout)关联的,Fprintf 则没有这个限制.
2. Sprintf()，是把格式字符串输出到指定字符串中，所以参数比printf多一个char*。那就是目标字符串地址。
3. Fprintf()， 是把格式字符串输出到指定文件设备中，所以参数笔printf多一个文件指针FILE*。主要用于文件操作。Fprintf()是格式化输出到一个stream，通常是到文件。

### 1.3、说说go语言中，数组与切片的区别？

1. 数组 

1）数组是具有固定长度且拥有零个或者多个相同数据类型元素的序列。 数组的长度是数组类型的一部分，所以[3]int 和 [4]int 是两种不同的数组类型。
2）数组需要指定大小，不指定也会根据初始化的自动推算出大小，不可改变 ;
3）数组是值传递;
4）数组是内置(build-in)类型,是一组同类型数据的集合，它是值类型，通过从0开始的下标索引访问元素值。在初始化后长度是固定的，无法修改其长度。当作为方法的参数传入时将复制一份数组而不是引用同一指针。数组的长度也是其类型的一部分，通过内置函数len(array)获取其长度。
5）数组定义：

```c++
var array [10]int
var array = [5]int{1,2,3,4,5}
```

2. 切片 

1）切片表示一个拥有相同类型元素的可变长度的序列。 切片是一种轻量级的数据结构，它有三个属性：指针、长度和容量。
2）切片不需要指定大小;
3）切片是地址传递;
4）切片可以通过数组来初始化，也可以通过内置函数make()初始化 .初始化时len=cap,在追加元素时如果容量cap不足时将按len的2倍扩容；
5）切片定义：
var slice []type = make([]type, len)

### 1.4、解释以下命令的作用？

1. go env: #用于查看go的环境变量
2. go run: #用于编译并运行go源码文件
3. go build: #用于编译源码文件、代码包、依赖包
4. go get: #用于动态获取远程代码包
5. go install: #用于编译go文件，并将编译结构安装到bin、pkg目录
6. go clean: #用于清理工作目录，删除编译和安装遗留的目标文件
7. go version: #用于查看go的版本信息

### 1.5、说说go语言中的协程？

1）协程和线程都可以实现程序的并发执行；
2）通过channel来进行协程间的通信；
3）只需要在函数调用前添加go关键字即可实现go的协程，创建并发任务；
4）关键字go并非执行并发任务，而是创建一个并发任务单元；

### 1.6、说说go语言中的协程与线程区别？

1. 协程

协程，英文名Coroutine。但在 Go 语言中，协程的英文名是：gorutine。它常常被用于进行多任务，即并发作业。没错，就是多线程作业的那个作业。

虽然在 Go 中，我们不用直接编写线程之类的代码来进行并发，但是 Go 的协程却依赖于线程来进行。

2. 协程与线程的区别

- 协程的特点

1）多个协程可由一个或多个线程管理，协程的调度发生在其所在的线程中。
2）可以被调度，调度策略由应用层代码定义，即可被高度自定义实现。
3）执行效率高。
4）占用内存少。

上面第 1 和 第 2 点
我们来看一个例子：

```go
func TestGorutine(t *testing.T) {
	runtime.GOMAXPROCS(1)  // 指定最大 P 为 1，从而管理协程最多的线程为 1 个
	wg := sync.WaitGroup{} // 控制等待所有协程都执行完再退出程序
	wg.Add(2)
	// 运行一个协程
	go func() {
		fmt.Println(1)
		fmt.Println(2)
		fmt.Println(3)
		wg.Done()
	}()

	// 运行第二个协程
	go func() {
		fmt.Println(65)
		fmt.Println(66)
		// 设置个睡眠，让该协程执行超时而被挂起，引起超时调度
		time.Sleep(time.Second)
		fmt.Println(67)
		wg.Done()
	}()
	wg.Wait()
}
```

上面的代码片段跑了两个协程，运行后，观察输出的顺序是交错的。可能是：

```
65
66
1
2
3
67
```

意味着在执行协程A的过程中，可以随时中断，去执协程行B，协程B也可能在执行过程中中断再去执行协程A。

看起来协程A 和 协程B 的运行像是线程的切换，但是请注意，这里的 A 和 B 都运行在同一个线程里面。它们的调度不是线程的切换，而是纯应用态的协程调度。

3. 关于上述代码中，为什么要指定下面两行代码？

```go
runtime.GOMAXPROCS(1)
time.Sleep(time.Second)
```

这需要您去看下 Go 的协程调度入门基础，请看我之前的另外一篇调度分析文章：

- Go 的协程调度机制

如果不设置 runtime.GOMAXPROCS(1)，那么程序将会根据操作系统的 CPU 核数而启动对应数量的 P，导致多个 M，即线程的启动。那么我们程序中的协程，就会被分配到不同的线程里面去了。为了演示，故设置数量 1，使得它们都被分配到了同一个线程里面，存于线程的协程队列里面，等待被执行或调度。

协程特点中的第 3 和 第 4 点。
执行效率高。
占用内存少。
因为协程的调度切换不是线程切换，而是由程序自身控制，因此，没有线程切换的开销，和多线程比，线程数量越多，协程的性能优势就越明显。调度发生在应用态而非内核态。

内存的花销，使用其所在的线程的内存，意味着线程的内存可以供多个协程使用。

其次协程的调度不需要多线程的锁机制，因为只有一个线程，也不存在同时写变量冲突，所以执行效率比多线程高很多。

4. 协程和线程的整体对比

|序号|比较的点|	线程|	协程|
|:---|:-------|:----|:-----|
|1   |数据存储|	内核态的内存空间|一般是线程提供的用户态内存空间|
|2   |切换操作|操作最终在内核层完成，应用层需要调用内核层提供的, syscall:底层函数	|应用层使用代码进行简单的现场保存和恢复即可|
|3   |任务调度|由内核实现，抢占方式，依赖各种锁|由用户态的实现的具体调度器进行。例如 go 协程的调度器|
|4   |语音支持程度|绝大部分编程语言|部分语言：Lua，Go，Python ...|
|5   |实现规范|按照现代操作系统规范实现|无统一规范。在应用层由开发者实现，高度自定义，比如只支持单线程的线程。不同的调度策略，等等|

### 1.7、说说go语言中的for循环？

for循环支持continue和break来控制循环，但是它提供了一个更高级的break，可以选择中断哪一个循环,for循环不支持以逗号为间隔的多个赋值语句，必须使用平行赋值的方式来初始化多个变量

### 1.8、说说go语言中的switch语句？

单个case中，可以出现多个结果选项

只有在case中明确添加fallthrough关键字，才会继续执行紧跟的下一个case

### 1.9、go语言中没有隐藏的this指针，这句话是什么意思？

方法施加的对象显式传递，没有被隐藏起来
golang的面向对象表达更直观，对于面向过程只是换了一种语法形式来表达
方法施加的对象不需要非得是指针，也不用非得叫this

### 1.10、 go语言中的引用类型包含哪些？

数组切片、字典(map)、通道（channel）、接口（interface）

### 1.11、go语言中指针运算有哪些？

可以通过“&”取指针的地址
可以通过“*”取指针指向的数据

### 1.12、说说go语言的main函数

1. main函数不能带参数
2. main函数不能定义返回值
3. main函数所在的包必须为main包
4. main函数中可以使用flag包来获取和解析命令行参数

### 1.13、说说go语言的同步锁？

1. 当一个goroutine获得了Mutex后，其他goroutine就只能乖乖的等待，除非该goroutine释放这个Mutex;
2. RWMutex在读锁占用的情况下，会阻止写，但不阻止读
3. RWMutex在写锁占用情况下，会阻止任何其他goroutine（无论读和写）进来，整个锁相当于由该goroutine独占

### 1.14、说说go语言的channel特性？

1. 给一个 nil channel 发送数据，造成永远阻塞
2. 从一个 nil channel 接收数据，造成永远阻塞
3. 给一个已经关闭的 channel 发送数据，引起 panic
4. 从一个已经关闭的 channel 接收数据，如果缓冲区中为空，则返回一个零值
5. 无缓冲的channel是同步的，而有缓冲的channel是非同步的

### 1.15、 go语言触发异常的场景有哪些？

1. 空指针解析
2. 下标越界
3. 除数为0
4. 调用panic函数

### 1.16、说说go语言的beego框架？

1. beego是一个golang实现的轻量级HTTP框架
2. beego可以通过注释路由、正则路由等多种方式完成url路由注入
3. 可以使用bee new工具生成空工程，然后使用bee run命令自动热编译

### 1.17、说说go语言的goconvey框架？

1. goconvey是一个支持golang的单元测试框架
2. goconvey能够自动监控文件修改并启动测试，并可以将测试结果实时输出到web界面
3. goconvey提供了丰富的断言简化测试用例的编写

### 1.18、go语言中，GoStub的作用是什么？

GoStub是一款轻量级的单元测试框架，接口友好，可以对全局变量、函数或过程进行打桩。

1. GoStub可以对全局变量打桩
2. GoStub可以对函数打桩
3. GoStub不可以对类的成员方法打桩
4. GoStub可以打动态桩，比如对一个函数打桩后，多次调用该函数会有不同的行为

### 1.19、golang有哪些测试框架？其区别

1. GoConvey
2. GoStub
3. GoMock

### 1.20、说说go语言的select机制？

1. select机制用来处理异步IO问题
2. select机制最大的一条限制就是每个case语句里必须是一个IO操作
3. golang在语言级别支持select关键字

### 1.21、说说进程、线程、协程之间的区别？

- 进程与线程区别

1. 进程是资源的分配和调度的一个独立单元，而线程是CPU调度的基本单元；
2. 同一个进程中可以包括多个线程；
3. 进程结束后它拥有的所有线程都将销毁，而线程的结束不会影响同个进程中的其他线程的结束；
4. 线程共享整个进程的资源（寄存器、堆栈、上下文），一个进程至少包括一个线程；
5. 进程的创建调用fork或者vfork，而线程的创建调用pthread_create；
6. 线程中执行时一般都要进行同步和互斥，因为他们共享同一进程的所有资源；
7. 进程是资源分配的单位，线程是操作系统调度的单位。

- 进程、线程与协程区别

1. 进程切换需要的资源很最大，效率很低；
线程切换需要的资源一般，效率一般 ；
协程切换任务资源很小，效率高。
2. 多进程、多线程根据cpu核数不一样,可能是并行的,也可能是并发的。
协程的本质就是使用当前进程在不同的函数代码中切换执行，可以理解为并行。 协程是一个用户层面的概念，不同协程的模型实现可能是单线程，也可能是多线程。
3. 进程拥有自己独立的堆和栈，既不共享堆，亦不共享栈，进程由操作系统调度。（全局变量保存在堆中，局部变量及函数保存在栈中）
线程拥有自己独立的栈和共享的堆，共享堆，不共享栈，线程亦由操作系统调度(标准线程是这样的)。
协程和线程一样共享堆，不共享栈，协程由程序员在协程的代码里显示调度。

4. 一个应用程序一般对应一个进程，一个进程一般有一个主线程，还有若干个辅助线程，线程之间是平行运行的，在线程里面可以开启协程，让程序在特定的时间内运行。

- 线程与协程区别

协程和线程的区别是：协程避免了无意义的调度，由此可以提高性能，但也因此，程序员必须自己承担调度的责任，同时，协程也失去了标准线程使用多CPU的能力。

### 1.22、map如何顺序读取

map不能顺序读取，是因为他是无序的，想要有序读取，首先的解决的问题就是，把key变为有序，所以可以把key放入切片，对切片进行排序，遍历切片，通过key取值。

### 1.23、Go语言的特性有哪些？

Go语言也称为 Golang，是由 Google 公司开发的一种静态强类型、编译型、并发型、并具有垃圾回收功能的编程语言。

接下来从几个方面来具体介绍一下Go语言的特性。

1. 语法简单

go语言将“++”、“--”从运算符降级为语句，保留指针，但默认阻止指针运算，带来的好处是显而易见的。还有，将切片和字典作为内置类型，从运行时的层面进行优化，这也算是一种“简单”。

2. 并发模型

可以说，Goroutine 是 Go 最显著的特征。它用类协程的方式来处理并发单元，却又在运行时层面做了更深度的优化处理。这使得语法上的并发编程变得极为容易，无须处理回调，无须关注线程切换，仅一个关键字，简单而自然。

搭配 channel，实现 CSP 模型。将并发单元间的数据耦合拆解开来，各司其职，这对所有纠结于内存共享、锁粒度的开发人员都是一个可期盼的解脱。若说有所不足，那就是应该有个更大的计划，将通信从进程内拓展到进程外，实现真正意义上的分布式。

3. 内存分配

将一切并发化固然是好，但带来的问题同样很多。如何实现高并发下的内存分配和管理就是个难题。好在 Go 选择了 **tcmalloc(Thread-Caching Malloc)，** 它本就是为并发而设计的高性能内存分配组件。

TCMalloc(Thread-Caching Malloc)与标准glibc库的malloc实现一样的功能，但是TCMalloc在效率和速度效率都比标准malloc高很多。

4. 垃圾回收

垃圾回收一直是个难题。早年间，Java 就因垃圾回收低效被嘲笑了许久，后来 Sun 连续收纳了好多人和技术才发展到今天。可即便如此，在 Hadoop 等大内存应用场景下，垃圾回收依旧捉襟见肘、步履维艰。

相比 Java，Go 面临的困难要更多。因指针的存在，所以回收内存不能做收缩处理。幸好，指针运算被阻止，否则要做到精确回收都难。

每次升级，垃圾回收器必然是核心组件里修改最多的部分。从并发清理，到降低 STW 时间，直到 Go 的 1.5 版本实现并发标记，逐步引入**三色标记和写屏障等等，** 都是为了能让垃圾回收在不影响用户逻辑的情况下更好地工作。尽管有了努力，当前版本的垃圾回收算法也只能说堪用，离好用尚有不少距离。

5. 静态链接

Go 刚发布时，静态链接被当作优点宣传。只须编译后的一个可执行文件，无须附加任何东西就能部署。这似乎很不错，只是后来风气变了。连着几个版本，编译器都在完善动态库 buildmode 功能，场面一时变得有些尴尬。

暂不说未完工的 buildmode 模式，静态编译的好处显而易见。**将运行时、依赖库直接打包到可执行文件内部，简化了部署和发布操作，无须事先安装运行环境和下载诸多第三方库。** 这种简单方式对于编写系统软件有着极大好处，因为库依赖一直都是个麻烦。

6. 标准库

功能完善、质量可靠的标准库为编程语言提供了充足动力。在不借助第三方扩展的情况下，就可完成大部分基础功能开发，这大大降低了学习和使用成本。最关键的是，标准库有升级和修复保障，还能从运行时获得深层次优化的便利，这是第三方库所不具备的。

Go 标准库虽称不得完全覆盖，但也算极为丰富。其中值得称道的是 net/http，仅须简单几条语句就能实现一个高性能 Web Server，这从来都是宣传的亮点。更何况大批基于此的优秀第三方 Framework 更是将 Go 推到 Web/Microservice 开发标准之一的位置。

当然，优秀第三方资源也是语言生态圈的重要组成部分。近年来崛起的几门语言中，Go 算是独树一帜，大批优秀作品频繁涌现，这也给我们学习 Go 提供了很好的参照。

7. 工具链

完整的工具链对于日常开发极为重要。Go 在此做得相当不错，无论是编译、格式化、错误检查、帮助文档，还是第三方包下载、更新都有对应的工具。其功能未必完善，但起码算得上简单易用。

1）Goconvey测试框架
2）内置完整测试框架，其中包括单元测试、性能测试、代码覆盖率、数据竞争，以及用来调优的 pprof，这些都是保障代码能正确而稳定运行的必备利器。

除此之外，还可通过环境变量输出运行时监控信息，尤其是垃圾回收和并发调度跟踪，可进一步帮助我们改进算法，获得更佳的运行期表现。

8. 更丰富的内置类型

其实作为一种新兴的语言，如果仅仅是为了某种特定的用途那么可能其内置类型不是很多，仅需要能够完成我的功能即可，但是Go语言“不仅支持几乎所有语言都支持的简单内置类型（比如整型和浮点型等）外，还支持一些其他的高级类型，比如字典类型，map要知道这些类型在其他语言中都是通过包的形式引入的外部数据类型。数组切片（Slice），类似于C++ STL中的vector，在Go也是一种内置的数据类型作为动态数组来使用。这里满有一个颇为简单的解释：”既然绝大多数开发者都需要用到这个类型，为什么还非要每个人都写一行import语句来包含一个库？”

9. 支持函数多返回值

在C，C++中，包括其他的一些高级语言是不支持多个函数返回值的。但是这项功能又确实是需要的，所以在C语言中一般通过将返回值定义成一个结构体，或者通过函数的参数引用的形式进行返回。而在Go语言中，作为一种新型的语言，目标定位为强大的语言当然不能放弃对这一需求的满足，所以支持函数多返回值是必须的，例如：

```go
func getName()(firstName, middleName, lastName, nickName string){
    return "May", "M", "Chen", "Babe" 
} //定义了一个多返回值的函数getName 

fn, mn, ln, nn := getName()      //调用赋值
_, _, lastName, _ := getName() //缺省调用
```

10. 错误处理

在传统的OOP编程中，为了捕获程序的健壮性需要捕获异常，使用的方法大都是try() catch{}模块，例如, 在下面的java代码中，可能需要的操作是：

```java
Connection conn = ...;
try {
    Statement stmt = ...;
    ...//别的一些异常捕获
finally {
    stmt.close();
    }
finally {
    conn.close(); 
}
```

而在Go中引入了三个关键字，分别是 `defer、panic和recover`，其中使用defer关键字语句的含义是不管程序是否出现异常，均在函数退出时自动执行相关代码。
所以上面你的java代码用Go进程重写只有两行：

```java
conn := ...
defer conn.Close()
```

另外两个关键词后面再讨论。所以“Go语言的错误处理机制可以大量减少代码量，让开发者也无需仅仅为了程序安全性而添加大量一层套一层的try-catch语句。这对于代码的阅读者和维护者来说也是一件很好的事情，因为可以避免在层层的代码嵌套中定位业务代码。”

11. 匿名函数和闭包

关于这个功能介绍的不多，大概就是说Go中的函数也可以作为参数进行传递：
“在Go语言中，所有的函数也是值类型，可以作为参数传递。Go语言支持常规的匿名函数和闭包，比如下列代码就定义了一个名为f的匿名函数，开发者可以随意对该匿名函数变量进行传递和调用：

```go
f := func(x, y int) int {
return x + y
}
```

12. 类型和接口

这个特性是Go在实现OPP时候的一些特性，主要有这么几点：

第一，Go语言没有很复杂的面向对象的概念，即没有继承和重载，其类型更像是C中的struct，并且直接使用了struct关键字，仅仅是最基本的类型组合功能。但是，尽管不支持这些语法特性，但是Go的接口却同样可以实现这些功能，只是实现的形式上会有不同而已。

即这里需要介绍的“非侵入型”接口的概念。

举个例子：

在C++中，一般会这样定义一个接口和类型的

```c++
// 抽象接口 
interface IFly 
{ 
	virtual void Fly()=0; 
}; 
// 实现类 
class Bird : public IFly 
{ 
	public: 
	Bird() {} 
	virtual ~Bird() {} 
	void Fly() 
	{ 
	// 以鸟的方式飞行 
	} 
}; 
//使用的时候 
void main() 
{ 
	IFly* pFly = new Bird(); 
	pFly->Fly(); 
	delete pFly; 
}
```

需要你自己以虚函数的形式定义一个接口，并且让类型继承这个接口并重写虚方法。在使用的时候需要进行动态绑定。

而在Go中实现相同的功能，你只需要

```go
type Bird struct { 
    … 
} 
func (b *Bird) Fly() { 
    // 以鸟的方式飞行 
} 
type IFly interface { 
    Fly() 
} 
func main() { 
    var fly IFly = new(Bird) 
    fly.Fly() 
} 
```

可以看出，“虽然Bird类型实现的时候，没有声明与接口IFly的关系，但接口和类型可以直
接转换，甚至接口的定义都不用在类型定义之前，这种比较松散的对应关系可以大幅降低因为接
口调整而导致的大量代码调整工作”。

13. 支持反射

这里的反射(reflecttion)和JAVA中的反射类似，可以用来获取对象类型的相信信息，并动态操作对象。因为反射可能会对程序的可读性有很大的干扰，所以，在Go中只是在特别需要反射支持的地方才实现反射的一些功能。“反射最常见的使用场景是做对象的序列化（serialization，有时候也叫Marshal & Unmarshal）。例如，Go语言标准库的encoding/json、encoding/xml、encoding/gob、encoding/binary等包就大量依赖于反射功能来实现。”

14. 语言的交互性

这里的交互性主要是和C的交互性，之所以这样是因为Go语言的开发者是最初贝尔实验室创建Unix系统以及C语言的一般人，包括：

肯·汤普逊（Ken Thompson，http://en.wikipedia.org/wiki/Ken_Thompson）：设计了B语言和C语言，创建了Unix和Plan 9操作系统，1983年图灵奖得主，Go语言的共同作者。

在Go语言中直接重用了大部份的C模块，这里称为Cgo.Cgo允许开发者混合编写C语言代码，然后Cgo工具可以将这些混合的C代码提取并生成对于C功能的调用包装代码。开发者基本上可以完全忽略这个Go语言和C语言的边界是如何跨越的。

例如书中一个例子，在Go语言中直接调用了C标准库的puts函数。

```go
package main
/*
#include <stdio.h>
*/
import "C"
import "unsafe"
func main() {
    cstr := C.CString("Hello, world")
    C.puts(cstr)
    C.free(unsafe.Pointer(cstr))
}
```


## 2、golang 语言面试题解析

### 2.1 写出下面代码输出内容。

```go
package main
import("fmt") 
func main() {
    defer_call()
}
func defer_call() {
    defer func() {
        fmt.Println("打印前")
    }()
    defer func() {
        fmt.Println("打印中")
    }() 
    defer func() {
        fmt.Println("打印后")
    }() 
    panic("触发异常")
}
```

考点：defer执行顺序

解答： defer 是后进先出。 panic 需要等defer 结束后才会向上传递。出现panic恐慌时候，会先按照defer的后入先出的顺序执行，最后才会执行panic。

输出结果：
打印后
打印中
打印前
panic: 触发异常

### 2.2 以下代码有什么问题，说明原因。

```go
type student struct {
    Name string Age int
}
func pase_student() {
    m := make(map[string] * student)
    stus := [] student {
        {
            Name: "zhou",
            Age: 24
        }, {
            Name: "li",
            Age: 23
        }, {
            Name: "wang",
            Age: 22
        },
    }
    for _, stu := range stus {
        m[stu.Name] = &stu
    }
}
```

考点：foreach

解答：这样的写法初学者经常会遇到的，很危险！与Java的foreach一样，都是使用副本的方式。所以m[stu.Name]=&stu实际上一致指向同一个指针，最终该指针的值为遍历的最后一个struct的值拷贝。就像想修改切片元素的属性：

```go
for _, stu := range stus {
    stu.Age = stu.Age + 10
}
```

也是不可行的。大家可以试试打印出来：

```go
func pase_student() {
    m := make(map[string] * student) 
    stus := [] student {
        {
            Name: "zhou",
            Age: 24
        }, {
            Name: "li",
            Age: 23
        }, {
            Name: "wang",
            Age: 22
        },
    }
    // 错误写法   
    for _, stu := range stus {
        m[stu.Name] = & stu
    }
    for k, v := range m {
        println(k, "=>", v.Name)
    }
    // 正确   
    for i := 0; i < len(stus); i++ {
        m[stus[i].Name] = & stus[i]
    }
    for k, v := range m {
        println(k, "=>", v.Name)
    }
}
```

### 2.3 下面的代码会输出什么，并说明原因

```go
func main() {
    runtime.GOMAXPROCS(1)
    wg := sync.WaitGroup {}
    wg.Add(20)
    for i := 0; i < 10; i++ {
        go func() {
            fmt.Println("A: ", i) 
            wg.Done()
        }()
    }
    for i := 0; i < 10; i++ {
        go func(i int) {
            fmt.Println("B: ", i)
            wg.Done()
        }(i)
    }
	wg.Wait()
```

考点：go执行的随机性和闭包

解答：谁也不知道执行后打印的顺序是什么样的，所以只能说是随机数字。但是A:均为输出10，B:从0~9输出(顺序不定)。

第一个go func中i是外部for的一个变量，地址不变化。遍历完成后，最终i=10。故go func执行时，i的值始终是10。

第二个go func中i是函数参数，与外部for中的i完全是两个变量。尾部(i)将发生值拷贝，go func内部指向值拷贝地址。

### 2.4 下面代码会输出什么？

```go
type People struct {}
func(p * People) ShowA() {
    fmt.Println("showA")
    p.ShowB()
}
func(p * People) ShowB() {
    fmt.Println("showB")
}
type Teacher struct {
    People
}
func(t * Teacher) ShowB() {
    fmt.Println("teacher showB")
}
func main() {
    t := Teacher {}
    t.ShowA()
}
```

考点：go的组合继承

解答：这是Golang的组合模式，可以实现OOP(Object Oriented Programming, 面向对象的程序设计)的继承。被组合的类型People所包含的方法虽然升级成了外部类型Teacher这个组合类型的方法（一定要是匿名字段），但它们的方法(ShowA())调用时接受者并没有发生变化。此时People类型并不知道自己会被什么类型组合，当然也就无法调用方法时去使用未知的组合者Teacher类型的功能。

输出结果：
showA
showB

### 2.5 下面代码会触发异常吗？请详细说明

```go
func main() {
    runtime.GOMAXPROCS(1) 
    int_chan := make(chan int, 1) 
    string_chan := make(chan string, 1) 
    int_chan <- 1 
    string_chan <- "hello"
    select {
        case value := <-int_chan:
            fmt.Println(value)
        case value := <-string_chan:
            panic(value)
    }
}
```

考点：select随机性

解答： select会随机选择一个可用通用做收发操作。**所以代码是有肯触发异常，也有可能不会。** 单个chan如果无缓冲时，将会阻塞。但结合 select可以在多个chan间等待执行。有三点原则：

1）当select 中只要有一个case能return，则立刻执行。 
2）当同一时间有多个case均能return则伪随机方式抽取任意一个执行。
3）如果没有一个case能return，则可以执行”default”块。

### 2.6 下面代码输出什么？

```go
func calc(index string, a, b int) int {
    ret := a + b 
    fmt.Println(index, a, b, ret) 
    return ret
}
func main() {
    a := 1
    b := 2 
    defer calc("1", a, calc("10", a, b))
    a = 0 
    defer calc("2", a, calc("20", a, b)) 
    b = 1
}
```

考点：defer执行顺序

解答：这道题类似第1题需要注意到defer执行顺序和值传递 index:1肯定是最后执行的，但是index:1的第三个参数是一个函数，所以最先被调用`calc("10",1,2)==>10,1,2,3` 执行index:2时,与之前一样，需要先调用`calc("20",0,2)==>20,0,2,2` 执行到b=1时候开始调用，`index:2==>calc("2",0,2)==>2,0,2,2` 最后执行`index:1==>calc("1",1,3)==>1,1,3,4`

输出结果：
10 1 2 3
20 0 2 2 
2 0 2 2
1 1 3 4

### 2.7 请写出以下输入内容

```go
func main() {
    s := make([] int, 0)
    s = append(s, 1, 2, 3)
    fmt.Println(s)
}
```

考点：make默认值和append

解答： make初始化是由默认值的哦，此处默认值为0,append以2倍内存空间进行扩容。

输出结果：
[0 0 0 0 0 1 2 3]

大家试试改为:

s := make([] int, 0) 
s = append(s, 1, 2, 3) 
fmt.Println(s) //[1 2 3]

### 2.8 下面的代码有什么问题?

```go
type UserAges struct {
    ages map[string] int 
    sync.Mutex
}
func(ua * UserAges) Add(name string, age int) {
    ua.Lock()
    defer ua.Unlock()
    ua.ages[name] = age
}
func(ua * UserAges) Get(name string) int {
    if age, ok := ua.ages[name]; ok {
        return age
    }
    return -1
}
```

考点：map线程安全

解答：可能会出现`fatal error: concurrent map read and map write.` 修改一下看看效果

```go
func(ua * UserAges) Get(name string) int {
    ua.Lock() 
    defer ua.Unlock() 
    if age, ok := ua.ages[name]; ok {
        return age
    }
    return -1
}
```

### 2.9 下面的迭代会有什么问题？

```go
func(set * threadSafeSet) Iter() <-chan interface {} {
    ch := make(chan interface {})
    go func() {
        set.RLock()
        for elem := range set.s {
            ch <-elem
        }
        close(ch) 
        set.RUnlock()
    }()
    return ch
}
```

考点：chan缓存池

解答：看到这道题，我也在猜想出题者的意图在哪里。 chan?sync.RWMutex?go?chan缓存池?迭代? 所以只能再读一次题目，就从迭代入手看看。既然是迭代就会要求set.s全部可以遍历一次。但是chan是为缓存的，那就代表这写入一次就会阻塞。我们把代码恢复为可以运行的方式，看看效果

package main
import(
    "sync"
    "fmt"
)
//下面的迭代会有什么问题？
type threadSafeSet struct {
    sync.RWMutex 
    s []interface {}
}
func(set * threadSafeSet) Iter() <-chan interface {} {
    // ch := make(chan interface{}) 
    // 解除注释看看！ 
    ch := make(chan interface {}, len(set.s)) 
    go func() {
        set.RLock()
        for elem, value := range set.s {
            ch <-elem 
            println("Iter:", elem, value)
        }
        close(ch)
        set.RUnlock()
    }()
    return ch
}
func main() {
    th := threadSafeSet {
        s :[]interface {} {
            "1", "2"
        },
    }
    v := <-th.Iter()
    fmt.Sprintf("%s%v", "ch", v)
}
10. 以下代码能编译过去吗？为什么？
package main
import (
    "fmt"
)
type People interface {
    Speak(string) string
}
type Stduent struct {}
func(stu * Stduent) Speak(think string)(talk string) {
    if think == "bitch" {
        talk = "You are a good boy"
    } else {
        talk = "hi"
    }
    return
}
func main() {
    var peo People = Stduent {}
    think := "bitch"
    fmt.Println(peo.Speak(think))
}
考点：golang的方法集

解答：编译不通过！做错了！？说明你对golang的方法集还有一些疑问。一句话：golang的方法集仅仅影响接口实现和方法表达式转化，与通过实例或者指针调用方法无关。

11. 以下代码打印出来什么内容，说出为什么。
package main
import(
    "fmt"
) 
type People interface {
    Show()
}
type Student struct {}
func(stu * Student) Show() {}
func live() People {
    var stu * Student
    return stu
}
func main() {
    if live() == nil {
        fmt.Println("AAAAAAA")
    } else {
        fmt.Println("BBBBBBB")
    }
}
考点：interface内部结构

解答：很经典的题！这个考点是很多人忽略的interface内部结构。 go中的接口分为两种一种是空的接口类似这样：

var in interface{}
另一种如题目：

type People interface {
    Show()
}
他们的底层结构如下：

type eface struct {      
    //空接口   
    _type *_type         //类型信息    
    data  unsafe.Pointer //指向数据的指针(go语言中特殊的指针类型unsafe.Pointer类似于c语言中的void*)
}
type iface struct {      
    //带有方法的接口    
    tab  *itab           //存储type信息还有结构实现方法的集合    
    data unsafe.Pointer  //指向数据的指针(go语言中特殊的指针类型unsafe.Pointer类似于c语言中的void*)}
type _type struct {   
    size       uintptr  //类型大小    
    ptrdata    uintptr  //前缀持有所有指针的内存大小    
    hash       uint32   //数据hash值    
    tflag      tflag    align      uint8    //对齐    
    fieldalign uint8    //嵌入结构体时的对齐    
    kind       uint8    //kind 有些枚举值kind等于0是无效的    
    alg        *typeAlg //函数指针数组，类型实现的所有方法    
    gcdata    *byte    str       nameOff    ptrToThis typeOff
}
type itab struct {    
    inter  *interfacetype  //接口类型   
    _type  *_type          //结构类型    
    link   *itab    bad    int32    inhash int32    fun    [1]uintptr      //可变大小 方法集合
}
可以看出iface比eface 中间多了一层itab结构。 itab 存储_type信息和[]fun方法集，从上面的结构我们就可得出，因为data指向了nil 并不代表interface 是nil，所以返回值并不为空，这里的fun(方法集)定义了接口的接收规则，在编译的过程中需要验证是否实现接口结果：

BBBBBBB

