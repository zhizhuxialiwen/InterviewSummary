package main

import (
	"fmt"
	"errors"
	//"sort"
)
//1 panic ： 结束执行，抛出异常。
//2 panic 可以再任何地方引发，但recover只有在defer调用的函数有效;recover监听panic异常,不会终止执行

func func1(){
	fmt.Println("fnuc1")
}

func func2(){
	defer func(){
		err := recover() //监听panic异常
		if err != nil{
			fmt.Println("err:",err) //err: 抛出异常
		}
	}()
	panic("抛出异常")
	
}

//a= 10, b= 0
func func3(a int, b int) int{
	defer func(){
		err := recover() //监听panic异常
		if err != nil{
			fmt.Println("err:",err) //err: runtime error: integer divide by zero
		}
	}()
	//panic("抛出异常")
	return a/b  //10/0
	
}

func readFile(fileName string) error{
	if fileName == "main.go"{
		return nil
	}else{
		return errors.New("读取文件失败")
	}
}

func myFunc(){
	defer func(){
		err := recover()
		if err != nil{
			fmt.Println("给管理员发送邮件")
		}
	}()
	err := readFile("main.go")
	if err != nil{
		panic(err)
	}
}

func main(){
	func1()
	func2()
	func3(10, 0)
	fmt.Println("结束")

	myFunc()
	fmt.Println("继续执行...")
}