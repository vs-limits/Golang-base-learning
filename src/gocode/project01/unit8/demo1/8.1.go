package main
import "fmt"
func main() {
	//切片  slice
	//声明一个切片，元素类型为int，长度为5
	//s := make([]int, 5)
	//切片是对数组一个连续片段的引用，可以方便的对数组进行操作，比如添加、删除元素、修改元素等。
	//方式一
	// var s []int = make([]int, 5)
	//方式二
	var intarr [5]int = [5]int{1, 2, 3, 4, 5}//数组
	var intslice []int = intarr[1:3]//从1开始到4结束，不包括4
	fmt.Println(intslice)
	fmt.Println(len(intslice))
	fmt.Println(cap(intslice))


	//内存分析
	//一个指向底层数组的指针(也就是开始的地址)，一个是切片的长度，一个是切片的容量
	intslice[1] = 16
	//修改切片也会修改原数组
	fmt.Println(intarr)

	//切片的定义
	//方式一
	//定义一个切片，让切片引用一个已经创建好的数组
	//上述方法就是方式一
	//该方式容量与引用数组的容量相等

	//方式二
	//通过make内置函数创建切片
	var slice2 = make([]int,5,16)
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))
	slice2[0] = 26
	fmt.Println(slice2)
	//make底层创建一个数组，不可以直接操作这个数组，要通过slice间接的访问各个元素
    //该方式自定义容量
	
	//方式三
	//定义一个切片，直接就指定具体数组
	slice3 := []int{1,4,7,8}
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
	//该方式的容量与长度相同
}