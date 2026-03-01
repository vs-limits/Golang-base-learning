package main
import "fmt"
type SayHello interface {
	Say()
}

type Chinese struct {
	name string
}
type British struct {
	name string
}
func (c Chinese) Say() {
	fmt.Println(c.name + "你好")
}
func (c Chinese) Conly() {
	fmt.Println("中国人独有")
}
func (b British) Say() {
	fmt.Println(b.name + "Hello")
}
func (b British) Bonly() {
	fmt.Println("英国人独有")
}


func greet(s SayHello) {
	s.Say()//只能调用共有方法
	//用断言判断s的类型，再调用对应的方法
	//断言的不同语法
	// 第一种形式
	// var ch Chinese = s.(Chinese)//s是否能转成chinese类型
	// ch.Conly()
	// 第二种形式
	// ch2,flag := s.(Chinese)
	// if flag == true {
	// 	ch2.Conly()
	// }else{
	// 	fmt.Println("不是中国人")
	// }//解决不同类型的返回值
	// 第三种形式
	// if ch,flag := s.(Chinese);flag == true{
	// 	ch.Conly()
	// }else{
	// 	fmt.Println("not Chinese")
	// }

	//switch type的用法
	switch s.(type){//type属于go中的一个关键字,固定写法
		case Chinese:
			ch := s.(Chinese)
			ch.Conly()
		case British:
			b := s.(British)
			b.Bonly()
	}
}
func main() {
	//断言的引入
	//可以直接判断是否是该类型的变量
	//例: value，ok = element.(T),value是变量的值，ok是一个bool类，element是interface变量，T是断言的类型
	c := Chinese{}
	b := British{}
	greet(c)
	greet(b)

	//断言的语法
	
}