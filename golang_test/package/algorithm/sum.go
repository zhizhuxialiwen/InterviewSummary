package algorithm

import (
	"fmt"
)

func init(){
    fmt.Println("sum.init()...")
}

func SumInt(x, y int) int{
	return x + y
} 

func SumFloat(f1, f2 float64) float64{
	return f1 + f2
} 