package main
import "fmt"
import "reflect"

//利用一个函数，函数的参数定义为空接口
func testReflect(arg interface{}) {//空接口没有任何方法
	v := reflect.ValueOf(arg)
	//通过SetInt()函数改变值
	v.Elem().SetInt(99)//通过Elem()函数获取结构体的指针，然后调用SetInt()函数改变值
}
func main() {
	var num = 100
	testReflect(&num)//需要传入指针地址，才可以更改
	//修改基本数据类型的值
	fmt.Println(num)
}