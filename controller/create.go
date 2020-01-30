package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"todo-api/model"
	"todo-api/views"
)

func create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			fmt.Println("Encountered POST Method inside controller")
			data := views.PostRequest{}
			json.NewDecoder(r.Body).Decode(&data)
			fmt.Println(data)
			if err := model.CreateTaskList(data.Name, data.Task); err != nil {
				w.Write([]byte("error occured"))
				return
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(data)
		}

	}
}
