package main

import (
	"fmt"
	//"sort"
	//"strings"
)
//1、自定义类型:函数和变量
type calcType func(int, int)int
type myInt int

func add(x,y int)int{
	return x + y
}

func sub(x,y int)int{
	return x - y
}

//2 函数作为另一个函数的参数 : 正常参数或自定义函数
func calcSum(x,y int, cb calcType )int{
	return cb(x,y)
}

func calcSub(x,y int, cb1 func(int, int)int)int{
	return cb1(x,y)
}

// 3 函数作为返回值
func doFuncType(outType string) calcType{
	switch outType{
	case "+":
		return add
	case "-":
		return sub
	case "*":
		return func(x,y int)int{
			return x * y
		}
	default:
		return nil
	}
}
func main(){
	var typeFunc calcType
	typeFunc = add
	fmt.Printf("typeFunc的类型： %T \n", typeFunc) //typeFunc的类型： main.calc

	sum := typeFunc(5,6)
	fmt.Println(sum)  //11

	var a1 int = 2
	var a2 myInt = 5
	fmt.Printf("a1类型：%T，a2类型：%T \n",a1,a2)  //a1类型：int，a2类型：main.myInt
	fmt.Println(a1+int(a2))  //7

	sum1 := calcSum(33, 44, add)
	fmt.Println(sum1)  //77

	sub1 := calcSub(33, 22,sub)
	fmt.Println(sub1)  //11

	// 匿名函数
	sum2 := calcSum(22,34, func(x,y int)int{
		return x + y
	})
	fmt.Println(sum2) //56
	//匿名自执行函数
    func(x,y int) {
		fmt.Println("test...")
		fmt.Println(x+y)
	}(3,5)
	// 匿名函数
	var fn = func(x,y int) int{
		return x + y
	}
	fmt.Println(fn(3,4))

	//函数作为返回值
	var outFunc = doFuncType("+")
	fmt.Println(outFunc(3,4)) //7

}