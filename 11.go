package main
import(
	"net/http"
	"log"
	"io/ioutil"
)
func main() {
	http.HandleFunc("/hello_world",func(w http.ResponseWriter,r *http.Request) {
		content,_:=ioutil.ReadFile("./hello_world.html")
		w.Write(content)
	})
	log.Fatal(http.ListenAndServe(":8000",nil))
}

	
