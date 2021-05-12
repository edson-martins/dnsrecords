package controllers

import (
	"html/template"
	"net/http"

	"github.com/edson-martins/dnsrecords/constants"
	"github.com/edson-martins/dnsrecords/network"
)

// - ------------------------------------------------------------------------------------------------------------------
// - Search handler used to manage the search form POST
// - ------------------------------------------------------------------------------------------------------------------
func SearchController(w http.ResponseWriter, r *http.Request) {

	// - --------------------------------------------------------------------------------------------------------------
	// - Retrieve the HTML form parameter of POST method
	// - --------------------------------------------------------------------------------------------------------------
	url := r.FormValue("entry-domain")

	log.Debugf("SearchController started to research the IP and MX data from %s domain", url)

	// - --------------------------------------------------------------------------------------------------------------
	// - Get the IP and MX data based on the URL parameter
	// - --------------------------------------------------------------------------------------------------------------
	Dns, err := network.GetDns(url)
	if err != nil {
		log.Errorf("The DNS get function failed with error %s and will return to the web client the error %v",
			err.Error(), http.StatusInternalServerError)

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t, _ := template.ParseFiles(constants.INDEX)
	t.Execute(w, Dns)
}
