package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	//"strconv"
)

func main(){
	//1 连接到redis: 建立tcp网络连接
	conn, err1 := redis.Dial("tcp","127.0.0.1:6379")
	if err1 != nil{
		fmt.Println("redis.Dial err = ", err1)
		return
	}
	defer conn.Close() //4关闭数据库
	//2.哈希hash: hset通过go向redis写入数据hash
	_,err2 := conn.Do("hset","user","name","liwen")
	if err2 != nil{
		fmt.Println("hset err = ", err2)
		return
	}
	_, err22 := conn.Do("hmset","user","age",20,"id",123456,"phone",18513278676)
	if err22 != nil{
		fmt.Println("hmset err = ", err22)
		return
	}
	//3. 哈希hash: hget 查找数据
	//redis.String()
	str,err3 := redis.String(conn.Do("hget","user","name"))
	if err3 != nil{
		fmt.Println("hget err = ", err3)
		return
	}
	fmt.Println(str)
    //redis.Strings
	str3,err33 := redis.Strings(conn.Do("hmget","user","name","id","phone"))
	if err33 != nil{
		fmt.Println("hmget err = ", err33)
		return
	}

	fmt.Println(str3)
	for k,v := range str3{
		fmt.Printf("hmget key: %v, value:%v \n",k,v)
	}
}