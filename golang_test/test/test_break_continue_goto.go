package main

import( 
	"fmt"
	//"strconv"
)

//import "strings"
//import "unsafe"

func main(){
	/*1、break
	break跳出当前循环for;
	break跳出在switch的case;
    break跳出多重循环，使用lable标签
	*/
	for i :=0; i < 10; i++{
		if i == 2 {
			break;
		}
		fmt.Println("i=",i)
	}
//break跳出多重循环，使用lable标签
lable1 :
	for i := 0; i< 2; i++{
		for j := 0 ; j < 5; j++{
			if j == 3{
				break lable1
			}
			fmt.Printf("i=%v,j=%v \n",i,j)
		}			
	}
	fmt.Println("lable1")
	//2、continue:结束当前循环，进行下一轮循环，仅用于for,也可以使用lable标签
	for i :=0; i < 10; i++{
		if i == 2 {
			continue;
		}
		fmt.Println("i=",i) //不打印2
	}

	for i := 0; i< 2; i++{
		for j := 0 ; j < 5; j++{
			if j == 3{
				continue
			}
			fmt.Printf("i=%v,j=%v \n",i,j)
		}			
	}
	fmt.Println("**************")
lable2:
	for i := 0; i< 2; i++{
		for j := 0 ; j < 5; j++{
			if j == 3{
				continue lable2
			}
			fmt.Printf("i=%v,j=%v \n",i,j)
		}			
	}
	fmt.Println("lable2")
	//3、goto通过标签进行跳转出循环、避免重复
	var n1 = 30
	if n1 > 24{
		fmt.Println("成年人")
		goto lable3
	}
	fmt.Println("aaa")
lable3:
	fmt.Println("bbb")
	fmt.Println("lable3")


}