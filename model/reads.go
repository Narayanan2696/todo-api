package model

import (
	"database/sql"
	"log"
	"todo-api/views"
)

func ReadAll() ([]views.PostRequest, error) {
	rows, err := con.Query("SELECT name, task FROM TASKLIST")
	defer rows.Close()
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

func Read(id int, name ...string) ([]views.PostRequest, error) {
	if name[0] != "" {
		// fmt.Printf("arr len:%d\n", len(name))
		// fmt.Printf("name:%s\n", name)
		rows, err := con.Query("SELECT name, task FROM TASKLIST where id=? and name=?", id, name[0])
		defer rows.Close()
		// fmt.Printf("Type of row:%T\n", rows)
		return prepareTask(rows, err)
	} else {
		rows, err := con.Query("SELECT name, task FROM TASKLIST where id=?", id)
		defer rows.Close()
		return prepareTask(rows, err)
	}
}

func prepareTask(rows *sql.Rows, err error) ([]views.PostRequest, error) {
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
