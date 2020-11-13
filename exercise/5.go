package main
import(
	"fmt"
	"net/http"
)
func main() {
	a,b:=http.Get("https://www.baidu.com")
	fmt.Printf("%v\n%v\n",a,b)
}
