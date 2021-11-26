package routes

import (
	"crud/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterRoutes = func(router *mux.Router) {
	router.HandleFunc("/book", controllers.Index).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.Show).Methods("GET")
	router.HandleFunc("/book", controllers.Store).Methods("POST")
	router.HandleFunc("/book/{bookId}", controllers.Update).Methods("POST")
	router.HandleFunc("/book/{bookId}", controllers.Destroy).Methods("DELETE")
}
