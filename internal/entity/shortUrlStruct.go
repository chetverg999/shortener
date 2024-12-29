package entity

import "gopkg.in/mgo.v2/bson"

type ShortURL struct {
	Id      bson.ObjectId `bson:"_id"`
	UserURL string        `bson:"UserURL"`
	Short   string        `bson:"Short"`
}
