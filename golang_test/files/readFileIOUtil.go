package main
import (
	"fmt"
	//"runtime"
	//"sync"
	//"time"
	//"reflect"
	//"os"
	"io/ioutil"
	//"bufio"
	
)
//ioutil读取文件
func main(){
	byteStr1,err1 := ioutil.ReadFile("./text.txt")
	if err1 != nil{
		fmt.Println("err1=",err1)
		return 
	}
	fmt.Println(byteStr1)
}