package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func handleFetch(w http.ResponseWriter, r *http.Request) {
	// get JSON body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// unmarshal JSON body
	var req FetchRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
