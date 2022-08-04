package controller

import (
    
    "net/http"
    "math/rand"
    "strconv"
    "encoding/json"
    "github.com/gorilla/mux"
	"github.com/malakhavam/medicineGoAPI/models"
	"github.com/malakhavam/medicineGoAPI/db"
  )

  // function to get all drugs

  func getDrugs(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(drugs)
}
   // function to get drug

   func getDrug(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    params := mux.Vars(r)
    for _, item := range drugs {
       if item.ID == params["id"] {
          json.NewEncoder(w).Encode(item)
          return
        }
    }
    json.NewEncoder(w).Encode("Sorry, Drug with that id not found")
}

// function to create new drug

func createDrug(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var drug Drug
    _ = json.NewDecoder(r.Body).Decode(&drug)
    drug.ID = strconv.Itoa(rand.Intn(1000000))
    drugs = append(drugs, drug) 
    json.NewEncoder(w).Encode(drug)
}

// function to update drug

func updateDrug(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for index, item := range drugs {
        if item.ID == params["id"] {
            drugs = append(drugs[:index], drugs[index+1:]...)
            var drug Drug
            _ = json.NewDecoder(r.Body).Decode(&drug)
            drug.ID = params["id"]
            drugs = append(drugs, drug) 
            json.NewEncoder(w).Encode(drug)
            return
        }
    }
    json.NewEncoder(w).Encode(drugs)
}

// function to delete drug

func deleteDrug(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for index, item := range drugs {
        if item.ID == params["id"] {
            drugs = append(drugs[:index], drugs[index+1:]...)
            break
        }
    }
    json.NewEncoder(w).Encode(drugs)
}
