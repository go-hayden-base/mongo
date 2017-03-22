package mongo

import (
	"errors"

	mgo "gopkg.in/mgo.v2"
)

// RegisterMongo register a mongodb connection.
// alias is connection's alias
// url is connection url
// db is connection's default database name
func RegisterMongo(alias, url, db string) error {
	if _, ok := sessionMap["default"]; !ok && alias != "default" {
		return errors.New(ErrNoDefaultConnection)
	}
	if _, ok := sessionMap[alias]; ok {
		return errors.New(ErrExistConnectionAlias)
	}
	if db == "" {
		db = "default"
	}
	aSession, err := mgo.Dial(url)
	if err != nil {
		return err
	}
	aMongo := new(tMongo)
	aMongo.session = aSession
	aMongo.db = db
	sessionMap[alias] = aMongo
	return nil
}

// NewMongo return a new mongodb operator
func NewMongo() Mongoer {
	aMogon := new(tMongo)
	aMogon.Use("default")
	return aMogon
}
