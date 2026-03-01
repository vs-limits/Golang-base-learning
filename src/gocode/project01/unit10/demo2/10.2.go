package main
import "fmt"
type Person struct {
	Name string
	Age int
}

type Student struct {
	Name string
	Age int
}

type Stu Student
func main() {
	//结构体转换
	//进行转换时需要有相同字段

	//方式一
	var s Student = Student{"Tom", 20}
	var p Person = Person{"Jerry", 25}
	s = Student(p)//结构体转换，要用强制转换
	fmt.Println(s)

	//方式二
	var stu Stu = Stu{"Tom", 20}
	fmt.Println(stu)
	s = Student(stu)
}