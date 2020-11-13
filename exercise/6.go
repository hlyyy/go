package main
import(
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
	"github.com/PuerkitoBio/goquery"
)
func main () {
	requestURL :="http://search.zongheng.com/s?keyword=火影&pageNo=1&sort="
	rp,err :=http.Get(requestURL)
	if err!=nil {
		fmt.Println(err)
	}
	body,err :=ioutil.ReadAll(rp.Body)
	if err!=nil {
		fmt.Println(err)
	}
	defer rp.Body.Close()
	content:=string(body)
	dom,err :=goquery.NewDocumentFromReader(strings.NewReader(content))
	if err !=nil {
		panic(err)
	}
	num:=1
	dom.Find(".search-tab").Each(func(i int, selection *goquery.Selection){
		selection.Find(".tit").Each(func(i int, title *goquery.Selection){
			fmt.Printf("%3d ",num)
			fmt.Println(title.Text())
			num++
		})
	})
}



