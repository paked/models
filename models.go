// Package models is a simple, lightweight ORM built in Go. It aims to provide a simple and clean API to improve testing and code readability in projects dependent on databases.
//
// Curretly models only supports MongoDB as a database driver.
//
// Getting started with models is as simple as satisfying the Modeller interface:
// 	type Modeller interface {
// 		BID() bson.ObjectId // The Modellers ID.
// 		C() string 	    // The collection the Modeller belongs to.
// 	}
//
// With a struct like:
// 	type Dog struct {
// 		ID    bson.ObjectId `bson:"_id"`
// 		Name  string        `bson:"name"`
// 		Owner string        `bson:"owner"`
// 		Age   int           `bson:"age"`
// 	}
// 	func (d Dog) BID() bson.ObjectId {
// 		return d.ID
// 	}
// 	func (d Dog) C() string {
// 		return "dogs"
// 	}
//
// Creating a model is as easy as creating a regular struct.
// 	d := Dog{ID: bson.ObjectID, Name: "Bucky", Owner:"Tony", Age: 3}
// and then, when you want to persist it:
// 	err := models.Persist(d)
package models
