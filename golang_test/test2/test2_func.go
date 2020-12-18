package main

import (
	"fmt"
	//"sort"
	//"strings"
)

//1.函数
func sumFn(x int, y int) int{
	sum := x + y
	return sum
}

//2.简写类型的函数
func subFn(x, y int) int{
	sub := x -y
	return sub
}
//3、可变参数的函数:切片
func sumFn1(x ...int) int {
	fmt.Printf("%v--%T \n",x,x) //[4 5 6 7]--[]int
	sum := 0
	for _, v := range x{
		sum += v
	}
	return sum
}

func sumFn2(x int, y...int) int {
	fmt.Printf("%v--%T \n",x,x) //[4 5 6 7]--[]int
	sum := x
	for _, v := range y{
		sum += v
	}
	return sum
}

//4、函数返回值：return 返回多个值
func sum_sub(x,y int) (int,int){
	sum := x + y
	sub := x - y
	return sum,sub
}
//4.1 前面参数可以不写类型，但是后面参数必须写类型，反之不行
func sum_sub1(x,y int) (sum int, sub int){
	sum = x + y
	sub = x - y
	return sum,sub
}
//4.2 前面返回值参数可以不写类型，但是后面返回值参数必须写类型，反之不行
func sum_sub2(x,y int) (sum , sub int){
	sum = x + y
	sub = x - y
	return sum,sub
}

func main(){
	x := 3
	y := 8
	sum := sumFn(x, y)
	fmt.Println(sum)
	
	x1 := 5
	y1 := 8
	sub1 := subFn(x1, y1)
	fmt.Println(sub1)

	sum1 := sumFn1(4,5,6,7)
	fmt.Println(sum1) //22
    sum2 := sumFn1(4,5,6,7,8)
	fmt.Println(sum2) //30

	sum3,sub3 := sum_sub(x,y)
	fmt.Println(sum3,sub3) //11 -5
	sum4,sub4 := sum_sub1(x,y)
	fmt.Println(sum4,sub4) //11 -5
	sum5,sub5 := sum_sub2(x,y)
	fmt.Println(sum5,sub5) //11 -5
	
	_, sub6 := sum_sub2(x,y)
	fmt.Println(sub6) 


}