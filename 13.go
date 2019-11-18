package main
import "fmt"

type Person struct {
	name string
	age  int
	sex  byte
}

func (tem *Person)PrintInfo() {
	fmt.Printf("name=%s,age=%d,sex=%c\n",tem.name,tem.age,tem.sex)
}

type Student struct {
	Person
	id int
	addr string
}

func main() {
	s :=Student{Person{"mike",18,'m'},666,"bj"}
	s.PrintInfo()
}
