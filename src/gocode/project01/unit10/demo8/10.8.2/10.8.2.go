package main
import "fmt"
type AInterface interface {//多个接口实现
	a()
}
type BInterface interface {
	b()
}
type Stu struct {
	
}
func (s Stu) a() {
	fmt.Println("a")
}
func (s Stu) b() {
	fmt.Println("b")
}

//继承接口
type CInterface interface {
	AInterface
	BInterface
	c()
}
type C struct {

}
func (c C) c() {
	fmt.Println("c")
}
func (c C) a() {
	fmt.Println("c-a")
}
func (c C) b() {
	fmt.Println("c-b")
}

//空接口
type E interface {

}


func main() {
	//注意事项二

	//一个自定义类型可以实现多个接口
	var s Stu
	var a AInterface = s
	var b BInterface = s
	//a.a() = s.a()
	//b.b() = s.b()
	a.a()
	b.b()

	//一个接口可以继承多个接口，但是要实现该接口，需要实现所有接口的方法
	var c C
	var CI CInterface = c
	CI.a()
	CI.b()
	CI.c()

	//interface类型默认是一个指针，若没有对interface初始化就使用，会输出nil

	//空接口没有任何方法，用途:可以用该方法定义接收任何类型的形参，做类似泛型的效果
	var e E = s
	fmt.Println(e)
	var e2 interface{} = s
	fmt.Println(e2)
}