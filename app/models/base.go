package models

import (
	"fmt"
	"ucl-epreuve-technique/app/utils"

	"xorm.io/xorm"
)

import _ "github.com/mattn/go-sqlite3"

var db *xorm.Engine

func init() {

	engine, err := xorm.NewEngine("sqlite3", "universite_demo.sqlite")
	if err != nil {
		fmt.Println("Database error", err)
	}
	db = engine
	if utils.GetEnv("APPLICATION_DEBUG") == "true" {
		db.ShowSQL(true)
	} else {
		db.ShowSQL(false)
	}

}

func GetDB() *xorm.Engine {
	return db
}
