package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from ClipVault"))
}

func clipView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Displa a specific clip..."))
}

func clipCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new clip..."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/clip/view", clipView)
	mux.HandleFunc("/clip/create", clipCreate)

	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
