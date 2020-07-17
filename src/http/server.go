package http
import (
	"net/http"
	"strconv"
)

var (
	host string = "127.0.0.1"
	port int = 9090

)

func Init() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		params := get(r, "test")
		w.Write([]byte(params))
	})

	http.ListenAndServe(host+":"+strconv.Itoa(port), nil)
}