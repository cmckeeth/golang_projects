package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)


func get() {
	controller.get()
}

func post() {
	controller.post()
}


func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", get).Methods(http.MethodGet)
	r.HandleFunc("/", post).Methods(http.MethodPost)
	r.HandleFunc("/", notFound)
	log.Fatal(http.ListenAndServe(":8090", r))
}
