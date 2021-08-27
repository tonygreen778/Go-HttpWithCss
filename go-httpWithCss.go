package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HomeHandler(rw http.ResponseWriter, r *http.Request) {
	http.ServeFile(rw, r, "templates/index.html")
}

func main() {
	r := mux.NewRouter()

	cssHandler := http.FileServer(http.Dir("./templates/css/"))

	http.Handle("/css/", http.StripPrefix("/css", cssHandler))
	r.HandleFunc("/", HomeHandler)
	http.Handle("/", r)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
