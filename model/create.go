package model

import (
	"log"
)

func CreateTaskList(name, task string) error {
	insertQ, err := con.Query("INSERT INTO TASKLIST(NAME, TASK) VALUES(?,?)", name, task)
	defer insertQ.Close()
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
