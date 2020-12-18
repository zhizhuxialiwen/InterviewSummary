package main

import( 
	"fmt"
	//"strconv"
)

//import "strings"
//import "unsafe"

func main(){
	//1、数组声明
	var arr1 [4]int
	var arrStr1 [4]string
	fmt.Println(arr1) //[0 0 0 0]
	fmt.Println(arrStr1)  //[   ]
	//1.1、方法一:数组初始化 
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	arr1[3] = 4
	fmt.Println(arr1)
	//1.2、方法二：声明且初始化
	var arr2 = [4]int{11,12,13}
	fmt.Println(arr2)
	var arrStr2 = [4]string{"liwen","lishan","lisi","wangwu"}
	fmt.Println(arrStr2)
	//1.3 方法三：让编译器根据初始值个数进行推断
	var arr3 = [...]int{11,12,13,14,15}
	fmt.Println("arr3=",arr3, " len(arr3)=",len(arr3))
	arr3[0] = 222
	fmt.Println("arr3=",arr3, " len(arr3)=",len(arr3))
	//1.4 方法四： 下标与值对应
	arr4 := [...]int{0:1,1:10,2:20,3:30}
	fmt.Println("arr4=",arr4, " len(arr4)=",len(arr4))

	//2 for 循环
	sum := 0
	for i :=0; i < len(arr4); i++{
		fmt.Println(arr4[i])
		sum += arr4[i]
	}
	fmt.Println("sum=",sum)
	//3 for ...range 
	for k,v := range arr4{
		fmt.Printf("k=%v,v=%v \n",k,v) 
	}
}