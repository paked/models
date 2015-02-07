package models

import (
	"gopkg.in/mgo.v2/bson"
	"os"
	"testing"
)

type Dog struct {
	ID    bson.ObjectId `bson:"_id"`
	Name  string        `bson:"name"`
	Owner string        `bson:"owner"`
	Age   int           `bson:"age"`
}

func (d Dog) BID() bson.ObjectId {
	return d.ID
}

func (d Dog) C() string {
	return "dogs"
}

var d *Dog

func TestMain(m *testing.M) {
	Init("localhost", "models-test")
	os.Exit(m.Run())
}

func TestModeller(t *testing.T) {
	d = &Dog{ID: bson.NewObjectId(), Name: "Doggy", Owner: "James", Age: 10}

	if err := Persist(d); err != nil {
		t.Error("Could not create that model")
		t.FailNow()
	}

	if err := Update(d, bson.M{"age": 5}); err != nil {
		t.Error("Could not udpate model", err)
		t.FailNow()
	}

	if d.Age != 5 {
		t.Error("Age should be 5, not ", d.Age)
	}
}

func TestPersist(t *testing.T) {
	e := &Dog{}
	err := RestoreByID(e, d.BID())

	if err != nil {
		t.Error("Error restoring model:", err)
	}

	if e.BID() != d.BID() {
		t.Error("This is not the same model...")
	}
}

func TestRemove(t *testing.T) {
	if err := Remove(d); err != nil {
		t.Error("Could not remove model:", err)
		t.FailNow()
	}

	e := &Dog{}

	if err := RestoreByID(e, d.BID()); err == nil {
		t.Error("Model not found.")
	}
}
