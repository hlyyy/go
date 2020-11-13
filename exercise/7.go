package main
import(
	"fmt"
	"net/http"
	"io/ioutil"
)
func main() {
	rp,err :=http.Get("http://www.baidu.com")
	if err!=nil {
		fmt.Println(err)
	}
	body,err :=ioutil.ReadAll(rp.Body)
	if err!=nil {
		fmt.Println(err)
	}
	defer rp.Body.Close()
	content :=string(body)
	fmt.Println(content)
}
