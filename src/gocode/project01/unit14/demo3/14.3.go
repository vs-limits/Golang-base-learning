package main
import "fmt"
import "reflect"

//利用一个函数，函数的参数定义为空接口
func testReflect(arg interface{}) {//空接口没有任何方法
	//先调用TypeOf()函数获取arg的类型
	t := reflect.TypeOf(arg)
	//再调用ValueOf()函数获取arg的实际值
	v := reflect.ValueOf(arg)

	//获取变量的类别
	k1 := t.Kind()
	fmt.Println("变量的类型：", k1)
	k2 := v.Kind()
	fmt.Println("变量的实际类型：", k2)
	//获取变量的类型
		//把v转成空接口
	i2 := v.Interface()
	//类型断言
	n,flag := i2.(Student)//不一定能成功
	if flag == true {
		fmt.Printf("结构体的类型是: %T",n)
		//等同于fmt.Println(t)
	}
}

//定义一个结构体
type Student struct {
	Name string
	Age  int
}


func main() {
	//反射的相关函数
	//3.Kind()：获取变量的类别
	//reflect.Type.Kind()
	//reflect.Value.Kind()

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