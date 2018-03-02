package db

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"log"
	"github.com/user/go_token/model"
)
//type User struct {
//	ID       bson.ObjectId `json:"id" bson:"_id"`
//	Username string        `json:"username" bson:"username"`
//	Password string        `json:"password" bson:"password"`
//	//Token     string        `json:"token,omitempty" 	bson:"-"`
//}

var db *mgo.Database

// Establish a connection to MongoDB database.
func init() {
	session, err := mgo.Dial("localhost:27018")
	if err != nil {
		fmt.Println("Error")
		log.Fatalf("Failed to connect to database: %v", err)
	}
	db = session.DB("authmongo")
	fmt.Println("Woow")
}
func CollectionUsers() *mgo.Collection {
	return db.C("usersdb")
}

func CreateUser(user model.User) error {
	return CollectionUsers().Insert(user)
}