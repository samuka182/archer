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

	"github.com/gorilla/mux"

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

	db.Save("responses", resp)
}

//CreateArchestration ...
func CreateArchestration(w http.ResponseWriter, r *http.Request) {
	jsonRequest, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	archerstration := &structs.Archerstration{}
	err = json.Unmarshal(jsonRequest, archerstration)

	archerstration.ID = bson.NewObjectId()
	archerstration.CreationTime = time.Now()

	err2 := db.Save("archerstrations", archerstration)
	if err2 != nil {
		panic(err2)
	}

}

//GetArchestration ...
func GetArchestration(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	archerstration := structs.Archerstration{}

	db.FindByID("archerstrations", id, &archerstration)

	if archerstration.ID == "" {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		enc := json.NewEncoder(w)
		err := enc.Encode(archerstration)
		if err != nil {
			panic(err)
		}
	}

}

//GetArchestrations ...
func GetArchestrations(w http.ResponseWriter, r *http.Request) {

	archerstrations := []structs.Archerstration{}

	db.FindAll("archerstrations", &archerstrations)

	if archerstrations == nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		enc := json.NewEncoder(w)
		err := enc.Encode(archerstrations)
		if err != nil {
			panic(err)
		}
	}

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
