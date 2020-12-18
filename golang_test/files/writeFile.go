package main
import (
	"fmt"
	//"runtime"
	//"sync"
	//"time"
	//"reflect"
	"os"
	//"io"
)

func main(){
	//os.O_CREATE|os.O_WRONLY|os.O_RDWR|os.O_TRUNC|os.O_APPEND|
	file,err1 := os.OpenFile("./textWrite.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
	defer file.Close()
	if err1 != nil{
		fmt.Println("err1=",err1)
		return 
	}

	// for i :=0 ;i < 10; i++{
	// 	file.WriteString("直接写入字符串22222 \r\n") // "\r"表示回车符
	// }

	var str = "byte******************"
	file.Write([]byte(str))
}