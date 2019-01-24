package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

// helloworkd returns a simple HTTP handler function which writes a response.
func helloWorld() http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		info := struct {
			Message string `json:"Message"`
		}{
			"Hello, World",
		}

		body, err := json.Marshal(info)
		if err != nil {
			log.Printf("Could not encode info data: %v", err)
			http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	}
}
