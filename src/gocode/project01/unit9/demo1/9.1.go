package main
import "fmt"
func main() {
	//映射，map，通过键值对相连
	//键值对：一对匹配的信息
	//学号(键 key)-姓名(值 value)

	//声明一个映射
	//key与value的类型可以是 bool string 数字 指针 channel
	//key通常为int与string类型，value通常为数字，string，map，结构体
	var m map[int]string
	//初始化映射
	m = make(map[int]string,10)//空映射,创造一个空间
	//make函数
	//make(Type,size) Type
	m[202400204042] = "小白"
	m[202400204043] = "小黑"
	//输出
	fmt.Println(m)
	//map集合使用前一定要make
	//map的key-value对是无序的，不能通过下标访问，只能通过键访问
	//value部分可以重复，key不能重复

	//make函数的第二个参数size，表示映射的容量，即可以存储多少个键值对，默认值为0，表示无限容量,可以省略

	//方式二
	b := make(map[int]string)
	b[202400204042] = "小白"
	b[202400204043] = "小黑"
	fmt.Println(b)

	//方式三
	c := map[int]string{
		202400204042 : "小白",
		202400204043 : "小黑",
	}
	fmt.Println(c)


}