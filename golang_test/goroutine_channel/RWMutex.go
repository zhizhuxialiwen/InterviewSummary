package main
import (
	"fmt"
	//"runtime"
	"sync"
	"time"
)


var wg sync.WaitGroup
var rwmutex sync.RWMutex

//写锁
func write(){
	rwmutex.Lock()
	fmt.Println("执行写操作")
	time.Sleep(time.Second*2)
	rwmutex.Unlock()
	wg.Done()
}
//读锁
func read(){
	rwmutex.RLock()
	fmt.Println("执行读操作")
	time.Sleep(time.Second*2)
	rwmutex.RUnlock()
	wg.Done()
}

func main(){

	for i := 2; i < 10; i++{
		wg.Add(1)
        go write()
        go read()	
	}
	wg.Wait()
	fmt.Println("主线程完成")

}