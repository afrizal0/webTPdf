package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func isValidUrl(urlWebsite string) bool {
	_, err := url.ParseRequestURI(urlWebsite)
	if err != nil {
		return false
	} else {
		return true
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", 301)
		return
	}

	urlWebsite := r.FormValue("url")
	if isValidUrl(urlWebsite) {
		fmt.Println(urlWebsite)
	} else {
		fmt.Errorf("Error: Website url wrong")
		http.Redirect(w, r, "/", 301)
	}
}
