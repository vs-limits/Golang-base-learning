package main
import "fmt"
import "net"//网络编程包

func process(conn net.Conn) {
	//用完连接一定要关闭
	defer func() {
		fmt.Printf("客户端 %s 断开连接\n", conn.RemoteAddr().String())
		conn.Close()
	}()

	for {
		//创建一个切片,将读取的数据放入切片
		buf := (make([]byte,1024))

		//从conn中读取数据
		n,err := conn.Read(buf)
		//n表示读取的字节数,err表示错误信息
		//因为不一定读满1024个字节，接收读取的n个字节
		//以便后续输出终止
		if err!= nil {
			fmt.Println("读取数据失败",err)
			return
		}
		//将读取的数据在服务器端输出
		fmt.Println("接收到数据:",string(buf[:n]))

	}
}

func main() {
	//TCP编程 创建服务器端
	fmt.Println("服务器端启动")
	//客户端要进行监听，需要listen函数,
	//参数 指定tcp协议，目标IP和端口
	listen,err := net.Listen("tcp","127.0.0.1:8080")
	if err!= nil {
		fmt.Println("服务器端启动失败",err)
		return
	}
	//等待客户端连接
	//循环等待客户端链接
	for{
		conn,err2 := listen.Accept()
		if err2!= nil {
			fmt.Println("客户端连接失败",err2)
			continue
		}else {
			//链接成功:
			fmt.Printf("客户端连接成功, conn类型: %v, 客户端信息: %v\n", conn, conn.RemoteAddr().String())//后者反应客户端的IP地址和端口号
		}

		//接收终端数据
		//准备一个协程，协程处理客户端服务请求
		go process(conn)//不同的客户端请求，链接的conn不一样
	}

	//链接测试
	//需要先启动服务器端，然后客户端连接服务器端，然后发送数据，服务器端接收数据并打印。
}