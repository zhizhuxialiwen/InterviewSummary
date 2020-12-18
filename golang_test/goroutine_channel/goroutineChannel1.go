package main
import (
	"fmt"
	//"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

//需求： 需统计1-120000的数字的素数
func putNum(intChan chan int){
	for i := 2; i < 120000; i++{
		intChan <- i
	}
	close(intChan)
	wg.Done()
}

//从通道读取素数,写入通道
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool){
	for i := range intChan{
		var flag = true
		for j := 2; j < i; j++{
			if i%j == 0{
				flag = false
				break
			}
		}
		if flag{
			primeChan <- i
			//fmt.Println("素数：",i)
		} 
	}
	//close(primeChan) //若一个channel 关闭，则无法给这个channel发送数据
	exitChan <- true  //exitChan里存放一条数据
	wg.Done()


}
//打印素数
func printPrime(primeChan chan int){
	for v := range primeChan{
		fmt.Println(v)
	}
	wg.Done()
}


func main(){
	start := time.Now().Unix()

	intChan := make(chan int, 1000)
	primeChan := make(chan int, 1000)
	exitChan := make(chan bool, 16) //表示primeChan退出
	//存放数字协程
    wg.Add(1)
	go putNum(intChan)
    //统计素数：16协程
	for i := 0; i < 16; i++{
		wg.Add(1)
		go primeNum(intChan, primeChan, exitChan)
	}

	//判断exitChan是否存满
	wg.Add(1)
    go func(){
        for i := 0; i < 16; i++{
		    <- exitChan
		}
		close(primeChan)
		wg.Done()
	}()
	//打印素数
    wg.Add(1)
	go printPrime(primeChan)

	wg.Wait()
	end :=  time.Now().Unix()
	fmt.Println("执行时间：",end -start)

	fmt.Println("主线程完成")


}