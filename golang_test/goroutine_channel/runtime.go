package main

import (
	"fmt"
	"runtime"
	//"sync"
)

func main(){
	//获取CPU的个数
	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum=",cpuNum) //cpuNum= 2
	//设置CPU的个数
	runtime.GOMAXPROCS(cpuNum -1 )
	fmt.Println("OK")

}