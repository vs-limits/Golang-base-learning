package main
import "fmt"
import "unsafe"
func main() {
	var a int8 = 23  //-128~127
	//查看字节数
	fmt.Println(unsafe.Sizeof(a))
	//定义浮点型变量
	var b float32 = 3.14
	var c float32 = -3.14
	var d float32 = 314e-2
	var e float32 = 314e+2
	fmt.Println(b, c, d, e)
	//定义字符型变量
	var ch1 byte = 'a'
	var ch2 byte = '6'
	var ch3 byte = ')'
	fmt.Println(ch1, ch2, ch3)
	//按照ASCII码进行存储
	var ch4 int = '中'
	//中文的Unicode码更大，所以需要用int类型存储
	fmt.Println(ch4)

	//想显示对应的字符，必须使用格式化输出
	fmt.Printf("%c\n%c\n%c\n%c\n", ch1, ch2, ch3, ch4)

	//定义布尔型变量
	var flag1 bool = true
	fmt.Println(flag1)

	var flag2 bool = 5 < 9
	fmt.Println(flag2)

	//定义字符串型变量
	var str1 string = "hello world"
	fmt.Println(str1)

	//字符串不可变,字符串定义好后，其中字符的值不能改变，like：str2[0] = 'H'  //错误
	var str2 string = "hello"
	str2 = "world"
	fmt.Println(str2)

	//字符串的表示形式
	//包含特殊字符时，要用反引号
	var str3 string = `hello
	world`
	fmt.Println(str3)
	//字符串的拼接
	var str4 string = "hello " + "world"
	fmt.Println(str4)
	str4 += "!"
	fmt.Println(str4)
	//字符串过长时，需要把 + 留到末尾
	
}