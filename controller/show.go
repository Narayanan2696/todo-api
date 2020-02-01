package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"todo-api/model"

	"github.com/gorilla/mux"
)

func showOrDelete() http.HandlerFunc {
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
		} else if r.Method == http.MethodDelete {
			fmt.Println("Encountered DELETE Method inside controller")
			ids := mux.Vars(r)
			id, err := strconv.Atoi(ids["id"])
			if model.DeleteTask(id) != true {
				fmt.Println(err)
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(struct {
				Report string `json:"report"`
			}{"task deleted"})
		}
	}
}
