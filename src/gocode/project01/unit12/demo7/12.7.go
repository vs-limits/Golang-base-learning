package main
import "fmt"
func main() {
	//管道 channel
	//管道特质:
	//本质时应该数据结构--队列
	//队列特点 先进先出
	//自身线程安全，多线程访问时，不需要加锁
	//管道具有类型，string类只能存放string类型的数据

	//案例
	//var Ich chan int //声明一个管道，类型为int
	//chan关键字
	//管道是一个引用类型，必须初始化才能写入数据，也就是make了之后才能使用

	//定义管道
	var Ich chan int
	Ich = make(chan int,4)//初始化，管道可以存放3个int类型数据
	//向管道中写入数据
	Ich <- 1 //写入数据
	Ich <- 2
	num := 10
	Ich <- num
	//不能存放大于容量的数据，否则会阻塞
	fmt.Printf("length:%v,cap:%v\n",len(Ich),cap(Ich))//实际长度是2，容量是3

	//从管道中读取数据
	num1 := <-Ich //读取数据
	fmt.Println(num1)
	num2 := <-Ich
	fmt.Println(num2)
	//符合先进先出
	//读取完毕后，管道为空，不能再读取数据，否则会阻塞

	//关闭管道
	close(Ich) //关闭管道，不再接收数据，但是可以读取已有数据
	num3 := <-Ich //读取已有数据
	fmt.Println(num3) //输出10
}