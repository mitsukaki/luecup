package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleTag(w http.ResponseWriter, r *http.Request) {
	// get tag from url
	vars := mux.Vars(r)
	tag := vars["tag"]

	// Handle with either get, put, or delete requests
	if r.Method == "GET" {
		HandleTagGet(w, r, tag)
	} else if r.Method == "PUT" {
		HandleTagPut(w, r, tag)
	} else if r.Method == "DELETE" {
		HandleTagDelete(w, r, tag)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func HandleTagGet(w http.ResponseWriter, r *http.Request, tag string) {
	data, err := db.Get([]byte(tag), nil)
	if err != nil {
		// 404 error
		http.Error(w, tag+" not found.", http.StatusNotFound)
		return
	}

	// encode data as JSON and write to response
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func HandleTagPut(w http.ResponseWriter, r *http.Request, tag string) {
	// get JSON body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// unmarshal JSON array body to check if valid JSON
	var dataArray []string
	if err := json.Unmarshal(body, &dataArray); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// save the body to the database
	if err := db.Put([]byte(tag), body, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// set status code to OK
	w.WriteHeader(http.StatusOK)
}

func HandleTagDelete(w http.ResponseWriter, r *http.Request, tag string) {
	if err := db.Delete([]byte(tag), nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// set status code to OK
	w.WriteHeader(http.StatusOK)
}
