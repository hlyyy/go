package main
import "fmt"
func main() {
	var n,i,j int
	fmt.Println("生成一个正方形 正方形有几层:")
	fmt.Scan(&n)
	for i=0;i<n;i++ {
		for j=0;j<n;j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}

