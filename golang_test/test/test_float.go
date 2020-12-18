package main

import "fmt"
import "unsafe"

func main(){
	//1 、定义float32类型,占用4个字节，保留6位小数
	var a1 float32 = 3.12
	fmt.Printf("值： %v--%f,类型：%T \n",a1,a1,a1)
	fmt.Println(unsafe.Sizeof(a1))
	//2、  定义float64类型,占用8个字节，保留6位小数
	var a2 float64 = 3.12
	fmt.Printf("值： %v--%f,类型：%T \n",a2,a2,a2)
	fmt.Println(unsafe.Sizeof(a2))
	//3 、%f 输出float类型，%.mf保留m为小数，四收入五原则
	var a3 float64 = 3.12456787
	fmt.Printf("值： %v--%.4f,类型：%T \n",a3,a3,a3)
	fmt.Println(unsafe.Sizeof(a3))
	//4 、64位系统默认go的负点默认类型为float64
	var a4 = 3.12456787
	fmt.Printf("值： %v--%f,类型：%T \n",a4,a4,a4)
	fmt.Println(unsafe.Sizeof(a4))
	//5、go科学计数法表示浮点类型
	var a5 float32 = 3.14e2 //表示3.14*10的2次幂方
	fmt.Printf("值： %v--%f,类型：%T \n",a5,a5,a5)
	fmt.Println(unsafe.Sizeof(a5))

	var a6 float32 = 3.14e-2 //表示3.14除以10的2次幂方
	fmt.Printf("值： %v--%f,类型：%T \n",a6,a6,a6)
	fmt.Println(unsafe.Sizeof(a6))

	//6、go 精度丢失问题？解决：引入第三方包decimal
	var b1 float64 = 1129.6
	fmt.Println(b1*100)  //期望是112960，但是是输出结果112959.99999999999

	b2  := 8.2
	b3  := 3.8
	fmt.Println(b2-b3) //期望输出4.4，但是输出结果4.3999999999999995

	//7、 int 类型转化float64类型
	b4 :=10
	b5 := float64(b4)
	fmt.Printf("b4类型：%T，b5类型：%T \n",b4,b5)
	//8、 float32 类型转化float64类型
	var b6 float32 = 11
	b7 := float64(b6)
	fmt.Printf("b6类型：%T，b7类型：%T \n",b6,b7)
	//8、 float32 类型转化int类型,不建议此做法，会丢失小数
	var b8 float32 = 12.14
	b9 := int(b8)
	fmt.Printf("b8类型：%T，b9类型：%T \n",b8,b9)
	fmt.Printf("b8：%v，b9：%v \n",b8,b9)

}