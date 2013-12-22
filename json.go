package tinyjson

import (
	"encoding/json"
	"io/ioutil"
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
	return json.NewDecoder(r.Body).Decode(v)
}

// Get provides a neat little wrapper for making GET requests to endpoints
// where you expect a JSON payload in response to save effort unmarshaling
// the response
func Get(url string, v interface{}) (*http.Response, error) {
	res, err := http.Get(url)
	defer res.Body.Close()
	if err != nil {
		return res, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(body, v)
	return res, err
}
