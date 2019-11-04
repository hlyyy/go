package main
import "fmt"
func main() {
	var i,j,n int
	fmt.Println("生成直角在左上方的等腰直角三角形")
	fmt.Printf("短边:")
	fmt.Scan(&n)
	for i=0;i<n;i++ {
		for j=1;j<=n-i;j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}

