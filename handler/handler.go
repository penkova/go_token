package handler

import (
	"fmt"
	"net/http"
)

const (
	// Key (Should come from somewhere else).
	Key = "secret"
)

func StartPage(w http.ResponseWriter, r *http.Request) {

}
func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello")
}
func SignUp(w http.ResponseWriter, r *http.Request) {

}
