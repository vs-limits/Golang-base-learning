package main
import "fmt"
import "os"
func main() {
	//文件
	//文件是保存数据的地方，是数据源的一种
	//like : word，txt，excel，jpg
	//主要作用是保存数据

	//File结构体---打开和关闭文件
	//打开文件---Open函数
//传入一个字符串，返回文件的指针，和是否打开成功
	file,err := os.Open("C:/Users/limits/Desktop/Golang/test.txt")
	if err!= nil {
		fmt.Println("打开错误为:",err)
	}else {
		fmt.Println("文件打开成功" ,file)//file接收的是指针，是一个地址
	}

	//关闭文件---close函数
	err2 := file.Close()
	if err2!= nil {
		fmt.Println("关闭错误为:",err)
	}else {
		fmt.Println("文件关闭成功")
	}

	//使文件不能读写，返回可能的错误
}