package main

import (
	"io"
	"net/http"
	"planify-api/Controllers"
)

func main() {
	// Welcome
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := io.WriteString(w, "Welcome to the Planify API - 1.0\n")
		if err != nil {
			return
		}
	})

	// Appointments
	http.HandleFunc("/appointments", Controllers.GetAppointments)

}
