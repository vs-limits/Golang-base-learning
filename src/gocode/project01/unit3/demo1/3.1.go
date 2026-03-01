package main
import "fmt"
func main() {
	//运算符
	//1.算术运算符，+，-，*，/，%，++，--
	//++，--只有后置，没有前置
	//% （取余）,等价公式: a%b = a - a/b * a

	//2.赋值运算符，=，+=，-=，*=，/=，%=，

	//3.关系运算符,==,!=,>,<,>=,<=
	//只会返回true或false

	//4.逻辑运算符，&&,||,!

	//5.位运算符

	//6.其他运算符
	//&:返回变量的储存地址
	//*：指针变量对应的数值
	var a = 10
	fmt.Println(&a)
	
	//获取用户终端输入
	//Scanln: 读取用户输入，直到遇到空格或换行符
	var name string
	var age int
	var score float32
	fmt.Scanln(&name,&age,&score)
	fmt.Println("name:",name,"age:",age,"score:",score)
	//Scanf: 读取用户输入，需要指定格式
	fmt.Scanf("%s %d %f",&name,&age,&score)
	fmt.Println("name:",name,"age:",age,"score:",score)
}