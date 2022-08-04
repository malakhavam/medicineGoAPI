package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/malakhavam/medicineGoAPI/routes"
   
  )

// main function

func main() {
  r := mux.drugRoutes()
    log.Fatal(http.ListenAndServe(":3000", r)) 
    
}