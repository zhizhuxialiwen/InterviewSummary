package main

import (
	"fmt"
	//"sort"
	"encoding/json"
)
//结构体转json：必须大写
/*
type PersonStru struct{
	name string
	age int
	sex string  //私有结构体字段无法转化json
}
*/
type PersonStru struct{
	Name string
	Age int
	Sex string
}

// 结构体的标签：结构体变量为小写
type PersonStruLable struct{
	Name string `json:"id"`
	Age int  `json:"age"`
	Sex string `json:"sex"`
	Sno string `json:"sno"`
}
func main(){
	//1 结构体对象转化json字符串
	
	var personStru1 = PersonStru{
		Name:"liwe1",
		Age:20,
		Sex:"男",
	}
    /*
	var personStru1 = PersonStru{
		name:"liwe1",
		age:20,
		sex:"男",
	}
	*/
	fmt.Printf("%#v \n", personStru1)
	//json的[]byte切片
	jsonByte,_ := json.Marshal(personStru1)
	//Byte切片转化json字符串
	jsonStr := string(jsonByte) //main.PersonStru{Name:"liwe1", Age:20, Sex:"男"}
	fmt.Printf("%#v \n", jsonStr) //"{\"Name\":\"liwe1\",\"Age\":20,\"Sex\":\"男\"}"
	fmt.Printf("%v \n", jsonStr)  //{"Name":"liwe1","Age":20,"Sex":"男"}
	
	//2 、json字符串转化结构体：使用反引号``
	var jsonStr1 = `{"Name":"liwe1","Age":20,"Sex":"男"}`
	var personStru2 PersonStru
	err := json.Unmarshal([]byte(jsonStr1), &personStru2)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Printf("%#v \n", personStru2) //main.PersonStru{Name:"liwe1", Age:20, Sex:"男"}
    fmt.Printf("%v \n", personStru2.Name) //liwe1
	// 3 结构体的标签：结构体变量为小写
	var personStruLable1 = PersonStruLable{
		Name:"liwe1",
		Age:23,
		Sex:"男",
		Sno:"123456",
	}
	//json的[]byte切片
	jsonByte3, _ := json.Marshal(personStruLable1)
	//Byte切片转化json字符串
	jsonStr3 := string(jsonByte3) 
	fmt.Printf("%v \n", jsonStr3)  //{"id":"liwe1","age":23,"sex":"男","sno":"123456"}
	
}