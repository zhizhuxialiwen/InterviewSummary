package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main(){
	//1 连接到redis: 建立tcp网络连接
	conn, err1 := redis.Dial("tcp","127.0.0.1:6379")
	if err1 != nil{
		fmt.Println("redis.Dial err = ", err1)
		return
	}
	defer conn.Close() //4关闭数据库
	//2.set通过go向redis写入数据string\
	//2.1 一个设置set
	_,err2 := conn.Do("set","name","liwen")
	if err2 != nil{
		fmt.Println("set err = ", err2)
		return
	}
    //2.2批量设置mset
	_,err22 := conn.Do("mset","name","zhangsan","id",324156, "address","beijing")
	if err22 != nil{
		fmt.Println("mset err = ", err22)
		return
	}
	//3. get 查找数据
	//3.1 一个获取get:String
	str,err3 := redis.String(conn.Do("get","name"))
	if err3 != nil{
		fmt.Println("get err = ", err3)
		return
	}
	fmt.Println(str)
	//3.2批量获取mget:Strings
	str3,err33 := redis.Strings(conn.Do("mget","name","id","address"))
	if err33 != nil{
		fmt.Println("mget err = ", err3)
		return
	}
	fmt.Println(str3)
}