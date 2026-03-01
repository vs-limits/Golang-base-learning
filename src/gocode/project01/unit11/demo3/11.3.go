package main
import "fmt"
import "os"
import "bufio"
import "io"

func main() {
	//写出文件
	//打开文件操作 OpenFile(name string, flag int, perm FileMode)
	//三个参数的含义 --> 要打开文件的路径，文件打开模式，权限控制(windows系统下设置无效)
	file,err := os.OpenFile("C:/Users/limits/Desktop/Golang/test.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	//返回内容: 该文件的指针，和 是否成功
	if err!= nil {
		fmt.Println("打开失败 err:", err)
		return
	}
	//及时关闭文件，防止内存泄漏
	defer file.Close()

	//写入文件操作 --> IO流 -->缓冲输出流
	writer := bufio.NewWriter(file)
	for i := 0;i < 5;i++ {
		writer.WriteString("写入测试 ")
	}//现在都写入缓冲区了
	//刷新数据后，才会写入文件
	writer.Flush()

	file0,_ := os.Open("C:/Users/limits/Desktop/Golang/test.txt")
	defer file0.Close()
	reader := bufio.NewReader(file0)
	for {
		str,err1 := reader.ReadString('\n')//读到换行就是一次读取
		fmt.Println(str)
		if err1 == io.EOF {//读取到文件末尾 EOF == end of file
			break
		}
	}

	fmt.Println("读取结束")

	//查看权限设置
	s := os.FileMode(0666).String()
	fmt.Println("权限设置:", s)
}