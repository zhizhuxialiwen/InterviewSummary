package main

import "fmt"

func main(){

	//1 常量 : 无法改变
	const pi = 3.14159
	fmt.Println(pi)
	//pi = 444 //错误
	//2 声明多个常量与变量一样，
	const (
		a1 = "liwen1"
		a2 = "liwen2"

	)
	fmt.Println(a1,a2)
	//3 const,若省略值与上面一样
	const (
		n1 = 101
		n2 
		n3
	)
	fmt.Println(n1,n2,n3)

	//4.1 const ,都会让iota初始化为0【自增长】
	const n4 = iota //n4=0
	fmt.Println(n4)
	//4.2 _跳过，但是增1
	const (
		n5 = iota //n5 =0
		_
		n6
		n7
	)
	fmt.Println(n5,n6,n7)
    //4.3 iota插队
	const (
		n8 = iota //n8 =0
		n9 = 100  //n9=100
		n10 = iota //n10 = 2
		n11  //n11 =3
	)
	fmt.Println(n8,n9,n10,n11)

	//4.4 多个iota定义在一行
	const (
		n12,n13  = iota + 1,  iota + 2//1 2
		n14,n15  //2 3
		n16,n17  //3 4
	)
	fmt.Println(n12,n13,n14,n15,n16,n17)
	//变量风格
	//1）没有分号“;”
	//2）区分变量大小写
	//3）变量风格：大小驼峰： userName (小驼峰：私有)，UserName（大驼峰：公有）
	//4) 等于 = 有空格,按保存就会有空格
}