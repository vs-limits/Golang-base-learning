package model

type Student struct {
	Name string
	Age  int
}

type person struct {
	Name string
	Age  int
}

//工厂模式
func NewPerson(name string, age int) *person {
	return &person{name,age}
}//工厂方法​​适合需要频繁扩展产品的场景