package main
import "fmt"

type stu struct {
	Name string
	Age  int
}

func main() {
	//创建构造体实例指定字段值
	//方式一 按照顺序赋值操作
	var s1 stu = stu{"小李",20}//根据结构体定义的顺序进行赋值
	fmt.Println(s1.Name, s1.Age)
	//方式二 按照指定类型
	var s2 stu = stu {
		Name:"alone",
		Age:20,
	}//直接指定赋值
	fmt.Println(s2)
	//方式三 想要返回结构体的指针类型
	var s3 *stu = &stu{"oh god",24}
	fmt.Println(*s3)
}