package structs

import "gopkg.in/mgo.v2/bson"
import "time"

//Request represents a request
type Request struct {
	ID        bson.ObjectId `json:"-" bson:"_id"`
	Timestamp time.Time     `json:"-" bson:"_timestamp"`
	Method    string        `json:"method" bson:"method"`
	Path      []struct {
		Key   string `json:"key" bson:"key"`
		Value string `json:"value" bson:"value"`
	} `json:"path" bson:"path"`
	Header []struct {
		Key   string `json:"key" bson:"key"`
		Value string `json:"value" bson:"value"`
	} `json:"header" bson:"header"`
	QueryString []struct {
		Key   string `json:"key" bson:"key"`
		Value string `json:"value" bson:"value"`
	} `json:"query-string" bson:"query-string"`
	URL     string `json:"url" bson:"url"`
	Payload string `json:"payload" bson:"payload"`
}
