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

func (p *Place) LoadOrCreate() error {
	err := p.Load()
	if err != nil && err == mgo.ErrNotFound {
		return p.Create()
	}

	return err
}

func (p *Place) Load() error {
	dbSession := db.Copy()
	defer dbSession.Close()

	collection := dbSession.DB("wiffee").C("places")
	query := bson.M{"placeID": p.PlaceID}

	return collection.Find(query).One(p)
}

func (p *Place) Create() error {
	dbSession := db.Copy()
	defer dbSession.Close()

	now := time.Now()
	p.ID = bson.NewObjectId()
	p.CreatedAt = now
	p.LastUpdatedAt = now

	return dbSession.DB("wiffee").C("places").Insert(p)
}
