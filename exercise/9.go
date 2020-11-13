package main
import(
	"fmt"
	"net/http"
	"log"
)
func main() {
	http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
		fmt.Fprintf(w,"Url.path=%q/n",r.URL.Path)
	})
	log.Fatal(http.ListenAndServe(":8000",nil))
}
