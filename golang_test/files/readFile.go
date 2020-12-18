package main
import (
	"fmt"
	//"runtime"
	//"sync"
	//"time"
	//"reflect"
	"os"
	"io"
)

func main(){
	//1 方法一：只读方式打开文件
	file1, err1 := os.Open("./text.txt")
	defer file1.Close()//关闭文件
    if err1 != nil{
		fmt.Println("err1=",err1)
		return 
	}
	//读物文件内
	var sliceStr []byte
	var sliceByte = make([]byte,128)
	
	for {
		n,err2 := file1.Read(sliceByte)
		if err2 == io.EOF{ //没有str1字符串
			fmt.Println("读取完毕")
			break
		}
		if err2 != nil{
			fmt.Println("err2=",err2)
			return 
		}
		//以分块128进行读取数据，最后一块的为n
		sliceStr = append(sliceStr, sliceByte[:n]...) //切片索引数据进行扩容添加
	}
	
	fmt.Println(string(sliceStr))

}

