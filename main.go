package main

import (
	"fmt"
	"net/http"
	"structs"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		Json_response := structs.

		if r.Method == http.MethodGet {
			fmt.Println("Encountered GET Method")
		}
	})
	http.ListenAndServe("localhost:3000", mux)
}
