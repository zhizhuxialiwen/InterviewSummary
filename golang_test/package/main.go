package main

import (
	"fmt"
	//"sort"
	//"encoding/json"
	//"package/algorithm"
	//"package/print"

	//匿名 ： _ "package/print"
	//别名： p "package/print"
	p "package/print"
)
//系统调用：优先于mian()执行
func init(){
    fmt.Println("main.init()...")
}

func main(){
	// retSubInt := algorithm.SubInt(5,4)  ////方法大写.SubInt
	// retSubFloat := algorithm.SubFloat(5.43,4.21)
	// fmt.Println(retSubInt)
	// fmt.Println(retSubFloat)
	//
	p.PrintInfo()
}
