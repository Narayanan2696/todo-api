package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"todo-api/structs"
)

func ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := structs.Response{
			Code: http.StatusOK,
			Body: "pinging pong",
		}
		json.NewEncoder(w).Encode(data)
		if r.Method == http.MethodGet {
			fmt.Println("Encountered GET Method inside controller")
		}
	}
}
