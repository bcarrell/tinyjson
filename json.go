package tinyjson

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func Write(w http.ResponseWriter, v interface{}) {
	payload, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(payload)))
	w.Write(payload)
}

// a handy one-liner to sugar up receiving JSON POST requests
func Read(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(&v)
}
