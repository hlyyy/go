package main
import "fmt"
func main() {
	var no,a int
	fmt.Println("请输入一个正整数:")
	fmt.Scan(&no)
	for a=1;a<=no;a++ {
		fmt.Printf("%d ",a)
	}
}

