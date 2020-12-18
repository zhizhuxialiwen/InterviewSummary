package main
import (
	//"github.com/garyburd/redigo/redis"
	"github.com/gomodule/redigo/redis"
	"fmt"
)


func main() {
    conn,err := redis.Dial("tcp","localhost:6379")
    if err != nil {
        fmt.Println("connect redis error :",err)
        return
    }
	defer conn.Close()
	//事务：多条语句一起执行，要么成功，要么失败
    conn.Send("MULTI") //事务启动
    conn.Send("INCR", "foo")
	conn.Send("INCR", "bar")
	// conn.Send("set", "student","liwen")
	// conn.send("set", "student","lisi")
    if r, err := conn.Do("EXEC"); err != nil{ //事务执行
		fmt.Println("err=" ,err)
		return
	}else{
		fmt.Println(r)//[1, 1]
	}
   
}