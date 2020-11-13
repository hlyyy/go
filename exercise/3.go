package main
import "fmt"
func main() {
	a :=[]int{1,2,3,4,5,6,7,8,9}
	b :=a[2:5]
	fmt.Println("b=",b)
	b[1]=666
	fmt.Println("a=",a)
	c :=b[2:7]
	fmt.Println("c=",c)
	c[2]=777
	fmt.Println("a=",a)
}

