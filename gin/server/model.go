package server

import "gopkg.in/mgo.v2"

type Person struct {
	FirstName string `form:"first_name" json:"first_name" binding:"required"`
	LastName  string `form:"last_name" json:"last_name" binding:"required"`
	Email     string `form:"email" json:"email" binding:"required"`
	Company   string `form:"company" json:"company" binding:"required"`
}

func fetchAllPersons(db *mgo.Database) []Person {
	persons := []Person{}
	err := db.C("persons").Find(nil).All(&persons)
	if err != nil {
		panic(err)
	}

	return persons
}
