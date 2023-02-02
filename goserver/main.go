package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", getRoot)

	http.HandleFunc("/hello", getHello)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Printf("error starting goserver: %s \n", err)
	}
	fmt.Printf("starting goserver at port 8000 \n")
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "Welcome to the Planify API - 1.0\n")
	if err != nil {
		return
	}
}

func getHello(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "Hello, HTTP!\n")
	if err != nil {
		return
	}
}
