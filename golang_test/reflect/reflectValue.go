package main
import (
	"fmt"
	//"runtime"
	//"sync"
	//"time"
	"reflect"
)

//1、通过反射获取空接口类型的值
func reflectType(x interface{}){
	v := reflect.ValueOf(x)
	fmt.Println(v)
	//var m = v + 12  //错误
	//var m = v.Int() + 12 //获取原始值
	kind := v.Kind()
	switch kind{
	case reflect.Int:
		fmt.Printf("int类型的原始值%v \n",v.Int() + 12 )
	case  reflect.Float32:
		fmt.Printf("Float32类型的原始值%v \n",v.Float() + 12.33 )
	case  reflect.Float64:
		fmt.Printf("Float64类型的原始值%v \n",v.Float() + 12.33 ) //Float64类型的原始值15.785
	case  reflect.String:
		fmt.Printf("string类型的原始值%v \n",v.String() ) //string类型的原始值liwen
	default:
		fmt.Printf("没有获得类型的原始值%v \n" )

	}

}
func main(){
    a1 := 11
	reflectType(a1)

	f1 := 3.455
    reflectType(f1)
	str1 := "liwen"
	reflectType(str1)


}