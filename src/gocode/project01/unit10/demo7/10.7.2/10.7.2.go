package main
import "fmt"
type A struct {
	a int
	b string
}
type B struct {
	c int
	d string
	a int//重复属性
}
type C struct {
	A
	B
	//基本数据类型
	int
}
type D struct {
	*A
	*B
	//基本数据类型
	int
}

type E struct {
	a int
	b string
	c A //组合模式

}

func main() {
	//注意事项二
	//一个结构体中嵌套了多个匿名结构体，就可以实现多重继承，尽量不使用多重继承
	//嵌入的匿名结构体中有相同属性与方法时，可以通过匿名结构体名进行区分
	c := C{A{10,"aaa"},B{12,"bbb",23},888}
	fmt.Println(c.b)
	fmt.Println(c.d)
	//fmt.Println(c.a)//出现了ambiguous selector错误，因为c.a是A的属性，B也有a属性，无法区分
	fmt.Println(c.A.a)
	fmt.Println(c.B.a)

	//注意事项三
	//结构体的匿名属性是基本数据类型
	fmt.Println(c.int)

	//嵌套结构体后，可以在创建实例后，直接指定各个匿名结构体属性的值
	cc := C{
		A{
			a : 10,
			b : "aaa"},
		B{
			c : 12,
			d : "bbb",
			a : 23},
		888}
	fmt.Println(cc)

	//嵌入匿名结构体的指针也可以
	c1 := D{&A{10,"aaa"},&B{12,"bbb",23},888}
	fmt.Println(c1)//会获取A和B的地址
	fmt.Println(*c1.A)
	fmt.Println(*c1.B)//因为A B都以指针形式嵌套，所以用实例的指针获取A，B的属性

	//结构体的属性可以是结构体
	e := E{10,"eee",A{11,"aaa"}}
	fmt.Println(e)//嵌套嵌套嵌套
	fmt.Println(e.c.a)
}