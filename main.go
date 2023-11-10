package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "aj"
		}
				
		fmt.Fprintf(w, "Hello %s", name)
	})

	//remove 'localhost' when deploying to prod
	http.ListenAndServe("localhost:8080", nil)
}
 