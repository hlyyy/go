package main
import "fmt"
func main() {
	var no int 
	fmt.Println("请输入一个正整数:")
	fmt.Scan(&no)
	for ;no>0;no-- {
		fmt.Printf("%d ",no)
	}
	fmt.Println("\n")
}
