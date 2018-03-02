package main

import (
	"net/http"
	"fmt"
	"github.com/gorilla/context"
	_ "github.com/dgrijalva/jwt-go"
	"github.com/user/go_token/handler"
)

const (
	port = ":8080"
)

func main() {
	fmt.Println("Server starting, point your browser to localhost:8080 to start")
	// Here we are instantiating the gorilla/mux router
	//r := mux.NewRouter()

	//r.HandleFunc("/", models.StartPage).Methods("POST")
	http.HandleFunc("/", handler.StartPage)
	http.HandleFunc("/signup", handler.SignUp)
	http.HandleFunc("/login", handler.Login)
	http.HandleFunc("/logout", handler.Logout)

	// Our application will run on port 3030. Here we declare the port and pass in our router.
	http.ListenAndServe(port, context.ClearHandler(http.DefaultServeMux))


	//r := mux.NewRouter()
	//
	//fmt.Println("Server starting, localhost:8000 to start")
	//r.HandleFunc("/", model.StartPage)
	//r.HandleFunc("/authenticate", model.CreateTokenEndpoint).Methods("POST")
	//r.HandleFunc("/protected", model.ProtectedEndpoint).Methods("GET")
	//
	//http.ListenAndServe(port, r)
	//
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
}



