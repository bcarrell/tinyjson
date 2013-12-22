package tinyjson

import (
	"encoding/json"
	"net/http"
)

func Write(w http.ResponseWriter, v interface{}) {
	payload, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func Read(r *http.Request, v interface{}) error {

}
