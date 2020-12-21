# golang的高级特性

## 1、函数

### 1.1、函数声明

函数声明包括函数名、形式参数列表、返回值列表（可省略）以及函数体。

```go
func name(parameter-list) (result-list) {
    body
}
```

1. 函数

```go
func sumFn(x int, y int) int{
    sum := x + y
    return sum
}
```

2. 简写类型的函数:由最后一个参数类型决定前面参数类型

```go
func subFn(x, y int) int{
    sub := x -y
    return sub
}
```

3. 可变参数的函数:切片，即输入任意个参数

```go
func sumFn1(x ...int) int {
    fmt.Printf("%v--%T \n",x,x) //[4 5 6 7]--[]int
    sum := 0
    for _, v := range x{
        sum += v
    }
    return sum
}
```

### 1.2、递归函数

递归函数：1、自己调用自己；2、结束条件；3、递减

```go
func recursionSum(n int) int{
    if n > 1{
    /* 第一次：100 + recursion(99);第二次：100 + 99 + recursion(98);...；100+99+...+2+recursionSum(2) */
        return n + recursionSum(n - 1) 
    }else{
        return 1  //100+99+...+2+1
    }
} 
```

### 1.3 函数多个返回值

1. 函数返回值：return 返回多个值

```go
func sum_sub(x,y int) (int,int){
    sum := x + y
    sub := x - y
    return sum,sub
}

func main{
    x := 3
    y := 8
    sum,sub := sum_sub(x,y)
    fmt.Println(sum,sub) //11 -5
}
输出结果：11 -5
```

2. 返回值类型变量

```go
func sum_sub1(x,y int) (sum int, sub int){
    sum = x + y
    sub = x - y
    return sum,sub
}
```

3. 前面返回值参数可以不写类型，但是后面返回值参数必须写类型，反之不行

```go
func sum_sub2(x,y int) (sum , sub int){
    sum = x + y
    sub = x - y
    return sum,sub
}
```

### 1.4、自定义函数类型、参数的函数、返回值的函数、匿名函数

#### 1.4.1、自定义函数类型

```go
type calcType func(int, int)int
func add(x,y int)int{
    return x + y
}

func sub(x,y int)int{
    return x - y
}

func main(){
    var typeFunc calcType
    typeFunc = add
    fmt.Printf("typeFunc的类型： %T \n", typeFunc) 
    sum := typeFunc(5,6)
    fmt.Println(sum)  //11
}
输出结果：
typeFunc的类型： main.calc
11
```

### 1.4.2、参数的函数 

函数作为另一个函数的参数 : 正常参数或自定义函数

```go
func calcSum(x,y int, cb calcType )int{
    return cb(x,y)
}

func calcSub(x,y int, cb1 func(int, int)int)int{
    return cb1(x,y)
}

func main(){
    sum1 := calcSum(33, 44, add)
    fmt.Println(sum1)  //77
    sub1 := calcSub(33, 22,sub)
    fmt.Println(sub1)  //11
}

输出结果：
77
11
```

### 1.4.3、返回值的函数 

```go
func doFuncType(outType string) calcType{
    switch outType{
    case "+":
        return add
    case "-":
        return sub
    case "*":
        return func(x,y int)int{
            return x * y
        }
    default:
        return nil
    }
}

func main(){
    var outFunc = doFuncType("+")
    fmt.Println(outFunc(3,4)) //7
}
输出结果：
7
```

### 1.4、匿名函数

1. 匿名函数

```go
var fn = func(x,y int) int{
        return x + y
}
fmt.Println(fn(3,4))
输出结果：
7
```

2. 匿名执行函数

```go
func(x,y int) {
    fmt.Println("test...")
    fmt.Println(x+y)
}(3,5)

输出结果：
8
```

### 1.5、闭包

全局变量特点：1、常驻内存；2、污染全局，即全局与局部有相同的变量
局部变量特点：1、不常住内存；2、不污染全局
**闭包特点：1、可以让一个变量常驻内存；2、可以让一个变量不污染全局**
(1) 闭包有权访问另一个函数作用域中的变量的函数
(2) 一个函数访问另一个函数
写法：函数里面嵌套一个函数，最后返回里面的函数

闭包的概念：闭包可以理解成“定义在一个函数内部的函数“。在本质上，闭包是将函数内部和函数外部连接起来的桥梁。或者说是函数和其引用环境的组合体。

Go语言中的闭包同样也会引用到函数外的变量。闭包的实现确保只要闭包还被使用，那么被闭包引用的变量会一直存在。

- 例1

```go
//函数片段
func add(base int) func(int) int {
    fmt.Printf("%p\n", &base)  //打印变量地址，可以看出来 内部函数时对外部传入参数的引用
    f := func(i int) int {
        //base变量常驻内存
        fmt.Printf("%p\n", &base)
        base += i
        return base
    }
    return f
}

//由 main 函数作为程序入口点启动
func main() {
    t1 := add(10)
    fmt.Println(t1(1), t1(2))
    t2 := add(100)
    fmt.Println(t2(1), t2(2))
}
输出结果：
0xc000064068
0xc000064068
0xc000064068
11 13
0xc0000640c0
0xc0000640c0
0xc0000640c0
101 103
```

内部函数是对外部变量引用。

- 例2

```go
func adder1() func() int {
    var i = 10
    return func() int {
        // i变量常驻内存没有累加变化
        return i + 1
    }
}
func adder2() func(y int) int {
    var i = 10
    return func(y int) int {
        // i变量常驻内存累加变化
        i += y
        return i
    }
}

func main(){
    var fn1 = adder1() //表示执行方法
    fmt.Println(fn1()) //11
    fmt.Println(fn1()) //11
    var fn2 = adder2() //表示执行方法
    fmt.Println(fn2(10)) //20
    fmt.Println(fn2(10)) //30
}
输出结果：
11
11
20
30
```

- 例3，延迟调用参数在求值或复制，指针或闭包会 "延迟" 读取。

```go
package main

func test() {
    x, y := "乔峰", "慕容复"

    defer func(s string) {
        println("defer:", s, y) // y 闭包引用 输出延迟和的值，即y+= 后的值=慕容复第二
    }(x) // 匿名执行函数调用，传送参数x 被复制,注意这里的x 是 乔峰,而不是下面的 x+= 后的值

    x += "第一"
    y += "第二"
    println("x =", x, "y =", y)
}

func main() {
    test()
}
输出结果：
x = 乔峰第一 
y = 慕容复第二
defer: 乔峰 慕容复第二
```

- 总结：

闭包并不是一门编程语言不可缺少的功能，但闭包的表现形式**一般是以匿名函数**的方式出现，就象上面说到的，能够动态灵活的创建以及传递，体现出函数式编程的特点。所以在一些场合，我们就多了一种编码方式的选择，适当的使用闭包可以使得我们的**代码简洁高效**。

- 使用闭包的注意点：

由于闭包会使得函数中的变量都被保存在内存中，内存消耗很大，所以不能滥用闭包

### 1.6、defer关键字

defer的延迟调用。

- defer特性：

1. 关键字 defer 用于注册延迟调用。
2. 这些调用直到 return 前才被执。因此，可以用来做资源清理。
3. 多个defer语句，按先进后出的方式执行。
4. defer语句中的变量，在defer声明时就决定了。
5. defer、return、返回值三者的执行逻辑应该是：return最先执行，return负责将结果写入返回值中；接着defer开始执行一些收尾工作；最后函数携带当前返回值退出。

- 结论：

1）a()int 函数的返回值没有被提前声名，其值来自于其他变量的赋值，而defer中修改的也是其他变量，而非返回值本身，因此函数退出时返回值并没有被改变。

2）b()(i int) 函数的返回值被提前声名，也就意味着defer中是可以调用到真实返回值的，因此defer在return赋值返回值 i 之后，再一次地修改了 i 的值，最终函数退出后的返回值才会是defer修改过的值。

- defer用途：

1. 关闭文件句柄
2. 锁资源释放
3. 数据库连接释放

- 匿名函数例子

1. 匿名函数执行顺序

```go
func deferFun1(){
    fmt.Println("开始")
    defer func(){
        fmt.Println(11)
        fmt.Println(12)
    }()  //匿名执行语句
    fmt.Println("结束")
}

输出结果：
开始
结束
11
12
```

2. 多个defer关键字

```go
fmt.Println("开始")
defer fmt.Println(1)
defer fmt.Println(2)
defer fmt.Println(3)
fmt.Println("结束")
输出结果：
开始
结束
3
2
1  
```

3. defer调用方法返回值为int类型

return -> defer -> defer的返回值(函数返回变量，必须声明变量名，defer修改其变量名才有效)

- 例1

```go
//变量a1为局部变量
func deferFun2() int{
    var a1 int //0
    defer func(){
        a1++
    }()  //defer 匿名执行语句
    return a1 //0
}
a2 := deferFun2()
fmt.Println(a2)  //0

输出结果：
0
```

- 例2

```go
func deferFun3()(a1 int) {
    defer func(){
        a1++
    }()  //defer 匿名执行语句
    return a1 //1
}
a3 := deferFun3()
fmt.Println(a3) //1

输出结果：
1
```

- 例3

```go
func deferFun4()(a1 int) {
    a1 = 5
    defer func(){
        a1++
    }()  //匿名执行语句
    return a1  //6
}
a4 := deferFun4()
fmt.Println(a4) //6

输出结果：
6
```

- 例4

```go
func deferFun5()(x int) {
    defer func(x int){
        //x = 0
        x++
    }(x)  //匿名执行语句
    return 5  //5
}  
a5 := deferFun5()
fmt.Println(a5) //5

输出结果：
5
```

- 例5

```go
func test() (ret int){
	ret = 10
	return 1
}

输出结果：
ret= 1
```

- 例6

```go
func test01() (ret int){
	defer func(){
		ret = 10
	}()
	return 1
}

输出结果：
ret1= 10
```

- 例7

```go
func test02() (ret int){
	defer func(){
		ret += 10
	}()
	return 1
}

输出结果：
ret2= 11
```

- 例8

```go
func test03() (ret int){
	ret = 10
	defer func(){
		
	}()
	return 1
}

输出结果：
ret3= 1
```

- 例9

```go
func test04() (ret int) {
    defer func(ret int) {
        ret = ret + 5
    }(ret)
    //defer匿名执行函数在return 1之前已经修改ret变量
    return 1
} 

输出结果：
ret4= 1
```

- 例10

```go
func test05() (ret int) {
    t := 5
    defer func() {
        t = t + 5
    }()
    return t
}

输出结果：
ret5= 5
```

5. defer 函数的嵌套
 
```go
func calc(index string, a, b int) int{
    ret := a + b
    fmt.Println(index,a,b,ret)
    return ret
}

x := 1
y := 2
defer calc("AA",x,calc("A", x, y))
x = 10
defer calc("BB",x,calc("B", x, y))
y = 20

输出结果：
A 1 2 3
B 10 2 12
BB 10 12 22
AA 1 3 4
```

注意:
注册顺序：
    defer calc("AA",x,calc("A", x, y))
    defer calc("BB",x,calc("B", x, y))
执行顺序：
    defer calc("BB",x,calc("B", x, y))
    defer calc("AA",x,calc("A", x, y))

### 1.7、panic 异常与recover监听

Panic： 抛出异常
Recover: 监听异常，继续执行； 在defer的匿名执行函数中。

1 panic ： 结束执行，抛出异常。
2 panic 可以再任何地方引发，但recover只有在defer调用的函数有效;recover监听panic异常,不会终止执行

```go
package main

import (
    "fmt"
    "errors"
)

func func2(){
    defer func(){
        err := recover() //监听panic异常
        if err != nil{
            fmt.Println("err:",err) //err: 抛出异常
        }
    }()
    panic("抛出异常")   
}

func readFile(fileName string) error{
    if fileName == "main.go"{
        return nil
    }else{
        return errors.New("读取文件失败")
    }
}

func myFunc(){
    defer func(){
        err := recover()
        if err != nil{
            fmt.Println("给管理员发送邮件")
        }
    }()
    err := readFile("main.go")
    if err != nil{
        panic(err)
    }
}

func main(){
    func2()
    fmt.Println("结束")
    myFunc()
    fmt.Println("继续执行...")
}

输出结果：
err: 抛出异常
结束
继续执行...
```

## 2、复合数据类型

### 2.1、time包及日期函数

#### 2.1.1、格式化输出

go诞生时间为2006年1月2号15点04分（记：2006 1 2 3 4 5）  03表示12小时制，15表示24小时制

```go
package main

import (
    "fmt"
    "time"
)
func main(){
    timeObj := time.Now()
    year := timeObj.Year()
    month := timeObj.Month()
    day := timeObj.Day()
    hour := timeObj.Hour()
    munite := timeObj.Minute()
    second := timeObj.Second()
    //1.1 %02d:表示2位，不够2位补0
    fmt.Printf("%d-%02d-%02d %02d:%02d:%02d",year,month,day,hour,munite,second) //
    /*1.2、格式化输出:go诞生时间为2006年1月2号15点04分（记：2006 1 2 3 4 5）  03表示12小时制，15表示24小时制*/
    var str1 = timeObj.Format("2006/01/02 03:04:05")
    fmt.Println(str1)
}

输出结果：
2020-07-20 15:44:09.9897106 +0800 CST m=+0.010970301
2020-07-20 15:45:19
```

#### 2.1.2、当前时间转化时间戳

```go
unixTime := timeObj.Unix()
fmt.Println("当前毫秒数时间戳：",unixTime) //
unixNanoTime := timeObj.UnixNano()
fmt.Println("当前纳秒秒数时间戳：",unixNanoTime) //
输出结果：
1595232070
1595232070113890600
```

#### 2.1.3、时间戳转化日期

```go
unixTime2 := 1595232070
timeObj2 := time.Unix(int64(unixTime2),0) //time.Unix(毫秒，纳秒)
var str2 = timeObj2.Format("2006/01/02 15:04:05")
fmt.Println(str2) //

输出结果：
2020/07/20 04:01:10
```

#### 2.1.4、日期字符串转化时间戳

```go
var str3 = "2020-02-28 15:23:24"
var tmp3 ="2006-01-02 15:04:05"  //模板
timeObj3,_ := time.ParseInLocation(tmp3, str3, time.Local )
fmt.Println(timeObj3)   //2020-02-28 15:23:24
fmt.Println(timeObj3.Unix())  //1582874604

输出结果：
2020-02-28 15:23:24
1582874604
```

#### 2.1.5、time包中定义时间间隔类型的常量

```go
fmt.Println(time.Nanosecond) 
fmt.Println(time.Microsecond) 
fmt.Println(time.Millisecond)  //1毫秒
fmt.Println(time.Second) //1秒
fmt.Println(time.Minute) 
fmt.Println(time.Hour) 
timeObj4 := timeObj.Add(time.Hour)
fmt.Println(timeObj4) //增加1小时

输出结果：
1ns
1µs
1ms
1s
1m0s
1h0m0s
timeObj4= 2020-10-17 16:09:56.111764 +0800 CST m=+3600.011967901
```

#### 2.1.6、定时器：time.NewTicker(time.Second)

```go
ticker := time.NewTicker(time.Second)
     n := 5
     for t := range ticker.C{
        n--
        fmt.Println(t)
        if n == 0{
            ticker.Stop() //终止定时器继续执行
            break
        }
     }
     //7.2 休眠：time.Sleep(time.Second) 每隔一秒打印一次
}

输出结果：
2020-10-17 15:09:57.4023304 +0800 CST m=+1.302534301
2020-10-17 15:09:58.3941943 +0800 CST m=+2.294398201
2020-10-17 15:09:59.3958202 +0800 CST m=+3.296024101
2020-10-17 15:10:00.4051134 +0800 CST m=+4.305317301
2020-10-17 15:10:01.3949763 +0800 CST m=+5.295180201
```

### 2.2 指针、make、new

- **值类型**

值类型包括基本数据类型，int,float,bool,string,以及数组和结构体(struct)。注意：sync.WaitGroup 对象是值类型，不是一个引用类型值类型变量声明后，不管是否已经赋值，编译器为其分配内存，此时该值存储于栈上。

- **引用类型**

引用类型包括slice切片，map ，chan，interface。
变量直接存放的就是一个内存地址值，这个地址值指向的空间存的才是值。所以修改其中一个，另外一个也会修改（同一个内存地址）。
引用类型必须申请内存才可以使用，make()是给引用类型申请内存空间。

#### 2.2.1、指针

```go
var a1 = 10
var p1 = &a1 
fmt.Printf("值：%v,类型：%T, 地址：%v \n",a1, a1,&a1) 
fmt.Printf("值：%v,类型：%T, 地址：%v, 指针指向地址的值：%v\n",p1, p1,&p1,*p1) 
*p1 = 21  
fmt.Println("*p1=",*p1) //
fmt.Println("a1=",a1) //

输出结果：
值：10,类型：int, 地址：0xc0000120b0
值：0xc0000120b0,类型：*int, 地址：0xc000006028, 指针指向地址的值：10
*p=21
 a1= 21
```

#### 2.2.2 make、new

- make与new区别

(1) make只有用于切片slice、map、通道channel的初始化，返回值还是**引用类型(T)**
`func make(Type, size IntegerType) Type`
`make(type, size, capacity)`

(2) new 用于类型的内存分配，并且内存对应的值为类型0值，返回值为**指向类存的指针(*T)**。
`func new(Type) *Type`

```go
var slice1 = make([]int, 4,4)
slice1[0] = 11
slice1[1] = 22
slice1 = append(slice1,1,2,4) //append进行扩容
fmt.Println(slice1)  //
// new 针对指针类型
var a3 *int
a3 = new(int)
fmt.Printf("值：%v,类型：%T, 地址：%v, 指针指向地址的值：%v\n",a3, a3,&a3,*a3) 
*a3 = 100
fmt.Println(*a3)  //100
输出结果：
[11 22 0 0 1 2 4]
值：0xc000012160,类型：*int, 地址：0xc000006038, 指针指向地址的值：0
100
```

### 2.3、结构体

#### 2.3.1 结构体定义

结构体：结构体（personStru）小写是私有的，结构体（PersonStru）大写是公有的

```go
type personStru struct{
    name string
    age int
    sex string
}
```

1. 实例化结构体

```go
var personStru1 personStru 
personStru1.name = "liwen"
personStru1.age = 20
personStru1.sex = "男"
fmt.Printf("值：%v,类型： %T \n",personStru1,personStru1)

输出结果：
值：{liwen 20 男},类型： main.personStru 
```

2. new实例化结构体： go 支持对结构体指针直接使用

```go
 //personStru2.name = "liwei"  等于 (*personStru2).name = "liwei" 
var personStru2 = new(personStru)
// personStru2.name = "liwei"  
// personStru2.age = 25
// personStru2.sex = "男"
(*personStru2).name = "liwei" 
(*personStru2).age = 25
(* personStru2).sex = "男"

fmt.Printf("值：%#v,类型： %T \n",personStru2,personStru2) 

输出结果：
值：&main.personStru{name:"liwei", age:25, sex:"男"},类型： *main.personStru
```

3. 引用地址实例化

```go
var personStru3 = &personStru{}
personStru3.name = "liwe"
personStru3.age = 23
personStru3.sex = "男"
fmt.Printf("值：%#v,类型： %T \n",personStru3,personStru3)

输出结果：
值：&main.personStru{name:"liwe", age:23, sex:"男"},类型： *main.personStru
```
 
4. 初始化的实例化

```go
var personStru4 = personStru{
    name:"liwe1",
    age:20,
    sex:"男",
}
fmt.Printf("值：%#v,类型： %T \n",personStru4,personStru4) 

输出结果：
值：main.personStru{name:"liwe1", age:20, sex:"男"},类型： main.personStru
```

5. 地址初始化的实例化

```go
var personStru5 = &personStru{
    name:"liwe1",
    age:20,
    sex:"男", //此逗号不能省略
}

fmt.Printf("值：%#v,类型： %T \n",personStru5,personStru5) 

输出结果：
值：&main.personStru{name:"liwe1", age:20, sex:"男"},类型： *main.personStru
```
  
总结： type 用于自定义类型、结构体

#### 2.3.2 结构体方法

```go
//1 、结构体是值类型，改变副本不影响主本: 
//go 的结构体是相互独立的，不会影响
type personStru struct{
    name string
    age int
    sex string
}
//2. go没有类class，但是定义结构体方法：结构体调用方法
func (personStruTmp personStru) printInfo(){
    fmt.Printf("name=%v, age=%v, sex=%v \n",personStruTmp.name,personStruTmp.age, personStruTmp.sex)
}
//2.1结构体无法修改结构体属性
func (personStruTmp1 personStru) setInfo1(name string, age int){
    personStruTmp1.name = name
    personStruTmp1.age = age
}
//2.2 结构体指针修改结构体属性
func (personStruTmp2 *personStru) setInfo2(name string, age int){
    (*personStruTmp2).name = name
    (*personStruTmp2).age = age
}
```

#### 2.3.3、结构体嵌套和结构体继承

1. 结构体嵌套

```go
type user struct{
    username string
    password string
    age  int
    //address1 address //user结构体嵌套address结构体
    address  //2.2 匿名嵌套结构体
    email email
}
type address struct{
    name string
    phone string
    city string
    age int
}
type email struct{
    name string
}
```

2. 结构体继承

```go
package main

import (
    "fmt"
)
type animal struct{
    name string
}
func (a animal) run(){
    fmt.Printf("%v 在运动\n", a.name)
}
//子结构体
type dog struct{
    age int
    *animal //结构体继承或嵌套
}
func (d dog) wang() {
    fmt.Printf("%v ,%v 在叫\n",d.age, d.name)
}

func main(){
    var d = dog{
        age:20,
        animal:&animal{
            name:"马",
        },
    }
    d.run()
    d.wang()
}

输出结果：
马 在运动
20 ,马 在叫
```

### 2.4、JSON

#### 2.4.1、结构体与json相互转换

```go
package main

import (
    "fmt"
    //"sort"
    "encoding/json"
)
//结构体转json：必须大写
//结构体变量、成员变量必须大写(public)，否则无法转化json
type PersonStru struct{
    Name string
    Age int
    Sex string
}
// 结构体的标签：结构体变量为小写
type PersonStruLable struct{
    Name string `json:"id"`
    Age int  `json:"age"`
    Sex string `json:"sex"`
    Sno string `json:"sno"`
}
func main(){
    //1 结构体对象转化json字符串
    var personStru1 = PersonStru{
        Name:"liwe1",
        Age:20,
        Sex:"男",
    }
    
    //json的[]byte切片
    jsonByte,_ := json.Marshal(personStru1)
    //Byte切片转化json字符串
    jsonStr := string(jsonByte) 
    fmt.Printf("%#v \n", jsonStr) //%#v表示字符串输出
    fmt.Printf("%v \n", jsonStr)  //%v表示默认输出
    
    //2 、json字符串转化结构体：使用反引号``
    var jsonStr1 = `{"Name":"liwe1","Age":20,"Sex":"男"}`
    var personStru2 PersonStru
    err := json.Unmarshal([]byte(jsonStr1), &personStru2)
    if err != nil{
        fmt.Println(err)
    }
    fmt.Printf("%#v \n", personStru2) 
}

输出结果：
"{\"Name\":\"liwe1\",\"Age\":20,\"Sex\":\"男\"}"
{"Name":"liwe1","Age":20,"Sex":"男"}
main.PersonStru{Name:"liwe1", Age:20, Sex:"男"}
```

#### 2.4.2、结构体嵌套与json转化

```go
package main

import (
    "fmt"
    "encoding/json"
)
//1 结构体嵌套转化json
type Student struct{
    ID int
    Gender string
    Name string
}
type Class struct{
    Title string
    Student []Student
}
func main(){
    //1 结构体嵌套转化json
    c := Class{
        Title: "001班",
        Student: make([]Student,0),
    }
    for i:=0; i < 5; i++{
        s := Student{
            ID : i,
            Gender: "男",
            Name : fmt.Sprintf("stu_%v",i),
        }
        c.Student = append(c.Student , s)
    }
    fmt.Println(c)
    strByte, err := json.Marshal(c)
    if err != nil{
        fmt.Println("json 转化失败")
    }else{
        strJson := string(strByte)
        fmt.Println(strJson)
    }
    
    //2 json字符串转化结构体嵌套
    strJson2 := `{"Title":"001班","Student":[{"ID":0,"Gender":"男","Name":"stu_0"},{"ID":1,"Gender":"男","Name":"stu_1"},{"ID":2,"Gender":"男","Name":"stu_2"},{"ID":3,"Gender":"男","Name":"stu_3"},{"ID":4,"Gender":"男","Name":"stu_4"}]}`
    var c2  = &Class{}
    err1 := json.Unmarshal([]byte(strJson2), c2)
    if err1 != nil{
        fmt.Println("err")
    }else{
        fmt.Printf("%#v \n", c2)
    }
    
}

输出结果：
{001班 [{0 男 stu_0} {1 男 stu_1} {2 男 stu_2} {3 男 stu_3} {4 男 stu_4}]}

{"Title":"001班","Student":[{"ID":0,"Gender":"男","Name":"stu_0"},{"ID":1,"Gender":"男","Name":"stu_1"},{"ID":2,"Gender":"男","Name":"stu_2"},{"ID":3,"Gender":"男","Name":"stu_3"},{"ID":4,"Gender":"男","Name":"stu_4"}]}

&main.Class{Title:"001班", Student:[]main.Student{main.Student{ID:0, Gender:"男", Name:"stu_0"}, main.Student{ID:1, Gender:"男", Name:"stu_1"}, main.Student{ID:2, Gender:"男", Name:"stu_2"}, main.Student{ID:3, Gender:"男", Name:"stu_3"}, main.Student{ID:4, Gender:"男", Name:"stu_4"}}}
```

## 3、接口

### 3.1、接口定义与实现

Type 接口名 interface{
方法名1（参数列表1）返回值列表1
方法名2（参数列表2）返回值列表2
}

```go
package main

import (
    "fmt"
)
//1接口
type Usb interface{
    start()
    stop()
}
//1.1 若接口里面的方法，必须使用结构体或自定义方法实现接口
type Phone struct{
    Name string
}
//1.2手机要实现usb接口，必须实现usb接口中的所有方法，也可以调用自己的方法
func (p Phone)start(){
    fmt.Println("启动：",p.Name)
}
func (p Phone)stop(){
    fmt.Println("关机:",p.Name)
}

func main(){
    p := Phone{
        Name:"华为手机",
    }
    p.start()
    var u Usb
    u = p
    u.stop()
    
}

输出结果：
启动： 华为手机
关机: 华为手机
```

### 3.2、空接口与类型断言

空接口：没有任何约束,即任何数据类型都可以实现空接口

```go
//空接口
//1.1空接口：任何数据类型都可以用空接口实现
type NullInterface interface{}
//1.2 空接口参数可以是任意类型
func show(a interface{}){
    fmt.Printf("值：%v,类型：%T \n", a,a)
}
//2类型断言：
//2.1 方法： 接口变量名.(类型)
func printInfo(b interface{}) {
    if _,ok := b.(string);ok{
        fmt.Println("string类型")
    }else if _,ok := b.(int);ok{
        fmt.Println("int")
    }else if _,ok := b.(float64);ok{
        fmt.Println("float64")
    }else{
        fmt.Println("null")
    }
}
//2.2 方法2：switch的b1.(type)表示变量判断类型，只能使用在switch
func printInfo1(b1 interface{}) {
    switch b1.(type){
    case string:
        fmt.Println("string类型")
    case int:
        fmt.Println("int")
    case float64:
        fmt.Println("float64")
    default:
        fmt.Println("null")
    }
}
```

### 3.3、结构体值与结构体指针实现接口的区别

结构体值：初始化值或初始化地址
结构体指针：初始化地址

```go
package main

import (
    "fmt"
)
//接口
type Usb interface{
    start()
    stop()
}
//若接口里面的方法，必须使用结构体或自定义方法实现接口
type Phone1 struct{
    Name string
}
type Phone2 struct{
    Name string
}
//2 结构体指针接收者
func (p *Phone1)start(){
    fmt.Println("启动1：",p.Name)
}
func (p *Phone1)stop(){
    fmt.Println("关机1：",p.Name)
}
func (p *Phone2)start(){
    fmt.Println("启动2：",p.Name)
}
func (p *Phone2)stop(){
    fmt.Println("关机2：",p.Name)
}
func main(){
    //1 结构体值接收者
    p1 := Phone1{
        Name:"华为手机",
    }
    var u1 Usb = &p1
    u1.start()
    //2 结构体指针接收者
    p2 := &Phone2{
        Name:"小米手机",
    }
    var u2 Usb = p2
    u2.start()
}

输出结果：
启动1： 华为手机
启动2： 小米手机
```

### 3.4 一个接口与2个或多个结构体，接口方法含有参数

```go
package main

import (
    "fmt"
)
//1个接口2个结构体
type Animals interface {
    SetName(string)
    GetName() string
}
// 狗
type Dog struct{
    Name string
}
//结构体是值类型，改变需要使用指针类型
func (d *Dog) SetName(name string){
    (*d).Name = name //d.Name = name
}
func (d Dog) GetName() string{
    return d.Name
}
//猫
type Cat struct{
    Name string
}
//结构体是值类型，改变需要使用指针类型
func (c *Cat) SetName(name string){
    (*c).Name = name //d.Name = name
}
func (c Cat) GetName() string{
    return c.Name
}
func main(){
    //Dog 
    var d = &Dog{
        Name : "金毛",
    }
    var a Animals = d 
    fmt.Println(a.GetName()) //金毛
    a.SetName("哈士奇")
    fmt.Println(a.GetName()) //哈士奇
    //cat
    var c = &Cat{
        Name : "黑猫",
    }
    var a1 Animals = c 
    fmt.Println(a1.GetName()) //金毛
    a1.SetName("橘猫")
    fmt.Println(a1.GetName()) //哈士奇
}

输出结果：
金毛
哈士奇
黑猫
橘猫
```

### 3.5 一个结构体实现多个接口

```go
package main

import (
    "fmt"
)
//1个结构体实现多个接口
type Animals1 interface {
    SetName(string)
}
type Animals2 interface {
    GetName() string
}
// 狗
type Dog struct{
    Name string
}
//结构体是值类型，改变需要使用指针类型
func (d *Dog) SetName(name string){
    (*d).Name = name //d.Name = name
}
func (d Dog) GetName() string{
    return d.Name
}
func main(){
    //Dog 
    var d = &Dog{
        Name : "金毛",
    }
    var a1 Animals1 = d 
    var a2 Animals2 = d 
    a1.SetName("小花狗")
    fmt.Println(a2.GetName()) //小花狗
}

输出结果：
小花狗
```

### 3.6  接口嵌套

允许接口嵌套使用

```go
package main

import (
    "fmt"
)
//1接口嵌套
type Animals1 interface {
    SetName(string)
    
}
type Animals2 interface {
    GetName() string
}
type AnimalsSum interface {
    Animals1
    Animals2
}
// 狗 :结构必须实现接口所有的方法
type Dog struct{
    Name string
}
//结构体是值类型，改变需要使用指针类型
func (d *Dog) SetName(name string){
    (*d).Name = name //d.Name = name
}
func (d Dog) GetName() string{
    return d.Name
}
func main(){
    //Dog 
    var d = &Dog{
        Name : "金毛",
    }
    var aSum AnimalsSum = d 
    aSum.SetName("小花狗")
    fmt.Println(aSum.GetName()) //小花狗
}

输出结果：
小花狗
```

### 3.7 空接口与类型判断的细节

1 空接口
（1）错误，空接口不能使用索引，
解决方法：类型判断获取其对象,索引

（2）错误，空接口没有对应的属性
解决方法：类型断言获取其对象

```go
package main

import (
    "fmt"
)
type Address struct{
    Name string
    Phone int
}
func main(){
    var userInfo = make(map[string] interface{})
    userInfo["username"] = "liwen"
    userInfo["age"] = 20
    userInfo["hobby"] = []string{"打篮球","羽毛球"}
    fmt.Println(userInfo["username"]) //liwen
    fmt.Println(userInfo["age"])  //20
    //fmt.Println(userInfo["hobby"][1])  //错误，空接口不能使用索引
    hobby2, _ := userInfo["hobby"].([] string) //类型判断获取其对象,索引
    fmt.Println(hobby2[1]) //羽毛球
    var address = Address{
        Name:"lwien",
        Phone: 123344,
    }
    userInfo["address"] = address
    fmt.Println(userInfo["address"]) //{lwien 123344}
    //fmt.Println(userInfo["address"].Name) //错误，空接口没有对应的属性
   address2 , _ := userInfo["address"].(Address)//类型断言获取其对象
   fmt.Println(address2.Name, address2.Phone) //{lwien 123344}
    
}

输出结果：
liwen
20
羽毛球
{lwien 123344}
lwien 123344
```

## 4、goroutine 和channel

### 4.1、goroutine 实现并发与并行

Golang的主线程（进程）：一个Golang的主线程（进程）可以起多个协程goroutine.，多个协程可以实现并行或并发。
协程：可以理解用户级别的线程，与正真的线程不同。Golang的协程就goroutine，通过go 关键词 + 方法名称.

#### 4.1.1 开启多个协程: sync.WaitGroup

//1、主线程与协程的执行顺产：若主线程执行完，则协程退出；若协程执行完，主线程进行执行。

`var wg sync.WaitGroup //监听协程完成 :g.Add(1) -》程序-》 wg.Done() -》wg.Wait()`

```go
package main

import (
    "fmt"
    "time"
    "sync"
)
//1、主线程与协程的执行顺产：若主线程执行完，则协程退出；若协程执行完，主线程进行执行.
var wg sync.WaitGroup //监听协程完成 ：   wg.Add(1) -》 wg.Done() -》wg.Wait() 
func test1(){
    for i := 0; i < 5; i++{
        fmt.Println("test: ", i)
        time.Sleep(time.Millisecond*10)
    }
    wg.Done()//协程计数器减1
}
func test2(){
    for i := 0; i < 5; i++{
        fmt.Println("test: ", i)
        time.Sleep(time.Millisecond*10)
    }
    wg.Done()//协程计数器减1
}
func main(){
    wg.Add(1)  //协程计数器加1
    go test1()
    wg.Add(1)  //协程计数器加1
    go test2()
    for i := 0; i < 5; i++{
        fmt.Println("main: ", i)
        time.Sleep(time.Millisecond*10)
    }
    //time.Sleep(time.Second)
    wg.Wait() //等待协程完毕
    fmt.Println("主线程退出!")
}

输出结果：
main:  0
test:  0
test:  0
test:  1
main:  1
test:  1
test:  2
main:  2
test:  2
main:  3
test:  3
test:  3
test:  4
test:  4
main:  4
主线程退出!
```

#### 4.1.2 Runtime.GOMAXPROCS()函数

runtime.GOMAXPROCS()函数设置程序并发占用的CPU的逻辑核心数，可以设置CPU的个数。
runtime.NumCPU 获取当前计算机上面的CPU个数

#### 4.1.3 Goroutine 与for实现多协程执行并发

```go
package main

import (
    "fmt"
    "sync"
    "time"
)
  
//1 开始多协程：并行
var wg sync.WaitGroup
func test(num int){
    defer wg.Done()
    for i := 0; i < 5; i++{
        fmt.Printf("协程%v打印第%v条数据\n",num, i)
        time.Sleep(time.Millisecond*100)
    }
    //wg.Done()//协程计数器减1
}
func main(){
    for i := 0; i < 5; i++{
        wg.Add(1)
        go test(i)
    }
    wg.Wait()
    fmt.Println("关闭主线程")
}
```

### 4.2 channel

Channel是管道，先入先出原则,是一种引用数据类型。
引用数据类型的总结：切片slice、集合map、管道channel
make的总结：切片slice、集合map、管道channel

声明管道channel：`var 变量 chan 元素类型`
例如：`var ch1 chan int`

- 创建管道channel

Ch : = make(chan int, 3)
管道有发送（send）、接收（receive）和关闭（close）三个操作
（1）发送（将数据放在管道内）
`ch <- 10`  //将10放到管道ch内
（2）接收（从管道内取值）
`x := <- ch` //从管道ch中接收到值赋值给变量x
`<- ch`
（3）关闭
`Close(ch)`

#### 4.2.1 管道基本使用

```go
package main
import (
    "fmt"
)

func main(){
    //1 创建管道
    ch := make(chan int, 3) //3表示容量
    //2 给管道传送值
    ch <- 10
    ch <- 20
    ch <- 30
    //3 接收管道的值
    a1 := <- ch
    fmt.Println("a1=",a1)//a1= 10 ，先入先出
    <- ch   //20
    a3 := <- ch
    fmt.Println("a3=",a3)  //a3= 30
    ch <- 55
    //4 管道长度与容量
    fmt.Printf("值：%v, 长度：%v,容量：%v\n", ch, len(ch),cap(ch)) //值：0xc000096080, 长度：1,容量：3
    //5 管道是引用类型:改变副本会改变主本
    ch1 := make(chan int, 4)
    ch1 <- 34
    ch1 <- 35
    ch1 <- 36
    ch2 := ch1
    ch2 <- 253
    <- ch1
    <- ch1
    <- ch1
    a4 := <- ch1
    fmt.Println(a4) //253
    //6 管道阻塞:
    //1）存放管道的数据大于管道容量会造成阻塞
    //2) 若管道的内容已经去完，再次取值i会造成阻塞
    ch3 := make(chan int, 2)
    ch3 <- 44
    ch3 <- 45
    //ch3 <- 46  //管道阻塞：fatal error: all goroutines are asleep - deadlock!
    a5 := <- ch3
    a6 := <- ch3
    //a7 := <- ch3 //超出管道取值范围，fatal error: all goroutines are asleep - deadlock!
    //fmt.Println(a5,a6,a7)
    fmt.Println(a5,a6)

    
}

输出结果：
a1= 10
a3= 30
值：0xc00001c180, 长度：1,容量：3
253
44 45
```

#### 4.2.2 for...range或for遍历获取管道的值

```go
package main
import (
    "fmt"
)
//1 管道循环遍历,
//1.1 若在使用for...range取出管道值,则之前必须关闭管道close(ch)
//1.2 若在使用for取出管道值,则之前不需要关闭管道close(ch)
func main(){
    ch1 := make(chan int, 10)
    for i :=1; i <= 10; i++{
        ch1 <- i 
    }
    defer close(ch1) //必须关闭管道
    // //管道没有key，只有value
    // for v := range ch1{
    //  fmt.Println(v)
    // }
    for i :=1; i <= 10; i++{
        fmt.Println(<- ch1)
    }
}

输出结果：
1
2
3
4
5
6
7
8
9
10
```

#### 4.2.3 goroutine与channel

创建通道make(chan int, 3) --> 开启计数器Add(1)-->go 协程 --> 关闭协程close() --> 关闭计时器Done() --> 关闭监控wait()
例1：

```go
package main
import (
    "fmt"
    "sync"
    "time"
)
//需求： goroutine 与 channel同时工作
var wg sync.WaitGroup
//写数据
func writeChan(writeCh chan int){
    defer wg.Done()
    for i := 1; i <= 10; i++{
        writeCh <- i
        fmt.Printf("【写入】数据%v到通道\n", i)
        time.Sleep(time.Millisecond*500)
    }
    close(writeCh)
}
//读数据
func readChan(readCh chan int){
    for v :=  range readCh {
        fmt.Printf("从通道【读取】数据%v\n", v)
        time.Sleep(time.Millisecond*50)
    }
    wg.Done()
}
func main(){
    var ch = make(chan int, 3)
    wg.Add(1)
    go writeChan(ch)
    wg.Add(1)
    go readChan(ch)
    wg.Wait()
    fmt.Println("关闭主线程\n")
}

输出结果：
从通道【读取】数据1
【写入】数据1到通道
【写入】数据2到通道
从通道【读取】数据2
【写入】数据3到通道
从通道【读取】数据3
关闭主线程
```

#### 4.2.4 channel方向

```go
package main
import (
    "fmt"
)
func main(){
    //1 双向管道
    ch1 := make(chan int, 3)
    ch1 <- 3
    ch1 <- 4
    a1 := <- ch1
    a2 := <- ch1
    fmt.Println(a1, a2)
    //2 只写入管道
    ch2 := make(chan<- int, 3)
    ch2 <- 5
    ch2 <- 6
    // <- ch2  //无法从 管道读取数据
    //3 只从管道读取数据
    //ch3 := make(<-chan int, 3)
    //ch3 <- 6
}
```

#### 4.2.5 select多路复用

Select用于同时从多个通道获取数据。

```go
Select {
Case <- ch1:
...
Case  data := <- ch2:
...
Default:
...
}
```

当select 随机选取case，当所有case 执行完，则结束select。

```go
package main
import (
    "fmt"
    "time"
)

func main(){
    intChan := make(chan int , 4)
    //defer close(intChan)
    for i :=0 ; i < 4; i++{
        intChan <- i;
    }
    strChan := make(chan string, 4)
    //defer close(strChan )
    for i :=0 ; i < 4; i++{
        strChan <- "hello" + fmt.Sprintf("%d",i);
    }
    //select 不需要关闭通道，多路复用：同时从多个通道获取数据
    for {
        select{
        case v := <- intChan:
            fmt.Printf("从inChan读取数据 % d \n", v)
            time.Sleep(time.Millisecond*50)
        case v := <- strChan:
            fmt.Printf("从strChan读取数据 % v \n", v)
            time.Sleep(time.Millisecond*50)
        default:
            fmt.Printf("读取数据完毕  \n", )
            return //注意退出...
        }
    }
}

输出结果：
从strChan读取数据 hello0
从inChan读取数据  0
从inChan读取数据  1
从inChan读取数据  2
从strChan读取数据 hello1
从strChan读取数据 hello2
从inChan读取数据  3
从strChan读取数据 hello3
读取数据完毕
```

### 4.3 互斥锁与读写锁

#### 4.3.1 互斥锁

同一时刻只有一个协程运行
互斥锁： sync.Mutex; Lock 锁住当前的共享资源：一个协程访问，另一个协程不能进行访问。Unlock进行解锁。
互斥锁解决资源竞争问题，并发竞争一个资源。

```go
package main
import (
    "fmt"
    "sync"
    "time"
)

var count = 0
var wg sync.WaitGroup
var mutex sync.Mutex
func test(){
    mutex.Lock()
    count++
    fmt.Println("the count is : ",count)
    time.Sleep(time.Millisecond)
    mutex.Unlock()
    wg.Done()
}
func main(){
    for i := 0; i < 5; i++{
        wg.Add(1)
        go test()
    }
    wg.Wait()
}

输出结果：
the count is :  1
the count is :  2
the count is :  3
the count is :  4
the count is :  5
```

#### 4.3.2 读写互斥锁

读锁可以让读操作并发，其他协程同时读取；但是写锁完全互斥。即当一个协程在写操作，其他协程不能进行读或写操作。

```go
Sync.RWMutex
写锁：
Func(* RWMutex)Lock 
Func(* RWMutex)Unlock
读锁：
Func(* RWMutex)RLock
Func(* RWMutex)RUnlock
```

例子

```go
package main
import (
    "fmt"
    "sync"
    "time"
)

var wg sync.WaitGroup
var rwmutex sync.RWMutex

//写锁
func write(){
    rwmutex.Lock()
    fmt.Println("执行写操作")
    time.Sleep(time.Second*2)
    rwmutex.Unlock()
    wg.Done()
}
//读锁
func read(){
    rwmutex.RLock()
    fmt.Println("执行读操作")
    time.Sleep(time.Second*2)
    rwmutex.RUnlock()
    wg.Done()
}
func main(){
    for i := 0; i < 3; i++{
        wg.Add(1)
        go write()
        wg.Add(1)
        go read()
    }
    wg.Wait()
    fmt.Println("主线程完成")
}

输出结果：
执行读操作
执行读操作
执行写操作
执行读操作
执行写操作
执行写操作
主线程完成
```

## 5、反射

既然反射就是用来检测存储在接口变量内部(值value；类型concrete type) pair对的一种机制。

- reflect.TypeOf：直接给到了我们想要的type类型，如float64、int、各种pointer、struct 等等真实的类型。

- reflect.ValueOf：直接给到了我们想要的具体的值，如1.2345这个具体数值。

- 场景

golang中反射最常见的使用场景是做对象的序列化（serialization，有时候也叫Marshal & Unmarshal）。

例如，Go语言标准库的encoding/json、encoding/xml、encoding/gob、encoding/binary等包就大量依赖于反射功能来实现。

### 5.1 反射类型reflect.TypeOf(变量名) 

判断空接口的类型？值是什么？
方法：
1、可以使用类型断言：变量名.(string)
2、可以使用反射实现:
V := reflect.TypeOf(变量名) 

```go
package main
import (
    "fmt"
    "reflect"
)
//1、通过反射获取空接口类型
func reflectType(x interface{}){
    v := reflect.TypeOf(x)  //通过反射获取空接口类型
    fmt.Printf("类型：%v， 类型名称：%v，种类：%v\n",  v,  v.Name(), v.Kind())
}
//2、自定义类型
//type: 自定义类型、结构体、接口定义
type myInt int
type Person struct{
    Name string
    Age int
}
func main(){
    a := 1
    f := 3.44
    b := true
    str := "liwen"
    reflectType(a) //类型：int， 类型名称：int，种类：intfloat64
    reflectType(f) //foat64
    reflectType(b) //bool
    reflectType(str)
    var a1 myInt = 11
    person1 := Person{
        Name:"lwien",
        Age:12,
    }
    reflectType(a1) //main.myInt
    reflectType(person1) //类型：main.Person， 类型名称：Person，种类：struct
    var arr1 = [3]int{1,2,3}
    reflectType(arr1) //类型：[3]int， 类型名称：，种类：array[]int
    var slice1 = []int{1,3,4}
    reflectType(slice1) //类型：[]int， 类型名称：，种类：slice
}

输出结果：
类型：int， 类型名称：int，种类：int
类型：float64， 类型名称：float64，种类：float64
类型：bool， 类型名称：bool，种类：bool
类型：string， 类型名称：string，种类：string
类型：main.myInt， 类型名称：myInt，种类：int
类型：main.Person， 类型名称：Person，种类：struct
类型：[3]int， 类型名称：，种类：array
类型：[]int， 类型名称：，种类：slice
```

### 5.2  反射值变量reflect.ValueOf(变量)

reflect.ValueOf()返回是reflect.Value类型，对类型值进行算数运算。

```go
package main
import (
    "fmt"
    "reflect"
)

//1、通过反射获取空接口类型的值
func reflectType(x interface{}){
    v := reflect.ValueOf(x)
    fmt.Println(v)
    //var m = v + 12  //错误
    //var m = v.Int() + 12 //获取原始值
    kind := v.Kind()
    switch kind{
    case reflect.Int:
        fmt.Printf("int类型的原始值%v \n",v.Int() + 12 )
    case  reflect.Float32:
        fmt.Printf("Float32类型的原始值%v \n",v.Float() + 12.33 )
    case  reflect.Float64:
        fmt.Printf("Float64类型的原始值%v \n",v.Float() + 12.33 ) 
    case  reflect.String:
        fmt.Printf("string类型的原始值%v \n",v.String()+"hello" )
    default:
        fmt.Printf("没有获得类型的原始值%v \n" )
    }
}
func main(){
    a1 := 11
    reflectType(a1)
    f1 := 3.455
    reflectType(f1)
    str1 := "liwen"
    reflectType(str1)
}

输出结果：
11
int类型的原始值23
3.455
Float64类型的原始值15.785
liwen
string类型的原始值liwenhello
```

### 5.3 通过反射设置变量的值

修改变量的值：

```go
Func (v Value) SetBool(x bool)
Func (v Value) SetInt(x int64)
Func (v Value) SetString(x string)
```

值传递类型：
`reflect.ValueOf(x).Kind()`
地址传递类型：
`reflect.ValueOf(x).Elem().Kind()`

```go
package main
import (
    "fmt"
    "reflect"
)
//空接口的值不能直接修改，可以反射修改
func reflectSetValue(x interface{}){
    //空接口的值不能修改
    //*x = 120  //invalid indirect of x (type interface {})
    v := reflect.ValueOf(x)
    //fmt.Println(v.Kind()) //ptr
    fmt.Println(v.Elem().Kind()) //int64
    if v.Elem().Kind() == reflect.Int64{
        v.Elem().SetInt(123)
    }else if  v.Elem().Kind() == reflect.String{
        v.Elem().SetString("wen")
    }
}
func main(){
    var a1 int64 = 100
    reflectSetValue(&a1)  //值修改需使用地址传入
    fmt.Println(a1) //123
    var str1 string = "liwen "
    reflectSetValue(&str1)
    fmt.Println(str1) //wen
}

输出结果：
int64
123
string
wen
```

### 5.4 结构体反射：属性

#### 5.4.1 通过反射显示结构体字段

```go
type Student struct{
    Name string `json:"name" form:"username"`
    Age int `json:"age"`
    Score int `json:"score"`
}stu1

//1 通过类型变量的Filed获取结构体字段:索引
field1 := reflect.TypeOf(stu1).Field(0) //Name
field1.Name //字段名称：Name
field1.Type //字段类型：string
field1.Tag.Get("json") // 字段Tag：name
field1.Tag.Get("form") //字段Tag：username

//2. 通过变量FieldByName获得结构体字段:key
field2, ok := reflect.TypeOf(stu1).FieldByName("Age")

//3. 通过变量NumField获得结构体字段有几个字段属性
var fieldCount =  t.NumField() //3个字段属性
```

例如

```go
package main
import (
    "fmt"
    "reflect"
)

//结构体
type Student struct{
    Name string `json:"name" form:"username"`
    Age int `json:"age"`
    Score int `json:"score"`
}
func (s Student) GetInfo() string{
    var str = fmt.Sprintf("姓名：%v, 年龄：%v, 分数： %v \n",s.Name, s.Age, s.Score)
    return str
}
func (s *Student) SetInfo(name string, age int, score int) {
    s.Name = name
    s.Age = age
    s.Score = score
}
func printStructFiled(stu1 interface{}){
    t := reflect.TypeOf(stu1)
    if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
        fmt.Println("传输的参数不是一个结构体")
        return
    }
    //1 通过类型变量的Filed获取结构体字段:索引
    field1 := t.Field(0) //Name
    fmt.Printf("%#v \n", field1) 
    fmt.Printf("字段名称：%v \n", field1.Name)  //字段名称：Name
    fmt.Printf("字段类型：%v \n", field1.Type) //字段类型：string
    fmt.Printf("字段Tag：%v \n", field1.Tag.Get("json")) // 字段Tag：name
    fmt.Printf("字段Tag：%v \n", field1.Tag.Get("form")) //字段Tag：username
    //2 通过变量FieldByName获得结构体字段:key
    field2, ok := t.FieldByName("Age")
    if ok{
        fmt.Printf("字段名称：%v \n", field2.Name)  //字段名称：Age
        fmt.Printf("字段类型：%v \n", field2.Type) //字段类型：int
        fmt.Printf("字段Tag：%v \n", field2.Tag.Get("json")) // age
    }
    //3 通过变量NumField获得结构体字段有几个字段属性
    var fieldCount =  t.NumField()
    fmt.Println("结构体有",fieldCount,"个属性") //结构体有 3 个属性
    //4 通过值变量获取结构体属性对应的值
    v := reflect.ValueOf(stu1)
    fmt.Println(v.FieldByName("Name")) //liwne
    fmt.Println(v.FieldByName("Age")) //12
    for i := 0 ; i< fieldCount; i++{
        fmt.Printf("属性名称：%v,属性值：%v,属性类型：%v, 属性Tag:%v \n",
        t.Field(i).Name,t.Field(i),t.Field(i).Type,t.Field(i).Tag)
    }
}

func main(){
    stu1 := Student{
        Name :"liwne",
        Age: 12,
        Score: 34,
    }
    printStructFiled(stu1)
}

输出结果：
reflect.StructField{Name:"Name", PkgPath:"", Type:(*reflect.rtype)(0x4e1d20), Tag:"json:\"name\" form:\"username\"", Offset:0x0, Index:[]int{0}, Anonymous:false}
字段名称：Name 
字段类型：string
字段Tag：name 
字段Tag：username
字段名称：Age
字段类型：int
字段Tag：age
结构体有 3 个属性
liwne
12
属性名称：Name,属性值：{Name  string json:"name" form:"username" 0 [0] false},属性类型：string, 属性Tag:json:"name" form:"username"
属性名称：Age,属性值：{Age  int json:"age" 16 [1] false},属性类型：int, 属性Tag:json:"age"
属性名称：Score,属性值：{Score  int json:"score" 24 [2] false},属性类型：int, 属性Tag:json:"score"
```

#### 5.4.2 打印结构体方法method()

例如

```go
func PrintStructMethod (stu1 interface{}){
    t := reflect.TypeOf(stu1)
    if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
        fmt.Println("传输的参数不是一个结构体")
        return
    }
    //1 通过method获取结构体方法
    method1 := t.Method(0) //获取第一个方法：GetInfo()
    fmt.Println(method1.Name) //GetInfo
    fmt.Println(method1.Type) //func(main.Student) string
    //2 通过结构体获取结构体有多少方法
    method2,ok := t.MethodByName("print")
    if ok{
        fmt.Println(method2.Name)
        fmt.Println(method2.Type)
    }
    //3 通过值变量执行方法,获取值
    v := reflect.ValueOf(stu1)
    fn0:= v.Method(0).Call(nil) //获取第一个方法：GetInfo()
    fmt.Println(fn0) // [姓名：liwne, 年龄：12, 分数： 34    ]
    fn1 := v.MethodByName("GetInfo").Call(nil) //GetInfo()
    fmt.Println(fn1)  //[姓名：liwne, 年龄：12, 分数： 34]
    //4 执行方法传入参数，修改结构体参数
    var params []reflect.Value
    params = append(params, reflect.ValueOf("liwen"))
    params = append(params, reflect.ValueOf(24))
    params = append(params, reflect.ValueOf(56))
    v.MethodByName("SetInfo").Call(params)
    fn2 := v.MethodByName("GetInfo").Call(nil) 
    fmt.Println(fn2) //[姓名：liwen, 年龄：24, 分数： 56]
    //5 获取方法数量
    fmt.Println("方法数量：",t.NumMethod()) //方法数量： 2
}

func main(){
    stu1 := Student{
        Name :"liwne",
        Age: 12,
        Score: 34,
    }
    PrintStructMethod (&stu1)

}

输出结果：
GetInfo
func(*main.Student) string
[姓名：liwne, 年龄：12, 分数： 34]
[姓名：liwne, 年龄：12, 分数： 34]
[姓名：liwen, 年龄：24, 分数： 56]
方法数量： 2
```

#### 5.4.3 反射修改结构体属性

```go
func reflectChangeStruct(stu1 interface{}){
    v := reflect.ValueOf(stu1)
    t := reflect.TypeOf(stu1)
    if t.Kind() != reflect.Ptr{
        fmt.Println("传输的参数不是一个执指针类型")
        return
    }else if t.Elem().Kind() != reflect.Struct {
        fmt.Println("传输的参数不是一个结构体")
        return
    }
    //修改结构体属性的值
    name := v.Elem().FieldByName("Name")
    name.SetString("lisi")
    age := v.Elem().FieldByName("Age")
    age.SetInt(23)
}
func main(){
    stu1 := Student{
        Name :"liwne",
        Age: 12,
        Score: 34,
    }

    reflectChangeStruct(&stu1)
    fmt.Println(stu1) //{lisi 23 56}
}

输出结果：
{lisi 23 56}
```

## 6、文件或目录读写操作

### 6.1、文件读取

1. 方法一：流方式读取（大文件）

- 只读方式打开文件 `file,err := os.Open(“文件”)`
- 读文件`file.Read`(读取存放的数据)
- 关闭文件 `defer file.Close()`

2. 方法二：流的方式bufio读取文件（大文件）

- 只读方式打开文件 `file,err := os.Open(“文件”)`
- 创建reader对象 `reader := bufio.NewReader(file)`
- ReadString 读取文件 `line,err := reader.ReadString(“\n”)`
- 关闭文件 `defer file.Close()`

3. 方法三：ioutil(In out utility,输入输出功能) 一次读完文件数据（小文件100~200M）

- `ioutil.ReadFile(“文件”)`
自动关闭文件

- 例1：方法一：流方式读取（大文件）

```go
package main
import (
    "fmt"
    "os"
    "io"
)

func main(){
    //1 方法一：只读方式打开文件
    file1, err1 := os.Open("./text.txt")
    defer file1.Close()//关闭文件
    if err1 != nil{
        fmt.Println("err1=",err1)
        return 
    }
    //读物文件内
    var sliceStr []byte
    var sliceByte = make([]byte,128) 
    for {
        n,err2 := file1.Read(sliceByte)
        if err2 == io.EOF{ //没有str1字符串
            fmt.Println("读取完毕")
            break
        }
        if err2 != nil{
            fmt.Println("err2=",err2)
            return 
        }
        //以分块128进行读取数据，最后一块的为n
        sliceStr = append(sliceStr, sliceByte[:n]...) //切片索引数据进行扩容添加
    }
    fmt.Println(string(sliceStr))
}
```

- 例2，方法二：流的方式bufio读取文件（大文件）

```go
package main
import (
    "fmt"
    "os"
    "io"
    "bufio"
)

func main(){
    //1方法二：以流的方式读文件
    file1, err1 := os.Open("./text.txt")
    defer file1.Close()//关闭文件
    if err1 != nil{
        fmt.Println("err1=",err1)
        return 
    }
    //bufio读取文件
    var fileStr string
    reader := bufio.NewReader(file1)
    for{
        str1,err2 := reader.ReadString('\n') //一次读取一行,使用
        if err2 == io.EOF{
            fileStr += str1  // 可能没有str1字符串
            fmt.Println("读取完毕")
            break
        }
        if err2 != nil{
            fmt.Println("err2=",err2)
            return 
        }
        fileStr += str1
    }
    fmt.Println("fileStr=",fileStr)
}
```

- 例3，方法三：ioutil 一次读完文件数据（小文件100~200M）

```go
package main
import (
    "fmt"
    "io/ioutil"   
)
//ioutil读取文件
func main(){
    byteStr1,err1 := ioutil.ReadFile("./text.txt")
    if err1 != nil{
        fmt.Println("err1=",err1)
        return 
    }
    fmt.Println(byteStr1)
}
```

### 6.2、文件写入

1. 方法一：

- 打开文件 `file,err := os.OpenFile(“文件”, os.O_CREATE | os.O_RDWR, 0666)`
Linux权限配置：文件权限：r(读)04，w(写)02,x（执行）01
- 写文件
`file.Write([] byte (str))` //写入字节切片数据
`file.Write(“字符串”)` //直接写入字符串数据
- 关闭文件 `defer file.Close()`

2. 方法二： bufio

- 打开文件 `file,err := os.OpenFile(“文件”, os.O_CREATE | os.O_RDWR, 0666)`
- 创建writer对象 `writer := bufio.NewWriter(file)`
- 将数据写入缓存 `writer.WriteString(“liwen\r\n”)`
- 将缓存内容写入文件 `writer.Flush()`
- 关闭文件 `defer file.Close()`

3. 方法三 ioutil

```go
str := “liwei”
err := ioutil.WriteFile(“文件”，[]byte(str), 0666)
```

- 例1，file.write()

```go
package main
import (
    "fmt"
    "os"
)

func main(){
    //os.O_CREATE|os.O_WRONLY|os.O_RDWR|os.O_TRUNC|os.O_APPEND|
    file,err1 := os.OpenFile("./textWrite.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
    defer file.Close()
    if err1 != nil{
        fmt.Println("err1=",err1)
        return 
    }
    // for i :=0 ;i < 10; i++{
    //  file.WriteString("直接写入字符串22222 \r\n") // "\r"表示回车符
    // }
    var str = "byte******************"
    file.Write([]byte(str))
}
```

- 例2，方法二： bufio

```go
package main
import (
    "fmt"
    "os"
    "bufio"
    "strconv"

)
func main(){
    //os.O_CREATE|os.O_WRONLY|os.O_RDWR|os.O_TRUNC|os.O_APPEND|
    file,err1 := os.OpenFile("./textWrite.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
    defer file.Close()
    if err1 != nil{
        fmt.Println("err1=",err1)
        return 
    }
    writer := bufio.NewWriter(file)
    //writer.WriteString("lisi is dog!") //将数据写入缓存
    for i :=0 ;i < 10; i++{
        writer.WriteString("直接写入字符串" + strconv.Itoa(i)+  "\r\n" )// "\r"表示回车符
    }
    writer.Flush() //将缓存数据写入文件
}
```

- 例3，方法三 ioutil

```go
package main
import (
    "fmt"
    "io/ioutil"
)
func main(){
    str := "helo liwen"
    err := ioutil.WriteFile("./textWrite.txt",[]byte(str), 0666)
    if err != nil{
        fmt.Println("Write file failed ,err:",err)
        return
    }
}
```

### 6.3、 文件基本操作

#### 6.3.1、文件重命名

```go
err := os.Rename(“原文件”，“目的文件”) //只能在磁盘文件
err2 := os.Rename(“text.txt”,"./renamefile.txt") 
    if err2 != nil{
        fmt.Println("重命名文件失败")
    }
```

#### 6.3.2、复制文件

```go
- 方法一：ioutil
input,err := ioutil.ReadFile(srcFileName)
err = ioutil.WriteFile(dstFileName, input, 0644)
- 方法二： os
source,_  := os.Open(srcFileName)
Defer source.close(source)
destination,_ := os.OpenFile(dstFileName, os.O_CREATE|os.O_WRONLY,0666)
Defer destination.close(destination)
n,err := source.Read(buf)
destination.Write(buf[:n])
```

- 例1,方法一复制文件

```go
package main
import (
    "fmt"
    "io/ioutil"
)

func copy(srcFileName string, dstFileName string)(err error){
    byteStr,err1 := ioutil.ReadFile(srcFileName)
    if err1 != nil{
        fmt.Println("err1=",err1)
        return err1
    }
    err2 := ioutil.WriteFile(dstFileName,byteStr,0666)
    if err2 != nil{
        fmt.Println("err2=",err2)
        return err2
    }
    return nil
}

func main(){
    srcFileName := "./text.txt"
    dstFileName := "./textcopy.txt"
    err := copy(srcFileName, dstFileName)
    
    if err != nil{
        fmt.Println("复制文件失败")
    }else{
        fmt.Println("复制文件成功")
    }
}
```

- 例2 ，复制文件os

```go
package main
import (
    "fmt"
    //"io/ioutil"
    "os"
    "io"

)
func copy(srcFileName string, dstFileName string)(err error){
    readFile,err1 := os.Open(srcFileName)
    if err1 != nil{
        fmt.Println("err1=",err1)
        return err1
    }
    defer readFile.Close()
    writeFile,err2 := os.OpenFile(dstFileName, os.O_CREATE|os.O_WRONLY,0666)
    if err2 != nil{
        fmt.Println("err2=",err2)
        return err2
    }
    defer writeFile.Close()
    var tempSlice = make([]byte, 12800)
    for{
        //读取数据
        n1,err3 := readFile.Read(tempSlice)
        if err3 == io.EOF{
            break
        }
        if err3 != nil{
            return err3
        }
        //写入数据
        if _,err4 := writeFile.Write(tempSlice[:n1]); err4 != nil{
            return err4
        }
    }   
    return nil
}
func main(){
    srcFileName := "./text.txt"
    dstFileName := "./textcopy.txt"
    err := copy(srcFileName, dstFileName)
    
    if err != nil{
        fmt.Println("复制文件失败")
    }else{
        fmt.Println("复制文件成功")
    }
    err2 := os.Rename(dstFileName,"./renamefile.txt") 
    if err2 != nil{
        fmt.Println("重命名文件失败")
    }
}
```

#### 6.3.3、创建目录

创建目录
`os.Mkdir(“文件夹”，0666)`
创建多个目录
`os.MkdirAll(“./dir1/dir2/dir3”,0666)`

```go
package main
import (
    "fmt"
    "os"
)

func main(){
    err1 := os.Mkdir("./dir",0666)
    if err1 != nil{
        fmt.Println("err1=",err1)
        return 
    }
    err2 := os.MkdirAll("./dir1/dir2/dir3", 0666)
    if err2 != nil{
        fmt.Println("err2=",err2)
        return 
    }
}
```

#### 6.3.4、删除文件与目录

os.Remove(“文件或目录”)
os.RemoveAll(“多目录”)

```go
package main
import (
    "fmt"
    "os"
)

func main(){
    //1删除文件
    err1 := os.Remove("./text.txt")
    if err1 != nil{
        fmt.Println("err1=",err1)    
    }
    //2 删除目录
    err2 := os.Remove("./dir")
    if err2 != nil{
        fmt.Println("err2=",err2)    
    }
    //3 删除所有目录
    err3 := os.RemoveAll("./dir1")
    if err3 != nil{
        fmt.Println("err3=",err3)    
    }
} 
```
