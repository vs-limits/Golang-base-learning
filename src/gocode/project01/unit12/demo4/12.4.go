//案例
package main
import "fmt"
import "sync"

var wg sync.WaitGroup//只定义，无需赋值
func main() {//主线程
	//使用WaitGroup控制协程退出
	//WaitGroup的作用是等待一组协程完成，然后再继续执行下面的代码
	//具有三个方法，Add，Done，Wait
	//Add方法用于添加协程的数量
	//Done方法用于完成一个协程
	//Wait方法用于等待所有协程完成

	for i := 0; i < 5; i++ {
		wg.Add(1)//协程开始时加1
		go func(n int){
			fmt.Println("协程",n)
			defer wg.Done()// 协程执行完成减1
		}(i)//外部变量i传递给匿名函数
	}
	//主线程一直在阻塞，当wg减为0时，就停止阻塞
	wg.Wait()

	//如果防止忘记计数器减一操作，可以结合defer使用

}