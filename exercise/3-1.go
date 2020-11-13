package main
import "fmt"
func main() {
	var a,b int
	fmt.Println("请输入两个整数:")
	fmt.Scan(&a,&b)
	if a%b==0 {
		fmt.Println("b是a的约数")
	} else { 
		fmt.Println("b不是a的约数")
	}
}
