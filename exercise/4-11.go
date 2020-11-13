package main
import "fmt"
func main() {
	var no int
	fmt.Println("请输入一个正整数:")
	fmt.Scan(&no)
	fmt.Printf("%d逆向显示结果是",no)
	for ;no!=0;no/=10  {
		fmt.Printf("%d",no%10)
	}
	fmt.Printf("\n")
}
 


