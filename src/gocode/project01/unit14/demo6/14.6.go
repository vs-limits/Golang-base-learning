package main
import "fmt"
import "reflect"

//定义一个结构体
type Student struct {
	Name string
	Age  int
}
func (s Student) Print() {
	fmt.Println("学生的姓名：", s.Name)
	fmt.Println("学生的年龄：", s.Age)
}
func (s Student) GetSum(n1,n2 int) int{
	fmt.Println("方法GetSum被调用了")
	return n1 + n2
}
func (s Student) Set(name string, age int) {
	s.Name = name
	s.Age = age
}

//定义函数操作结构体
func testStruct(a interface{}) {//a为接口类型,是一个空接口
	//a转换成reflect.Value类型,获取实际值
	val := reflect.ValueOf(a)
	fmt.Println(val)

	//Elem()方法获取指针的地址,然后通过NumField()方法获取结构体的属性数目
	n := val.Elem().NumField()
	fmt.Println("结构体的属性数目：", n)

	//修改属性的值
	val.Elem().Field(0).SetString("张三")
	val.Elem().Field(1).SetInt(25)
}
func main() {
	//反射修改变量的值
	stu := Student{
		Name:"joker",
		Age:20,
	}
	testStruct(&stu)//要改变值，需要传指针地址
}