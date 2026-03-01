package main
import "fmt"
import "sync"
import "time"
//加入读写锁
var wg sync.WaitGroup//只定义，无需赋值
var lock sync.RWMutex

func read() {
	defer wg.Done()
	lock.RLock()//加读锁,只是读取数据时，这个锁不产生影响，当读写同时进行时，就会产生影响
	defer lock.RUnlock()//释放读锁
	fmt.Println("Reading...")
	time.Sleep(1 * time.Second)
	fmt.Println("Reading 成功...")
}

func write() {
	defer wg.Done()
	lock.Lock()//加写锁，当读写同时进行时，写锁会阻塞，直到读锁释放
	defer lock.Unlock()//释放写锁
	fmt.Println("Writing...")
	time.Sleep(1 * time.Second)
	fmt.Println("Writing 成功...")
}
func main() {
	wg.Add(5)
	//启动协程-- > 读多写少
	for  i := 0; i < 4; i++ {
		go read()
	}
	go write()

	wg.Wait()
	fmt.Println("读完了...")
	//读写锁

	//读写锁(RWMutex)
	//经常用于读 >>(远大于) 写的场景
	//在读的时候，数据间不产生影响，写和读之间才会产生影响

}