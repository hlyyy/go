package main
import(
	"fmt"
	"math/rand"
	"time"
)

func rev_intary(s []int,n int){
	for i:=0;i<(n-1)/2;i++ {
		s[i],s[n-1-i]=s[n-1-i],s[i]
	}
}

func main () {
	var n int
	fmt.Println("请输入元素的个数:")
	fmt.Scan(&n)
	rand.Seed(time.Now().UnixNano())
	a :=[]int{}
	for i:=0;i<n;i++ {
		a=append(a,rand.Intn(1000))
	}
	fmt.Println("a=",a)
	rev_intary(a,n)
	fmt.Println("a=",a)
}
