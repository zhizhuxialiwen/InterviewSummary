package main

import "fmt"

import "strings"
//import "unsafe"

func main(){
	//1.字符串定义
	var str1 string = "liwen1"
	var str2 = "liwen2"
	str3 := "liwen3"
	fmt.Printf("%v--%T\n",str1,str1)
	fmt.Printf("%v--%T\n",str2,str2)
	fmt.Printf("%v--%T\n",str3,str3)
	//2.字符串转义符
	var str4 = "liwen is \n good."   //换行\n
	fmt.Println(str4)
	var str5 = "C:\\go\\bin" //反斜杠\\
	fmt.Println(str5)
	var str6 = "C\"bin" //双引号
	fmt.Println(str6)
	//3 多行字符串：反引号字符
	var str7 =`liwen1
	liwen2
	liwen3`
	fmt.Println(str7)
	//4 字符串长度len(str)
	var str8 = "liwen"
	var str9 = "李文"
	fmt.Println(len(str8))   //长度为5
	fmt.Println(len(str9)) //长度为6,一个汉字占3个字节
	//5 拼接字符串： + 或fmt.Sprintf
	str10 := str8 + str9
	fmt.Println(str10)
	str11 := fmt.Sprintf("%v %v",str8, str9)
	fmt.Println(str11)
	str12 := "liwen1" +
	    "liwen2" +
	   	"liwen3"
	fmt.Println(str12)
	//6、分割字符串：strings.Split,需要导入strings包
	var str13 = "123-456-789"
	arr1 := strings.Split(str13, "-")
	fmt.Println(arr1)  //切片：[123 456 789]
	//7、切片连接成字符串：strig.Join(a[] string, sep string)
	str14 := strings.Join(arr1, "*")
	fmt.Println(str14) //123*456*789

	arr2 := []string{"php","java","golang"}
	str15 := strings.Join(arr2, "-")
	fmt.Println(str15) //php-java-golang
	fmt.Printf("%v %T \n",str15,str15)
	//8、判断str1是否包含str2: strings.contains
	str16 := "this is str"
	str17 := "this"
	flag1 := strings.Contains(str16,str17)
	fmt.Println(flag1)  //true
	//9、判断str1是否包含str2的前缀与后缀：strings.HasPrefix,strings.HasSuffix
	flag2 := strings.HasPrefix(str16,str17)
	flag3 := strings.HasSuffix(str16,str17)
	fmt.Println(flag2) //true
	fmt.Println(flag3)  //false

	//10、字符串出现的位置：strings.Index(), strings.LastIndex()，查找不到返回-1，查找到下标，从0开始
	str18 := "this is str"
	str19 := "is"
	num1 := strings.Index(str18,str19)
	fmt.Println(num1) //2
	num2 := strings.LastIndex(str18,str19)
	fmt.Println(num2) //5
	



}