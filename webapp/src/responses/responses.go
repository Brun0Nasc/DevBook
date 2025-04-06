package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// APIError represents the error response
type APIError struct {
	Error string `json:"error"`
}

// JSON returns a json response to the request
func JSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}

}

// StatusCodeError handler the requests with status code greater than 400
func StatusCodeError(w http.ResponseWriter, r *http.Response) {
	var err APIError
	json.NewDecoder(r.Body).Decode(&err)
	JSON(w, r.StatusCode, err)
}
