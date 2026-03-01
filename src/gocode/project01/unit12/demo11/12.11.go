package main
import "fmt"
import "time"
func main() {
	//select功能(多路复用)，解决多个管道的选择问题
	//可以从多个管道中随机公平地选择一个来执行
	
	//定义一个int管道
	Ich := make(chan int,3)
	go func() {
		time.Sleep(time.Second * 5)
		Ich <- 10
	}()

	//定义一个string管道
	Sch := make(chan string,3)
	go func() {
		time.Sleep(time.Second * 2)
		Sch <- "limit"
	}()

	//fmt.Println(<-Ich)//本身取数据就是阻塞
	select{
		case v := <-Ich:
			fmt.Println(v)
		case v := <-Sch:
			fmt.Println(v)
		//case之间为并列关系，只要有一个case可以执行，就会执行
		default:
			fmt.Println("防止select被堵塞")
	}
}
// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(2) // 等待两个goroutine完成

// 	// 定义int通道
// 	Ich := make(chan int, 3)
// 	go func() {
// 		defer wg.Done()
// 		time.Sleep(5 * time.Second)
// 		Ich <- 10
// 		fmt.Println("Ich 数据已发送")
// 	}()

// 	// 定义string通道
// 	Sch := make(chan string, 3)
// 	go func() {
// 		defer wg.Done()
// 		time.Sleep(2 * time.Second)
// 		Sch <- "limit"
// 		fmt.Println("Sch 数据已发送")
// 	}()

// 	// 独立goroutine处理通道数据
// 	go func() {
// 		counter := 0
// 		for counter < 2 { // 确保两个通道都处理
// 			select {
// 			case v := <-Ich:
// 				fmt.Println("Ich 内容:", v)
// 				counter++
// 			case v := <-Sch:
// 				fmt.Println("Sch 内容:", v)
// 				counter++
// 			}
// 		}
// 	}()

// 	wg.Wait()      // 等待发送goroutine完成
// 	time.Sleep(1) // 确保处理goroutine有足够时间
// }