package main
import (
	"fmt"
	//"io/ioutil"
	"os"
	//"io"
)


func main(){
	err1 := os.Mkdir("./dir",0666)
	if err1 != nil{
		fmt.Println("err1=",err1)
		return 
	}
	err2 := os.MkdirAll("./dir1/dir2/dir3", 0666)
	if err2 != nil{
		fmt.Println("err2=",err2)
		return 
	}

}