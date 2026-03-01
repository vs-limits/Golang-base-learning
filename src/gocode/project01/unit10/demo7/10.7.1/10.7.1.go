package main
import "fmt"
//继承引入
type Animal struct {
	Weight int
	Age int
}
func (an *Animal) Shout(){
	fmt.Println("Animal Shout")
}
func (an *Animal) Show(){
	fmt.Printf("Age : %d, Weight : %d\n", an.Age, an.Weight)
}
type Cat struct {
	//加入匿名结构体
	Animal 
	Age int//重复属性
}

//重复属性调用
func (c *Cat) Show(){
	fmt.Printf("····Cat Age : %d, Weight : %d\n", c.Age, c.Weight)
}

func (c *Cat) Scratch(){
	fmt.Println("Cat Scratch")
}
func main() {
	cat := &Cat{}
	cat.Animal.Weight = 5
	cat.Age = 2
	//等同于
	//cat.Weight = 5
	//cat.Animal.Age = 2
	cat.Show()
	cat.Shout()
	cat.Scratch()
	//继承的优点:代码复用，以及代码拓展

	//注意事项一
	//结构体可以使用嵌套匿名结构体的所有属性与方法

	//匿名结构体字段访问可以简化
	//cat.Weight = 5

	//当结构体与匿名结构体有相同的属性与方法时，采用就近访问原则
	//就近访问原则:
	cat2 := &Cat{}
	cat2.Weight = 5
	cat2.Age = 10//会赋值给cat结构体里的Age属性
	cat2.Show()//重名方法也会先调用cat结构体的Show方法，类似重载
	//可以通过匿名结构体名区分
	cat.Animal.Age = 12//会赋值给Animal结构体里的Age属性
	cat.Animal.Show()//调用Animal结构体的Show方法
	cat.Show()//调用cat结构体的Show方法
}