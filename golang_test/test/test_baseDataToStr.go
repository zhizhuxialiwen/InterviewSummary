package main

import( 
	"fmt"
	"strconv"
)

//import "strings"
//import "unsafe"

func main(){

	//1、整型之间的转化
	var a1 int8 = 20
	var a2 int16 = 40
	fmt.Println(int16(a1)+a2) //60
	//2、浮点型之间的转换
	var f1 float32 = 20.345
	var f2 float64 = 40.3456
	fmt.Println(float64(f1) + f2) //60.69059931335449

	//3、整型与浮点型之间转化：建议浮点转化整型
	var a3 int = 20
	var f3 float32 = 30.342
	fmt.Println(float32(a3) + f3) //50.342
    //注意:转化建议低位转换高位，若高位转化低位，则出现溢出情况，输出结果不是我们预期的
	//4、其他类型转化string类型
	var i int = 20
	var f float64 = 12.3456
	var t bool = true
	var b byte = 'a'
	//4.1 fmt.Sprintf(): 将int,float64,bool,byte转化string类型
	str1 := fmt.Sprintf("%d", i)
	fmt.Printf("值： %v, 类型： %T \n",str1,str1)//值： 20, 类型： string
	str2 := fmt.Sprintf("%f", f)
	fmt.Printf("值： %v, 类型： %T\n", str2,str2) //值： 12.345600, 类型： string
	str3 := fmt.Sprintf("%t", t)
	fmt.Printf("值： %v, 类型： %T\n", str3,str3) //值： true, 类型： string
	str4 := fmt.Sprintf("%c", b)
	fmt.Printf("值： %v, 类型： %T\n",str4,str4) //值： a, 类型： string
	//4.2 strconv包： 将int,float64,bool,byte转化string类型
	//4.2.1 FormatInt: 参数1为int64，参数2为传值int类型的进制
	var i1 int = 20
	str5 := strconv.FormatInt(int64(i1), 10)
	fmt.Printf("值： %v, 类型： %T \n",str5,str5)//值： 20, 类型： string
	/*4.2.2 FormatFloat:
	参数1：转化类型float64
	参数2：格式化类型 'f' (-ddd.dddd)、'b' (-ddddp+-ddd,指数为二进制)、'e' (-d.dddde+-ddd,十进制指数)、
	'E' (-d.ddddE+-ddd,十进制指数)、'g' (指数很大时候用'e'格式，否则'f'格式)、'G' (指数很大时候用'e'格式，否则'f'格式)
	参数3：保留的小数点 -1(不对小数点格式化)
	参数4: 格式化的类型传入64或32
	*/
	var f4 float32 = 30.22424
	str6 := strconv.FormatFloat(float64(f4), 'f',4,64)
	fmt.Printf("值： %v, 类型： %T \n",str6,str6) //值： 30.2242, 类型： string
	//4.2.3 FormaTBool(ture)
	//4.2.3 FormaTUint(uint64(a),10)

}