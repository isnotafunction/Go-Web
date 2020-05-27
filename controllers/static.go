package controllers

import (
	"Go-Web/views"
)

func NewStatic() *Static {

	return &Static{
		Home:    views.NewView("base", "views/static/home.gohtml"),
		Contact: views.NewView("base", "views/static/contact.gohtml"),
	}
}

type Static struct {
	Home    *views.View
	Contact *views.View
}
