package controllers

import (
	"Go-Web/views"
	"fmt"
	"net/http"
)

//NewUsers is used to create a new Users controller
//This function will panic if the templates are not parsed correctly
//and should only be used during the initial setup
func NewUsers() *Users {

	return &Users{
		New: views.NewView("base", "views/users/newUser.gohtml"),
	}
}

type Users struct {
	New *views.View
}

type SignupForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

//Create is used to process the signup form and create a new user account
//POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var SignupForm SignupForm
	err := parseForm(r, &SignupForm)
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(w, SignupForm)
}
