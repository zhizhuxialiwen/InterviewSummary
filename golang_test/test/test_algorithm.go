package main

import( 
	"fmt"
	//"strconv"
)

//import "strings"
//import "unsafe"

func main(){
	//1算法： +- * / %   ++ --
	//1.1、除法，
	//1.1.1、若被除数是整数，除数是整数，商为整数，余数删除；
    var a1 = 10
	var b1 = 3
	fmt.Println(a1/b1)  //3
	//1.1.2若被除数是浮点数，除数是浮点数，商为浮点数，余数删除；
	var a2 = 10.0
	var b2 = 3.0
	fmt.Println(a2/b2) //3.3333333333333335
	//1.2、求余数， 余数= 被除数-（被除数/除数）/除数
	var a3 = 11
	var b3 = 3
	fmt.Println(a3/b3)   //3
	//1.3 ++ -- 在go语言只能单独语句，不是算数运算符，（1）只能单独使用 ；(2)没有前缀++/--a,只有后缀a ++ / --
	//var a4 int = 8
    var b4 int
	//b4 = a4 ++  //错误，只能单独使用,不能使用于复制运算符
	//b4 = a4 --  //错误，只能单独使用
	//++ b4  //错误，没有前缀++/--a,只有后缀a ++ / --
	b4 ++
	fmt.Println(b4)   //1
	//2、关系运算符： ==、！=、>、>=、<、<= ;若成功，则true，反之为false.
	var a5 = 9
	var b5 = 8
	if a5 > b5{
		fmt.Println("true") 
	}else{
		fmt.Println("false") 
	}
	//3、逻辑运算符： &&与、||或、 ！非
	//逻辑与：a1&&a2, 若a1为ture，则执行a2; 若a1为false，则不执行a2;
	//逻辑或：a1 || a2, 若a1为ture，则不执行a2; 若a1为false，则执行a2;
	flag1 := true
	fmt.Println(!flag1) 
	//4、赋值运算：= 、+=、-=、*=、/=、%=

	//练习：两个变量a6,b6进行交换，不使用中间变量，最终打印
	var a6 = 8
	var b6 = 11
	a6 = a6 + b6 // a6=8+11
	b6 = a6 - b6 // b6为a6值,b6=19-8=11
	a6 = a6 - b6 // a6为b6值,a6=19-11=8
	fmt.Printf("a6=%v,b6=%v \n",a6,b6)

	//5、位运算符：&、|、^、<<、>>
	//<<表示左移n位是乘以2的n次方；>>表示右移n位是除以2的n次方；
	var a7 = 5 //0101
	var b7 = 2 //0010
	fmt.Println("a7&b7=",a7&b7) //a7&b7= 0
	fmt.Println("a7|b7=",a7|b7) //a7|b7= 7
	fmt.Println("a7^b7=",a7^b7)  //a7^b7= 7
	fmt.Println("a7<<b7=",a7<<b7)  //a7<<b7= 20 = 5*(2^2)
	fmt.Println("a7>>b7=",a7>>b7)  //a7>>b7= 1 = 5/(2^2)


}
	