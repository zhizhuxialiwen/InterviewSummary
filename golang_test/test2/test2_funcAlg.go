package main

import (
	"fmt"
	"sort"
	//"strings"
)

//1、升序排序：选择排序
func SortIntAsc(slice1 []int) []int{
	for i:=0; i < len(slice1); i++{
		for j := i+1; j < len(slice1); j++{
			if slice1[i] > slice1[j]{
				temp := slice1[i]
				slice1[i] = slice1[j]
				slice1[j] = temp
			}
		}
	}
	return slice1
}

//1、降序排序：选择排序
func SortIntDesc(slice1 []int) []int{
	for i:=0; i < len(slice1); i++{
		for j := i+1; j < len(slice1); j++{
			if slice1[i] < slice1[j]{
				temp := slice1[i]
				slice1[i] = slice1[j]
				slice1[j] = temp
			}
		}
	}
	return slice1
}

func mapSort(mapStr1 map[string] string) string{
	var sliceKey []string
	//1将map的key放到切片
	for k,_ := range mapStr1{
		sliceKey = append(sliceKey, k)
	}
	//2 切片：对key进行升序排序
	sort.Strings(sliceKey)

	var str string
	for _,v := range sliceKey{
        str += fmt.Sprintf("%v--%v->",v,mapStr1[v])
	}
	return str
}

func main(){
	var sliceInt = []int{1,5,8,3,4}
	arr1 := SortIntAsc(sliceInt)
	fmt.Println(arr1)  //[1 3 4 5 8]

	arr2 := SortIntDesc(sliceInt)
	fmt.Println(arr2) //[8 5 4 3 1]


	var mapStr1 = map[string]string{
		"username": "liwen",
		"age": "20",
		"sex": "man",
	}
	str1 := mapSort(mapStr1)
	fmt.Println(str1)
}