package main

import (
	"fmt"
	"net/http"
	"todo-api/controller"
	"todo-api/model"

	_ "github.com/go-sql-driver/mysql"
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
	db := model.Connect()
	defer db.Close() // defer is used to execute the statement end of the scope here last line of main()
	fmt.Println("Connected to MySql DB")
	http.ListenAndServe("localhost:3000", mux)
}
