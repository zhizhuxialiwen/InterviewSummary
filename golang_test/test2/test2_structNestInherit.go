package main

import (
	"fmt"
	//"sort"
)
//1 结构体匿名字段:用于嵌套结构体。类型必须唯一，不能出现重复
//结构体：基本类型、切片、map、结构体
type Person struct{
	name string
	age int
	hobby []string
	map1 map[string]string
}
//2 结构体嵌套
type user struct{
	username string
	password string
	age  int
	//address1 address //user结构体嵌套address结构体
	address  //2.2 匿名嵌套结构体
	email email
}

type address struct{
	name string
	phone string
	city string
	age int
}
//2.2 匿名结构体嵌套：父结构体嵌套子结构体，不写结构体变量，调用时可以不写子结构体名
//2.3 若父结构体与子结构体有相同成员变量，优先调用父结构体变量

type email struct{
	name string
}

//3 结构体继承： 结构体自定义方法, 
/*
//3.1 父结构嵌套到子结构体进行继承
//父结构体
type animal struct{
	name string
}
func (a animal ) run(){
	fmt.Printf("%v 在运动\n", a.name)
}
//子结构体
type dog struct{
	age int
	animal //结构体继承或嵌套
}
func (d dog) wang() {
	fmt.Printf("%v ,%v 在叫\n",d.age, d.name)
}
*/
//3.2 指针结构体继承或指针结构体嵌套
type animal struct{
	name string
}
func (a animal ) run(){
	fmt.Printf("%v 在运动\n", a.name)
}
//子结构体
type dog struct{
	age int
	*animal //结构体继承或嵌套
}
func (d dog) wang() {
	fmt.Printf("%v ,%v 在叫\n",d.age, d.name)
}

func main(){
	var p Person
	p.name = "liwen"
	p.age = 20
	p.hobby = make([]string, 3,5)
	p.hobby[0] = "打篮球"
	p.hobby[1] = "羽毛球"
	p.hobby[2] = "兵兵球"

	p.map1 = make(map[string] string)
	p.map1["address"] = "北京"
	p.map1["phone"] = "18513278676"
	fmt.Printf("%#v \n", p)
	/*main.Person{name:"liwen", age:20, hobby:[]string{"打篮球", "羽毛球", "兵兵球"}, 
	map1:map[string]string{"address":"北京", "phone":"18513278676"}}*/
	var user1 user
	user1.username = "liwei"
	user1.password = "1826347932"
	user1.address.name = "lisi"
	user1.address.phone = "888888"
	user1.address.city = "北京"
	fmt.Printf("%#v \n", user1) 
	/*
	main.user{username:"liwei", password:"1826347932", 
	address1:main.address{name:"lisi", phone:"888888", city:"北京"}}    */
	user1.city = "上海"  //匿名嵌套结构体
	fmt.Printf("%#v \n", user1) 

	user1.age = 20
	fmt.Printf("%#v \n", user1) 

	user1.email.name = "zhangsan"
	fmt.Printf("%#v \n", user1) 
	/*
    //3.1 父结构嵌套到子结构体进行继承
	var d = dog{
		age:20,
		animal:animal{
			name:"马",
		},
	}
	d.run()
	d.wang()
	*/
	//3.2指针结构体继承或指针结构体嵌套
    var d = dog{
		age:20,
		animal:&animal{
			name:"马",
		},
	}
	d.run()
	d.wang()
   
}