package structs

import (
	"encoding/json"
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"
)

//Response represents a response of request
type Response struct {
	ID             bson.ObjectId `json:"-" bson:"_id"`
	Timestamp      time.Time     `json:"-" bson:"_timestamp"`
	Headers        []Header      `json:"headers,omitempty" bson:"headers"`
	HTTPStatus     string        `json:"http-status,omitempty" bson:"http-status"`
	HTTPStatusCode int           `json:"http-status-code,omitempty" bson:"http-status-code"`
	Payload        string        `json:"payload,omitempty" bson:"payload"`
	ExecutionTime  string        `json:"response-time,omitempty" bson:"response-time"`
}

//Header represents a http header
type Header struct {
	Chave string `json:"key" bson:"key"`
	Valor string `json:"value" bson:"value"`
}

//PayloadIsArray ...
func (r *Response) PayloadIsArray() (b bool) {
	return strings.HasPrefix(r.Payload, "[")
}

//GetJSONArray ...
func (r *Response) GetJSONArray() (a *JSONArray) {
	data := &JSONArray{}

	json.Unmarshal([]byte(r.Payload), &data.Values)

	return data
}

//GetJSONObject ...
func (r *Response) GetJSONObject() (o *JSONObject) {
	data := &JSONObject{}

	json.Unmarshal([]byte(r.Payload), &data.Value)

	return data
}
