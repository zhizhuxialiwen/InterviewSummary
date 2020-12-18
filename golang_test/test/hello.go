package main

import "fmt"

func main(){
	//1 Print与 Println区别
	//fmt.Print("hello word!\n")  //无法换行；一次输出多次数据，无空格
	//fmt.Println("hello word!") //可以换行；一次输出多次数据，空格	
	//fmt.Printf("hello word!\n") //
	// ctrl +/ 表示注释快捷键
    /*
	fmt.Print("a")
	fmt.Print("b")
	fmt.Print("c\n")

	fmt.Println("a")
	fmt.Println("b")
	fmt.Println("c")
	
	fmt.Print("a","b","c\n")
	fmt.Println("a","b","c")
  
	//2 Println 与Printf的区别
	//go定义必须使用
	var a = "aaaa"
	fmt.Println(a)   //非格式化输出
	fmt.Printf("%v\n",a)  //格式化输出
     */
	var b int = 1
	var c int = 2


	var d int = 3
	fmt.Println("b=", b, "c=", c, "d=",d) 
	fmt.Printf("b=%v c=%v d=%v \n",b,c,d) //相应值的默认格式

	//类型推到方式定义变量
	e := 10
	f := 20
	g := 30
	fmt.Printf("e=%v f=%v g=%v \n",e,f,g) //相应值的默认格式
	fmt.Printf("e=%v e的类型是%T \n", e,e)
	
}
	


