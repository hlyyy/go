:package main
import "fmt"

type student struct {
	name string
	age  int
	addr string
}

func main() {
	s1 :=student{"mike",16,"bj"}
	fmt.Println("s1=",s1)
	s2 :=student{name:"helen"}
	fmt.Println("s2=",s2)
}


