package main
import (
	"fmt"
	//"runtime"
	"sync"
	"time"
)


//需求： 需统计1-120000的数字的素数？goroutine for
// 1 协程： 1-30000
// 2 协程：30001-60000
// 3协程：60001-90000
// 4 协程：90001-120000
//start:(n-1)3000 + 1

var wg sync.WaitGroup
func test(num int){
	defer wg.Done()
	for i := (num - 1)*30000 + 1; i < num*30000 ; i++{
		//var flag = true
		for j := 2; j < i; j++{
			if i%j == 0{
				//flag = false
				break
			}
		} 
		/*
		if flag{
			fmt.Println("素数：",i)
		}
		*/
	}
}
func main(){

	start := time.Now().Unix()
	for i := 1; i <= 4; i++{
		wg.Add(1)
		test(i)
	}
	wg.Wait()
	end :=  time.Now().Unix()
	fmt.Println("执行时间：",end -start) //执行时间： 7 毫秒
	fmt.Println("关闭主线程")
}

