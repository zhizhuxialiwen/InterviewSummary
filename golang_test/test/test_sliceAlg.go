package main

import (
	"fmt"
	"sort"
)

func main(){
	//1、简单排序：每一轮选出最小值； 从小到大排序
	var slice1 = []int{3, 5, 4, 33, 22, 55}
	for i := 0; i < len(slice1); i ++ {
		for j := i + 1; j < len(slice1); j++ {
            if slice1[i] > slice1[j]{  //从小到大排序 ;  if slice1[i] < slice1[j] :从大到小
				temp := slice1[i]
				slice1[i] = slice1[j]
				slice1[j] = temp
			}
		}
	}
	fmt.Println(slice1)  //[3 4 5 22 33 55]
	//2、冒泡排序：每一轮冒泡最大值，放在最后位置
	var slice2 = []int{3, 5, 15, 11, 22, 55}
	for i := 0; i < len(slice2); i ++ {
		for j := 0; j < len(slice2)-1-i; j++ {
            if slice2[j] > slice2[j+1]{  //从小到大排序 ;  if slice2[j] < slice2[j+1] :从大到小
				temp := slice2[j]
				slice2[j] = slice2[j+1]
				slice2[j] = temp
			}
		}
	} 
	fmt.Println(slice2)  
	//3、包排序 sort:升序
	var slice3 = []int{3, 5, 15, 13, 12, 55}
	var sclieFloat = []float64{3.2,4.5,5.4 , 2.3, 1.2}
	var scliceStr = []string{"a","c","b","g","e"}
	sort.Ints(slice3)
	sort.Float64s(sclieFloat)
	sort.Strings(scliceStr)

	fmt.Println(slice3) //[3 5 12 13 15 55]
	fmt.Println(sclieFloat) //[1.2 2.3 3.2 4.5 5.4]
	fmt.Println(scliceStr)  //[a b c e g]
	//3.2反转：降序排序
	var slice4 = []int{3, 5, 15, 13, 12, 55}
	var sclieFloat4 = []float64{3.2,4.5,5.4 , 2.3, 1.2}
	var scliceStr4 = []string{"a","c","b","g","e"}
	sort.Sort(sort.Reverse(sort.IntSlice(slice4)))
	sort.Sort(sort.Reverse(sort.Float64Slice(sclieFloat4)))
	sort.Sort(sort.Reverse(sort.StringSlice(scliceStr4)))
	fmt.Println(slice4) //[55 15 13 12 5 3]
	fmt.Println(sclieFloat4) //[5.4 4.5 3.2 2.3 1.2]
	fmt.Println(scliceStr4)  //[g e c b a]
}