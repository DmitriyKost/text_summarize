package pkg

import (
	"net/http"
	"text/template"
)


var templates = template.Must(template.ParseGlob("static/*.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, data any) {
    err := templates.ExecuteTemplate(w, tmpl, data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}
