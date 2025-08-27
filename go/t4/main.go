package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		log.Print(r.URL)
		res := []byte("nothing at all")
		w.Write(res)
	})

	port := ":4242"
	err := http.ListenAndServe(port, nil)

	if err != nil {
		log.Fatal("error running server")
	}

}
