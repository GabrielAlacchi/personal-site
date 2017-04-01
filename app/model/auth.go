package model

import (
	"gopkg.in/mgo.v2/bson"
	"errors"
)

type Auth struct {
	Email string `json:"email" bson:"email"`
	BcryptDigest string `json:"hash" bson:"hash"`
	BcryptSalt string `json:"salt" bson:"salt"`
}

func (auth *Auth) QueryByEmail(email string) error {

	sess := getSession()
	cAuth := sess.DB("alacchi-com").C("authentication")
	query := cAuth.Find(bson.M{"email": email})

	if ct, err := query.Count(); ct == 1 {
		query.One(auth)
		return nil
	} else if ct == 0 {
		return errors.New("Email not found: " + email)
	} else {
		return err
	}

}

func (auth Auth) Insert() error {
	sess := getSession()
	cAuth := sess.DB("alacchi-com").C("authentication")
	return cAuth.Insert(auth)
}