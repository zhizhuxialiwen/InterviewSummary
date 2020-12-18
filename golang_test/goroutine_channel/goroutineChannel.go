package main
import (
	"fmt"
	//"runtime"
	"sync"
	"time"
)



//需求： goroutine 与 channel同时工作
var wg sync.WaitGroup

//写数据
func writeChan(writeCh chan int){
	defer wg.Done()

	for i := 1; i <= 10; i++{
		writeCh <- i
		fmt.Printf("【写入】数据%v到通道\n", i)
		time.Sleep(time.Millisecond*500)
	}
	close(writeCh)

}
//读数据
func readChan(readCh chan int){
	for v :=  range readCh {
		fmt.Printf("从通道【读取】数据%v\n", v)
		time.Sleep(time.Millisecond*50)
	}
	wg.Done()
}

func main(){

	var ch = make(chan int, 10)

	wg.Add(1)
	go writeChan(ch)
	wg.Add(1)
	go readChan(ch)

	wg.Wait()
	fmt.Println("关闭主线程\n")

}


