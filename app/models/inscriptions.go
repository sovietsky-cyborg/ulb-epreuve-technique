package models

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

type ListeInscriptions struct {
	Matricule  string    `json:"matricule";xorm:"primaryKey;autoIncrement:false"`
	Nom        string    `json:"nom"`
	Prenom     string    `json:"prenom"`
	AnneeEtude int       `json:"annee_etude";xorm:"primaryKey;autoIncrement:false"`
	CoursJson  CoursJSON `json:"cours_json";xorm:"varchar(256)"`
}

type CoursJSON []string

func (c *CoursJSON) UnmarshalJSON(b []byte) error {

	var f []string
	unquoted, _ := strconv.Unquote(string(b))
	_ = json.Unmarshal([]byte(unquoted), &f)
	*c = f
	return nil
}

func (c *CoursJSON) Scan(value interface{}) error {
	var ba []byte

	switch v := value.(type) {
	case []byte:
		ba = v
	case string:
		ba = []byte(v)
	default:
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	var t []string
	rd := bytes.NewReader(ba)
	decoder := json.NewDecoder(rd)
	decoder.UseNumber()
	err := decoder.Decode(&t)
	*c = t
	return err
}

func (c CoursJSON) Value() (driver.Value, error) {
	return json.Marshal(c)
}

func GetInscriptions() ([]ListeInscriptions, error) {

	var inscriptions []ListeInscriptions
	err := db.Find(&inscriptions)
	if err != nil {
		return nil, err
	}

	return inscriptions, err
}
