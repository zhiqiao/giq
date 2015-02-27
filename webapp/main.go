package webapp


import (
	"net/http"
)


func init() {
	http.HandleFunc("/form", Form)
	http.HandleFunc("/sign", Sign)
	http.HandleFunc("/question", Question)
}
