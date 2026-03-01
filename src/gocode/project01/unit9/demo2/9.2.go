package main
import "fmt"
func main() {
	//增加与更新操作
	//map[key] = value

	m := map[int]string{
		1 : "张三",
		2 : "李四",
	}
	fmt.Println(m)

	//删除 delete函数
	delete(m, 1)
	fmt.Println(m)

	//清空,遍历key 逐个删除
	for k := range m {
		delete(m, k)
	}

	//或者make一个新map
	m = make(map[int]string)
	m[2] = "奥创"
	fmt.Println(m)

	//查找操作
	value,exist := m[2]
	fmt.Println(value,exist)
	
	//获取长度len()函数
	fmt.Println(len(m))
	
	for key,value := range m {
		fmt.Println(key,value)
	}//只有for range遍历，不能for遍历

	a := make(map[string]map[int]string)
	//嵌套map value为map[int]string
	a["法师"] = map[int]string{102:"甄姬",202:"高渐离"}
	for key,value := range a {
		fmt.Println(key)
		for key2,value2 := range value {
			fmt.Println(key2,value2)
		}
	}
}