package print

import (
	"fmt"
	"package/algorithm"
)

func init(){
    fmt.Println("print.init()...")
}

func PrintInfo(){
	retSubInt := algorithm.SubInt(5,4)  
	fmt.Println(retSubInt)
	fmt.Println("打印信息")
}