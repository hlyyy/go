package main
import "fmt"
func min3(a,b,c int) (min int) {
	min=a
	if b<min {
		min=b
	}
	if c<min {
		min=c
	}
	return 
}
func main() {
	var a,b,c,min int
	fmt.Println("请输入三个整数:")
	fmt.Scan(&a,&b,&c)
	min=min3(a,b,c)
	fmt.Println("min=",min)
} 

