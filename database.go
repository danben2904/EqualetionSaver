package main

import (
	"fmt"
	"net/http"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Request struct {
	abc string
	ans string
}

func runServer(addr string) {
	db, err := sql.Open("sqlite3", "Go.sqlite3")
	if err != nil {
        panic(err)
    }
    defer db.Close()
	mux := http.NewServeMux()
	mux.HandleFunc("/ask",
		func(w http.ResponseWriter, r *http.Request) {
			abc := r.FormValue("abc")
			row := db.QueryRow("select * from counted where abc = $1", abc)
			res := Request{}
			err := row.Scan(&res.abc, &res.ans)
			if err != nil {
			    fmt.Fprintln(w, "false")
			    return
			}
			fmt.Fprintln(w, res.ans)
		})
	mux.HandleFunc("/add",
		func(w http.ResponseWriter, r *http.Request) {
			abc := r.FormValue("abc")
			ans := r.FormValue("ans")
			row := db.QueryRow("select * from counted where abc = $1", abc)
			res := Request{}
			err := row.Scan(&res.abc, &res.ans)
			if err == nil {
			    return
			}
			_, err = db.Exec("insert into counted (abc, ans) values ($1, $2)", abc, ans)
			if err != nil {
		        panic(err)
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
	runServer(":8082")
}
