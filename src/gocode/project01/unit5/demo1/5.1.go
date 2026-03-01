package main
import "fmt"
//自定义函数
func cal(a int, b int) (int,int) {//若返回值只有一个类型，可以省略括号
	return a + b,a * b
}//函数与函数之间为并列关系，可以直接调用
//没有返回值就可以不写返回类型
func swap(a int, b int) (int,int) {
	temp := a
	a = b
	b = temp
	return a,b
}
func test(args...int) {
	for i := 0;i < len(args);i++ {
		fmt.Println(args[i])
	}
}
func test2(num *int) {
	*num = 100 //通过指针修改值
}
func test3(num int) {
	fmt.Println(num)
}
//定义一个函数，把另一个函数作为参数传入
func test01(num1 int,num2 int, testFunc func(int) ) {
	fmt.Println("done")
}
type MyFunc func(int)//自定义函数类型
func test02(num1 int,num2 int, testFunc MyFunc ) {
	fmt.Println("done")
}
func test03(num1 int,num2 int) (sum int,sub int) {
	sum = num1 + num2
	sub = num1 - num2
	return//可以省略返回值
}
func main() {
	//函数引入
	var num1 int = 10
	var num2 int = 20
	sum,result := cal(num1,num2)
	//sum,_ := cal(num1,num2) //如果不需要第二个返回值，可以用_代替
	fmt.Println("result:", result,"sum:", sum)

	//内存分析
	num3,num4 := swap(num1,num2)//只能赋值新变量，未被定义的变量
	fmt.Println("num3:", num3,"num4:", num4)
	//函数不支持重载

	//func test(args...int) 可以传入任意个数的参数
	test(1,2,3,4,5,6)

	//通过指针修改值
	var num5 int = 10
	test2(&num5)
	fmt.Println("num5:", num5)

	//
	a := test3
	a(10)//通过该变量对函数进行调用
	

	//自定义类型
	type MyInt int
	var num6 MyInt = 10//虽然是自定义类型，但仍然是int类型，且编译识别时认为int 与 MyInt 不是一种类型
	fmt.Println(num6)

	test02(1,2,test3)

	p,m := test03(10,5)
	fmt.Println("p:", p,"m:", m)
}