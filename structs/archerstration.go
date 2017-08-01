package structs

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//Archerstration ...
type Archerstration struct {
	ID           bson.ObjectId `json:"_id" bson:"_id"`
	CreationTime time.Time     `json:"_creation_time" bson:"_creation_time"`
	UpdateTime   string        `json:"_update_time" bson:"_update_time"`
	Name         string        `json:"name" bson:"name"`
	Flows        []Flow        `json:"flows" bson:"flows"`
}
