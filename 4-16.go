package main
import "fmt"
func main() {
	var no,a int
	fmt.Println("请输入一个整数:")
	fmt.Scan(&no)
	for a=1;a<=no;a=a+2 {
		fmt.Printf("%d ",a)
	}
}
		
