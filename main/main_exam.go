package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// 1) Struct for a Route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// 2) Array of Route
type Routes []Route

// 3) our first handler
func getIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}

// 4) initiate the router
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	// declare routes
	var routes = Routes{
		Route{
			"getIndex",
			"GET",
			"/",
			getIndex,
		},
		Route{
			"getJson",
			"GET",
			"/json",
			getJson,
		},
	}
	// bind routes
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}

// 5) init main func
func main() {
	router := NewRouter()
	log.Fatal(
		// start on port 3000 by default
		http.ListenAndServe(":3000", router),
	)
}

// 6) Basic response struct
type BasicResponse struct {
	Error   int    `json:"error"`
	Message string `json:"message"`
}

// 7) our first json response handler
func getJson(w http.ResponseWriter, r *http.Request) {
	// set JSON response header
	w.Header().Set("Content-type", "application/json; charset=UTF-8;")
	json.NewEncoder(w).Encode(BasicResponse{
		0,
		"Hello JSON!",
	})
}
