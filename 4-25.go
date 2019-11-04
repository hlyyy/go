package main
import "fmt"
func main() {
	var i,j,n int
	fmt.Println("让我们来画一个向下的金字塔")
	fmt.Printf("金字塔有几层:")
	fmt.Scan(&n)
	for i=1;i<=n;i++ {
		for j=1;j<i;j++ {
			fmt.Printf(" ")
		}
		for j=1;j<=2*n-2*i+1;j++ {
			fmt.Printf("%d",i%10)
		}
		fmt.Printf("\n")
	}
}
