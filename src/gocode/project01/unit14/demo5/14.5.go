//通过反射操作结构体的属性与方法
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

	//通过reflect.Value类型操作结构体的属性
	n1 := val.NumField()//得到结构体的属性数量
	//遍历-获取具体的属性
	for i := 0; i < n1; i++ {
		fmt.Printf("第%d个属性的值是：%v\n",i,val.Field(i))
		//通过Field()方法得到属性的值，通过Set()方法设置属性的值
	}

	//通过reflect.Value类型操作结构体的方法
	n2 := val.NumMethod()//得到结构体的方法数量
	fmt.Println("结构体的方法数量：",n2)	
	//遍历-获取具体的方法
	val.Method(1).Call(nil)//调用第3个方法，并传入nil参数
	//方法的排序按照首字母的ASCLL码排序
	//首字母要大写，不然无法调用

	//定义一个value的切片
	//因为Call函数需要value切片类型的参数
	var params []reflect.Value
	params = append(params,reflect.ValueOf(10))
	params = append(params,reflect.ValueOf(20))
	//调用第1个方法，并传入参数
	result := val.Method(0).Call(params)//返回一个value类型切片，包含方法的返回值
	fmt.Println("方法返回值为:",result[0].Int())//value类型转成int类型

}
func main() {
	//4.NumField()：获取结构体的字段数量
	//5.Field()：获取结构体的字段
	//6.Method()：获取结构体的方法
	//7.Call()：调用结构体的方法

	//对结构体类型反射
	stu := Student{
		Name:"joker",
		Age:20,
	}
	testStruct(stu)//传入函数，转成空接口

	
}