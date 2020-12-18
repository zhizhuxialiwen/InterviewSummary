package main

import (
	"fmt"
)

//1接口
type Usb interface{
	start()
	stop()
}
//1.1 若接口里面的方法，必须使用结构体或自定义方法实现接口
type Phone struct{
	Name string
}
//1.2手机要实现usb接口，必须实现usb接口中的所有方法，也可以调用自己的方法
func (p Phone)start(){
	fmt.Println("启动：",p.Name)
}
func (p Phone)stop(){
	fmt.Println("关机",p.Name)
}


type Computer struct{
}


// 2 自定义方法调用接口的方法
func (c Computer) work(u Usb){
	if _,ok := u.(Phone); ok{  //类型断言
		u.start()
	}else{
		u.stop()
	}
}
func main(){
	p := Phone{
		Name:"华为手机",
	}
	p.start()

	var u Usb
	u = p
	u.stop()

	var c = Computer{}
	var p1 = Phone{
		Name :"小米",
	}
	c.work(p1)  //通过phone结构体调用Usb接口


}