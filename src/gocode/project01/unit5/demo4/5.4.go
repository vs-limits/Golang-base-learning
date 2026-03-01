package main
import "fmt"
import "strconv"
import "strings"
func main() {
	//字符串函数
	//len() - 返回字符串的长度
	str := "Hello World"
	fmt.Println(len(str))
	//字符串遍历
	//方式一：for range
	// for i, v := range str {
	// 	fmt.Printf("%d:%c\n", i, v)
	// }
	//方式二：r := []rune(str)
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Println(r[i])
	}
	//字符串转整数，整数转字符串
	strint,_ := strconv.Atoi("123")//返回整数
	intstr:= strconv.Itoa(123)//返回字符串
	fmt.Println(strint,intstr)

	//查找与统计指定子串

	count := strings.Count(str,"l")
	fmt.Println(count)

	//不区分大小写的字符串比较
	result := strings.EqualFold("hello","HELLO")//返回true
	fmt.Println(result)

	//区分大小写的字符串
	fmt.Println("hello" == "HELLO")

	//返回子串在字符串第一次出现的索引值，不存在返回-1
	index := strings.Index(str,"W")
	fmt.Println(index)

	//字符串替换
	//strings.Replace(str,old,new,n)
	str1 := strings.Replace(str,"l","L",-1)//-1表示全部替换,替换几个n就为几
	fmt.Println(str1)

	//字符串拆分
	//strings.Split(str,sep)
	arr := strings.Split(str,"l")
	fmt.Println(arr)

	//字符串大小写替换
	fmt.Println(strings.ToLower(str))//全部小写
	fmt.Println(strings.ToUpper(str))//全部大写

	//字符串两边删除指定字符
	fmt.Println(strings.Trim(str,"H"))//删除左边l
	fmt.Println(strings.TrimRight(str,"d"))//删除右边l
	fmt.Println(strings.TrimSpace(" 6 "))//删除两边空格
	fmt.Println(strings.Trim("~6~","~"))

}