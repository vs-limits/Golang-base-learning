package main
import "fmt"
import "io/ioutil"
func main() {
	//文件复制操作
	//定义源文件
	filePath := "C:/Users/limits/Desktop/Golang/test.txt"
	file2Path := "C:/Users/limits/Desktop/Golang/test2.txt"

	//打开文件进行读取
	content,err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("读取错误:",err)
		return
	}
	//直接写出文件
	err = ioutil.WriteFile(file2Path,content,0666)//写入哪个路径里，什么内容，权限设置
	if err!= nil {
		fmt.Println("写入错误:",err)
		return
	}
	fmt.Println("文件复制成功")
}