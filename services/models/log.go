package models

import (
	"errors"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ActionLog struct {
	ID         bson.ObjectId `bson: "_id"`
	RecordID   string        `bson:"recordID"`
	Collection string        `bson:"collection"`
	Action     string        `bson:action`
	CreatedAt  time.Time     `bson:"createdAt"`
}

func (l *ActionLog) Create(db *mgo.Session) error {

	l.ID = bson.NewObjectId()
	l.CreatedAt = time.Now()

	if l.Action == "" {
		l.Action = "Create"
	}

	if l.RecordID != "" && l.Collection != "" {
		return db.DB("wiffee").C("log").Insert(l)
	}

	return errors.New("Missing attributes: RecordID or Collection")
}
