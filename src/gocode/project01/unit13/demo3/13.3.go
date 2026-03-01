package main
import "fmt"
import "net"//网络编程包
import "bufio"
import "os"
func main() {
	//TCP编程 创建客户端
	//先调用Dial()函数，传入目标IP地址和端口号，建立TCP连接
	fmt.Println("客户端启动")
	//Dial参数指定tcp协议，目标IP地址和端口号，最后是发送的内容
	conn,err := net.Dial("tcp", "127.0.0.1:8080")
	//返回参数conn是一个net.Conn接口，用于读写数据
	if err != nil {
		fmt.Println("连接失败",err)
		return
	}
	fmt.Println("连接成功",conn)

	//发送终端数据
	//通过客户端发送单行数据，然后退出
	reader := bufio.NewReader(os.Stdin)//os.Stdin代表终端标准输入

	//从终端读取一行用户输入的信息
	str,err := reader.ReadString('\n')//以换行符为结尾
	if err != nil {
		fmt.Println("读取失败",err)
	}

	//输入成功就将str数据发送给服务器
	n, err2 := conn.Write([]byte(str))//Write函数需要传一个切片
	//所以强制转换成[]byte类型
	//n代表发送的字节数，err2代表错误信息
	if err2 != nil {
		fmt.Println("发送失败",err2)
	}
	fmt.Printf("数据通过客户端发送成功，一共发送了%d字节数据,并退出\n",n)
}