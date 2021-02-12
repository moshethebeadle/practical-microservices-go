package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("hi")

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("OK")); err != nil {
			log.Fatalf("error writing bytes: %v", err)
		}
		w.WriteHeader(200)
	}))

	if err := http.ListenAndServe(":8080", http.DefaultServeMux); err != nil {
		log.Fatalf("listening and serving: %v", err)
	}
}