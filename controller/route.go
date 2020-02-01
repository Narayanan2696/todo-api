package controller

import (
	"github.com/gorilla/mux"
)

// func Register() *http.ServeMux {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/ping", ping())
// 	mux.HandleFunc("/", crud())
// 	mux.HandleFunc("/{id}", show())
// 	return mux
// }

func Register() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/ping", ping())
	router.HandleFunc("/", crud())
	router.HandleFunc("/{id}", showOrDelete())
	return router
}
