package main

import (
	"fmt"
	"time"
	"sync"
)

//1、主线程与协程的执行顺产：若主线程执行完，则协程退出；若协程执行完，主线程进行执行.

var wg sync.WaitGroup //监听协程完成 ：   wg.Add(1) -》 wg.Done() -》wg.Wait() 
func test1(){
	for i := 0; i < 5; i++{
		fmt.Println("test: ", i)
		time.Sleep(time.Millisecond*10)
	}
	wg.Done()//协程计数器减1
}

func test2(){
	for i := 0; i < 5; i++{
		fmt.Println("test: ", i)
		time.Sleep(time.Millisecond*10)
	}
	wg.Done()//协程计数器减1
}

func main(){
    wg.Add(1)  //协程计数器加1
	go test1()
	wg.Add(1)  //协程计数器加1
	go test2()
	for i := 0; i < 5; i++{
		fmt.Println("main: ", i)
		time.Sleep(time.Millisecond*10)
	}
	//time.Sleep(time.Second)
	wg.Wait() //等待协程完毕
	fmt.Println("主线程退出: ")
}