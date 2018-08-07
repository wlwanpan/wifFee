package models

import (
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Wifi struct {
	ID        bson.ObjectId `bson:"_id"`
	PlaceID   string        `bson:"placeID"`
	CreatedAt time.Time     `bson:"createdAt"`
	Ping      uint64        `bson:"ping"`
	UpSpeed   uint64        `bson:"upSpeed"`
	DownSpeed uint64        `bson:downSpeed`
}

func (wif *Wifi) Create(db *mgo.Session) error {

	wif.CreatedAt = time.Now()

	collection := db.DB("wiffee").C("wifi")
	return collection.Insert(wif)
}
