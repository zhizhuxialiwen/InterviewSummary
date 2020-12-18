package main

import (
	"fmt"
	//"sort"
	//"strings"
	"time"
)
func main(){
	timeObj := time.Now()
	fmt.Println(timeObj) //2020-07-20 15:44:09.9897106 +0800 CST m=+0.010970301
	//1.1、正常输出：%02d:表示2位，不够2位补0
	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	munite := timeObj.Minute()
	second := timeObj.Second()
	//%02d:表示2位，不够2位补0
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d",year,month,day,hour,munite,second) //2020-07-20 15:45:19
	//1.2、格式化输出:go诞生时间为2006年1月2号15点04分（记：2006 1 2 3 4 5）  03表示12小时制，15表示24小时制
	var str1 = timeObj.Format("2006/01/02 03:04:05")
	fmt.Println(str1)
	//2、获取当前时间转化时间戳：从1997开始计时,到现在的毫秒
	unixTime := timeObj.Unix()
	fmt.Println("当前毫秒数时间戳：",unixTime) //1595232070
	unixNanoTime := timeObj.UnixNano()
	fmt.Println("当前纳秒秒数时间戳：",unixNanoTime) //1595232070113890600
	//3、时间戳转化日期
	unixTime2 := 1595232070
	timeObj2 := time.Unix(int64(unixTime2),0)  //time.Unix(毫秒，纳秒)
	var str2 = timeObj2.Format("2006/01/02 15:04:05")
	fmt.Println(str2) //2020/07/20 04:01:10
	//4、日期字符串转化时间戳
	var str3 = "2020-02-28 15:23:24"
	var tmp3 ="2006-01-02 15:04:05"  //模板
	timeObj3,_ := time.ParseInLocation(tmp3, str3, time.Local )
	fmt.Println(timeObj3)   //2020-02-28 15:23:24
	fmt.Println(timeObj3.Unix())  //1582874604
	//5、time包中定义时间间隔类型的常量
	fmt.Println(time.Nanosecond) 
	fmt.Println(time.Microsecond) 
	fmt.Println(time.Millisecond)  //1毫秒
	fmt.Println(time.Second) //1秒
	fmt.Println(time.Minute) 
	fmt.Println(time.Hour) 
	//6、时间操作函数
	timeObj4 := timeObj.Add(time.Hour)
	fmt.Println(timeObj4) //增加1小时
	//7、go的定时器
	//7.1 定时器：time.NewTicker(time.Second)
	 ticker := time.NewTicker(time.Second)
	 n := 5
	 for t := range ticker.C{
		n--
		fmt.Println(t)
		if n == 0{
			ticker.Stop() //终止定时器继续执行
			break
		}
	 }
	 //7.2 休眠：time.Sleep(time.Second) 每隔一秒打印一次

}