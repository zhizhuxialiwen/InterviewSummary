package main
import (
	"fmt"
	//"runtime"
	//"sync"
	//"time"
	"reflect"
)

//1、通过反射获取空接口类型
func reflectType(x interface{}){
	v := reflect.TypeOf(x)  //通过反射获取空接口类型
	fmt.Println(v)
	v.Name() //类型名称,底层的类型
	v.Kind() //种类
	fmt.Printf("类型：%v， 类型名称：%v，种类：%v",	v,	v.Name(), v.Kind())
}

//2、自定义类型
//type: 自定义类型、结构体、接口定义
type myInt int
type Person struct{
	Name string
	Age int
}
func main(){
	a := 1
	f := 3.44
	b := true
	str := "liwen"
	reflectType(a)  //类型：int， 类型名称：int，种类：intfloat64
	reflectType(f) //foat64
	reflectType(b)//bool
	reflectType(str)

	var a1 myInt = 11
	person1 := Person{
		Name:"lwien",
		Age:12,
	}
	reflectType(a1) //main.myInt
	reflectType(person1) //类型：main.Person， 类型名称：Person，种类：struct
	var arr1 = [3]int{1,2,3}
    reflectType(arr1) //类型：[3]int， 类型名称：，种类：array[]int
	var slice1 = []int{1,3,4}
	reflectType(slice1) //类型：[]int， 类型名称：，种类：slice
}
