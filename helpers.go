package models

import (
	"gopkg.in/mgo.v2/bson"
	"reflect"
)

// set the values of pointer x, to those specified in the bson.M map.
func setValues(x interface{}, values bson.M) {
	v := reflect.ValueOf(x).Elem()

	for i := 0; i < v.NumField(); i++ {
		f := v.Type().Field(i)
		tag := f.Tag.Get("bson")

		val := reflect.ValueOf(values[tag])

		if !val.IsValid() || empty(val) {
			continue
		}

		v.Field(i).Set(val)
	}
}

// empty checks if the gived interface is equivelent to it's zero value.
func empty(x interface{}) bool {
	return reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
}

// updateValues updates a model in the MongoDB.
func updateValues(m Modeller, values bson.M) error {
	c := conn.collection(m.C())

	return c.UpdateId(m.BID(), bson.M{"$set": values})
}
