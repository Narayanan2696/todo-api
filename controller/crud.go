package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"todo-api/model"
	"todo-api/views"
)

func crud() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			fmt.Println("Encountered POST Method inside controller")
			data := views.PostRequest{}
			json.NewDecoder(r.Body).Decode(&data)
			fmt.Println(data)
			err := model.CreateTaskList(data.Name, data.Task)
			if err != nil {
				fmt.Println("creation error")
				w.Write([]byte("error occured"))
				return
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(data)
		} else if r.Method == http.MethodGet {
			fmt.Println("Encountered GET Method inside controller")
			data, err := model.ReadAll()
			if err != nil {
				w.Write([]byte(err.Error()))
				return
			}
			fmt.Println(data)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		}

	}
}
