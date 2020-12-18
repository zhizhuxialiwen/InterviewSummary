package main

import (
	"fmt"
	"sort"
	"strings"
)

func main(){
	//1声明map
	//1.1、方法一：make创建map类型的数据
	var mapStr1 = make(map[string]string)
	mapStr1["userName"] = "李文"
	mapStr1["age"] = "30"
	mapStr1["sex"] = "男"
	fmt.Println(mapStr1["userName"])
	fmt.Println(mapStr1) //map[age:30 sex:男 userName:李文]

	//1.2 方法二：
	var mapStr2 = map[string]string{
		"userName" : "zhangsan",
		"age":"20",
		"sex":"man", 
	}
	fmt.Println(mapStr2) //map[age:20 sex:man userName:zhangsan]
	//1.3 方法三
	mapStr3 := map[string]string{
		"userName" : "zhangsan",
		"age":"20",
		"sex":"man", 
	}
	fmt.Println(mapStr3) 
	//2、循环遍历for ... range
	for k,v := range mapStr3{
        fmt.Printf("key=%v, value=%v \n",k,v)
	}

	//3、修改map数据
	mapStr3["userName"] = "lisi"
	fmt.Println(mapStr3) //map[age:20 sex:man userName:lisi]
	//4、查找
	fmt.Println(mapStr3["userName"])
	//5 、判断key是否查存在？
	v,ok := mapStr3["userName"]
	fmt.Println(v,ok)
	//6、delete(map, key):删除map的key和value
	delete(mapStr3,"userName")
	fmt.Println(mapStr3)  //map[age:20 sex:man]
    //7、定义一个map类型key的切片
	var mapUserInfo = make([]map[string] string, 3, 3)
	fmt.Println(mapUserInfo[0])
	fmt.Println(mapUserInfo[0] == nil) //不初始化为nil
	if mapUserInfo[0] == nil{
		mapUserInfo[0] =make(map[string] string)
		mapUserInfo[0]["userName"] = "李文"
		mapUserInfo[0]["age"] = "30"
		mapUserInfo[0]["sex"] = "男"
	}
	if mapUserInfo[1] == nil{
		mapUserInfo[1] =make(map[string] string)
		mapUserInfo[1]["userName"] = "李为"
		mapUserInfo[1]["age"] = "39"
		mapUserInfo[1]["sex"] = "男"
	}
	fmt.Println(mapUserInfo) //[map[age:30 sex:男 userName:李文] map[age:39 sex:男 userName:李为] map[]]

	for _,v := range mapUserInfo{
		//fmt.Println(v)
		for key,value := range  v{
			fmt.Printf("key=%v, value=%v \n",key,value)
		}
	}

	//8、定义map类型值value的切片slice
	var mapUserInfo1 = make(map[string][]string)
	mapUserInfo1["hobby"] = []string{
		"吃饭",
		"旅游",
		"看书",
	}
	mapUserInfo1["work"] = []string{
		"java",
		"php",
		"go",
	}
	fmt.Println(mapUserInfo1)  //map[hobby:[吃饭 旅游 看书]]
	for _,v := range mapUserInfo1{
		//fmt.Println(v)
		for key,value := range  v{
			fmt.Printf("key=%v, value=%v \n",key,value)
		}
	}

	// 9、map：引用类型:改变副本导致主本也改变
	mapUserInfo2 := make(map[string] string)
	mapUserInfo2["username"] = "liwen"
	mapUserInfo2["age"] = "20"
	mapUserInfo3 := mapUserInfo2
	mapUserInfo3["username"] = "lisan"
	fmt.Println(mapUserInfo2) //map[age:20 username:lisan]
	fmt.Println(mapUserInfo3)  //map[age:20 username:lisan]

	//10、map排序
	var mapInt = make(map[int]int)
	mapInt[10] = 100
	mapInt[2] = 5
	mapInt[6] = 8
	mapInt[4] = 99
	fmt.Println(mapInt) //map[2:5 4:99 6:8 10:100]
	
	for k,v := range mapInt{
		fmt.Println(k,v) 
	}
	/*输出结果：
	4 99
	10 100
	2 5
	6 8
	*/
	// 10.1 按照key升序输出map
	//把map的key放在切片里
	var sclie10 []int
	for k, _ := range mapInt{
	    sclie10 = append(sclie10, k)
	}
	fmt.Println(sclie10) //[10 2 6 4]
	//key升序排序
	sort.Ints(sclie10)
	fmt.Println(sclie10) //[2 4 6 10]
	for k, v := range sclie10{
	    fmt.Printf("key=%v,value=%v \n",k,v)
	}
	//算法：统计字符串出现单词的次数
	var str1 = "liwen is good is good good"
	var strSlice = strings.Split(str1, " ")
	fmt.Println(strSlice) //[liwen is good is good good]

	var strMap = make(map[string] int)
	for _, v := range strSlice{
	    strMap[v]++
	}
	fmt.Println(strMap) //map[good:3 is:2 liwen:1]

}