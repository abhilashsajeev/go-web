package controllers

import (
    "gopkg.in/mgo.v2"
    "github.com/abhilashsajeev/go-web/taskmanager/common"
)

type Context struct {
    MongoSession *mgo.Session
    User string
}

func (c *Context) Close(){
    c.MongoSession.Close()
}

// returns DB collection for given name
func (c *Context) DBCollection(name string) *mgo.Collection {
    return c.MongoSession.DB(common.Appconfig.Database).C(name)
}

func NewContext() *Context{
    session := common.GetSession().Copy()
    context := &Context {
        MongoSession: session,
    }
    return context
}
    


