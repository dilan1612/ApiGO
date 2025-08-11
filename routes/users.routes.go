package routes

import (
	"encoding/json"
	"net/http"

	"github.com/dilan1612/Carg_Go_Api/db"
	"github.com/dilan1612/Carg_Go_Api/models"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	db.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)

}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {

}

func PostUsersHandler(w http.ResponseWriter, r *http.Request) {

	var user models.User

	json.NewDecoder(r.Body).Decode(&user)

	createdUser := db.DB.Create(&user)
	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&user)
}

func DeleteUsersHandler(w http.ResponseWriter, r *http.Request) {
	// Handler logic to get users
	w.Write([]byte("delete"))
}
