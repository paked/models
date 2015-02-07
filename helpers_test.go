package models

import (
	"gopkg.in/mgo.v2/bson"
	"testing"
)

func TestSetValues(t *testing.T) {
	d := Dog{Age: 1, Name: "Bucky"}
	v := bson.M{"age": 2, "name": "Buck"}

	setValues(&d, v)

	if d.Age != v["age"] {
		t.Errorf("Age should be %v not %v", v["age"], d.Age)
	}

	if d.Name != v["name"] {
		t.Errorf("Name should be %v not %v", v["name"], d.Name)
	}
}

func TestEmpty(t *testing.T) {
	if !empty(Dog{}) {
		t.Error("Dog{} should be empty, not not empty")
	}
}
