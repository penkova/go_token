package model

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID        bson.ObjectId `json:"id" 					bson:"_id,omitempty"`
	Username  string        `json:"username" 			bson:"username"`
	Password  string        `json:"password,omitempty" 	bson:"password"`
	Token     string        `json:"token,omitempty" 	bson:"-"`
}

type TokenJwt struct {
	Token string `json:"token"`
}