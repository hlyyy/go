package main
import "fmt"
func main() {
	var no,a int
	fmt.Println("请输入一个正整数:")
	fmt.Scan(&no)
	for ;no!=0;no/=10 {
		a++
	}
	fmt.Printf("这个数的位数是%d\n",a)
}
