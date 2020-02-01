package model

import (
	"log"
)

func DeleteTask(id int) bool {
	deleteQ, err := con.Query("DELETE FROM TASKLIST WHERE ID=?", id)
	// fmt.Println(deleteQ)
	defer deleteQ.Close()
	if err != nil {
		log.Fatal(err.Error)
		return false
	}
	return true
}
