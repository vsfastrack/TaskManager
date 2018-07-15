package controllers

import (
	"gopkg.in/mgo.v2"
)

//GetCollection from DBUtils
func GetCollection(collectionName string) *mgo.Session {
	context := NewContext()
	defer context.Close()
	c := context.DbCollection(collectionName)
	return context
}
