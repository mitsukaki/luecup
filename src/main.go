package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/syndtr/goleveldb/leveldb"
)

// Datastore
var db *leveldb.DB

func main() {
	// open the database
	var err error
	db, err = leveldb.OpenFile("path/to/db", nil)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// create a new router
	r := mux.NewRouter()

	// serve public directory
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	// api endpoints
	r.HandleFunc("/api/fetch/{count:[0-9]+}/", handleFetch)
	r.HandleFunc("/api/tags/{tag}/", handleTag)

	// start server
	http.ListenAndServe(":8080", r)
}
