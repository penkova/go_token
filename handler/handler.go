package handler

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
	"encoding/json"
	"github.com/user/go_token/db"
	"github.com/user/go_token/model"
	"gopkg.in/mgo.v2/bson"
	"github.com/dgrijalva/jwt-go"
)

const (
	// Key (Should come from somewhere else).
	Key = "secret"
)
//Fetch all templates
var tmpls, tmpls_error = template.ParseGlob("tmpl/*")

// StartPage - start page
func StartPage(w http.ResponseWriter, r *http.Request) {

}

// SignUpUser - user registration
func SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.ServeFile(w, r, "templates/signup.html")
		return
	}

	username := r.FormValue("userName")
	password := r.FormValue("password")
	//signup := r.FormValue("signup")

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

// LoginUser - authentication user
func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login!")
	if r.Method != "POST" {
		http.ServeFile(w, r, "tmpl/login.html")
		return
	}

	//get user from put request
	username := r.FormValue("username")
	password := r.FormValue("password")
	//login := r.FormValue("login")

	//get request to database
	usDB, err :=  db.FindUser(username, password)
	if err != nil {
		log.Println("can’t find user:", err)
		http.Redirect(w, r, "/login", 301)
		return
	}
	fmt.Println(usDB)

	//create and return token
	if (username == usDB.Username && password == usDB.Password) {

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": usDB.Username,
			"password": usDB.Password,
		})
		tokenString, error := token.SignedString([]byte(Key))
		if error != nil {
			fmt.Println(error)
		}
		json.NewEncoder(w).Encode(model.TokenJwt{Token: tokenString})

		log.Println("You just have got token")

	} else {
		log.Println("password or user isn't correct", err)
		http.Error(w, "password or user isn't correct", http.StatusBadRequest)
	}

	//hmacSampleSecret := []byte("my_secret")
	//// Create a new token object, specifying signing method and the claims
	//// you would like it to contain.
	//token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	//	"foo": "bar",
	//	"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	//	//"nbf": time.Now().Unix(),
	//})
	//
	//// Sign and get the complete encoded token as a string using the secret
	//tokenString, err := token.SignedString(hmacSampleSecret)
	//
	//fmt.Println(tokenString, err)






	//token := jwt.New(jwt.SigningMethodHS256)
	//
	//// Set claims
	//claims := token.Claims.(jwt.MapClaims)
	//claims["id"] = u.ID
	//claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	//
	//// Generate encoded token and send it as response
	//u.Token, err = token.SignedString([]byte(Key))
	//if err != nil {
	//	return err
	//}
	//
	//u.Password = "" // Don't send password
	//return c.JSON(http.StatusOK, u)

}

// LogoutUser
func Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Bye")
}
