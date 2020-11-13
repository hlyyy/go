package main
import "fmt"
func main() {
	var no int
	fmt.Println("正整数:")
	fmt.Scan(&no)
	for ;no>0;no-- {
		fmt.Println("*")
	}
}
