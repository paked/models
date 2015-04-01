package models

import (
	"errors"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Init creates a connection to the database
func Init(host, db string) error {
	conn = &connection{}
	return conn.connect(host, db)
}

// Modeller is an interface for use with the ORM, describing a model.
type Modeller interface {
	BID() bson.ObjectId
	C() string
}

// Persist creates a copy of the model and persists it in the DB.
func Persist(m Modeller) error {
	c := conn.collection(m.C())
	if err := c.Insert(m); err != nil {
		return err
	}

	return nil
}

// Update updates a Modeller interface with the provided values in persistent storage.
func Update(m Modeller, values bson.M) error {
	if err := updateValues(m, values); err != nil {
		return err
	}

	setValues(m, values)

	return nil
}

// Remove removes a model from the MongoDB.
func Remove(m Modeller) error {
	c := conn.collection(m.C())

	return c.RemoveId(m.BID())
}

// RemoveAll removes all models matching which match the values query
func RemoveAll(collection string, values bson.M) error {
	c := conn.collection(collection)
	_, err := c.RemoveAll(values)
	return err
}

// RestoreByID restores a model using the specified bson ObjectId.
func RestoreByID(m Modeller, id bson.ObjectId) error {
	return Restore(m, bson.M{"_id": id})
}

// Restore restores a model using specified values as a search key
func Restore(m Modeller, values bson.M) error {
	c := conn.collection(m.C())

	err := c.Find(values).One(m)
	if err != nil {
		return err
	}

	if empty(m) {
		return errors.New("Could not find that model")
	}

	return nil
}

// Fetch retrieves all models matching a set of values in a collections
func Fetch(collection string, values bson.M, sort string) (*mgo.Iter, error) {
	c := conn.collection(collection)

	iter := c.Find(values).Sort(sort).Iter()
	return iter, nil
}
