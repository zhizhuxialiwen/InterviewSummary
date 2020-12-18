package main

import "fmt"

func getUserInfo() (string, int){
	return "liwen11", 11
}

func main(){

	/*
	1 var 声明
	var 变量名称 类型
	go变量名称为字母、下划线、数字组成，首个字符不能为数字
	go关键字不能作为变量名称
	*/

	//初始化1：声明再定义，不支持重复声明
	var username1 string
	username1 = "liwen"
	fmt.Println(username1)
    //初始化2：直接定义
	var username2 string = "liwen2"
	fmt.Println(username2)
    //初始化3 :类型推到
	var username3 = "liwen3"
	fmt.Println(username3)
	/*
	2 一次定义多个变量
	var 变量名称 ，变量名称 类型
	 或
	 var(
		变量名称 类型
		变量名称 类型
	 )
	*/

	var a1, a2 string
	a1 = "liwen1"
	a2 = "liwen2"
	//a2 = 123 //错误写法，定义类型与赋值类型一致
	fmt.Println(a1, a2)

	var(
		a3 string 
		a4 int
	)
	a3 = "liwen3"
	a4 = 444
	fmt.Println(a3, a4)

	var(
		a5 = "liwen5" 
		a6 =  666
	)
	fmt.Println(a5, a6)

	/*
	3 短变量声明法：在函数内部使用更为简略 := 方式声明初始化变量
	注意：只能用于局部变量，不能用于全局变量
	*/
	a7 := "liwen7"
	a8 := 777
	fmt.Println(a7, a8)
	fmt.Printf("a7的类型是%T \n",a7)
	
	a9,a10 := "liwen9", 1000
	fmt.Println(a9, a10)

	/*
	4 匿名变量 ：在使用多重赋值是，若忽略某个值，可以使用匿名变量（anonymous varaible）
	匿名变量使用下划线 _
	func getUserInfo() (string, int){
		return "liwen11", 11
	}
	*/
	var username11, age11 =  getUserInfo()
	fmt.Println(username11, age11)
	
    //匿名变量不占用命名空间，不分配空间，因此匿名变量之间不存在重复声明
	var username12,_ =  getUserInfo()
	fmt.Println(username12)
    
	var  _, age12 =  getUserInfo()
	fmt.Println(age12)
	



}
