package main

import (
	"fmt"
	//"sort"
	//"strings"
)
//1、defer语句：先被defer的语句最后执行，反之最先执行。
//若含有多个defer的语句，执行顺序：先执行最后一个defer语句，再执行n-1个defer语句
//若一个defer的函数，先执行第一个语句，再执行2个语句
/*
defer： 当函数内含有defer的 匿名函数，首先执行函数内的语句; 
若return a1 与 函数返回值RET a1变量一致，则继续执行defer 匿名执行语句，最后执行return a1变量；
反之，则return 结束,不执行defer匿名语句。
*/

func deferFun1(){
	fmt.Println("开始")
    defer func(){
		fmt.Println(11)
		fmt.Println(12)
	}()  //匿名执行语句
	fmt.Println("结束")
	/*
	开始
	结束
	11
	12
	*/
}

//调用方法返回值为0 
func deferFun2() int{
    var a1 int //0
    defer func(){
		a1++
	}()  //defer 匿名执行语句
	return a1 //0
}

//匿名返回值a1 int为1
func deferFun3()(a1 int) {
    defer func(){
		a1++
	}()  //defer 匿名执行语句
	return a1 //1
}

func deferFun4()(a1 int) {
	a1 = 5
    defer func(){
		a1++
	}()  //匿名执行语句
	return a1  //6

}


func deferFun5()(x int) {
    defer func(x int){
		//x = 0
		x++
	}(x)  //匿名执行语句
	return 5  //5
}

func calc(index string, a, b int) int{
	ret := a + b
	fmt.Println(index,a,b,ret)
	return ret
}

func main(){
	//1 defer 演示
	// fmt.Println("开始")
	// defer fmt.Println(1)
	// defer fmt.Println(2)
	// defer fmt.Println(3)
	// fmt.Println("结束")
	/* 输出结果：
	开始
	结束
	3
	2
	1
	*/
	//2 、defer在命名返回值和匿名返回

	//deferFun1()
	a2 := deferFun2()
	fmt.Println(a2)  //0
   
	a3 := deferFun3()
	fmt.Println(a3) //1

	a4 := deferFun4()
	fmt.Println(a4) //5

	a5 := deferFun5()
	fmt.Println(a5)//5

	x := 1
	y := 2
	defer calc("AA",x,calc("A", x, y))
	x = 10
	defer calc("BB",x,calc("B", x, y))
	y = 20
	/*
	注册顺序：
	defer calc("AA",x,calc("A", x, y))
	defer calc("BB",x,calc("B", x, y))
	执行顺序：
	defer calc("BB",x,calc("B", x, y))
    defer calc("AA",x,calc("A", x, y))
	*/
	//1、calc("A", x, y) : A 1 2  --> 3 
	//2、calc("B", x, y) ：B 10 2 --> 12
	//3、defer calc("BB",x,calc("B", x, y))：BB 10 12-->22
	//4、defer calc("AA",x,calc("A", x, y)): AA 1 3 --> 4
}