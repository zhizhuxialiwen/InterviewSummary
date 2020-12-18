package main

import (
	"fmt"
)

//1接口嵌套
type Animals1 interface {
	SetName(string)
	
}

type Animals2 interface {
	GetName() string
}

type AnimalsSum interface {
	Animals1
	Animals2
}

// 狗 :结构必须实现接口所有的方法
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
	var aSum AnimalsSum = d 
	aSum.SetName("小花狗")
	fmt.Println(aSum.GetName()) //小花狗
}