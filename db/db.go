package db

import (
	"github.com/user/go_token/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"log"
)

var db *mgo.Database

// Establish a connection to MongoDB database.
func init() {
	session, err := mgo.Dial("localhost:27018")
	if err != nil {
		fmt.Println("Error")
		log.Fatalf("Failed to connect to database: %v", err)
	}
	db = session.DB("authmongo")
	fmt.Println("Connect to databses - Successful")
}

// CollectionUsers - connection to usersdb collections
func CollectionUsers() *mgo.Collection {
	return db.C("usersdb")
}

// CreateUser - create User
func CreateUser(user model.User) error {
	return CollectionUsers().Insert(user)
}

// FindUser - finding user by username and password
func FindUser(username string, password string) (*model.User, error) {
	res := model.User{}
	err := CollectionUsers().Find(bson.M{
		"username": username,
		"password": password,
	}).One(&res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
