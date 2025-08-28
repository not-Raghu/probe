package main

import (
	"log"
	"net/http"
)

// func main() {
// 	cmd := exec.Command("cat", "main.go")
// 	out, err := cmd.Output()
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	fmt.Print(string(out))
// }

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/yes", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("yes route"))

	})

	mux.HandleFunc("/no", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("no route"))
	})

	mux.HandleFunc("/no/yes", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("no yes route"))
	})

	if http.ListenAndServe(":3000", mux) != nil {
		log.Fatal("fatal error runing the server ")
	}
}
