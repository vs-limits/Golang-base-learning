package model

type person struct {
	Name string
	Age  int
}

//工厂模式
func NewPerson(name string, age int) *person {
	return &person{name,age}
}

//定义set与get方法，对属性进行封装
func (p *person) SetName(name string) {
	p.Name = name
	fmt.Println("set name success")
}

func (p *person) GetName() int {
	return p.Name
}