package main
import (
	"fmt"
	//"runtime"
	//"sync"
	//"time"
)

func test(){
	defer func(){
		if err := recover(); err != nil{
			fmt.Println("test 发送异常",err)
		}
	}()

	//var myMap map[int]string
	//myMap[0] = "golang"  //error
	panic("异常。。。")
}
func main(){
	go test()
	fmt.Printf("主线程完成。。。\n")
}