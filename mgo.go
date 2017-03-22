package mongo

import (
	"errors"

	"gopkg.in/mgo.v2"
)

var sessionMap = make(map[string]*tMongo)

type Mongoer interface {
	Use(alias string) error
	UseDB(db string)
	Session() *mgo.Session
	CurrentDB() string
	Collection(c string) *mgo.Collection
	// CURD
	Insert(collection string, docs ...interface{}) error
	Update(collection string, selector interface{}, update interface{}) error
	Find(collection string, query interface{}) ([]interface{}, error)
	FindOne(collection string, query interface{}) (interface{}, error)
	Remove(collection string, selector interface{}) error
}

type tMongo struct {
	db      string
	session *mgo.Session
}

func (s *tMongo) Use(alias string) error {
	c, ok := sessionMap[alias]
	if ok {
		s.db = c.db
		s.session = c.session.Clone()
	}
	return errors.New(ErrNoConnection + " named " + alias)
}

func (s *tMongo) UseDB(db string) {
	if db == "" {
		return
	}
	s.db = db
}

func (s *tMongo) Session() *mgo.Session {
	return s.session
}

func (s *tMongo) CurrentDB() string {
	return s.db
}

func (s *tMongo) Collection(c string) *mgo.Collection {
	return s.session.DB(s.db).C(c)
}

// CURD
func (s *tMongo) Insert(collection string, docs ...interface{}) error {
	c := s.Collection(collection)
	if c == nil {
		return errors.New(ErrCannotSwitchCollection + " '" + collection + "' in db '" + s.db + "'")
	}
	return c.Insert(docs)
}

func (s *tMongo) Update(collection string, selector interface{}, update interface{}) error {
	c := s.Collection(collection)
	if c == nil {
		return errors.New(ErrCannotSwitchCollection + " '" + collection + "' in db '" + s.db + "'")
	}
	return c.Update(selector, update)
}

func (s *tMongo) Find(collection string, query interface{}) ([]interface{}, error) {
	c := s.Collection(collection)
	if c == nil {
		return nil, errors.New(ErrCannotSwitchCollection + " '" + collection + "' in db '" + s.db + "'")
	}
	var result []interface{}
	err := c.Find(query).All(&result)
	return result, err
}

func (s *tMongo) FindOne(collection string, query interface{}) (interface{}, error) {
	c := s.Collection(collection)
	if c == nil {
		return nil, errors.New(ErrCannotSwitchCollection + " '" + collection + "' in db '" + s.db + "'")
	}
	var result interface{}
	err := c.Find(query).One(&result)
	return result, err
}

func (s *tMongo) Remove(collection string, selector interface{}) error {
	c := s.Collection(collection)
	if c == nil {
		return errors.New(ErrCannotSwitchCollection + " '" + collection + "' in db '" + s.db + "'")
	}
	return c.Remove(selector)
}
