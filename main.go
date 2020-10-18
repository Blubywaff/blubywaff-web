package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpls *template.Template

func init() {
	tpls = template.Must(template.New("").ParseGlob("pages/*.gohtml"))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/marblegame/", func(w http.ResponseWriter, req *http.Request) {
		_ = tpls.ExecuteTemplate(w, "marblegame.gohtml", nil)
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		_, _ = fmt.Fprint(w, "Blubywaff's Test Platform")
	})

	handlerfunc := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		mux.ServeHTTP(w, req)
	})

	log.Fatal(http.ListenAndServeTLS(":443", "TLS/cert.pem", "TLS/privkey.pem", handlerfunc))
}
