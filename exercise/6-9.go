package main
import "fmt"

func main() {
	var a [10000]int
	var n int
	fmt.Println("请输入数组中元素的个数:")
	fmt.Scan(&n)
	for i:=0;i<n;i++ {
		fmt.Printf("a[%d]=",i)
		fmt.Scan(&a[i])
	}
	for i:=0;i<n;i++ {
		a[n-i-1],a[i]=a[i],a[n-i-1]
	}
	fmt.Println("反转之后的数组为:")
	for i:=0;i<n;i++ {
		fmt.Printf("a[%d]=%d\n",i,a[i])
	}
}

