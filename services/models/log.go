package models

import (
	"errors"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type ActionLog struct {
	ID         bson.ObjectId `bson: "_id"`
	RecordID   string        `bson:"recordID"`
	Collection string        `bson:"collection"`
	Action     string        `bson:action`
	CreatedAt  time.Time     `bson:"createdAt"`
}

func (l *ActionLog) Create() error {
	dbSession := db.Copy()
	defer dbSession.Close()

	l.ID = bson.NewObjectId()
	l.CreatedAt = time.Now()

	if l.Action == "" {
		l.Action = "Create"
	}

	if l.RecordID != "" && l.Collection != "" {
		return dbSession.DB("").C("log").Insert(l)
	}

	return errors.New("Missing attributes: RecordID or Collection")
}
