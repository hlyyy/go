package main
import "fmt"
func main() {
	var a,b int
	fmt.Println("请输入两个整数:")
	fmt.Scan(&a,&b)
	if a<b {
		a,b=b,a
	}
	if a-b<=10 {
		fmt.Println("它们的差小于等于10")
	} else {
		fmt.Println("它们的差大于10")
	}
}
