package main
import "fmt"
import "io"//io包 --> io.EOF
import "io/ioutil"//io流包下的ioutil包
import "os"//os包
import "bufio"//bufio包
func main() {
	//io流的引入
	//流 对文件进行操作
	//针对程序定义方向

	//读取文件1，适用于文件较小的情况
	content,err := ioutil.ReadFile("C:/Users/limits/Desktop/Golang/test.txt")
	if err!= nil {
		fmt.Println("读取错误为 ",err)
	} else {
		fmt.Println("内容为",string(content))
	}//读取操作不用打开和关闭文件，直接读取即可
	//Open和Close操作已经被封装进ReadFile函数内部了

	//读取文件2，适用于文件较大或者需要流式读取的情况
	//带缓冲的流式读取(分批次读取)
	//这种方式需要自己操作打开和关闭文件
	//不需要io/ioutil包的帮助，需要os包和bufio包
	file0,err0 := os.Open("C:/Users/limits/Desktop/Golang/test.txt")
	if err0 != nil {
		fmt.Println("打开文件错误为 ",err0)
	}
	//当函数退出时，让file0关闭，防止内存泄漏
	defer file0.Close()
	//创建一个流
	reader := bufio.NewReader(file0)
	for {
		str,err1 := reader.ReadString('\n')//读到换行就是一次读取
		fmt.Println(str)
		if err1 == io.EOF {//读取到文件末尾 EOF == end of file
			break
		}
	}

	fmt.Println("读取结束")
}