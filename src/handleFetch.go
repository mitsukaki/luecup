package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func HandleFetch(w http.ResponseWriter, r *http.Request) {
	// get tag from url
	vars := mux.Vars(r)
	countStr := vars["count"]
	count, err := strconv.Atoi(countStr)

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

	// create string to int map
	itemCounts := make(map[string]int)

	// iterate over tags in the request
	for _, tag := range req.Tags {
		// get data from database
		data, err := db.Get([]byte(tag), nil)
		if err != nil {
			// 404 status
			http.Error(w, tag+"not found.", http.StatusNotFound)
			return
		}

		// parse the json array
		var dataArray []string
		err = json.Unmarshal(data, &dataArray)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// iterate over the array and add to map
		for _, item := range dataArray {
			itemCounts[item]++
		}
	}

	// sort the map by value
	sortedItems := make([]string, 0, len(itemCounts))
	for item := range itemCounts {
		sortedItems = append(sortedItems, item)
	}

	sort.Slice(sortedItems, func(i, j int) bool {
		return itemCounts[sortedItems[i]] > itemCounts[sortedItems[j]]
	})

	// make sure we don't pull more than we have
	if count > len(sortedItems) {
		count = len(sortedItems)
	}

	// respond with the count of items
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf("[%s]", strings.Join(sortedItems[:count], ","))))
}
