package main
import "fmt"
func main() {
	p :=new(int)
	*p=666
	fmt.Println("*p=",*p)
}
