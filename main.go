package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"todo-api/structs"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {

		data := structs.Response{
			Code: http.StatusOK,
			Body: "pong",
		}
		json.NewEncoder(w).Encode(data)
		if r.Method == http.MethodGet {
			fmt.Println("Encountered GET Method")
		}
	})
	http.ListenAndServe("localhost:3000", mux)
}
