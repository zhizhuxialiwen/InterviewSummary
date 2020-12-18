package main

import (
	"fmt"
	//"runtime"
	"sync"
	"time"
)

//1 开始多协程：并行

var wg sync.WaitGroup
func test(num int){
	defer wg.Done()
	for i := 0; i < 5; i++{
		fmt.Printf("协程%v打印第%v条数据\n",num, i)
		time.Sleep(time.Millisecond*100)
	}
	//wg.Done()//协程计数器减1
}

func main(){
	for i := 0; i < 5; i++{
		wg.Add(1)
		go test(i)
	}
	wg.Wait()
	fmt.Println("关闭主线程")
}