package main
import "fmt"

type Humaner interface{
	sayhi()
}

type Student struct {
	name string
	age  int
}

func (tem *Student)sayhi() {
	fmt.Println("student",tem.name,tem.age)
}

type Teacher struct {
	addr string
	name string
}

func (tem *Teacher)sayhi() {
	fmt.Println("teacher",tem.addr,tem.name)
}

func whosayhi (i Humaner) {
	i.sayhi()
}

func main() {
	s :=&Student{"mike",18}
	whosayhi(s)
	a :=&Teacher{"bj","helen"}
	whosayhi(a)
}
