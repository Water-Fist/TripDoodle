package chat

import (
	"flag"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "http server address")

func main() {
	flag.Parse()

	http.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {
		ServeWs(w, r)
	})

	log.Fatal(http.ListenAndServe(*addr, nil))
}