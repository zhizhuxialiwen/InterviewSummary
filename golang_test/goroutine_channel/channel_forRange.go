package main
import (
	"fmt"
	//"runtime"
	//"sync"
	//"time"
)

//1 管道循环遍历,
//1.1 若在使用for...range取出管道值,则之前必须关闭管道close(ch)
//1.2 若在使用for取出管道值,则之前不需要关闭管道close(ch)

func main(){
	ch1 := make(chan int, 10)
	for i :=1; i <= 10; i++{
		ch1 <- i 
	}

	// close(ch1) //必须关闭管道

	// //管道没有key，只有value
	// for v := range ch1{
	// 	fmt.Println(v)
	// }

	for i :=1; i <= 10; i++{
		fmt.Println(<- ch1)
	}
}