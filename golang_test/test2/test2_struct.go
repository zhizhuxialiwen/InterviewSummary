package main

import (
	"fmt"
	//"sort"
)
//1、自定义类型
type myInt int
// type myFn func(int, int )int

//2、类型别名
type myFloat = float64
//3、结构体：结构体（personStru）小写是私有的，结构体（PersonStru）大写是公有的
type personStru struct{
	name string
	age int
	sex string
}

func main(){
	var a1 myInt = 10
    fmt.Printf("值：%v,类型： %T \n",a1,a1) //值：10,类型： main.myInt
	var f1 myFloat = 12.33
	fmt.Printf("值：%v,类型： %T \n",f1,f1) //值：12.33,类型： float64
	// 1 实例化结构体
	//1.1方法一： 实例化结构体
	var personStru1 personStru 
	personStru1.name = "liwen"
	personStru1.age = 20
	personStru1.sex = "男"
	//值：{liwen 20 男},类型： main.personStru
	fmt.Printf("值：%v,类型： %T \n",personStru1,personStru1) 
    //值：main.personStru{name:"liwen", age:20, sex:"男"},类型： main.personStru
	fmt.Printf("值：%#v,类型： %T \n",personStru1,personStru1) 
	//1.2 方法二： new实例化结构体： go 支持对结构体指针直接使用:
	//personStru2.name = "liwei"  等于 (*personStru2).name = "liwei" 
	var personStru2 = new(personStru)
	// personStru2.name = "liwei"  
	// personStru2.age = 25
	// personStru2.sex = "男"
	(*personStru2).name = "liwei" 
	(*personStru2).age = 25
	(* personStru2).sex = "男"
	//值：&main.personStru{name:"liwei", age:25, sex:"男"},类型： *main.personStru
	fmt.Printf("值：%#v,类型： %T \n",personStru2,personStru2) 
	//1.3 方法三：引用地址实例化
	var personStru3 = &personStru{}
	personStru3.name = "liwe"
	personStru3.age = 23
	personStru3.sex = "男"
    //值：&main.personStru{name:"liwe", age:23, sex:"男"},类型： *main.personStru
	fmt.Printf("值：%#v,类型： %T \n",personStru3,personStru3) 
	// 1.4 方法四：初始化的实例化
	var personStru4 = personStru{
		name:"liwe1",
		age:20,
		sex:"男",
	}
	//值：main.personStru{name:"liwe1", age:20, sex:"男"},类型： main.personStru
	fmt.Printf("值：%#v,类型： %T \n",personStru4,personStru4) 
	// 1.5 方法五：地址初始化的实例化
	var personStru5 = &personStru{
		name:"liwe1",
		age:20,
		sex:"男", //此逗号不能省略
	}
	//值：&main.personStru{name:"liwe1", age:20, sex:"男"},类型： *main.personStru
	fmt.Printf("值：%#v,类型： %T \n",personStru5,personStru5) 

	//总结： type 用于自定义类型、结构体
}