package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.FileServer(http.Dir("./public"))
	http.Handle("/", server)

	http.HandleFunc("/healths", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	err := http.ListenAndServe(":8080", http.DefaultServeMux)
	if err != nil {
		fmt.Print("Error while opening the server")
	}
}
