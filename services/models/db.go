package models

import (
	"log"
	"os"

	mgo "gopkg.in/mgo.v2"
)

var (
	db     *mgo.Session
	dbName string = "wiffee"
)

func InitDb() {
	// To add secure password
	mgoAddr := os.Getenv("MONGO_DB_ADDRESS")
	var err error

	info := &mgo.DialInfo{
		Addrs:    []string{mgoAddr},
		Database: dbName,
	}

	db, err = mgo.DialWithInfo(info)
	if err != nil {
		log.Fatal("Cannot connect to mongodb", err)
	}
}
