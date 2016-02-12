package server

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

type DatabaseSession struct {
	*mgo.Session
	DatabaseName string
}

func NewSession(name string) *DatabaseSession {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	addIndexToPersonEmails(session.DB(name))
	return &DatabaseSession{session, name}
}

func addIndexToPersonEmails(db *mgo.Database) {
	index := mgo.Index{
		Key:      []string{"email"},
		Unique:   true,
		DropDups: true,
	}
	indexErr := db.C("persons").EnsureIndex(index)
	if indexErr != nil {
		panic(indexErr)
	}
}

func (session *DatabaseSession) Database() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := session.Clone()
		c.Set("db", s.DB(session.DatabaseName))
		c.Next()
		defer s.Close()
	}
}
