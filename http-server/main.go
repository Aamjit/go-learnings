package main

import (
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/", homeResponse)
	// http.HandleFunc("/hello", getHello)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		println("Failed to start server")
	}
}

func homeResponse(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, "This is my website!\n")
	// _, err := w.Write([]byte("Got response"))
	// return err
}
