package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)
var pool * redis.Pool
//初始化连接池
func init(){
	pool = &redis.Pool{
		MaxIdle:8,//最大空闲连接
		MaxActive:0,//表示和数据库的最大连接数，0表示没有限制
		IdleTimeout:100, //最大空闲时间
		Dial:func() (redis.Conn,error) { //初始化链接代码，链接哪个ip的redis
			return redis.Dial("tcp","localhost:6379")
			//return  redis.Dial("tcp","127.0.0.1:6379")
		},
	}

}
func main(){
	conn := pool.Get()//从连接池获取一个连接
	defer pool.Close()//关闭一个连接池，不能从连接池取出链接

	if _,err := conn.Do("set","name","liwen");err != nil{
		fmt.Println("set err=",err)
		return
	}

	if res,err := redis.String(conn.Do("get","name")); err != nil{
		fmt.Println("get err=",err)
		return
	}else{
		fmt.Println("res=",res)
	}
	
	
}