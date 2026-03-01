package main
import "fmt"
import "time"
func main() {
	//日期与时间函数，~时间戳
	now := time.Now()
	fmt.Println("Now:", now)//返回值是一个结构体，类型是time
	fmt.Println("Year:", now.Year())//年份
	fmt.Println("Month:", now.Month())//月份
	fmt.Println("Day:", now.Day())//日期
	fmt.Println("Hour:", now.Hour())//小时
	fmt.Println("Minute:", now.Minute())//分钟
	fmt.Println("Second:", now.Second())//秒

	datestr := fmt.Sprintf("%d/%02d/%02d %02dh/%02dm/%02ds", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println("Date:", datestr)
	datestr2 := now.Format("2006 15:04")
	fmt.Println("Date2:", datestr2)

	//内置函数
	//len()函数，返回字符串的长度
	str02 := "golang"
	fmt.Println(len(str02))
	//new()函数，分配内存
	num := new(int)//返回值是对应类型的指针,也就是int指针
	*num = 100
	fmt.Println(*num)
	//make()函数，创建切片、map、通道
}