package main
import "fmt"
//结构体
//方式一
type Person struct {
	Name string
	Age int
	School string
}

//方式二
type Student struct {
	Name string
	Age int
	School string
}
func main() {
	//面向对象设计
	//不是纯粹的面向对象
	//Golang没有class类，只有struct结构体

	//方式一
	var p1 Person
	fmt.Println(p1)
	p1.Name = "Rapeter"
	p1.Age = 18
	p1.School = "Harvard"
	fmt.Println(p1)

	//方式二
	var s1 Student = Student{"Tom",20,"Stanford"}
	fmt.Println(s1)

	var s2 Student = Student{
		Name: "Jerry",
		Age: 21,
	}
	fmt.Println(s2)

	//方式三 返回结构体指针
	var s *Student = new(Student)
	//s 为指向地址
	(*s).Name = "Jane"
	(*s).Age = 25
	(*s).School = "MIT"
	fmt.Println(*s)
	s.Name = "Alice"//实际转化成（*s）.Name = "Alice"

	//方式四
	var s0 *Student = &Student{"Bob",22,"Caltech"}


	//结构体转换
	//进行转换时需要有相同字段
	//结构体

}