package main

import (
	"Go-Web/views"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	// fmt.Fprint(w, "<h1>Welcome to my awesome site</h1>")
	fmt.Fprint(w, "<h1>Contact page</h1>")
}

//temporary global var
var homeView *views.View
var contactView *views.View

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	err := homeView.Template.ExecuteTemplate(w, homeView.Layout, nil)
	if err != nil {
		//not a good idea to panic here, will clean up later
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := contactView.Template.ExecuteTemplate(w, contactView.Layout, nil)
	if err != nil {
		panic(err)
	}
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>Such empty...</h1>")
}

func main() {

	homeView = views.NewView("base", "views/home.gohtml")
	contactView = views.NewView("base", "views/contact.gohtml")

	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notFound)
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	http.ListenAndServe(":3000", r)
}
