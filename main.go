package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/malakhavam/medicineGoApi/models"
    "github.com/malakhavam/medicineGoApi/controller"
  )

// main function

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/drugs", getDrugs).Methods("GET")
    r.HandleFunc("/drugs/{id}", getDrug).Methods("GET")
    r.HandleFunc("/drugs", createDrug).Methods("POST")
    r.HandleFunc("/drugs/{id}", updateDrug).Methods("PUT")
    r.HandleFunc("/drugs/{id}", deleteDrug).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":3000", r))
}