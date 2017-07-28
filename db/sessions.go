package db

import (
	"archer/config"

	mgo "gopkg.in/mgo.v2"
)

var conf = config.Get()

//GetSession ...
func getSession() *mgo.Session {

	s, err := mgo.Dial(conf.DB.Local)

	if err != nil {
		panic(err)
	}
	return s
}

//Save ...
func Save(collection string, u interface{}) {
	getSession().DB(conf.DB.Name).C(collection).Insert(u)
}
