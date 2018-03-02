package main

import (
	"fmt"
	_ "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/user/go_token/handler"
	"net/http"
)

const (
	port = ":3030"
)

func main() {
	fmt.Println("Server starting, point your browser to localhost:3030 to start")
	// Here we are instantiating the gorilla/mux router
	r := mux.NewRouter()

	//r.HandleFunc("/", models.StartPage).Methods("POST")
	r.HandleFunc("/", handler.StartPage)
	r.HandleFunc("/signup", handler.SignUp)
	r.HandleFunc("/login", handler.Login)
	r.HandleFunc("/logout", handler.Logout)

	// Our application will run on port 3030. Here we declare the port and pass in our router.
	http.ListenAndServe(port, r)
}
