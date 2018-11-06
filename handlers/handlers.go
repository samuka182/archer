package handlers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/samuka182/archer/archttp"
	"github.com/samuka182/archer/structs"

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

	resp, err := archttp.Request(dados)
	if err != nil {
		panic(err)
	}

	fullDispatch(w, resp, err)

}

func fullDispatch(w http.ResponseWriter, resp *structs.Response, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(w)
	err = enc.Encode(resp)
	if err != nil {
		panic(err)
	}

	resp.ID = bson.NewObjectId()
	resp.Timestamp = time.Now()
}

//GetTest ...
func GetTest(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(w)
	err := enc.Encode(&map[string]string{"nome": "teste"})
	if err != nil {
		panic(err)
	}

}
