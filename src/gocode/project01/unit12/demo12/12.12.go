package main
import "fmt"
import "sync"
func printNum() {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}
//除法操作
func divide(){
	defer func(){
		err := recover()
		if err != nil{
			fmt.Println("错误为",err)
		}
	}()//recover捕获panic
	defer wg.Done()
	num1 := 10
	num2 := 0
	result := num1 / num2
	fmt.Println(result)
}//除数为零时出现panic

var wg sync.WaitGroup
func main() {
	wg.Add(2)
	//启动两个协程
	go printNum()
	go divide()
	wg.Wait()
}