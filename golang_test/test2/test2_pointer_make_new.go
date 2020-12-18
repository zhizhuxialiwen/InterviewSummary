package main

import (
	"fmt"
	//"errors"
	//"sort"
)
func func2(x int){
	x = 10
}
func func3(x *int){
	*x = 40
}

func main(){
	//指针：引用数据类型
	//1.1、go 的变量都有对应的地址，包含指针变量含有对应地址
	var a1 = 10
	var p1 = &a1 // 值：0xc0000120b0,类型：*int, 地址：0xc000006028, 指针指向地址的值：10
	fmt.Printf("值：%v,类型：%T, 地址：%v \n",a1, a1,&a1) //值：10,类型：int, 地址：0xc0000120b0
	//值：0xc0000120b0,类型：*int, 地址：0xc000006028, 指针指向地址的值：10
	fmt.Printf("值：%v,类型：%T, 地址：%v, 指针指向地址的值：%v\n",p1, p1,&p1,*p1) 
	fmt.Println(p1)  //a的地址：0xc0000120b0
	fmt.Println(*p1) //*p对应a的值：10
	*p1 = 21
	fmt.Println(*p1) //*p对应a的值：21
	fmt.Println("a1=",a1) // a1= 21

	//1.2 改变指针的地址
	var a2 = 5
	func2(a2)
	fmt.Println(a2)  //5
	func3(&a2)
	fmt.Println(a2)  //40: 改变指针指向地址的值
	//3 make：切片、map集合
	//3.1 make:切片
	//引用数据类型：切片、map集合、指针；
	//引用数据类型必须声明，且分配内存空间，否则无法存储，使用make 和new 创建存储空间
	var slice1 = make([]int, 4,4)
	slice1[0] = 11
	slice1[1] = 22
	slice1 = append(slice1,1,2,4) //append进行扩容
	fmt.Println(slice1)  //[11 22 0 0 1 2 4]

	//3.2 new 针对指针类型
	var a3 *int
	a3 = new(int)
	//值：0xc000012160,类型：*int, 地址：0xc000006038, 指针指向地址的值：0
	fmt.Printf("值：%v,类型：%T, 地址：%v, 指针指向地址的值：%v\n",a3, a3,&a3,*a3) 
    *a3 = 100
	fmt.Println(*a3)  //100
	/*3.3 make与new区别
	(1) make只有用于切片slice、集合map、通道channel的初始化，返回值还是引用类型
	(2) new 用于类型的内存分配，并且内存对应的值为类型0值，返回值为指向类存的指针。
	*/
}