package structs

import (
	"encoding/json"
	"strings"
)

//Response represents a response of request
type Response struct {
	Headers        []Header `json:"headers"`
	HTTPStatus     string   `json:"http-status"`
	HTTPStatusCode int      `json:"http-status-code"`
	Payload        string   `json:"payload"`
	ExecutionTime  string   `json:"response-time"`
}

//Header represents a http header
type Header struct {
	Chave string `json:"key"`
	Valor string `json:"value"`
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
