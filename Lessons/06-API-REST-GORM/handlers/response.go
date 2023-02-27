package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func sendData(rw http.ResponseWriter, data interface{}, status int) {
	// Headers
	rw.Header().Set("content-type", "application/json")
	rw.WriteHeader(status)

	// Respuesta
	output, _ := json.Marshal(&data)
	fmt.Fprintln(rw, string(output))
}

func sendError(rw http.ResponseWriter, status int) {
	rw.WriteHeader(status)                         // Header.
	fmt.Fprintln(rw, string("Resource not found")) // Respuesta.
}
