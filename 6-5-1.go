package main
import "fmt"

func sum2(n int) int {
	if n==1 {
		return 1
	} else  {
		return n+sum2(n-1)
	}
}

func main() {
	var n int
	fmt.Printf("请输入一个整数:")
	fmt.Scan(&n)
	fmt.Printf("1到%d的和为%d\n",n,sum2(n))
}
