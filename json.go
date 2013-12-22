package tinyjson

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// Write marshals JSON, sets appropriate headers, and writes to a response
// If an error happens, we'll return a 500 error
func Write(w http.ResponseWriter, v interface{}) {
	payload, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(payload)))
	w.Write(payload)
}

// Read provides some sugar for decoding a JSON payload from a HTTP request
func Read(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(&v)
}
