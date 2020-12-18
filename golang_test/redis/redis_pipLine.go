
package main
import (
	//"github.com/garyburd/redigo/redis"
	"fmt"
	"github.com/gomodule/redigo/redis"
)


func main() {
    conn,err := redis.Dial("tcp","localhost:6379")
    if err != nil {
        fmt.Println("connect redis error :",err)
        return
    }
	defer conn.Close()
	//1 send发送到缓存
	
    conn.Send("HSET", "student","Score","100")
	conn.Send("HGET", "student","Score")
	
	conn.Send("HMSET", "student","name", "wd","age","22")
	conn.Send("HMGET", "student","name","age")
	//2一次性发送服务器
    conn.Flush()
    //3一次接收相应结果
    if res1, err := conn.Receive(); err != nil{
		fmt.Println("res1 err1 = ",err)
	}else{
		fmt.Printf("Receive res1:%v \n", res1) //Receive res1:0
	}

	if res2, err := conn.Receive(); err != nil{
		fmt.Println("res2 err1 = ",err)
	}else{
		fmt.Printf("Receive res2:%s \n", res2) //Receive res2:100
	}

	if res3, err := conn.Receive(); err != nil{
		fmt.Println("res3 err1 = ",err)
	}else{
		fmt.Printf("Receive res3: %v \n", res3) //Receive res3: OK
	}

	if res4, err := conn.Receive(); err != nil{
		fmt.Println("res4 err1 = ",err)
	}else{
		fmt.Printf("Receive res4: %s \n", res4)  //Receive res4: [wd 22]
	}

	
   
}//Receive res1:0 //Receive res2:0//Receive res3:22
