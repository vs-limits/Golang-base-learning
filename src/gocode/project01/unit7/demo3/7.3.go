package main
import "fmt"
func main() {
	//多维数组
	var arr [2][3][4]int
	fmt.Println(arr)

	fmt.Printf("%p\n",&arr)
	fmt.Printf("%p\n",&arr[0])
	fmt.Printf("%p\n",&arr[0][0])

	fmt.Printf("%p\n",&arr[1])
	fmt.Printf("%p\n",&arr[1][0])

	var arr2 [2][3]int = [2][3]int{{1,3,5},{2,4,6}}
	fmt.Println(arr2)

	//遍历二维数组
	//for循环
	for i := 0; i < len(arr2); i++ {
		for j := 0;j < len(arr2[i]);j++ {
			fmt.Println(arr2[i][j])
		}
	}

	//for range循环
	for index,arr01 := range arr2 {
		for _,value := range arr01 {
			fmt.Println(index,arr01,value)
		}
	}
}