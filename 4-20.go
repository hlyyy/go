package main
import "fmt"
func main() {
	var i,j int
	fmt.Printf("   ")
	for i=1;i<=9;i++ {
		fmt.Printf("%3d",i)
	}
	fmt.Printf("\n")
	for i=1;i<=9;i++ {
		fmt.Printf("%3d",i)
		for j=1;j<=9;j++ {
			fmt.Printf("%3d",i*j)
		}
		fmt.Printf("\n")
	}
}

