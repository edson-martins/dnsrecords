package controllers

import (
	"html/template"
	"net/http"

	"github.com/edson-martins/dnsrecords/constants"
)

// - ------------------------------------------------------------------------------------------------------------------
// - Root handler used to show the index html page
// - ------------------------------------------------------------------------------------------------------------------
func RootController(w http.ResponseWriter, r *http.Request) {

	log.Debug("RootController started and redirecting to the index page")

	t, _ := template.ParseFiles(constants.INDEX)
	t.Execute(w, nil)
}
