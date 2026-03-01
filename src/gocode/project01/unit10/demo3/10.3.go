package main
import "fmt"
type Person struct {
	Name string
	Age int
}

func (p Person) test() {
	fmt.Println(p.Name, p.Age)
}
//指针传递
func (p *Person) test2() {
	p.Name = "Jerry"
	p.Age = 25
	fmt.Println(p.Name, p.Age)
}
//基本数据类型方法
type integer int
func (i integer) test3() {
	i = 20
	fmt.Println(i)
}

type Student struct {
	Name string
	Age int
}
//指针绑定，实现String()方法
func (s *Student) String() string{
	str := fmt.Sprintf("Name: %s, Age: %d", s.Name, s.Age)//拼成一个字符串
	return str
}
//函数指定什么类型，传入参数就是什么类型
func method1(s *Student) {
	fmt.Println(s.Name)
}
//方法不一样，接收参数为值类型，可以传入指针类型
//接收参数为指针类型，可以传入值类型

func main() {
	//方法引入
	var p Person = Person{"Tom", 20}
	p.test()//是一个值传递，调用的是值类型的方法，不改变原本值
	(&p).test()//是一个指针传递，依旧可以调用方法，该形式虽然是指针传递，但是本质上还是值传递，不能改变原本值

	//注意事项一
	//在方法调用中,是值拷贝传递
	//可以通过结构体指针的方式处理
	p.test2()//是一个指针传递，调用的是指针类型的方法，改变原本值
	(&p).test2()//是一个指针传递，调用的是指针类型的方法，改变原本值

	//注意事项二
	//方法绑定在指定的数据类型上，而不是在结构体上
	//基本数据类型也可以有方法
	var i integer = 10
	i.test3()//调用的是值类型的方法，不改变原本值

	//注意事项三
	//方法访问控制权限，跟函数一样,方法首字母小写，只能在本包访问，大写则可被其他包访问
	stu := Student{"Mike", 25}
	fmt.Println(stu)
	fmt.Println(&stu)//传入地址才和指针绑定，绑定了String方法，会自动调用

	//方法与函数的区别
	//前者需要绑定数据类型
	//后者不需要绑定数据类型

	//调用方法不一样
	//方法调用需要传递接收者，函数调用不需要
	//函数名(参数)  变量.方法名(参数)
	method1(&stu)
}