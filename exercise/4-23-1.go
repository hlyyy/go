package main
import "fmt"
func main() {
	var i,j,n int
	fmt.Println("生成直角在右上方的等腰直角三角形")
	fmt.Println("短边:")
	fmt.Scan(&n)
	for i=1;i<=n;i++ {
		for j=1;j<i;j++ {
			fmt.Printf(" ")
		}
		for j=0;j<=n-i;j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}
