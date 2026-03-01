package main
import "fmt"
func main() {
	//管道 channel
	//管道特质:
	//本质时应该数据结构--队列
	//队列特点 先进先出
	//自身线程安全，多线程访问时，不需要加锁
	//管道具有类型，string类只能存放string类型的数据

	var Ich chan int
	Ich = make(chan int,100)
	for i := 0; i < 100; i++ {
		Ich <- i
	}
	//管道的遍历
	//支持for range遍历
	//遍历时，若管道没有关闭，会出现deadlock错误
	//若管道已经关闭，则会正常遍历管道，完成后会自动退出
	close(Ich)
	for value := range Ich {//管道没有索引
		fmt.Println(value)
	}
}