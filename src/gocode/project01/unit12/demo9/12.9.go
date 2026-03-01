package main
import "fmt"
import "time"
import "sync"
//写
func writeDate(Ich chan int) {
	defer wg.Done()
	for i := 1; i <= 50; i++ {
		Ich <- i
		fmt.Println("writeDate:", i)
		time.Sleep(time.Second * 0.5)
	}

	close(Ich)//写入数据后关闭管道
}

//读
func readDate(Ich chan int) {
	defer wg.Done()
	for value := range Ich {
		fmt.Println("readDate:", value)
		time.Sleep(time.Second * 0.5)
	}
}
var wg sync.WaitGroup
func main() {
	wg.Add(2)
	//协程与管道协同工作
	//定义管道
	Ich := make(chan int,50)
	//开启一个writeDate协程，将数据写入管道
	go writeDate(Ich)
	//开始一个readDate协程，从管道读取数据
	go readDate(Ich)
	//主线程需要等待两个协程结束，才能退出
	wg.Wait()
	fmt.Println("main end")
}