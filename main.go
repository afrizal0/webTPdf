package main

import (
	"net/http"

	. "github.com/afrizal0/webTPdf/handler"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)

	http.ListenAndServe(":8080", r)
}
