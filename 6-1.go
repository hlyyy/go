package main
import "fmt"

func jiaoxiao(a,b int)(min int) {
	if a>b {
		min=b
	} else {
		min=a
	}
	return 
}

func main() {
	var a,b,min int
	fmt.Println("请输入两个整数:")
	fmt.Scan(&a,&b)
	min=jiaoxiao(a,b)
	fmt.Println("min=",min)
}


