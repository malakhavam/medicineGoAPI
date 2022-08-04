package routes

import (
	
	"net/http"
	"github.com/malakhavam/medicineGoAPI/controller"
	"github.com/gorilla/mux"
)

func DrugRoutes () *mux.Router {
	var router = mux.NewRouter()
	router = mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/",func(rw http.ResponseWriter, r *http.Request) {
	

	})

	// routes
	router.HandleFunc("/api/drugs", controller.GetDrugs).Methods("GET")
    router.HandleFunc("/api/drugs/{id}", controller.GetDrug).Methods("GET")
    router.HandleFunc("/api/drugs", controller.CreateDrug).Methods("POST")
    router.HandleFunc("/api/drugs/{id}", controller.UpdateDrug).Methods("PUT")
    router.HandleFunc("/api/drugs/{id}", controller.DeleteDrug).Methods("DELETE")

	return router
}