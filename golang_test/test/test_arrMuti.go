package main

import( 
	"fmt"
	//"strconv"
)

//import "strings"
//import "unsafe"

func main(){
	/*1
	值类型：改变变量副本的值，不会改变变量本身的值
	引用数据类型: 改变变量副本的值，会改变变量本身的值
	*/
	//1.1、值类型：基本数据类型和数组 ; 
	var a1 = 10
	a2 := a1
	a1 = 20
	fmt.Printf("a1=%v,a2=%v \n",a1,a2)//a1=20,a2=10    
	var arr1 =[...]int{1,2,3}
	arr2 := arr1
	arr1[0] = 11
	fmt.Printf("arr1=%v,arr2=%v \n",arr1,arr2) //arr1=[11 2 3],arr2=[1 2 3]
	//1.2、引用数据类型: 切片 ;
	var arr3 = []int{11,12,13}
	arr4 := arr3
	arr3[0] = 111
	fmt.Printf("arr3=%v,arr4=%v \n",arr3,arr4)  //arr3=[111 12 13],arr4=[111 12 13]

	//2多维数组：二维数组
	//方法一：定义二维数组
	/*
	var arrMuti1 = [3][2]string{
		{"北京","上海"},
		{"郴州","长沙"},
		{"成都","重庆"},
	}
	*/
	//方法二：定义二维数组
	var arrMuti1 = [...][2]string{
		{"北京","上海"},
		{"郴州","长沙"},
		{"成都","重庆"},
	}
	
	/* 不支持以下定义
	var arrMuti1 = [3][...]string{
		{"北京","上海"},
		{"郴州","长沙"},
		{"成都","重庆"},
	}
	*/
	
	for _, v1 := range arrMuti1{
        for _,v2 := range v1{
			fmt.Println(v2)
		}
	}
	fmt.Println("*************")
	for i :=0; i < len(arrMuti1);i++{
		for j := 0; j < len(arrMuti1[i]); j++{
			fmt.Println(arrMuti1[i][j])
		}
	}

}