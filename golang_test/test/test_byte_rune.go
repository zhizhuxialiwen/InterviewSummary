package main

import "fmt"

//import "strings"
//import "unsafe"

func main(){
	//1.go定义字符 ,字符属于int类型
	var a1 = 'a'
	fmt.Printf("值：%v,类型：%T \n",a1,a1) //值：97,类型：int32

	//2、原样输出字符
	var a2 = 'a'
	fmt.Printf("原样输出：%c,类型：%T\n",a2,a2) //值：a,类型：int32
	//3、定义字符串输出字符
	var str1 = "this"
	fmt.Printf("值：%v,原样输出：%c,类型：%T\n",str1[2],str1[2],str1[2]) //值：105,原样输出：i,类型：uint8
	//4、utf-8:一个汉字占3个字节，一个字母占一个字符
	//unsafe.Sizeof()无法查看string类型数据占用的内存空间
	//fmt.Println(unsafe.Sizeof(str1))  //错误
	fmt.Println(len(str1)) //4
	//5、定义一个字符，字符的值是汉字
	//go汉字使用Utf-8编码，编码后的值就是int类型
	var a3 = '国'
	fmt.Printf("值：%v,类型：%T \n",a3,a3) //值：22269,类型：int32
	fmt.Printf("原样输出：%c,类型：%T\n",a3,a3) 
	//6、循环遍历，go字符类型：uint8和rune, uint8类型称为byte，代表ASCII码的一个字符；rune类型表示一个UTF-8字符
	str2 := "你好 golang"
	for i := 0; i < len(str2); i++{  //byte 或uint8
		fmt.Printf("%v(%c)  ",str2[i],str2[i])
	}
	
	for _, v := range str2{ //rune
		fmt.Printf("%v(%c)  ",v,v)
	}
	fmt.Println("\n")
	//7、修改字符串
	str3 := "big"
	//str3[0] = 'p' //错误
	byteStr1 := []byte(str3)  //byte:英文
	byteStr1[0] = 'p'
	fmt.Println(string(byteStr1))

	str4 := "李文big"
	//str3[0] = 'p' //错误
	runeStr2 := []rune(str4)  //rune:中文、英文
	runeStr2[0] = '红'
	fmt.Println(string(runeStr2)) //红文big




}