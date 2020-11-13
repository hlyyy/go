package main
import "fmt"
func main() {
	var a,b int
	fmt.Println("请输入两个整数:")
	fmt.Scan(&a,&b)
	fmt.Printf("a的值是b的%d%%",a*100/b)
}
