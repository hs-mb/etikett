package main

import (
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/hs-mb/label/webprint/views"
)

var StaticDir = "./static"
var ServeAddr = ":8080"

func main() {
	mux := http.NewServeMux()

	mux.Handle("GET /static/", http.StripPrefix("/static", http.FileServer(http.Dir(StaticDir))))

	mux.Handle("GET /{$}", templ.Handler(views.Index()))

	log.Printf("Listening on %s", ServeAddr)
	http.ListenAndServe(ServeAddr, mux)
}
