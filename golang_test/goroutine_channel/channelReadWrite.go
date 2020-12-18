package main
import (
	"fmt"
	//"runtime"
	//"sync"
	//"time"
)

func main(){
	//1 双向管道
	ch1 := make(chan int, 3)
	ch1 <- 3
	ch1 <- 4
	a1 := <- ch1
	a2 := <- ch1

	fmt.Println(a1, a2)
    //2 只写入管道
	ch2 := make(chan<- int, 3)
	ch2 <- 5
	ch2 <- 6
	// <- ch2  //无法从 管道读取数据
	//3 只从管道读取数据
	//ch3 := make(<-chan int, 3)
	//ch3 <- 6


	

}
