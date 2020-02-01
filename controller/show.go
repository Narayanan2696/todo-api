package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"todo-api/model"

	"github.com/gorilla/mux"
)

func show() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			fmt.Println("Encountered GET Method inside controller")
			// ids := r.URL.Query().Get("id")
			ids := mux.Vars(r)
			// fmt.Printf("\nstring id:%s\n", ids["id"])
			id, err := strconv.Atoi(ids["id"])
			name := r.URL.Query().Get("name")
			// fmt.Printf("\nquery name:%s\n", name)
			data, err := model.Read(id, name)
			if err != nil {
				w.Write([]byte("error occured"))
				return
			}
			w.WriteHeader(http.StatusOK)
			// fmt.Println(data)
			json.NewEncoder(w).Encode(data)
		}
	}
}
