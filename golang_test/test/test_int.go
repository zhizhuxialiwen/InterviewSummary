package main

import "fmt"
import "unsafe"

func main(){
	//1 定义int类型
	var num = 10
	num = 11
	fmt.Printf("num=%v,num类型是%T \n",num,num)
	//2 int8的范围 (-128, 127)
	//var n1 int8 = 130 //错误
	var n1 int8 = 127 
	fmt.Printf("n1=%v, n1 类型是 %T\n",n1,n1)
	//2.1 unsafe.Sizeof :内存空间
	fmt.Println(unsafe.Sizeof(n1))

	//3 uint8 的范围(0,255)
	//var n2 uint8 = 256 //错误
	var n2 uint8 = 255
	fmt.Printf("n2=%v, n2 类型是 %T\n",n2,n2)
	fmt.Println(unsafe.Sizeof(n2))

	//4 int不同长度直接转化
	var n3 int32 = 10
	var n4 int64 = 20
	fmt.Println(int64(n3) + n4)
	fmt.Println(n3 + int32(n4))

	//5 高位向低位转化
	var n5 int16 = 130
	fmt.Println( int8(n5)) //-126 ；错误int8 (-127,128), 130不在范围之内

	var n6 int16 = 110
	fmt.Println( int8(n6)) 
	//6 数字字面量语法：%V表示原样， %d 表示十进制，%b表示二进制，%o表示八进制，%x表示十六进制
	n7 := 12
	fmt.Printf("n7=%v, n7 类型是 %T\n",n7,n7)  
	fmt.Println(unsafe.Sizeof(n7)) //8字节为64位，因此类型为int64
	fmt.Printf("n7=%d\n",n7)  
	fmt.Printf("n7=%b\n",n7)  
	fmt.Printf("n7=%o\n",n7)  
	fmt.Printf("n7=%x\n",n7)  



}