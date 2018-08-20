package models

import (
	"log"

	mgo "gopkg.in/mgo.v2"
)

var (
	db *mgo.Session
)

func InitDb(mgoAddr string) {
	// To add secure password
	var err error
	db, err = mgo.Dial(mgoAddr)
	if err != nil {
		log.Fatal("Cannot connect to mongodb", err)
	}
}
