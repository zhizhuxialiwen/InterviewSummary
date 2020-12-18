package main

import (
	"fmt"
	//"sort"
)

//1 、结构体是值类型，改变副本不影响主本: 
//go 的结构体是相互独立的，不会影响
type personStru struct{
	name string
	age int
	sex string
}
//2. go没有类class，但是定义结构体方法：结构体调用方法
func (personStruTmp personStru) printInfo(){
	fmt.Printf("name=%v, age=%v, sex=%v \n",personStruTmp.name,personStruTmp.age, personStruTmp.sex)
}
//2.1结构体无法修改结构体属性
func (personStruTmp1 personStru) setInfo1(name string, age int){
	personStruTmp1.name = name
	personStruTmp1.age = age
}
//2.2 结构体指针修改结构体属性
func (personStruTmp2 *personStru) setInfo2(name string, age int){
	(*personStruTmp2).name = name
	(*personStruTmp2).age = age
	
}

// 3 自定义类型添加方法 : 自定义变量调用方法
type myInt int
func (myInt1 myInt) printTypeInfo(){
	fmt.Println("自定义类型的自定义方法")
}

func main(){
    //1 、结构体是值类型，改变副本不影响主本
	var personStru1 = personStru{
		name:"liwe1",
		age:20,
		sex:"男",
	}
	personStru2 := personStru1
	personStru2.name = " 李文"
	//值：main.personStru{name:"liwe1", age:20, sex:"男"},类型： main.personStru
	fmt.Printf("值：%#v,类型： %T \n",personStru1,personStru1) 
	//值：main.personStru{name:" 李文", age:20, sex:"男"},类型： main.personStru
	fmt.Printf("值：%#v,类型： %T \n",personStru2,personStru2) 
	//2. go没有类class，定义结构体方法
	personStru1.printInfo()  //name=liwe1, age=20, sex=男	
	var personStru3 = personStru{
		name:"lisi",
		age:29,
		sex:"男",
	}
	personStru3.printInfo() //name=lisi, age=29, sex=男  

	//setInfo
	personStru3.setInfo1("李威", 32)
	personStru3.printInfo() //name=lisi, age=29, sex=男

	personStru3.setInfo2("李威", 32)
	personStru3.printInfo()  //name=李威, age=32, sex=男

	var a1 myInt 
	a1.printTypeInfo() //自定义类型的自定义方法
}