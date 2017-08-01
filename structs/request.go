package structs

import "gopkg.in/mgo.v2/bson"
import "time"

//Request represents a request
type Request struct {
	ID        bson.ObjectId `json:"-" bson:"_id,omitempty"`
	Timestamp time.Time     `json:"-" bson:"_timestamp,omitempty"`
	Method    string        `json:"method" bson:"method,omitempty"`
	Path      []struct {
		Key   string `json:"key" bson:"key"`
		Value string `json:"value" bson:"value"`
	} `json:"path,omitempty" bson:"path,omitempty"`
	Header []struct {
		Key   string `json:"key" bson:"key"`
		Value string `json:"value" bson:"value"`
	} `json:"header,omitempty" bson:"header,omitempty"`
	QueryString []struct {
		Key   string `json:"key" bson:"key"`
		Value string `json:"value" bson:"value"`
	} `json:"query-string" bson:"query-string,omitempty"`
	URL     string `json:"url" bson:"url,omitempty"`
	Payload string `json:"payload,omitempty" bson:"payload,omitempty"`
}
