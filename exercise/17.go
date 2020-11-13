package main
import(
	"errors"
	"fmt"
)

func Div(a,b float64)(result float64,err error) {
	err=nil
	if b==0 {
		err=errors.New("分母为0")
	} else {
		result=a/b
	}
	return
}

func main() {
	result,err :=Div(2,5)
	if err!=nil {
		fmt.Println(err)
	} else {
		fmt.Printf("result=%f\n",result)
	}
}
