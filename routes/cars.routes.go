package routes

import (
	"encoding/json"
	"net/http"

	"github.com/dilan1612/Carg_Go_Api/db"
	"github.com/dilan1612/Carg_Go_Api/models"
	"github.com/gorilla/mux"
)

func GetCarsHandler(w http.ResponseWriter, r *http.Request) {

	var cars []models.Car
	db.DB.Find(&cars)
	json.NewEncoder(w).Encode(cars)

}

func CreateCarsHandler(w http.ResponseWriter, r *http.Request) {
	var car models.Car
	json.NewDecoder(r.Body).Decode(&car)
	createdCar := db.DB.Create(&car)
	err := createdCar.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(car)

}

func GetCarHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var car models.Car

	db.DB.First(&car, params["id"])

	if car.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Car not found"))
		return
	}
	json.NewEncoder(w).Encode(&car)

}

func DeleteCarsHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	var car models.Car

	db.DB.First(&car, params["id"])

	if car.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Car not found"))
		return
	}
	db.DB.Unscoped().Delete(&car)
	w.WriteHeader(http.StatusNoContent)

}
