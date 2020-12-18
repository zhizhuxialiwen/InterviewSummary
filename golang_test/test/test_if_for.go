package main

import( 
	"fmt"
	//"strconv"
)

//import "strings"
//import "unsafe"

func main(){
	//if 建议没有小括号（）
	//1.1
	n1 := 4
	if n1 < 8{
		fmt.Println(n1)	
	}

    if n2 :=5; n2 < 8{
		fmt.Println(n2)
	}
	//1.2
	n3 := 6
	if n3 < 8{
		fmt.Println("n3小于8",n3)	
	}else{
        fmt.Println("n3大于8",n3)	
	}
	//1.3
	n4 := 7
	if n4 < 8{
		fmt.Printf("n4=%v小于8 \n",n4)	
	}else if n4 > 8 {
        fmt.Printf("n4=%v大于8 \n",n4)	
	}else {
		fmt.Printf("n4=%v等于8 \n",n4)	
	}
	/*
	注意：1、if的大括号不能省略;2.if/else的需要紧挨着大括号
	*/

	/*2、for 初始语句；条件语句；结束语句{
		循环语句
	}
	注意:for循环的死循环,golang没有while循环
	*/
	//2.1
	for i := 0 ;i < 10 ; i++{
		fmt.Printf("i=%v ",i)
	}
	fmt.Printf("\n")
	//2.2
	j := 0 
	for ;j < 10 ; j++{
		fmt.Printf("j=%v ",j)
	}
	fmt.Printf("\n")
	//2.3
	j1 := 0
	for j1 < 10 {
		fmt.Printf("j1=%v ",j1)
		j1 ++
	}
	fmt.Printf("\n")
	//2.4
	j2 := 0
	for{
		if j2 < 10{
			fmt.Printf("j2=%v ",j2)
		}else{
			break
		}
		j2++
	}
	fmt.Printf("\n")

	//3 for循环的嵌套
	for i := 1; i <= 9; i++{
		for j :=1; j <= i; j++{
			fmt.Printf("%v*%v=%v \t", i,j,i*j)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")

}