package main
import "fmt"

func min_of(a [10000]int,n int)(min int) {
	min=a[0]
	for i:=1;i<n;i++ {
		if a[i]<min {
			min=a[i]
		}
	}
	return
}

func main() {
	var n,min int
	var s [10000]int
	fmt.Println("请输入数组中元素的个数:")
	fmt.Scan(&n)
	for i:=0;i<n;i++ {
		fmt.Printf("s[%d]=",i)
		fmt.Scan(&s[i])
	}
	min=min_of(s,n)
	fmt.Println("数组中的最小值为",min)
}


