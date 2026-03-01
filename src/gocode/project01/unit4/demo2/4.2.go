package main
import "fmt"
func main() {
	//关键字
	//break,continue
	var sum int = 0
	label1:
	for i := 0;i < 10;i++ {
		sum += i
		label2://标签的使用
		for j := 0;j < 10;j++ {
			sum += j
			if sum > 30 {
				fmt.Println("break")
				break label2
			}else {
				fmt.Println("continue")
				continue label1
			}
		}
	}

	for i := 0;i < 100;i++ {
        if i % 11 != 0 {
			continue//结束当前循环，直接进行下一次循环
		}
		fmt.Println(i)
	}

	//goto,不建议使用，

	//return,结束当前函数

}