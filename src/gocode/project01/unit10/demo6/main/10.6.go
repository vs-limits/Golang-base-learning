package main
import "fmt"
import "gocode/project01/unit10/demo6/model"
func main() {
	//封装的引入
	//封装：就是把抽象的字段和对字段的操作封装在一起，对外提供一个接口，隐藏内部的实现细节。
	
	//好处：
	//1.隐藏实现细节，提高代码的可读性和可维护性。
	//2.提高代码的复用性。

	//封装的实现：
	//将结构体，属性的首字母小写
	//给结构体所在包提供一个工厂模式的函数，首字母大写
	//提供一个首字母大写的Set方法，用对应属性并赋值
	//提供一个首字母大写的Get方法，返回对应属性的值

	//代码实现:
	p := model.NewPerson("Tom", 20)
	p.SetName("Jerry")
	fmt.Println(p.GetName())
	fmt.Println(*p)
	fmt.Println(p.Name)
	//都可以访问
	
}