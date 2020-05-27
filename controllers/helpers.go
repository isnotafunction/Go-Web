package controllers

import (
	"net/http"

	"github.com/gorilla/schema"
)

func parseForm(r *http.Request, form interface{}) error {

	err := r.ParseForm()
	if err != nil {
		return (err)
	}
	var decoder = schema.NewDecoder()
	if err := decoder.Decode(form, r.PostForm); err != nil {
		return err
	}
	return nil
}
