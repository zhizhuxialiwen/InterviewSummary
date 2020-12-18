//package tcpIpSocket
package main

import (
	"fmt"
	"net"
	//"io"
)

func process(conn net.Conn){
	//循环接收客户端数据
	defer conn.Close() //关闭
	for {
		//创建一个新切片
		sliceBuf := make([]byte, 1024)
		//1/等待客户端通过conn发送数据
		//若客户端没有write[发送]，那么协程就阻塞在这里
		//fmt.Println("服务器等待客户端发送信息 + \n", conn.RemoteAddr().String())
		n, err := conn.Read(sliceBuf)
		// if err == io.EOF{
		// 	fmt.Println("客户端已退出")
		// 	return
		// }
		if err != nil{
			fmt.Println("服务器的read err=",err)
			return
		}
		//2 显示客户端发送的无内容到服务端的终端
		fmt.Println(string(sliceBuf[:n])) //注意：sliceBuf[:n] 切片所有数据
	}
}

func main(){
	fmt.Println("服务器开始监听")
	//1 监听
	listen, err := net.Listen("tcp","0.0.0.0:8888")
	if err != nil{
		fmt.Println("listen err= ",err)
		return
	}
	defer listen.Close() //延时关闭监听

	//循环等待客户端连接服务器
	for{
		fmt.Println("等待客户端来来链接...")
		conn,err := listen.Accept()
		if err != nil{
			fmt.Println(" listen.Accpet() err= ",err)
			return
		}else{
			fmt.Println(" listen.Accpet() successful conn=%v, client ip=%v ",conn, conn.RemoteAddr())
		}

		go process(conn)
	}
}

