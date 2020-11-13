package main
import "fmt"
func cube(a int)(b int) {
	b=a*a*a
	return 
}
func main() {
	var a int
	fmt.Println("请输入一个整数:")
	fmt.Scan(&a)
	fmt.Println("这个整数的三次方为",cube(a))
}

