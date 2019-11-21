package main
import "fmt"

type Humaner interface {
	sayhi()
}

type Personer interface {
	Humaner
	sing(l string)
}

/*相当于
type Personer interface{
	sayhi()
	sing(l string)
}
*/

type Student struct {
	name string
	age  int
}

func (tem *Student) sayhi() {
	fmt.Println("sayhi",tem.name,tem.age)
}

func (tem *Student)sing(l string) {
	fmt.Println("sing:",l)
}

func main() {
	var i Personer
	a :=&Student{"mike",16}
	i=a
	i.sayhi()
	i.sing("you need to calm down")

	var Pro Personer=&Student{"yoyo",20}
	var pro Humaner
	pro=Pro
	pro.sayhi()
//超集可以转化成子集，反过来不可以,且只能使用子集中的方法


}






