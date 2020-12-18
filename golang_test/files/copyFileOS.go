package main
import (
	"fmt"
	//"io/ioutil"
	"os"
	"io"

)
func copy(srcFileName string, dstFileName string)(err error){
	readFile,err1 := os.Open(srcFileName)
	if err1 != nil{
		fmt.Println("err1=",err1)
		return err1
	}
	defer readFile.Close()

	writeFile,err2 := os.OpenFile(dstFileName, os.O_CREATE|os.O_WRONLY,0666)
	if err2 != nil{
		fmt.Println("err2=",err2)
		return err2
	}
	defer writeFile.Close()
	var tempSlice = make([]byte, 12800)
    for{
		//读取数据
		n1,err3 := readFile.Read(tempSlice)
		if err3 == io.EOF{
			break
		}
		if err3 != nil{
			return err3
		}
        //写入数据
		if _,err4 := writeFile.Write(tempSlice[:n1]); err4 != nil{
			return err4
		}
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

	err2 := os.Rename(dstFileName,"./renamefile.txt")  
	if err2 != nil{
		fmt.Println("重命名文件失败")
	}
	
}
