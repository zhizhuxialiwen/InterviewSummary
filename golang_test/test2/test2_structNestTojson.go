package main

import (
	"fmt"
	//"sort"
	"encoding/json"
)
//1 结构体嵌套转化json
type Student struct{
	ID int
	Gender string
	Name string
}

type Class struct{
	Title string
	Student []Student
}

func main(){

	//1 结构体嵌套转化json
    c := Class{
		Title: "001班",
		Student: make([]Student,0),
	}

	for i:=0; i < 5; i++{
		s := Student{
			ID : i,
			Gender: "男",
			Name : fmt.Sprintf("stu_%v",i),
		}
		c.Student = append(c.Student , s)
	}

	fmt.Println(c)
	strByte, err := json.Marshal(c)
	if err != nil{
		fmt.Println("json 转化失败")
	}else{
		strJson := string(strByte)
		fmt.Println(strJson)
	}
	
	
	//2 json字符串转化结构体嵌套
	strJson2 := `{"Title":"001班","Student":[{"ID":0,"Gender":"男","Name":"stu_0"},{"ID":1,"Gender":"男","Name":"stu_1"},{"ID":2,"Gender":"男","Name":"stu_2"},{"ID":3,"Gender":"男","Name":"stu_3"},{"ID":4,"Gender":"男","Name":"stu_4"}]}`

	var c2  = &Class{}
	err1 := json.Unmarshal([]byte(strJson2), c2)

	if err1 != nil{
		fmt.Println("err")
	}else{
		fmt.Printf("%#v \n", c2)
		fmt.Printf("%v \n", c2.Title)
	}
	
}