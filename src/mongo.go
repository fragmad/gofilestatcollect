package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	//	"gopkg.in/mgo.v2/bson"
	"os"
)

type Book struct {
	Name string `bson:"Name"`
}

func main() {
	var url = "localhost:8080"
	session, err := mgo.Dial(url)

	if err != nil {
		fmt.Printf("Cannot connect, go error %v", err)
		os.Exit(1)
	}
	c := session.DB("store").C("books")

	book := Book{
		Name: "Neuromancer",
	}

	err = c.Insert(book)

	//	c := session.DB(database).C(collection)
	//	err := c.Find(query).One(&result)
}
