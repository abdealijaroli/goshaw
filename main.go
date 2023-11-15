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
		if name == "" {
			http.Error(w, "hello, ...?", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "hello %s", name)
	})

	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		if name == "" {
			http.Error(w, "goodbye, ...?", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "goodbye %s", name)
	})

	//remove 'localhost' when deploying to prod
	http.ListenAndServe("localhost:8080", nil)
}
