package models

type Person struct {
	Id          int64  `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	Phonenumber string `json:"phonenumber"`
	City        string `json:"city"`
	Address     string `json:"address"`
	Province    string `json:"province"`
}
