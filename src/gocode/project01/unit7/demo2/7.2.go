package main
import "fmt"
func main() {
	//数组初始化方式

	//方式一
	var arr1 [3]int = [3]int{1,2,3}
	fmt.Println(arr1)

	//方式二
	var arr2 = [3]int{1,4,7}
	fmt.Println(arr2)

	//方式三
	var arr3 = [...]int{1,4,7,11}//省略长度
	fmt.Println(arr3)

	//方式四
	var arr4 = [...]int{2:66,0:33,1:0}//指定索引值
	fmt.Println(arr4)

	//注意事项
	//1.数组的长度必须是常量或者变量，不能是表达式
	//数组的长度属于类型的一部分

	//数组默认下是值传递，值拷贝，所以修改数组元素的值不会影响到原数组
	arr5 := [3]int{1,2,3}
	test(arr5)
	fmt.Println(arr5)

	//要在其他函数中，改变原来数组，可以使用引用传递
	test2(&arr5)
	fmt.Println(arr5)
}

func test(arr [3]int)  {
	arr[0] = 7
}

func test2(arr *[3]int) {
	arr[0] = 7
}