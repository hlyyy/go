package main
import "fmt"
func main() {
	var a int 
	fmt.Println("请输入一个整数")
	fmt.Scan(&a)
	switch a%2 {
	case 0: fmt.Println("该整数是偶数")
	case 1,-1: fmt.Println("该整数是奇数")
	}
}
