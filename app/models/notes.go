package models

import (
	"fmt"
)

type ListeNotes struct {
	Id         int    `json:"id";xorm:"primaryKey;`
	Matricule  string `json:"matricule"`
	Mnemonique string `json:"mnemonique"`
	Note       int    `json:"note"`
}

func GetNotes() ([]ListeNotes, error) {

	var notes []ListeNotes
	err := db.Find(&notes)

	fmt.Println("notes", notes)
	if err != nil {
		return nil, err
	}
	return notes, nil
}
