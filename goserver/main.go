package main

import (
	"fmt"
	"net/http"
)

func main() {

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Printf("error starting goserver: %s \n", err)
	}
	fmt.Printf("starting goserver at port 8000 \n")
}
