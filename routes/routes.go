package routes

import (
	"encoding/json"
	"net/http"
	"github.com/malakhavam/medicineGoAPI/models"
    "github.com/malakhavam/medicineGoAPI/controller"
	"github.com/gorilla/mux"
)

func drugRoutes () {
	r := mux.NewRouter()


	// routes
	r.HandleFunc("/drugs", getDrugs).Methods("GET")
    r.HandleFunc("/drugs/{id}", getDrug).Methods("GET")
    r.HandleFunc("/drugs", createDrug).Methods("POST")
    r.HandleFunc("/drugs/{id}", updateDrug).Methods("PUT")
    r.HandleFunc("/drugs/{id}", deleteDrug).Methods("DELETE")


}