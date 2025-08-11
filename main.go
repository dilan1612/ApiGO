package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/dilan1612/Carg_Go_Api/db"
	"github.com/dilan1612/Carg_Go_Api/models"
	"github.com/dilan1612/Carg_Go_Api/routes"
)

func main() {

	db.DBConnection()

	db.DB.AutoMigrate(models.User{})
	db.DB.AutoMigrate(models.Car{})

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUsersHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.DeleteUsersHandler).Methods("DELETE")

	http.ListenAndServe(":3000", r)

}
