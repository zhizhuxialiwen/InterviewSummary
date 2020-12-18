package main

import (
	"fmt"
)

//接口
type Usb interface{
	start()
	stop()
}
//若接口里面的方法，必须使用结构体或自定义方法实现接口
type Phone struct{
	Name string
}

// //1 结构体值接收者
// func (p Phone)start(){
// 	fmt.Println("启动：",p.Name)
// }
// func (p Phone)stop(){
// 	fmt.Println("关机",p.Name)
// }
//2 结构体指针接收者
func (p *Phone)start(){
	fmt.Println("启动：",p.Name)
}
func (p *Phone)stop(){
	fmt.Println("关机",p.Name)
}

func main(){
	//1 结构体值接收者
	// p1 := Phone{
	// 	Name:"华为手机",
	// }
	// var u1 Usb = p1
	// u1.start()

	// p2 := &Phone{
	// 	Name:"苹果手机",
	// }
	// var u2 Usb = p2
	// u2.start()
  
	//2 结构体指针接收者
	p3 := &Phone{
		Name:"小米手机",
	}
	var u3 Usb = p3
	u3.start()

}