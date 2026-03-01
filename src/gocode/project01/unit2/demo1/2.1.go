package main
import "fmt"
//全局变量的统一声明
var (
	n6 = 500
	n7 = "King"
)

func main() {
	var a int = 10
	fmt.Println("a =",a)
	var b int //声明但是不赋值，使用默认值0
	fmt.Println("b =",b)
	//自动类型推断
	var c = 10.2
	var d = "King"
	fmt.Println(c,"+",d)
	//省略var声明，将”=“写为”:=“
	e := 20
	fmt.Println("e =",e)
    //声明多个变量
	var n1,n2,n3 int
	fmt.Println(n1,"+",n2,"+",n3)
	//一次性赋值，且可以是不同类型
	var n4,name,n5 = 10,"Joker",20.2
	fmt.Println(n4,"+",name,"+",n5)

	fmt.Println(n6,"+",n7)
}