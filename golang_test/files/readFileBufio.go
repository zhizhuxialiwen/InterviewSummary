package main
import (
	"fmt"
	//"runtime"
	//"sync"
	//"time"
	//"reflect"
	"os"
	"io"
	"bufio"
	
)

func main(){
	//1方法二：以流的方式读文件
	file1, err1 := os.Open("./text.txt")
	defer file1.Close()//关闭文件
    if err1 != nil{
		fmt.Println("err1=",err1)
		return 
	}
	//bufio读取文件
	var fileStr string
	reader := bufio.NewReader(file1)
	for{
		str1,err2 := reader.ReadString('\n') //一次读取一行,使用
		if err2 == io.EOF{
			fileStr += str1  // 可能没有str1字符串
			fmt.Println("读取完毕")
			break
		}
		if err2 != nil{
			fmt.Println("err2=",err2)
			return 
		}
		fileStr += str1
		
	}
	fmt.Println("fileStr=",fileStr)
	
}