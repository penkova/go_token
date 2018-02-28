package main

import (
	"net/http"
	"fmt"

	"github.com/gorilla/mux"
	"github.com/user/go_token/handler"
)
const (
	port = ":8000"
)

func main(){
	r := mux.NewRouter()

	fmt.Println("Server starting, localhost:8000 to start")
	r.HandleFunc("/", handler.StartPage)
	r.HandleFunc("/login", handler.Login).Methods("POST")
	r.HandleFunc("/signup", handler.SignUp).Methods("GET")

	http.ListenAndServe(port, r)




	//r := mux.NewRouter()
	//
	//fmt.Println("Server starting, localhost:8000 to start")
	//r.HandleFunc("/", model.StartPage)
	//r.HandleFunc("/authenticate", model.CreateTokenEndpoint).Methods("POST")
	//r.HandleFunc("/protected", model.ProtectedEndpoint).Methods("GET")
	//
	//http.ListenAndServe(port, r)

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



