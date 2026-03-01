//案例
package main
import "fmt"
import "strconv"
import "time"

func test(){
	for i := 0; i < 10; i++ {
		fmt.Println("test + " + strconv.Itoa(i))
		//阻塞一秒
		time.Sleep(time.Second)
	}
}

func main() {//主线程
	go test() //需要开启协程才能交替运行


	for i := 0; i < 10; i++ {
		fmt.Println("main + " + strconv.Itoa(i))
		//阻塞一秒
		time.Sleep(time.Second)

		//主死从随
		//主线程结束，则协程也会强制结束
		if i == 6 {
			break
		}
	}

	//主线程和协程的执行流程

}