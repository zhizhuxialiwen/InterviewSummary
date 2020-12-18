package main

import( 
	"fmt"
	"strconv"
)

//import "strings"
//import "unsafe"

func main(){
	//1、string 类型转换整型
	
	
	/*1.1 ParseInt
	参数1：string数据；参数2：进制；参数3：位数32或64
	*/
	str1 := "123456"
	fmt.Printf("%v--%T \n",str1,str1)  //123456--string
	num1,_ := strconv.ParseInt(str1,10,64)
	fmt.Printf("%v--%T\n",num1,num1) //123456--int64
	/*1.2 ParseFloat
	参数1：string数据；参数2：位数32或64
	*/
	str2 := "123456.3452341221"
	f1,_ := strconv.ParseFloat(str2,64)
	fmt.Printf("%v--%T\n",f1,f1) //123456.3452341221--float64
	//1.3 不建议使用ParseBool
}