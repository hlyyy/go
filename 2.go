package  main
import "fmt"
func main() {
	a:=[5]int {1,2,3,4,5}
	s:=a[1:4:5]
	fmt.Println("s=",s)
	fmt.Println("len(s)=",len(s))
	fmt.Println("cap(s)=",cap(s))
}
