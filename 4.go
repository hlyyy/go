package main
import(
	"fmt"
	"net/http"
	"io/ioutil"
)
func main() {
	requestUrl :="http://search.zongheng.com/s?keyword=火影&pageNo=1&sort="
	rp,err :=http.Get(requestUrl)
	if err !=nil {
		fmt.Println("http.Get err=",err)
		panic(err)
	}
	body,err :=ioutil.ReadAll(rp.Body)
	if err !=nil {
		panic(err)
	}
	defer rp.Body.Close()
	fmt.Println(string(body))
}
 
