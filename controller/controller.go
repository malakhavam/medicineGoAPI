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

  func GetDrugs(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    alldrugs := db.DrugList
    json.NewEncoder(w).Encode(alldrugs)
}
   // function to get one drug

   func GetDrug(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    params := mux.Vars(r)
    for _, item := range db.DrugList {
       if item.ID == params["id"] {
          json.NewEncoder(w).Encode(item)
          return
        }
    }
    json.NewEncoder(w).Encode("Sorry, Drug with that id not found")
}

// function to create new drug

func CreateDrug(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var drug models.Drug
    _ = json.NewDecoder(r.Body).Decode(&drug)
    drug.ID = strconv.Itoa(rand.Intn(1000000))
    db.DrugList = append(db.DrugList, drug) 
    json.NewEncoder(w).Encode(drug)
}

// function to update drug

func UpdateDrug(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for index, item := range db.DrugList {
        if item.ID == params["id"] {
            db.DrugList = append(db.DrugList[:index], db.DrugList[index+1:]...)
            var drug models.Drug
            _ = json.NewDecoder(r.Body).Decode(&drug)
            drug.ID = params["id"]
            db.DrugList = append(db.DrugList, drug) 
            json.NewEncoder(w).Encode(drug)
            return
        }
    }
    json.NewEncoder(w).Encode(db.DrugList)
}

// function to delete drug

func DeleteDrug(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for index, item := range db.DrugList {
        if item.ID == params["id"] {
            db.DrugList = append(db.DrugList[:index], db.DrugList[index+1:]...)
            break
        }
    }
    json.NewEncoder(w).Encode(db.DrugList)
}
