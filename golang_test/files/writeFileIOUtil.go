package main
import (
	"fmt"
	"io/ioutil"

)

func main(){
	str := "helo liwen"
	err := ioutil.WriteFile("./textWrite.txt",[]byte(str), 0666)
	if err != nil{
		fmt.Println("Write file failed ,err:",err)
		return
	}
}