package main
import (
	"fmt"
	//"runtime"
	"sync"
	"time"
)

var count = 0
var wg sync.WaitGroup
var mutex sync.Mutex

func test(){
	mutex.Lock()
	count++
	fmt.Println("the count is : ",count)
	time.Sleep(time.Millisecond)
	mutex.Unlock()
	wg.Done()
}

func main(){
	for i := 0; i < 40; i++{
		wg.Add(1)
		go test()
	}
	wg.Wait()
}