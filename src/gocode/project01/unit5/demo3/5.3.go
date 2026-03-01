package main
import "fmt"

var num int = test()//全局变量的定义
func test() int {
	fmt.Println("test function")
	return 10
}
func init() {//init函数定义
	fmt.Println("init function")
}//初始化函数，用于进行一些初始化的操作，会在main函数之前执行

func getSum() func (int) int {
	var sum int = 0
	return func (num int) int {
		sum += num
		return sum
	}
}

func main() {//main函数的调用
	//init函数
	fmt.Println("main function")

	//匿名函数
	result := func(num1 int ,num2 int) int{//只使用一次
		fmt.Println("anonymous function")
		return num1 + num2
	}(10,20)
	fmt.Println(result)

	//可以将函数赋给变量,此时该变量代表函数,可以多次调用
	sub := func(num1 int,num2 int) int{
		fmt.Println("sub function")
		return num1 - num2
	}
	result01 := sub(10,5)
	fmt.Println(result01)

	//闭包，类似递归操作。本质时一个匿名函数，只是引入外界的变量
	f := getSum()
	fmt.Println(f(10))
	fmt.Println(f(20))//每次调用都会累加,sum 并不会清零
	//匿名函数 + 引用的变量 = 闭包
	//闭包的特点：函数嵌套，函数引用了外部变量，返回一个函数,不清除内存

	fmt.Println("-----------")
	//defer关键字
	//defer关键字可以让函数在return之后执行，即使函数发生异常也会执行,为了释放资源
	fmt.Println(add(10,20))
}

func add(num1 int,num2 int) int {
	//程序遇到defer关键字，不会立即执行函数体，而是将函数的调用存入一个栈中，等到函数返回时再依次执行
	defer fmt.Println(num1)
	defer fmt.Println(num2)

	var result int = num1 + num2
	fmt.Println(result)
	return result
}