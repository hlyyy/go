package main 
import "fmt"
func main() {
	var a,b,c int
	fmt.Println("请输入三个整数:")
	fmt.Scan(&a,&b,&c)
	if a==b && b==c {
		fmt.Println("三个值都相等")
	} else if a!=b && b!=c && a!=c {
		fmt.Println("三个值各不相同")
	} else {
		fmt.Println("有两个值相等")
	}
}

