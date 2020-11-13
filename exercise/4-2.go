package main
import "fmt"
func main() {
	var a,b,sum int
	fmt.Println("请输入两个整数")
	fmt.Scan(&a,&b)
	if  a<b {
		a,b=b,a
	}
	for sum=0;b<=a;b++ {
		sum+=b
	}
	fmt.Println("它们之间包括它们所有整数的和是",sum)
}
