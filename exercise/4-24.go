package main
import "fmt"
func main() {
	var i,j,n int
	fmt.Println("让我们来画一个金字塔")
	fmt.Println("金字塔有几层:")
	fmt.Scan(&n)
	for i=1;i<=n;i++ {
		for j=1;j<=n-i;j++ {
			fmt.Printf(" ")
		}
		for j=1;j<=2*i-1;j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}


