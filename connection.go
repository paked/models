package models

import (
	"gopkg.in/mgo.v2"
)

type connection struct {
	session *mgo.Session
	db      string
}

func (conn *connection) connect(host, db string) error {
	var err error

	conn.db = db
	conn.session, err = mgo.Dial(host)

	return err
}

func (conn *connection) collection(name string) *mgo.Collection {
	return conn.session.DB(conn.db).C(name)
}

var (
	conn *connection
)
