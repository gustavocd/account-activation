package templates

import (
	"html/template"
	"log"
	"net/http"
)

var t *template.Template

// Render ...
func Render(w http.ResponseWriter, name string, data interface{}) {
	t = template.Must(template.ParseGlob("templates/*.html"))
	if err := t.ExecuteTemplate(w, name, nil); err != nil {
		log.Fatal(err.Error())
	}
}
