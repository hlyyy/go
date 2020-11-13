package main
import "fmt"
func  sum1(n int)(sum int) {
	for i:=1;i<=n;i++ {
		sum+=i
	}
	return 
}
func main() {
	var n int
	fmt.Printf("请输入n的值:")
	fmt.Scan(&n)
	fmt.Printf("1到%d之间所有整数的和为%d\n",n,sum1(n))
}
