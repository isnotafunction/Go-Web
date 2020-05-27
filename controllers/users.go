package controllers

import (
	"Go-Web/views"
	"fmt"
	"net/http"

	"github.com/gorilla/schema"
)

//NewUsers is used to create a new Users controller
//This function will panic if the templates are not parsed correctly
//and should only be used during the initial setup
func NewUsers() *Users {

	return &Users{
		NewView: views.NewView("base", "views/users/new.gohtml"),
	}
}

type Users struct {
	NewView *views.View
}

//New is used to render the signup form used to create a new user account
//GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	u.NewView.Render(w, nil)
}

type SignupForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

//Create is used to process the signup form and create a new user account
//POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var decoder = schema.NewDecoder()
	var SignupForm SignupForm

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	err = decoder.Decode(&SignupForm, r.PostForm)

	fmt.Fprintln(w, SignupForm)
}
