package models

import (
	"fmt"
)

type ListeCours struct {
	Mnemonique string `json:"mnemonique"`
	Intitule   string `json:"intitule"`
	Credit     int    `json:"credit"`
	Titulaire  string `json:"titulaire"`
}

func GetCours() ([]ListeCours, error) {

	cours := []ListeCours{}
	err := db.
		Select("*").
		Find(&cours)
	fmt.Println("err", err)
	fmt.Println("cours", cours)

	if err != nil {
		return nil, err
	}
	return cours, nil
}
