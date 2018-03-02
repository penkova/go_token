package handler

import (
	"fmt"
	"net/http"
	"github.com/user/go_token/model"
	"github.com/user/go_token/db"
	"gopkg.in/mgo.v2/bson"
)

const (
	// Key (Should come from somewhere else).
	Key = "secret"
)

func StartPage(w http.ResponseWriter, r *http.Request) {
}
func SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.ServeFile(w, r, "templates/signup.html")
		return
	}

	username := r.FormValue("userName")
	password := r.FormValue("password")
	signup := r.FormValue("signup")

	var user model.User
	user.Username = username
	user.Password = password

	user.ID = bson.NewObjectId()

	if err := db.CreateUser(user); err != nil {
		fmt.Println("error with insert")
		http.Error(w, http.StatusText(500), 500)
		return
	}

	fmt.Println("Person with ID %v created successfully \n" + user.ID)

}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello")
}

func Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Bye")
}
