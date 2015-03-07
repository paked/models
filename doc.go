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
// In order to use the models package a connection to the database (local or remote) has to be made.
//  models.Init("localhost", "database")
//
// Creating a model is just like creating a regular struct.
// 	d := Dog{ID: bson.ObjectID, Name: "Bucky", Owner:"Tony", Age: 3}
// Once you have a struct that satisfies the Modeller interface you are able to Persist the model.
// 	err := models.Persist(d)
//
// A persisted version of a model will only update when a models.Update(m, values) called.
// The keys in the values parameter should be whatever you tagged that field of the struct with.
// Eg. "name" for Dog.Name, "age" for Dog.Age and "_id" for Dog.ID
// 	err := models.Update(&d, bson.M{"age": 4})
//
// If you want to restore your model from a persisted record call models.Restore(m, values).
// This will use the values as a search query for mgo, and re-hydrate the passed in Pointer.
//	err := models.Restore(d, bson.M{"name": "Bucky"})
// There is a built-in alias for restoring a model based on ID called models.RestoreByID(m).
// err := models.RestoreById(d, bson.NewObjectId())
//
// Finally, if you want to remove a persisted model from the database, call models.Remove(m).
//	err := models.Remove(d)
//
// Basic options for manipulating groups of retrieving and deleting model groups are available.
// Fetching is the recommended way of pulling a group of models from the db.
//  var dogs []Dog
//  dog := Dog{}
//  iter, err := models.Fetch(dog.C(), bson.M{"age": 0})
//  if err != nil {
//		// handle the error
//  }
//  for iter.Next(&dog) {
// 		dogs = append(dogs, dog)
//	}
// Removing groups is also doable through the RemoveAll(collection, values) function.
//  err := models.RemoveAll("collection", bson.M{"x": -1})
package models
