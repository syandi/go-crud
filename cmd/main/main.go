package main

import (
	"log"
	"net/http"

	"crud/pkg/routes"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", r))
}
