package controllers

import (
	"github.com/vsfastrack/TaskManager/common"
	"gopkg.in/mgo.v2"
)

//Context for session
type Context struct {
	MongoSession *mgo.Session
}

//Close for session
func (c *Context) Close() {
	c.MongoSession.Close()
}

//DbCollection for session
func (c *Context) DbCollection(name string) *mgo.Collection {
	return c.MongoSession.DB(common.AppConfig.Database).C(name)
}

//NewContext for session
func NewContext() *Context {
	session := common.GetSession().Copy()
	context := &Context{
		MongoSession: session,
	}
	return context
}
