package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func handleTag(w http.ResponseWriter, r *http.Request) {
	// get tag from url
	vars := mux.Vars(r)
	tag := vars["tag"]

	// handle with either get, put, or delete requests
	if r.Method == "GET" {
		handleTagGet(w, r, tag)
	} else if r.Method == "PUT" {
		handleTagPut(w, r, tag)
	} else if r.Method == "DELETE" {
		handleTagDelete(w, r, tag)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleTagGet(w http.ResponseWriter, r *http.Request, tag string) {
	data, err := db.Get([]byte(tag), nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// encode data as JSON and write to response
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func handleTagPut(w http.ResponseWriter, r *http.Request, tag string) {
	// get JSON body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// unmarshal JSON body
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// encode data as JSON and write to database
	dataBytes, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := db.Put([]byte(tag), dataBytes, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// set status code to OK
	w.WriteHeader(http.StatusOK)
}

func handleTagDelete(w http.ResponseWriter, r *http.Request, tag string) {
	if err := db.Delete([]byte(tag), nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// set status code to OK
	w.WriteHeader(http.StatusOK)
}
