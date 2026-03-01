package main
import "fmt"
import "C:\Users\limits\Desktop\Golang\unit2\demo4\test" //绝对路径
//import "unit2/demo4/test" //相对路径
func main() {
	//指针类型
	//定义指针变量
	var age int = 19
	var p *int = &age
	fmt.Println(p)
	fmt.Println(&p)
	fmt.Println(*p)

	//标识符
	//组成规则
	//字母，数字，下划线, 汉字
	//下划线开头代表忽略，不能用作标识符
	//belike: _"fmt",会忽略fmt库

	//若变量名，函数名，常量名首字母大写，则可以被其他包访问
	//若变量名，函数名，常量名首字母小写，则只能在当前包内访问
	fmt.Println(test.Age) //访问test包中的Age变量

	//关键字
	//break,default,func,interface,select
	//case,defer,go,map,struct
	//chan,else,goto,package,swutch
	//const,fallthrough,if,range,type
	//continue,for,import,return,var
}