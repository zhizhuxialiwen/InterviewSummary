package main

import (
	"fmt"
)

type Address struct{
	Name string
	Phone int
}

func main(){
	var userInfo = make(map[string] interface{})
	userInfo["username"] = "liwen"
	userInfo["age"] = 20
	userInfo["hobby"] = []string{"打篮球","羽毛球"}

	fmt.Println(userInfo["username"]) //liwen
	fmt.Println(userInfo["age"])  //20
	//fmt.Println(userInfo["hobby"][1])  //错误，空接口不能使用索引
    hobby2, _ := userInfo["hobby"].([] string) //类型判断获取其对象,索引
	fmt.Println(hobby2[1]) //羽毛球

	var address = Address{
		Name:"lwien",
		Phone: 123344,
	}
	userInfo["address"] = address
	fmt.Println(userInfo["address"]) //{lwien 123344}

	//fmt.Println(userInfo["address"].Name) //错误，空接口没有对应的属性
   address2 , _ := userInfo["address"].(Address)//类型断言获取其对象
   fmt.Println(address2.Name, address2.Phone) //{lwien 123344}
	
}