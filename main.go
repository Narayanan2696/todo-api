package main

import (
	"net/http"
	"todo-api/controller"
)

func main() {
	// mux := http.NewServeMux()
	// mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {

	// 	data := structs.Response{
	// 		Code: http.StatusOK,
	// 		Body: "pong",
	// 	}
	// 	json.NewEncoder(w).Encode(data)
	// 	if r.Method == http.MethodGet {
	// 		fmt.Println("Encountered GET Method")
	// 	}
	// })
	mux := controller.Register()
	http.ListenAndServe("localhost:3000", mux)
}
