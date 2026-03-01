package main
import "fmt"
func main() {
	//流程控制
	//条件判断
	//1.if分支
	var a int = 68
	if a > 60 {
		fmt.Println(a)
	}else if a <= 60 {
		fmt.Println("不及格")
	}
	//if后面可以并列的加入变量的定义
	//用;进行并列
	if b := 20;b < 60 {
		fmt.Println(b)
	}
	//2.switch分支
	switch a {
	case 60://不能放判断式
		fmt.Println("及格")
	case 70:
		fmt.Println("良好")
	case 80:
		fmt.Println("优秀")
	default:
		fmt.Println("错误分数")
	}
	switch a / 10 {
	case 6:
		fmt.Println("及格")
	case 7:
		fmt.Println("良好")
	case 8:
		fmt.Println("优秀")
	}
	//switch相关细节
	//switch后面可以跟表达式
	//case的值需要与switch后的表达式类型一致
	//case可以跟多个值用逗号间隔开
	//case后不需要break
	//default分支可以省略，也可以写在任意位置
	//switch穿透，利用fallthrough关键字
	//fallthrough表示case后面的代码可以执行，而不用break跳出switch

	//for循环
	var sum int = 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	//灵活形式
	i := 1
	for i < 10 {
		sum += i
		i++
	}
	fmt.Println(sum)
	
	//for range循环
	//for range循环可以遍历数组、切片、map、字符串等
	//遍历数组或切片时，索引和值都可以获取
	
	//方式一
	var str string = "hello golang"
	for index,value := range str {
		fmt.Printf("%d %d %c\n",index,value,value)
	}//对字节进行遍历，不能输出汉字(占用多个字节)
	
	//方式二,一定要将下标与值一起使用，这个可以输出汉字
	for i , value := range str {
		fmt.Printf("%d %c\n", i, value) //对字符进行遍历
	}
}