package main

import (
	"fmt"
	//"sort"
	//"strings"
)

//1 、递归函数：1自己调用自己；2结束条件 ；3 递减
func recursionSum(n int) int{
	if n > 1{
		return n + recursionSum(n - 1) //第一次：100 + recursion(99);第二次：100 + 99 + recursion(98);...
	}else{
		return 1
	}
}

func recuMult(n int) int{
	if n > 1{
		return n*recuMult(n - 1) //第一次：5* recuMult(4);第二次：5*4*recuMult(3);...
	}else{
		return 1
	}
}
/*2、闭包特点：1、可以让一个变量常驻内存；2、可以让一个变量不污染全局
(1) 闭包有权访问另一个函数作用域中的变量的函数
(2) 一个函数访问另一个函数
写法：函数里面嵌套一个函数，最后返回里面的函数
*/
func adder1() func() int {
	var i = 10 
	return func() int {
		return i + 1
	}
}

func adder2() func(y int) int {
	var i = 10 
	return func(y int) int {
		i += y // i为常住内存: 第二次调用i会使用第一次i的值，不污染全局
		return i 
	}
}

func main(){
	//1、循环：1-100的和
	sum := 0
	for i := 1; i <= 100; i++{
		sum += i
	}
	fmt.Println("sum=",sum)  //5050
	//2、递归：1-100的和:100+99+98+...+1
	sum1 := recursionSum(100)
	fmt.Println("sum1=",sum1)
	//2.1 递归： 实现的5阶乘：5*4*3*2*1
	rec1 := recuMult(5)
	fmt.Println("rec1=",rec1)

	//闭包
	var fn1 = adder1() //表示执行方法
	fmt.Println(fn1()) //11
	fmt.Println(fn1()) //11

	var fn2 = adder2() //表示执行方法
	fmt.Println(fn2(10)) //20
	fmt.Println(fn2(10)) //30
}