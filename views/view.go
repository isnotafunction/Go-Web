package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var LayoutDir string = "views/layouts/"
var LayoutExt string = ".gohtml"

func NewView(layout string, files ...string) *View {
	files = append(files, layoutFiles()...)

	t, err := template.ParseFiles(files...)

	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

type View struct {
	Template *template.Template
	Layout   string
}

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := v.Render(w, nil)
	if err != nil {
		panic(err)
	}
}

//Render method renders a view with predefined layout
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

//layoutFiles returns a slice of strings of layout file names
func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + LayoutExt)
	if err != nil {
		panic(err)
	}

	return files
}
