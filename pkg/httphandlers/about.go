package httphandlers

import (
	"html/template"
	"net/http"
)

// GetAboutHandler displays the "About" page
func GetAboutHandler(w http.ResponseWriter, r *http.Request) {

	page := template.Must(template.ParseFiles(
		"static/html/_base.html",
		"static/html/about.html",
	))

	page.Execute(w, nil)

}
