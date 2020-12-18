package main

import (
	"fmt"
)

//1个接口2个结构体
type Animals interface {
	SetName(string)
	GetName() string
}
// 狗
type Dog struct{
	Name string
}

//结构体是值类型，改变需要使用指针类型
func (d *Dog) SetName(name string){
	(*d).Name = name //d.Name = name
}

func (d Dog) GetName() string{
	return d.Name
}

//猫

type Cat struct{
	Name string
}

//结构体是值类型，改变需要使用指针类型
func (c *Cat) SetName(name string){
	(*c).Name = name //d.Name = name
}

func (c Cat) GetName() string{
	return c.Name
}

func main(){
	//Dog 
	var d = &Dog{
		Name : "金毛",
	}
	var a Animals = d 
	fmt.Println(a.GetName()) //金毛
	a.SetName("哈士奇")
	fmt.Println(a.GetName()) //哈士奇
	//cat
	var c = &Cat{
		Name : "黑猫",
	}
	var a1 Animals = c 
	fmt.Println(a1.GetName()) //金毛
	a1.SetName("橘猫")
	fmt.Println(a1.GetName()) //哈士奇
}