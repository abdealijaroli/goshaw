package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		fmt.Fprintf(w, "hello %s", name)
	})

	//remove 'localhost' when deploying to prod
	http.ListenAndServe("localhost:8080", nil)
}
 