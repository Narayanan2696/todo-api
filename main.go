package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodGet {
			fmt.Println("Encountered GET Method")
		}
	})
	http.ListenAndServe("localhost:3000", mux)
}
