//案例
package main
import "fmt"
import "sync"

var wg sync.WaitGroup//只定义，无需赋值
var total int
//加入互斥锁
var lock sync.Mutex

func add() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		lock.Lock()//加锁
		total = total + 1
		//解锁
		lock.Unlock()
	}
}

func sub() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		lock.Lock()//加锁
		total = total - 1
		//解锁
		lock.Unlock()
	}
}

func main() {//主线程
	//多个协程操纵同一数据
	wg.Add(2)
	go add()
	go sub()
	wg.Wait()
	fmt.Println(total)
	//理论结果为0，但实际结果不为0
	//问题分析:
	//两个协程操作同一数据total，导致数据竞争，导致结果不准确

	//使用互斥锁同步协程，可以避免错误
	//确保一个协程在执行逻辑的时候其他协程不执行
	
	//锁的机制
	//互斥锁：同一时刻只允许一个协程访问共享资源
	//同步机制：保证共享资源的正确性
	//此时就可以得到理论答案

	//互斥锁(Mutex)
	//Lock()上锁，Unlock()解锁
	//适用于读写不确定的场景，读写次数没有明显的区别
	//性能效率比较低
}