package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
)


type Request struct {
	a  int
	b  int
	c  int
}


func runServer(addr string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			abc := r.FormValue("abc")
			a, erra := strconv.Atoi(strings.Split(abc, " ")[0])
			b, errb := strconv.Atoi(strings.Split(abc, " ")[1])
			c, errc := strconv.Atoi(strings.Split(abc, " ")[2])
			var res string
			if erra == nil && errb == nil && errc == nil {
				reqq := Request{a: a, b: b, c: c}
				_ = reqq
				x1, x2, Flag := sq_ur(float64(a), float64(b), float64(c))
				if Flag == false {
					res = "No solution!"
				} else {
					res = fmt.Sprint(x1) + " " + fmt.Sprint(x2)
				}
				_, _ = fmt.Fprint(w, res)
			}
			fmt.Println("starting server at", addr)
		})
	server := http.Server{
		Addr:    addr,
		Handler: mux,
	}
	server.ListenAndServe()
}


func main() {
	runServer(":8081")
}


func sq_ur(a, b, c float64) (float64, float64, bool) {
	if b*b-4*a*c < 0 {
		return 0, 0, false
	} else {
		d := math.Sqrt(b*b - 4*a*c)
		x1 := (-b + d) / 2
		x2 := (-b - d) / 2
		return x1, x2, true
	}
}
