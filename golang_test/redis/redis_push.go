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
	//2.push
	// //2.1 左输入数据lpush
	// if _, err := conn.Do("lpush", "listDemo", "world"); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	if _, err := conn.Do("lpush", "listDemo", "hello","china","USA"); err != nil {
		fmt.Println(err)
		return
	}
	// // 1.2 linsert 左插入： 
    // if value, err := conn.Do("linsert", "listDemo", "BEFORE", "world", "---"); err != nil {
	// 	fmt.Println(err)
	// 	return
	// } else {
	// 	fmt.Println("LINSERT=",value)
	// }
	//2.pop
	//2.1 lpop
	if value,err := conn.Do("lpop","listDemo"); err != nil{
		fmt.Println(err)
		return
	}else{
		fmt.Println("lpop value=",value)
	}
	//2.2 rpop
	if value,err := conn.Do("rpop","listDemo"); err != nil{
		fmt.Println(err)
		return
	}else{
		fmt.Println("rpop value=",value)
	}
   
	//2 显示列表数据
	//2.1 lrange
	if values, err := redis.Values(conn.Do("lrange", "listDemo", "0", "-1")); err != nil {
		fmt.Println(err)
		return
	} else {
		for _, v := range values {
			value := v.([]uint8)
			fmt.Println(string(value))
		}
	}
}