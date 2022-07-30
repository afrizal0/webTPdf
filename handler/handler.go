package handler

import (
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
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

func getHtmlFromUrl(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", errors.New("can't get content from url, please check the url")
	}
	defer resp.Body.Close()

	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("can't read body content from url")
	}
	return string(html), nil
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		return
	}

	urlWebsite := r.FormValue("url")

	if isValidUrl(urlWebsite) {
		fmt.Println(urlWebsite)
	} else {
		fmt.Errorf("Error: Website url wrong")
		http.Redirect(w, r, "/", 301)
		return
	}

	htmlCode, err := getHtmlFromUrl(urlWebsite)
	if err != nil {
		fmt.Errorf("Error Get Content: %w", err)
		return
	}
	fmt.Println(htmlCode)

}
