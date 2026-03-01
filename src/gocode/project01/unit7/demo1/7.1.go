package main
import "fmt"
func main() {
	//数组
	//定义数组
	var arr [5]int
	//初始化数组
	arr[0] = 89
	arr[1] = 97
	arr[2] = 79
	arr[3] = 85
	arr[4] = 94
	sum := 0
	for i := 0;i < len(arr);i++ {
		sum += arr[i]
	}
	avg := sum / len(arr)
	fmt.Println("数组元素的平均值是：", avg)

	//数组内存分析
	fmt.Println("数组的长度是：", len(arr))

	//数组遍历
	for i := 0;i < len(arr);i++ {
		fmt.Scanln(&arr[i])
	}
	for i := 0;i < len(arr);i++ {
		fmt.Println("数组元素", i, "是：", arr[i])
	}

	//方式二
	for key , value := range arr {
		fmt.Println("数组元素", key, "是：", value)
	}
	for _,value := range arr {
		fmt.Println("数组元素", value)
	}//省略索引

}