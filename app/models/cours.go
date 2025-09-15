package models

import (
	"fmt"
)

type ListeCours struct {
	Mnemonique string `json:"mnemonique"`
	Intitule   string `json:"intitule"`
	Credit     int    `json:"credit"`
	Titulaire  string `json:"titulaire"`
	//ListeNotes `xorm:"extends"`
}

func GetCours() ([]ListeCours, error) {

	cours := []ListeCours{}
	err := db.
		Join("", "liste_notes", "liste_notes.mnemonique = liste_cours.mnemonique").
		Select(" liste_cours.*, liste_notes.*").
		Find(&cours)
	fmt.Println("err", err)
	fmt.Println("cours", cours)

	if err != nil {
		return nil, err
	}
	return cours, nil
}
