package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"sync"
)

type Request struct {
	a  int
	b  int
	c  int
	id int
}

func requestToServer(reqq Request, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.PostForm("http://172.20.10.3:8080", url.Values{"a": {strconv.Itoa(reqq.a)}, "b": {strconv.Itoa(reqq.b)}, "c": {strconv.Itoa(reqq.c)}, "id": {strconv.Itoa(reqq.id)}})
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
	}
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
			var wg sync.WaitGroup
			a, erra := strconv.Atoi(r.FormValue("a"))
			b, errb := strconv.Atoi(r.FormValue("b"))
			c, errc := strconv.Atoi(r.FormValue("c"))
			if erra == nil && errb == nil && errc == nil {
				wg.Add(1)
				reqq := Request{a: a, b: b, c: c}
				_ = reqq
				go requestToServer(reqq, &wg)
			}
			wg.Wait()
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
