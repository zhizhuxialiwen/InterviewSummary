package main
import (
	"fmt"
	//"runtime"
	//"sync"
	//"time"
)

func main(){
	//1 创建管道
	ch := make(chan int, 3) //3表示容量
	//2 给管道传送值
	ch <- 10
	ch <- 20
	ch <- 30
	//3 接收管道的值
	a1 := <- ch
	fmt.Println("a1=",a1)//a1= 10 ，先入先出
	<- ch   //20
	a3 := <- ch
	fmt.Println("a3=",a3)  //a3= 30
	ch <- 55
	//4 管道长度与容量
	fmt.Printf("值：%v, 长度：%v,容量：%v\n", ch, len(ch),cap(ch)) //值：0xc000096080, 长度：1,容量：3
	//5 管道是引用类型:改变副本会改变主本
	ch1 := make(chan int, 4)
	ch1 <- 34
	ch1 <- 35
	ch1 <- 36
	ch2 := ch1
	ch2 <- 253
	<- ch1
	<- ch1
	<- ch1
	a4 := <- ch1
	fmt.Println(a4) //25
	//6 管道阻塞:
	//1）存放管道的数据大于管道容量会造成阻塞
	//2) 若管道的内容已经去完，再次取值i会造成阻塞
	ch3 := make(chan int, 2)
	ch3 <- 44
	ch3 <- 45
	//ch3 <- 46  //管道阻塞：fatal error: all goroutines are asleep - deadlock!
	a5 := <- ch3
	a6 := <- ch3
	//a7 := <- ch3 //超出管道取值范围，fatal error: all goroutines are asleep - deadlock!
	//fmt.Println(a5,a6,a7)
	fmt.Println(a5,a6)
	
}