package main
import(
	"net/http"
	"io/ioutil"
	"log"
)

func main() {
	http.HandleFunc("/zhuzhu",func(w http.ResponseWriter,r *http.Request) {
		content,_ :=ioutil.ReadFile("./zhuzhu.html")
		w.Write(content)
	})
	log.Fatal(http.ListenAndServe(":8000",nil))
}
