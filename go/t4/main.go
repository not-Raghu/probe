package main

import (
	"fmt"
	"net/http"
	"os"
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

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		args := os.Args
		fmt.Print(args)
		fmt.Print(args[0])
		fmt.Print(args[1])
		fmt.Print(args[2])
		w.Write([]byte(args[0] + " " + args[1] + " " + args[2]))

	})

	mux.HandleFunc("/yes", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("yes route"))

	})

	mux.HandleFunc("/no", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("first no route"))
		w.Write([]byte("second no route"))
	})

	mux.HandleFunc("/no/yes", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("no yes route"))
	})

	go main2()
	http.ListenAndServe(":3000", mux)

}

func main2() {
	mux2 := http.NewServeMux()
	mux2.HandleFunc("/mux2", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("this is mux 2 route"))
	})
	http.ListenAndServe(":3001", mux2)
}
