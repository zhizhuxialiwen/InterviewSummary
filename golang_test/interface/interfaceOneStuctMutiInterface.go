package main

import (
	"fmt"
)

//1个结构体实现多个接口
type Animals1 interface {
	SetName(string)
}

type Animals2 interface {
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

func main(){
	//Dog 
	var d = &Dog{
		Name : "金毛",
	}
	var a1 Animals1 = d 
	var a2 Animals2 = d 
	a1.SetName("小花狗")
	fmt.Println(a2.GetName()) //小花狗
}