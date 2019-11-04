package main
import "fmt"
func main() {
	var no,a,b int
	fmt.Println("请输入一个整数:")
	fmt.Scan(&no)
	a=no/10
	no%=10
	for ;a>0;a-- {
		fmt.Printf("1234567890")
	}
	for b=1;b<=no;b++ {
		fmt.Printf("%d",b)
	}
	fmt.Printf("\n")
}
