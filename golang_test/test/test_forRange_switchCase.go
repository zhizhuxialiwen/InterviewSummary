package main

import( 
	"fmt"
	//"strconv"
)

//import "strings"
//import "unsafe"

func main(){
	//1、for...range
	//1.1
	var str1 = "liwen is good!"
	for k,v := range str1{
		fmt.Printf("key=%v, value=%c \n",k,v)
	}
	//1.2切片
	var arr1 = []string{"liwen1", "liwei", "lishan"}
	for _, var2 := range arr1{
		fmt.Println("value=",var2)
	}

	//2、switch ... case...break
	//2.1
	var ext1 = ".html"
	switch ext1{
	case ".html":
		fmt.Println("text/html")	
	case ".css":
		fmt.Println("text/css")	
	case ".js":
		fmt.Println("text/js")		
	default:
		fmt.Println("找不到此后缀")
		
	}
	//2.2
	switch ext1 := ".css1";ext1{
	case ".html":
		fmt.Println("text/html")		
	case ".css":
		fmt.Println("text/css")		
	case ".js":
		fmt.Println("text/js")	
	default:
		fmt.Println("找不到此后缀")		
	}
	//2.3 一个分支含有多个值
	var score  = "d" //abc及格，d不及格
	switch score{
	case "a","b","c":
		fmt.Println("及格")	
	case "d":
		fmt.Println("不及格")		
	default:
		fmt.Println("没有找到。")	
	}
	//2.4 分支可以使用表达式，switch没有连接变量age
	var age = 40
	switch {
	case age <24:
		fmt.Println("好好学习")	
	case age >= 24 && age <= 60:
		fmt.Println("好好挣钱")
	case age > 60:
		fmt.Println("注意身体")	   
	default:
		fmt.Println("输入错误")
	}

}