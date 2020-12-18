package main
import (
	"fmt"
	"io/ioutil"

)
func copy(srcFileName string, dstFileName string)(err error){
	byteStr,err1 := ioutil.ReadFile(srcFileName)
	if err1 != nil{
		fmt.Println("err1=",err1)
		return err1
	}
	err2 := ioutil.WriteFile(dstFileName,byteStr,0666)
	if err2 != nil{
		fmt.Println("err2=",err2)
		return err2
	}
	return nil
}
func main(){
	srcFileName := "./text.txt"
	dstFileName := "./textcopy.txt"
	err := copy(srcFileName, dstFileName)
	
	if err != nil{
		fmt.Println("复制文件失败")
	}else{
		fmt.Println("复制文件成功")
	}
	
}
