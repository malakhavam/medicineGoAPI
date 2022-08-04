package db

import "github.com/malakhavam/medicineGoApi/models"

var drugs = []models.Drug {
	{
		ID: "1", 
		Name: "Zyrtec", 
		Class:
			{
				Form: "Tablet", 
				Dosage: "10mg",
			},
	},

	{
		ID: "2", 
		Name: "Cortisone", 
		Class: 
			{
				Form: "Injections", 
				Dosage: "2ml",
			},
	},

	{
		ID: "3", 
		Name: "Omegawel", 
		Class: 
			{
				Form: "Softgels", 
				Dosage: "2g",
			},
	},
	
}
