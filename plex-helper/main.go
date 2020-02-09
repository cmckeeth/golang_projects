package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getStatus(c *controller, r *http.Request) {
	c.controllerGet()
}

func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "You've made a post request"}`))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", getStatus).Methods(http.MethodGet)
	r.HandleFunc("/", post).Methods(http.MethodPost)
	r.HandleFunc("/", notFound)
	log.Fatal(http.ListenAndServe(":8080", r))
}
