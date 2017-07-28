package handlers

import (
	"archer/archttp"
	"archer/db"
	"archer/structs"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"gopkg.in/mgo.v2/bson"
)

//FullRequest ...
func FullRequest(w http.ResponseWriter, r *http.Request) {

	jsonRequest, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	dados := &structs.Request{}
	err = json.Unmarshal(jsonRequest, dados)

	dados.ID = bson.NewObjectId()
	dados.Timestamp = time.Now()

	db.Save("requests", dados)

	resp, err := archttp.Request(dados)
	if err != nil {
		panic(err)
	}

	FullDispatch(w, resp, err)

}

//FullDispatch ...
func FullDispatch(w http.ResponseWriter, resp *structs.Response, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(w)
	err = enc.Encode(resp)
	if err != nil {
		panic(err)
	}

	resp.ID = bson.NewObjectId()
	resp.Timestamp = time.Now()

	db.Save("responses", resp)
}
