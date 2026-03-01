package main
import "fmt"
//接口定义
type SayHello interface {//interface关键字定义接口
	//声明没有实现的方法
	Say()
}

//实现接口 -> 定义一个结构体
type Chinese struct {

}
type British struct {

}//将方法绑定到结构体上，进行实现
func (c Chinese) Say() {//Say是方法需要传参，所以需要加括号
	fmt.Println("你好")
}
func (b British) Say() {
	fmt.Println("Hello")
}

//定义一个函数，专门用来测试接口的使用,用来接收具备Say()接口能力的变量
func greet(s SayHello) {//接收接口变量
	s.Say()
}

//自定义数据类型
type intType int
func (i intType) Say() {
	fmt.Println("intType")
}

func main() {
	//接口的引入，定义规则，定义规范
	//接口中可以定义一组方法，但不需要实现，不需要方法体，且接口中不包含任何变量
	//实现接口要实现所有的方法才是实现接口
	//Golang中的接口，不需要显式的实现接口且没有implement关键字
	//接口的目的是为了规范定义，具体由别人实现即可

	c := Chinese{}
	b := British{}

	greet(c)
	c.Say()
	greet(b)//等效的
	b.Say()

	//注意事项一
	//接口本身不能创建实例，但可以指向一个实现该接口的变量
	//var s SayHello
	//s.Say() //报错，接口没有实现方法
	var s SayHello = c
	s.Say()//该方法可以调用，因为c实现了Say()方法

	//只要是自定义数据类型就可以实现接口，不一定非得是结构体
	var i intType = 10
	var s1 SayHello = i
	i.Say()//调用自定义数据类型的方法
	s1.Say()//调用自定义数据类型的方法

	//一个自定义类型可以实现多个接口
	
	//一个接口可以继承多个接口，但是要实现该接口，需要实现所有接口的方法

	//interface类型默认是一个指针，若没有对interface初始化就使用，会输出int

	//空接口没有任何方法，
}