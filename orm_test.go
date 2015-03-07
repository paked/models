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

func TestPersist(t *testing.T) {
	d = &Dog{ID: bson.NewObjectId(), Name: "Doggy", Owner: "James", Age: 10}

	if err := Persist(d); err != nil {
		t.Error("Could not create that model")
		t.FailNow()
	}
}

func TestUpdate(t *testing.T) {
	v := bson.M{"age": 5, "name": "dogg"}
	if err := Update(d, v); err != nil {
		t.Error("Could not udpate model", err)
		t.FailNow()
	}

	if d.Age != v["age"] {
		t.Error("Age should be 5, not ", d.Age)
	}

	if d.Name != v["name"] {
		t.Errorf("Name should be %v not %v", v["name"], d.Name)
	}
}

func TestRestore(t *testing.T) {
	e := &Dog{}
	err := Restore(e, bson.M{"name": d.Name})

	if err != nil {
		t.Error("Error restoring model:", err)
	}

	if e.BID() != d.BID() {
		t.Error("This is not the same model...")
	}
}

func TestRestoreByID(t *testing.T) {
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

func TestFetch(t *testing.T) {
	for i := 0; i < 10; i++ {
		e := Dog{ID: bson.NewObjectId(), Name: "bill" + string(i), Owner: "paked"}
		if err := Persist(e); err != nil {
			t.Fatal("Unable to persist model", i)
		}
	}

	var es []Dog
	dog := Dog{}
	dogs, err := Fetch(dog.C(), bson.M{"owner": "paked"})
	if err != nil {
		t.Error("Unable to fetch all of those doggies :(")
	}

	for dogs.Next(&dog) {
		es = append(es, dog)
	}

	if err := RemoveAll(dog.C(), bson.M{"owner": "paked"}); err != nil {
		t.Error("Unable to cleanup!")
	}
}
