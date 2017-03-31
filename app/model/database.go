package model

import (
	"gopkg.in/mgo.v2"
	"time"
	"log"
	"os"
)

var (
	session *mgo.Session
	url string
)

func getSession() *mgo.Session {
	var err error
	reconnect := false

	if url == "" {
		url = os.Getenv("MONGO_URL")
		if url == "" {
			url = "mongodb://localhost:27017"
		}
	}

	// Connect to MongoDB
	if session == nil {
		log.Println("[database] Connecting to mongodb at URL: " + url)
		reconnect = true
	} else if session.Ping() != nil {
		log.Println("[database] Mongo ping failed")
		reconnect = true
	}

	if reconnect {
		session, err = mgo.DialWithTimeout(url, 5 * time.Second)
		session.SetSocketTimeout(1 * time.Second)

		if err != nil {
			log.Fatal("[database] Cannot connect to MongoDB:", err.Error())
		}
	}

	// Return a copy of the global session, can be reconfigured after if necessary
	return session.Copy()
}
