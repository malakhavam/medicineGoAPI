package models

// Drug Structure

type Drug struct {
    ID      string `json:"id"`
    Name   string `json:"name"`
    Class *Class `json:"class"`
  }
  
  type Class struct {
    Form string `json:"form"`
    Dosage  string `json:"dosage"`
  }
