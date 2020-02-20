package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Request struct {
	a int
	b int
	c int
}

func runServer(addr string) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	mux := http.NewServeMux()
	mux.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				tmpl.Execute(w, nil)
				return
			}
		})

	server := http.Server{
		Addr:    addr,
		Handler: mux,
	}

	fmt.Println("starting server at", addr)
	server.ListenAndServe()
}

func main() {
	runServer(":8081")
}
