package main
import "fmt"
func main() {
	var n int
	fmt.Println("请输入月份:")
	fmt.Scan(&n)
	switch n {
	case 3,4,5 :fmt.Println("春季")
	case 6,7,8 :fmt.Println("夏季")
	case 9,10,11:fmt.Println("秋季")
	case 1,2,12 :fmt.Println("冬季")
	default :fmt.Println("该月不存在")
	}
}
