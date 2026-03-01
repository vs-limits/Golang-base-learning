//案例
package main
import "fmt"
import "time"

func main() {//主线程
	//启动多个协程
	//匿名函数加外部变量 --> 闭包
	for i := 0; i < 5; i++ {
		//先启动一个协程
		//使用匿名函数，直接调用
		go func(n int){
			fmt.Println("协程",n)
		}(i)//外部变量i传递给匿名函数
	}
	time.Sleep(time.Second * 2)//阻塞几秒就乘以几
}