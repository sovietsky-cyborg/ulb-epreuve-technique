package models

import (
	"fmt"

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
	db.ShowSQL(true)

}

func GetDB() *xorm.Engine {
	return db
}
