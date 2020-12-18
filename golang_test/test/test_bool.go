package main

import "fmt"

func main(){
	/*
	go语言布尔型只有true和false两个值
	注意：
	1.布尔类型变量默认值为false
	2.go语言不允许将整数类型转化为布尔型
	3.布尔型无法参与数值运算，也无法与其他按类型进行转化
	*/
	var flag1 = true
	fmt.Printf("%v--%T\n",flag1,flag1)
    //1.布尔类型变量默认值为false
	var flag2 bool
	fmt.Printf("%v--%T\n",flag2,flag2)
	//2. string 默认值为空
	var str string
	fmt.Printf("%v--%T\n",str,str)
	//3. int 默认值为0
	var i int
	fmt.Printf("%v--%T\n",i,i)
	//4. float64 默认值为0
	var f float64
	fmt.Printf("%v--%T\n",f,f)
	//5.go语言不允许将整数类型转化为布尔型
	var a = true  
	if a{   //a=1,错误写法
		fmt.Println("true")
	}

	//6.布尔型无法参与数值运算，也无法与其他按类型进行转化
	var b = false  
	if b{   //a = "liwen"  ,错误写法
		fmt.Println("true")
	}else{
		fmt.Println("false")
	}
}