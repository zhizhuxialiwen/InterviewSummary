package algorithm  //放在最上面
import (
	"fmt"
)

func init(){
    fmt.Println("Sub.init()...")
}

//SubInt: 公有的； subInt:私有的
func SubInt(x, y int)int{
	return x - y
}

func SubFloat(f1, f2 float64)float64{
	return  f1 - f2
}