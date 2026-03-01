package main
import "fmt"
func main() {
	//切片的遍历
	slice0 := []int{66,88,99}
	//1.for循环遍历
	for i := 0;i < len(slice0);i++ {
		fmt.Println(slice0[i])
		fmt.Printf("slice0[%v] = %v\n",i,slice0[i])
	}
	//2.for range 遍历
	fmt.Println("-------------------")
	for i,value := range slice0 {
		fmt.Printf("slice0[%v] = %v\n",i,value)
	}

	//注意事项1
	//切片定义后不能直接使用，需要让其引用一个数组，或者make一个空间供slice使用
	//切片不能越界
	//简写
	//var slice1 = arr[0:end] --> var slice1 = arr[:end]
	//简写
	//var slice2 = arr[start:len(arr)] --> var slice2 = arr[start:]
	//简写
	//var slice3 = arr[0:len(arr)] --> var slice3 = arr[:]
	//切片可以继续切片

	//注意事项2
	//切片可以动态增长，但是不能缩短，只能在末尾添加元素，或者使用append函数添加元素
	slice1 := append(slice0,55,22)
	slice0 = append(slice0,77)
	fmt.Println(slice1)
	fmt.Println(slice0)
	//底层追加元素，是对数组进行扩容，老数组变新数组
	//append函数返回一个新的切片，原切片不变
	//底层数组的指针指的是新数组
	//可以通过append函数将切片追加给另一个切片
	slice2 := append(slice1,slice0...)
	fmt.Println(slice2)

	//注意事项3
	var a []int = []int{2,5,8,11,9}
	var b []int = make([]int,3)
	//拷贝
	copy(b,a)//不论n有多大，都只拷贝a的前n个元素到b
	fmt.Println(b)

}