package main

import (
    "os"
    "log"
    "net/http"
    "github.com/malakhavam/medicineGoAPI/routes"

   
  )

// main function

func main() {
  router := routes.DrugRoutes()
  http.Handle("api", router)


    log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), router)) 
    
}