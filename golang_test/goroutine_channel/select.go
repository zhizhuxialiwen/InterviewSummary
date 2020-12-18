package main
import (
	"fmt"
	//"runtime"
	//"sync"
	"time"
)

func main(){

	intChan := make(chan int , 4)
	for i :=0 ; i < 4; i++{
		intChan <- i;
	}

	strChan := make(chan string, 4)
	for i :=0 ; i < 4; i++{
		strChan <- "hello" + fmt.Sprintf("%d",i);
	}
	//select 不需要关闭通道，多路复用：同时从多个通道获取数据
	for {
		select{
		case v := <- intChan:
			fmt.Printf("从inChan读取数据 % d \n", v)
			time.Sleep(time.Millisecond*50)
		case v := <- strChan:
			fmt.Printf("从strChan读取数据 % v \n", v)
			time.Sleep(time.Millisecond*50)
		default:
			fmt.Printf("读取数据完毕  \n", )
			return //注意退出...
		}
	}
}