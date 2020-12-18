package main

import "fmt"

func test() (ret int){
	ret = 10
	return 1
}

func test01() (ret int){
	defer func(){
		ret = 10
	}()
	return 1
}

func test02() (ret int){
	defer func(){
		ret += 10
	}()
	return 1
}

func test03() (ret int){
	ret = 10
	defer func(){
		
	}()
	return 1
}

func test04() (ret int) {
    defer func(ret int) {
        ret = ret + 5
    }(ret)
    return 1
} 

func test05() (ret int) {
    t := 5
    defer func() {
        t = t + 5
    }()
    return t
}

func main(){
	
	ret := test()
	fmt.Println("ret=", ret) 

	ret1 := test01()
	fmt.Println("ret1=", ret1) 

	ret2 := test02()
	fmt.Println("ret2=", ret2) 

	ret3 := test03()
	fmt.Println("ret3=", ret3) 

	ret4 := test04()
	fmt.Println("ret4=", ret4) 

	ret5 := test05()
	fmt.Println("ret5=", ret5) 
}