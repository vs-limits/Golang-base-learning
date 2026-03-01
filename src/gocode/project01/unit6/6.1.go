package main
import "fmt"
import "errors"
func main() {
	//defer + recover机制处理错误
	err := test()
	fmt.Println("+++++++++++")
	//自定义错误
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

}

func test() (err error) {
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		fmt.Println("Error:", err)
	// 	}//recover()函数用于捕获panic()抛出的异常，并恢复程序的正常运行。
	// }()
	num1 := 10
	num2 := 0
	if num2 == 0 {
		return errors.New("除数不能为0")
	}else {
		result := num1 / num2
		fmt.Println(result)
		return nil
	}
}