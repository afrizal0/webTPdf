package handler

import (
	"fmt"
	"net/http"
	"text/template"

	. "github.com/afrizal0/webTPdf/utils"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseGlob("templates/*.html"))
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		return
	}

	urlWebsite := r.FormValue("url")

	if IsValidUrl(urlWebsite) {
		fmt.Println(urlWebsite)
	} else {
		fmt.Errorf("Error: Website url wrong")
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		return
	}

	GeneratePDF(urlWebsite)

}
