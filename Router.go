package main

import (
	"fmt"
	"net/http"
	"strings"
)

func count(reqq string) {
	resp, err := http.PostForm("http://172.20.10.6:8080/count", url.Values{"abc" : {reqq}})
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
	}
	return bodyString
}

func makeData(reqq string, s string) {
	defer wg.Done()
	resp, err := http.PostForm("http://172.20.10.6:8080/make", url.Values{"abc": {reqq}, "ans": {s}})
	if err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()

}

func askData(reqq string) string{
	resp, err := http.PostForm("http://172.20.10.3:8080/ask", url.Values{"abc" : {reqq}})
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
	}
	return bodyString
}


func runServer(addr string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request){
			a := r.FormValue("a")
			b := r.FormValue("b")
			c := r.FormValue("c")
			reqq := a + " " + b + " " + c
			s := askData(reqq)
			if s == "false" {
				s = count(reqq)
				makeData(reqq, s)
			}
			fmt.Fprintln(w, s)
			})

	server := http.Server{
		Addr:    addr,
		Handler: mux,
	}
	server.ListenAndServe()
}

func main() {
	runServer(":8080")
}
