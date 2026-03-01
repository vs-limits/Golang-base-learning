package main
import "fmt"
//接口定义
type SayHello interface {//interface关键字定义接口
	//声明没有实现的方法
	Say()
}

//实现接口 -> 定义一个结构体
type Chinese struct {
	name string
}
type British struct {
	name string
}//将方法绑定到结构体上，进行实现
func (c Chinese) Say() {//Say是方法需要传参，所以需要加括号
	fmt.Println(c.name + "你好")
}
func (b British) Say() {
	fmt.Println(b.name + "Hello")
}

//定义一个函数，专门用来测试接口的使用,用来接收具备Say()接口能力的变量
func greet(s SayHello) {//接收接口变量
	s.Say()
}//调用接口方法，体现多态


func main() {
	//多态
	//多态通过接口实现，可以按照统一的接口调用不同的实现，此时接口变量就呈现不同的形态
	c := Chinese{}
	b := British{}
	greet(c)
	greet(b)

	//多态参数
	//s SayHello -> 多态的参数
	//多态数组
	var arr [3]SayHello
	arr[0] = Chinese{"琳"}
	arr[1] = British{"Lily"}
	arr[2] = Chinese{"盖"}
	fmt.Println(arr)
}