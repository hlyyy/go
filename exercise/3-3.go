package main
import "fmt"
func main() {
	var a int
	fmt.Println("请输入一个整数:")
	fmt.Scan(&a)
	if a>=0 {
		fmt.Println("绝对值为",a)
	} else {
		fmt.Printf("绝对值为%d\n",-a)
	}
}
