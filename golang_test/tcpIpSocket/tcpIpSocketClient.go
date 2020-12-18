//package tcpIpSocketClient
package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
	"strings"
)

func main(){
	//1
	//conn,err := net.Dial("tcp","192.168.211.160:8888")
	conn,err1 := net.Dial("tcp","192.168.1.102:8888")  //建立网络连接
	if err1 != nil{
		fmt.Println("client dail err1=",err1)
		return
	}

	//1 客户端单行发送数据
	reader := bufio.NewReader(os.Stdin) //os.Stdin标准输入【终端】
	for{
		//从终端读取一行数据，并发送给服务器
		line, err2 := reader.ReadString('\n')  //必须单引号,读取一行数据
		if err2 != nil{
			fmt.Println("ReadString err2=",err2)
			//return
		}
		//输入exit表示退出
		line = strings.Trim(line,"\r\n") //Trim去掉换行
		if line == "exit"{
			fmt.Println("客户端退出")
			break
		}
		// 再将line发送给服务器
		//n, err3 := conn.Write([]byte(line))
		n, err3 := conn.Write([]byte(line + "\n"))
		if err3 != nil{
			fmt.Println("conn.Write err2=",err2)
			//return
		}
		fmt.Printf("客户端发送了%d字节数，并且退出\n",n)
	}
	
}