package main
import "fmt"
func main() {
	var no int
	fmt.Println("正整数:")
	fmt.Scan(&no)
	if no>0 {
		for {
			fmt.Printf("+")
			no--
			if no==0 {
				break
			}
			fmt.Printf("-")
			no--
			if no==0 {
				break
			}
		}
	}
}
