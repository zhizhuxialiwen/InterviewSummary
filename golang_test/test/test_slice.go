package main

import( 
	"fmt"
	//"strconv"
)

//import "strings"
//import "unsafe"

func main(){
	//1、声明切片
	//1.1 方法一：
	var slice1 []int
	fmt.Printf("值：%v, 类型：%T，长度：%v \n", slice1, slice1, len(slice1)) //值：[], 类型：[]int，长度：0
	//1.2 方法二：
	var slice2 = []int{1,2,3}
	fmt.Printf("值：%v, 类型：%T，长度：%v \n", slice2, slice2, len(slice2))  //值：[1 2 3], 类型：[]int，长度：3
	//1.3 方法三：
	var slice3 = []int{1:11,2:22,4:33}
	fmt.Printf("值：%v, 类型：%T，长度：%v \n", slice3, slice3, len(slice3))  //值：[0 11 22 0 33], 类型：[]int，长度：5

	//2、go的默认值为nil
	fmt.Println(slice1 == nil)  //true
	fmt.Println(slice2 == nil)  //false

	//3、切片循环遍历
	var strSlice1 = []string{"liwen","liwei","lisan"}
	for i :=0; i< len(strSlice1); i++{
		fmt.Println(strSlice1[i])
	}
	for _,v := range strSlice1{
		fmt.Println(v)
	}
	//4、数组赋值于切片：基于数组的切片
	a1 := [5]int{1,2,3,4,5}
	a2 := a1[:]
	fmt.Printf("%v-%T \n", a2,a2) //[1 2 3 4 5]-[]int
	a3 := a1[1:4] //包含下标为1的元素，不包含下标为4的元素
	fmt.Printf("%v-%T \n", a3,a3 ) //[2 3 4]-[]int
	a4 := a1[2:]
	fmt.Printf("%v-%T \n", a4,a4) //[3 4 5]-[]int
	a5 := a1[:4]
	fmt.Printf("%v-%T \n", a5,a5) //[1 2 3 4]-[]int
	//5、基于切片的切片
	strSlice2 := []string{"liwen","zhangshan","lisi","wangwu"}
	strSlice3 := strSlice2[1:]
	fmt.Printf("%v-%T \n", strSlice3,strSlice3)  //[zhangshan lisi wangwu]-[]string
	//6、切片的长度和容量len():包含元素的个数；cap():从第开始下标的第一个元素,到底层数组最后一个元素末尾的个数
	slice4 := []int{1,2,3,4,5}
	fmt.Printf("长度：%d, 容量： %d \n", len(slice4), cap(slice4)) //长度：5, 容量： 5
    slice5 := slice4[2:]
	fmt.Printf("长度：%d, 容量： %d \n", len(slice5), cap(slice5)) //长度：3, 容量： 3
	slice6 := slice4[2:3] //下标为4的袁术不包含
	fmt.Printf("长度：%d, 容量： %d \n", len(slice6), cap(slice6)) //长度：1, 容量： 3
	slice7 := slice4[:3] 
	fmt.Printf("长度：%d, 容量： %d \n", len(slice7), cap(slice7)) //长度：3, 容量： 5

	//7、make函数创建切片： make([]T,size, cap)
	var slice8 = make([]int , 4, 5)
	fmt.Printf("长度：%d, 容量： %d \n", len(slice8), cap(slice8)) //长度：4, 容量： 5
	slice8[0] = 10
	slice8[1] = 11
	slice8[2] = 12
	slice8[3] = 13
	fmt.Println(slice8)  //[10 11 12 13]
	
	//8、append扩容：
	//8.1go切片无法通过下标方法进行扩容，只能通过append进行扩容
	var slice9 []int
	slice9 = append(slice9, 10)
	slice9 = append(slice9, 20)
	slice9 = append(slice9, 30, 40, 50)
	fmt.Printf("值：%v, 长度：%v,容量： %d \n", slice9, len(slice9),cap(slice9)) //值：[10 20 30 40 50], 长度：5,容量： 6

	//8.2、append合并切片
	strSlice4 :=[]string{"liwen1", "liwne2"}
	strSlice5 :=[]string{"liwen3", "liwne4"}
	strSlice4 = append(strSlice4, strSlice5...)
	fmt.Println(strSlice4) //[liwen1 liwne2 liwen3 liwne4]
	//8.3 append扩容测量
	var slice10 []int
	for i :=0; i <= 10; i++{
		slice10 = append(slice10, i)
		fmt.Printf("值：%v, 长度：%v,容量： %d \n", slice10, len(slice10),cap(slice10))
	}
	//9 切片引用类型：改变副本影响主本值
	var slice11 = []int{1,2,3}
	var slice12 = slice11
	slice12[0] = 11
	fmt.Println(slice11)  //[11 2 3]
	fmt.Println(slice12)  //[11 2 3]
	//10 copy(目的切片，原切片)函数复制切片
	slice13 := make([]int, 4, 5)
	copy(slice13, slice11)
	slice13[0] = 22
	fmt.Println(slice11)  //[11 2 3]
	fmt.Println(slice13)  //[22 2 3 0]
	//11 go语言没有专门的方法进行删除元素
	//删除索引为2的元素为3，注意：最后一个元素需要添加"...
	var slice14 = []int{1,2,3,4,5,6}
	slice14 = append(slice14[:2],slice14[3:]...)
	fmt.Println(slice14)

}