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
	//2无序集合 set通过go向redis写入数据string\
	//2.1 一个设置set
	_, err := conn.Do("sadd", "myset", "mobike", "foo", "ofo", "bluegogo")
	if err != nil {
		fmt.Println("set add failed", err.Error())
	}
	//查看set集合
	value, err := redis.Values(conn.Do("smembers", "myset"))
	if err != nil {
		fmt.Println("set get members failed", err.Error())
	} else {
		fmt.Printf("myset members :")
		for _, v := range value {
			fmt.Printf("%s \n", v.([]byte))
		}
		fmt.Printf("\n")
	}
	fmt.Println(value)

	

}