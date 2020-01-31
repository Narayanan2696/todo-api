package model

import (
	"log"
	"todo-api/views"
)

func ReadAll() ([]views.PostRequest, error) {
	rows, err := con.Query("SELECT * FROM TASKLIST")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	tasks := []views.PostRequest{}
	for rows.Next() {
		data := views.PostRequest{}
		rows.Scan(&data.Name, &data.Task)
		tasks = append(tasks, data)
	}
	return tasks, nil
}
