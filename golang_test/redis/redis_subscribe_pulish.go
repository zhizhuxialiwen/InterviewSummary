package main
import (
	//"github.com/garyburd/redigo/redis"
	"github.com/gomodule/redigo/redis"
    "fmt"
    "time"
)

func Subs() {  //订阅者
    conn, err := redis.Dial("tcp", "localhost:6379")
    if err != nil {
        fmt.Println("connect redis error :", err)
        return
    }
    defer conn.Close()
    psc := redis.PubSubConn{conn}
    psc.Subscribe("channel1") //订阅channel1频道
    for {
        switch v := psc.Receive().(type) {
        case redis.Message:
            fmt.Printf("%s: message: %s\n", v.Channel, v.Data)  //channel1: message: this is wd
        case redis.Subscription:
            fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count) //channel1: subscribe 1
        case error:
            fmt.Println(v)
            return
        }
    }
}
func Push(message string)  { //发布者
    conn, _ := redis.Dial("tcp", "localhost:6379")
    _,err1 := conn.Do("PUBLISH", "channel1", message)
	if err1 != nil {
		fmt.Println("pub err: ", err1)
		return
	}

}
func main()  {
    go Subs()
    go Push("this is wd")
    time.Sleep(time.Second*3)
}//channel1: subscribe 1//channel1: message: this is wd