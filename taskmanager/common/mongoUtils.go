package common

import (
    "log"
    "time"
    "gopkg.in/mgo.v2"
)

var session *mgo.Session

func GetSession() *mgo.Session {
    if session == nil {
        var err error
        session, err = mgo.DialWithInfo(&mgo.DialInfo{
            Addrs: []string{Appconfig.MongoDBHost},
            Timeout : 60 * time.Second,
            });
        if err != nil {
            log.Fatalf("[GetSession]: %s\n", err)
        }
    }
    return session
}

func createDbSession() {
    var err error
    session, err = mgo.DialWithInfo(&mgo.DialInfo{
        Addrs: []string{Appconfig.MongoDBHost},
        Timeout : 60 * time.Second,
        });
    if err != nil {
        log.Fatalf("[CreateDBSession]: %s\n", err)
    }
}

// Add indexes into MongoDB
func addIndexes() {
    var err error
    userIndex := mgo.Index{
        Key:
        []string{"email"},
        Unique:
        true,
        Background: true,
        Sparse:
        true,
    }
    taskIndex := mgo.Index{
        Key:
        []string{"createdby"},
        Unique:
        false,
        Background: true,
        Sparse:
        true,
    }
    noteIndex := mgo.Index{
        Key:
        []string{"taskid"},
        Unique:
        false,
        Background: true,
        Sparse:
        true,
    }
    // Add indexes into MongoDB
    session := GetSession().Copy()
    defer session.Close()
    userCol := session.DB(Appconfig.Database).C("users")
    taskCol := session.DB(Appconfig.Database).C("tasks")
    noteCol := session.DB(Appconfig.Database).C("notes")
    err = userCol.EnsureIndex(userIndex)
    if err != nil {
        log.Fatalf("[addIndexes]: %s\n", err)
    }
    err = taskCol.EnsureIndex(taskIndex)
    if err != nil {
        log.Fatalf("[addIndexes]: %s\n", err)
    }
    err = noteCol.EnsureIndex(noteIndex)
    if err != nil {
        log.Fatalf("[addIndexes]: %s\n", err)
    }
}