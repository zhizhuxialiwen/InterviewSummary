package main

import (
	"fmt"
)
//空接口
//1.1空接口：任何数据类型都可以实现空接口
type NullInterface interface{}

//1.2 空接口参数可以是任意类型
func show(a interface{}){
	fmt.Printf("值：%v,类型：%T \n", a,a)
}

//2类型断言：
//2.1 if 方法： 接口变量名.(类型)
func printInfo(b interface{}) {
	if str1,ok := b.(string);ok{
		fmt.Println("string类型:",str1)
	}else if _,ok := b.(int);ok{
		fmt.Println("int")
	}else if _,ok := b.(float64);ok{
		fmt.Println("float64")
	}else{
		fmt.Println("null")
	}
}
//2.2 方法2：switch的b1.(type)表示变量判断类型，只能使用在switch
func printInfo1(b1 interface{}) {

	switch b1.(type){
	case string:
		fmt.Println("string类型")
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("float64")
	default:
		fmt.Println("null")
	}
	
}


func main(){
	var nullInterface1 NullInterface
	var str1 = "liwen"
	nullInterface1 = str1
	fmt.Printf("值：%v,类型：%T \n", nullInterface1,nullInterface1)

	var n1 = 11
	nullInterface1 = n1
	fmt.Printf("值：%v,类型：%T \n", nullInterface1,nullInterface1)

	//2、空接口也是一个类型
	var nullInterface2 interface{}
	var f1 = 2.3443
	nullInterface2= f1
	fmt.Printf("值：%v,类型：%T \n", nullInterface2,nullInterface2)
	
	// 空接口参数可以是任意类型
	show(1)
	show("lwww")
	var slice1 = []interface{}{1,2.33,"liwen"}
	fmt.Println(slice1)
	
	//2类型断言：对类型判断
	//传入不同类型，不同类型实现不同的功能
	var nullInterface3 interface{}
	nullInterface3 = "厉害"
	v, ok := nullInterface3.(string )
	if ok{
		fmt.Println("nullInterface3 is string, value is :", v)
	}else{
		fmt.Println("断言失败")
	}

	printInfo("liwen")

	printInfo1(1111)
}