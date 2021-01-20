package cyoa

import (
	"html/template"
	"log"
	"net/http"
)

// TemplateHandler doc
func TemplateHandler(adventure Adventure) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.ParseFiles("template.html"))
		err := t.Execute(w, adventure)
		if err != nil {
			log.Println(err)
		}
	}
}
