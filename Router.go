package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

func count(reqq string, wg *sync.WaitGroup) string {
	defer wg.Done()
	resp, err := http.PostForm("http://127.0.0.1:5000", url.Values{"abc" : {reqq}})
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyString := ""
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString = string(bodyBytes)
	}
	fmt.Println(bodyString)
	return bodyString
}


func makeData(reqq string, s string, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.PostForm("http://127.0.0.1:5050/make", url.Values{"abc": {reqq}, "ans": {s}})
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
}

func askData(reqq string, wg *sync.WaitGroup) string{
	defer wg.Done()
	resp, err := http.PostForm("http://127.0.0.1:5050/ask", url.Values{"abc" : {reqq}})
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
  	bodyString := ""
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString = string(bodyBytes)
	}
	return bodyString
}


func runServerRouter(addr string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request){
			var wg sync.WaitGroup
			wg.Add(1)
			a := r.FormValue("a")
			b := r.FormValue("b")
			c := r.FormValue("c")
			reqq := a + " " + b + " " + c
			s := askData(reqq, &wg)
			wg.Wait()
			if strings.TrimSpace(s) == "false" {
				wg.Add(1)
				new_s := count(reqq, &wg)
				fmt.Fprintln(w, new_s)
				wg.Wait()
				wg.Add(1)
				makeData(reqq, new_s, &wg)
				wg.Wait()
			} else {
				fmt.Fprintln(w, s)
			}
			})

	server := http.Server{
		Addr:    addr,
		Handler: mux,
	}
	server.ListenAndServe()
}

func main() {
	runServerRouter(":8081")
}
