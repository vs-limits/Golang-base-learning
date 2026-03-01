package main
import "fmt"
import "gocode/project01/unit10/demo5/model"
func main() {
	//跨包创建结构体实例
	var s model.Student = Student{"alone",23}

	//工厂模式调用
	p := model.NewPerson("alone",25)
}