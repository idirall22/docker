package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Hello golang"))
	})
	log.Println("Server listen on port :5000")
	if err := http.ListenAndServe(":5000", r); err != nil {
		log.Println("----------", err)
		os.Exit(1)
	}
}
