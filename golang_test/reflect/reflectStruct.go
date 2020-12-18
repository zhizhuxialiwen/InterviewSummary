package main
import (
	"fmt"
	//"runtime"
	//"sync"
	//"time"
	"reflect"
)

//结构体
type Student struct{
	Name string `json:"name" form:"username"`
	Age int `json:"age"`
	Score int `json:"score"`
}
func (s Student) GetInfo() string{
	var str = fmt.Sprintf("姓名：%v, 年龄：%v, 分数： %v \n",s.Name, s.Age, s.Score)
	return str
}

func (s *Student) SetInfo(name string, age int, score int) {
	s.Name = name
	s.Age = age
	s.Score = score
}

func printStructFiled(stu1 interface{}){
	t := reflect.TypeOf(stu1)
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("传输的参数不是一个结构体")
		return
	}
	//1 通过类型变量的Filed获取结构体字段:索引
	field1 := t.Field(0) //Name
	fmt.Printf("%#v \n", field1) //reflect.StructField{Name:"Name", PkgPath:"", Type:(*reflect.rtype)(0x4aad00), Tag:"json:\"name\" from :\"username\"", Offset:0x0, Index:[]int{0}, Anonymous:false}
	fmt.Printf("字段名称：%v \n", field1.Name)  //字段名称：Name
	fmt.Printf("字段类型：%v \n", field1.Type) //字段类型：string
	fmt.Printf("字段Tag：%v \n", field1.Tag.Get("json")) // 字段Tag：name
	fmt.Printf("字段Tag：%v \n", field1.Tag.Get("form")) //字段Tag：username
	fmt.Println("*************************************")
	//2 通过变量FieldByName获得结构体字段:key
	field2, ok := t.FieldByName("Age")
	if ok{
		fmt.Printf("字段名称：%v \n", field2.Name)  //字段名称：Age
		fmt.Printf("字段类型：%v \n", field2.Type) //字段类型：int
		fmt.Printf("字段Tag：%v \n", field2.Tag.Get("json")) // age
	}
	//3 通过变量NumField获得结构体字段有几个字段属性
	var fieldCount =  t.NumField()
	fmt.Println("结构体有",fieldCount,"个属性") //结构体有 3 个属性
	//4 通过值变量获取结构体属性对应的值
	v := reflect.ValueOf(stu1)
	fmt.Println(v.FieldByName("Name")) //liwne
	fmt.Println(v.FieldByName("Age")) //12

	for i := 0 ; i< fieldCount; i++{
		fmt.Printf("属性名称：%v,属性值：%v,属性类型：%v, 属性Tag:%v \n",
	    t.Field(i).Name,t.Field(i),t.Field(i).Type,t.Field(i).Tag)
	}

}

//
func PrintStructMethod (stu1 interface{}){
	t := reflect.TypeOf(stu1)
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("传输的参数不是一个结构体")
		return
	}
	//1 通过method获取结构体方法
	method1 := t.Method(0)
	fmt.Println(method1.Name) //GetInfo
	fmt.Println(method1.Type) //func(main.Student) string
	//2 通过结构体获取结构体有多少方法
	method2,ok := t.MethodByName("print")
	if ok{
		fmt.Println(method2.Name)
		fmt.Println(method2.Type)
	}
	//3 通过值变量执行方法,获取值
	v := reflect.ValueOf(stu1)
	fn0:= v.Method(0).Call(nil) //GetInfo()
	fmt.Println(fn0) // [姓名：liwne, 年龄：12, 分数： 34	]
	fn1 := v.MethodByName("GetInfo").Call(nil) //GetInfo()
	fmt.Println(fn1)  //[姓名：liwne, 年龄：12, 分数： 34]

	//4 执行方法传入参数，修改结构体参数
	var params []reflect.Value
	params = append(params, reflect.ValueOf("liwen"))
	params = append(params, reflect.ValueOf(24))
	params = append(params, reflect.ValueOf(56))
	v.MethodByName("SetInfo").Call(params)
	fn2 := v.MethodByName("GetInfo").Call(nil) 
	fmt.Println(fn2) //[姓名：liwen, 年龄：24, 分数： 56]

	//5 获取方法数量
	fmt.Println("方法数量：",t.NumMethod()) //方法数量： 2
}

//反射修改结构体属性
func reflectChangeStruct(stu1 interface{}){
	v := reflect.ValueOf(stu1)
	t := reflect.TypeOf(stu1)
	if t.Kind() != reflect.Ptr{
		fmt.Println("传输的参数不是一个执指针类型")
		return
	}else if t.Elem().Kind() != reflect.Struct {
		fmt.Println("传输的参数不是一个结构体")
		return
	}
	//修改结构体属性的值
	name := v.Elem().FieldByName("Name")
	name.SetString("lisi")

	age := v.Elem().FieldByName("Age")
	age.SetInt(23)
    
	

}
func main(){
	stu1 := Student{
		Name :"liwne",
		Age: 12,
		Score: 34,
	}
	printStructFiled(stu1)

	PrintStructMethod(&stu1)

	reflectChangeStruct(&stu1)
	fmt.Println(stu1) //{lisi 23 56}

}