package db

import (
	"archer/config"
	"archer/structs"
	"time"

	mgo "gopkg.in/mgo.v2"
)

var conf = &structs.Config{}
var info = &mgo.DialInfo{}
var session = &mgo.Session{}

func init() {
	conf = config.Get()
	info = &mgo.DialInfo{
		Addrs:    []string{conf.DB.Host},
		Timeout:  time.Duration(conf.DB.Timeout) * time.Second,
		Database: "admin",
		Username: conf.DB.User,
		Password: conf.DB.Pass,
	}
	session = getSession()
}

//GetSession ...
func getSession() *mgo.Session {
	s, err := mgo.DialWithInfo(info)

	if err != nil {
		panic(err)
	}
	return s
}

//Save ...
func Save(collection string, u interface{}) {
	err := session.DB(conf.DB.Name).C(collection).Insert(u)
	if err != nil {
		panic(err)
	}
}
