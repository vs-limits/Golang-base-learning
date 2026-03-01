package main
import "fmt"
import "strconv"
func main() {
	//基本数据类型默认值
	var a int
	var b float32
	var c byte
	var d bool
	var e string
	fmt.Println(a, b, c, d, e)
	//int(0),float32/64(0),bool(false),string(空串)

	//基本数据类型转换(强制转换)
	var a1 int = 11
	fmt.Println(float32(a1)) //11.0
	var a2 float32 = 3.14
	fmt.Println(int(a2)) //3

	//将int64转换为int8，编译不出错，但会产生溢出
	var a3 int64 = 888888
	var a4 int8 = int8(a3)
	fmt.Println(a4)//56
	
	//基本数据类型转为string
	var n1 int = 100
	var n2 float32 = 3.14
	var n3 bool = true
	var n4 byte = 'A'

	//方式一 使用Sprintf函数将基本数据类型转为string
	var s1 string = fmt.Sprintf("n1:%d",n1)
	fmt.Println(s1) //n1:100
	var s2 string = fmt.Sprintf("n2:%f",n2)
	fmt.Println(s2) //n2:3.140000
	var s3 string = fmt.Sprintf("%t",n3)
	fmt.Printf("n3:%q\n",s3) //n3:true
	var s4 string = fmt.Sprintf("n4:%q",n4)
	fmt.Println(s4) //n4:A

	//方式二 
	var s5 string = strconv.FormatInt(int64(n1),10) //将int转为string
	fmt.Println(s5)

	//将string转为基本数据类型
	var str1 string = "true"
	var bb bool
	bb, _ = strconv.ParseBool(str1)
	// 注意：ParseBool返回两个值，第一个是转换后的值，第二个是错误信息
	// 如果转换失败，第二个值会包含错误信息
	//而err可以用_忽略
	fmt.Printf("bb:%t\n", bb) //bb:true

	var str2 string = "123"
	var n5 int64
	n5, _ = strconv.ParseInt(str2,10,64)
	fmt.Printf("n5:%d\n", n5) //n5:123
	//若str2不是合法的数字字符串，则会返回错误信息，输出默认值
}