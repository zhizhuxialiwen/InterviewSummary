package main
import (
	"fmt"
	//"io/ioutil"
	"os"
	//"io"
)


func main(){
	//1删除文件
	err1 := os.Remove("./text.txt")
	if err1 != nil{
		fmt.Println("err1=",err1)	 
	}
	//2 删除目录
	err2 := os.Remove("./dir")
	if err2 != nil{
		fmt.Println("err2=",err2)	 
	}
	//3 删除所有目录
	err3 := os.RemoveAll("./dir1")
	if err3 != nil{
		fmt.Println("err3=",err3)	 
	}
}