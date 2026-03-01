package main
import "fmt"
func main() {
	//管道可以声明为只读或只写性质
	//默认情况下，管道是双向的
	//声明为只写
	var ch1 chan<- int//管道具备<-(在chan后) 只写性质
	ch1 = make(chan int,3)
	ch1 <- 20
	//num := <-ch1  会报错，因为ch1只写
	//fmt.Println("num:",num)
	fmt.Println("ch1:",ch1)


	var ch chan int
	ch = make(chan int,3)
	ch <- 10
	ch <- 20

	//声明为只读
	var ch2 <-chan int//管道具备<-(chan前) 只读性质
	ch2 = make(chan int,3)
	ch2 = ch//通过一个双向管道赋值给只读管道，再进行输出
	if ch2 != nil {
		num := <-ch2
		fmt.Println("num:",num)
	}
	//ch2 <- 30  会报错，因为ch2只读
	fmt.Println("ch2:",ch2)

	//管道的阻塞
	//只写不读，输入的个数超过管道的容量，会导致阻塞
	//或者写的快，读的慢，会导致输出的个数超过管道的容量，导致阻塞
}