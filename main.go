package main

import (
    "log"
    "net/http"
    "math/rand"
    "strconv"
    "encoding/json"
    "github.com/gorilla/mux"
  )

 // drug structure

  type Drug struct {
    ID      string `json:"id"`
    Name   string `json:"name"`
    Class *Class `json:"class"`
  }
  
  type Class struct {
    Form string `json:"form"`
    Dosage  string `json:"dosage"`
  }

  var drugs []Drug

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

// main function

func main() {
    r := mux.NewRouter()
    drugs = append(drugs, Drug{ID: "1", Name: "Zyrtec", Class: &Class{Form: "Tablet", Dosage: "10mg"}})
    drugs = append(drugs, Drug{ID: "2", Name: "Cortisone", Class: &Class{Form: "Injections", Dosage: "2ml"}})
	drugs = append(drugs, Drug{ID: "3", Name: "Omegawel", Class: &Class{Form: "Softgelss", Dosage: "2g"}})
    r.HandleFunc("/drugs", getDrugs).Methods("GET")
    r.HandleFunc("/drugs/{id}", getDrug).Methods("GET")
    r.HandleFunc("/drugs", createDrug).Methods("POST")
    r.HandleFunc("/drugs/{id}", updateDrug).Methods("PUT")
    r.HandleFunc("/drugs/{id}", deleteDrug).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":3000", r))
}