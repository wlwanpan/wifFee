package models

import (
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Place struct {
	ID            bson.ObjectId `bson:"_id"`
	PlaceID       string        `bson:"placeID"`
	CreatedAt     time.Time     `bson:"createdAt"`
	LastUpdatedAt time.Time     `bson:"lastUpdatedAt"`
	AvgUp         uint32        `bson:"avgUp"`
	AvgDown       uint32        `bson:"avgDown"`
}

func (p *Place) LoadOrCreate(db *mgo.Session) error {

	err := p.Load(db)
	if err != nil && err.Error() == "not found" {
		return p.Create(db)
	}

	return err
}

func (p *Place) Load(db *mgo.Session) error {

	collection := db.DB("wiffee").C("places")
	query := bson.M{"placeID": p.PlaceID}

	return collection.Find(query).One(p)
}

func (p *Place) Create(db *mgo.Session) error {

	now := time.Now()
	p.ID = bson.NewObjectId()
	p.CreatedAt = now
	p.LastUpdatedAt = now

	return db.DB("wiffee").C("places").Insert(p)
}
