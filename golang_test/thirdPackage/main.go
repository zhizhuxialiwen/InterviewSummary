package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/tidwall/gjson"
)

const json = `{"name":{"first":"Janet","last":"liwen"},"age":47}`

func main() {
	// quantity := decimal.NewFromInt(3)
	// fmt.Println(quantity)
	var f1 float64 = 2.34
	var f2 float64 = 4.56
	fmt.Println(f1 + f2) //6.8999999999999995
    //甲方
	sum1 := decimal.NewFromFloat(f1).Add( decimal.NewFromFloat(f2))
	fmt.Println(sum1)  //4.9
	//相减
	sub1 := decimal.NewFromFloat(f1).Sub( decimal.NewFromFloat(f2))
	fmt.Println(sub1)  //-2.22
	//相乘
	mul1 := decimal.NewFromFloat(f1).Mul( decimal.NewFromFloat(f2))
	fmt.Println(mul1) //10.6704


	value := gjson.Get(json, "name.last")
	println(value.String()) //liwen


}