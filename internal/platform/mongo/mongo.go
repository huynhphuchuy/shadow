package mongo

import (
	"server/internal/config"
	"time"

	"gopkg.in/mgo.v2"
)

var Session *mgo.Database

func Init() {

	config := config.GetConfig()
	host := config.GetString("mongo.host")
	database := config.GetString("mongo.database")
	username := config.GetString("mongo.username")
	password := config.GetString("mongo.password")

	// Tạo phiên kết nối với MongDB
	info := &mgo.DialInfo{
		Addrs:    []string{host},
		Timeout:  60 * time.Second,
		Database: database,
		Username: username,
		Password: password,
	}
	session, err := mgo.DialWithInfo(info)
	if err != nil {
		panic(err)
	}

	Session = session.DB(database)
}
