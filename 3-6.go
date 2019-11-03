package main
import "fmt"
func main() {
	var a,b,c,min int 
	fmt.Println("请输入三个整数")
	fmt.Scan(&a,&b,&c)
	min=a
	if min>b {
		min=b
	}
	if min>c {
		min=c
	}
	fmt.Println("这三个整数中最小值为",min)
}
