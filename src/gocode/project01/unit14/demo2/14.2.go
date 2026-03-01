package main
import "fmt"
import "reflect"

//利用一个函数，函数的参数定义为空接口
func testReflect(arg interface{}) {//空接口没有任何方法
	//先调用TypeOf()函数获取arg的类型
	t := reflect.TypeOf(arg)
	//再调用ValueOf()函数获取arg的实际值
	v := reflect.ValueOf(arg)
	fmt.Println("Type:", t)
	fmt.Println("Value:", v)//v是value类型，无法与int类型相加

	//把v转成空接口
	i2 := v.Interface()
	//类型断言
	n,flag := i2.(Student)//不一定能成功
	if flag == true {
		fmt.Println("Name:", n.Name)
		fmt.Println("Age:", n.Age)
	}
}

//定义一个结构体
type Student struct {
	Name string
	Age  int
}


func main() {
	//反射

	//反射的引入
	//1.反射可以在运行时动态获取变量的各种信息，比如变量的类型，类别等信息
	//2.若为结构体变量，还可以获取结构体本身的属性与方法
	//3.可以通过反射修改变量的值，可以调用相关的方法
	//4.需要import “reflect”包

	//反射的相关函数
	//1.TypeOf()：获取变量的类型
	//2.ValueOf()：获取变量的实际值
	//前两种需要的参数是空接口

	//3.Kind()：获取变量的类别
	//4.NumField()：获取结构体的字段数量
	//5.Field()：获取结构体的字段
	//6.Method()：获取结构体的方法
	//7.Call()：调用结构体的方法

	//对结构体类型反射
	stu := Student{
		Name:"joker",
		Age:20,
	}
	testReflect(stu)
}