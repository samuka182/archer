package db

import (
	"archer/config"
	"archer/structs"
	"os"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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
	if env := os.Getenv("APP_ENV"); env != "" {
		s, err := mgo.DialWithInfo(info)

		if err != nil {
			panic(err)
		}
		return s
	}

	s, err := mgo.Dial(conf.DB.Host)

	if err != nil {
		panic(err)
	}
	return s

}

//Save ...
func Save(collection string, u interface{}) error {
	return session.DB(conf.DB.Name).C(collection).Insert(u)
}

//FindByID ...
func FindByID(collection string, id string, data interface{}) error {
	return session.DB(conf.DB.Name).C(collection).Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(data)
}

//FindAll ...
func FindAll(collection string, data interface{}) error {
	return session.DB(conf.DB.Name).C(collection).Find(nil).All(data)
}
