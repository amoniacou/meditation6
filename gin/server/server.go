package server

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"net/http"
)

func NewServer(session *DatabaseSession) *gin.Engine {
	s := gin.Default()
	s.Use(session.Database())
	s.GET("/persons", func(c *gin.Context) {
		db := c.MustGet("db").(*mgo.Database)
		c.JSON(http.StatusOK, fetchAllPersons(db))
	})
	s.POST("/persons", func(c *gin.Context) {
		var json Person
		if c.BindJSON(&json) == nil {
			db := c.MustGet("db").(*mgo.Database)
			err := db.C("persons").Insert(json)
			if err == nil {
				c.JSON(201, json)
			} else {
				c.JSON(400, gin.H{"error": err.Error()})
			}
		}
	})
	return s
}
