package main
import (
	"fmt"
	//"runtime"
	//"sync"
	//"time"
	"reflect"
)
//空接口的值不能直接修改，可以反射修改
func reflectSetValue(x interface{}){
	//空接口的值不能修改
	//*x = 120  //invalid indirect of x (type interface {})
	v := reflect.ValueOf(x)
	fmt.Println(v.Kind()) //ptr

	fmt.Println(v.Elem().Kind()) //int64
	if v.Elem().Kind() == reflect.Int64{
		v.Elem().SetInt(123)
	}else if  v.Elem().Kind() == reflect.String{
		v.Elem().SetString("wen")
	}
}

func main(){
	var a1 int64 = 100
	reflectSetValue(&a1)
	fmt.Println(a1) //123

	var str1 string = "liwen "
	reflectSetValue(&str1)
	fmt.Println(str1) //wen


}