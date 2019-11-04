package main
import "fmt"
func main() {
	var n,a int
	fmt.Println("请输入n的值:")
	fmt.Scan(&n)
	for a=1;a<=n;a++ {
		fmt.Printf("%d的二次方是%d\n",a,a*a)
	}
}
