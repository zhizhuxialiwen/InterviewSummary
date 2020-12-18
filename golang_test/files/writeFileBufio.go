package main
import (
	"fmt"
	//"runtime"
	//"sync"
	//"time"
	//"reflect"
	"os"
	//"io"
	"bufio"
	"strconv"

)



func main(){
	//os.O_CREATE|os.O_WRONLY|os.O_RDWR|os.O_TRUNC|os.O_APPEND|
	file,err1 := os.OpenFile("./textWrite.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
	defer file.Close()
	if err1 != nil{
		fmt.Println("err1=",err1)
		return 
	}
	writer := bufio.NewWriter(file)
	//writer.WriteString("lisi is dog!") //将数据写入缓存
	for i :=0 ;i < 10; i++{
		file.WriteString("直接写入字符串" + strconv.Itoa(i)+  "\r\n" )// "\r"表示回车符
	}

	writer.Flush() //将缓存数据写入文件
}