# tinyjson

This is a tiny Go package for reducing repetition and headache in handling JSON.

**Installation**: `go get github.com/bcarrell/tinyjson`

More instructions coming soon.

### Examples

##### Writing JSON to `http.ResponseWriter`

	package main

	import (
		"net/http"

		"github.com/bcarrell/tinyjson"
	)

	type person struct {
		Name    string `json:"name"`
		Gender  string `json:"gender"`
		IsHuman bool   `json:"is_human"`
	}

	func main() {
		http.HandleFunc("/", foo)
		http.ListenAndServe(":3000", nil)
	}

	func foo(w http.ResponseWriter, r *http.Request) {
		tricia := &person{
			Name:    "Cylon Number Six",
			Gender:  "Female",
			IsHuman: false,
		}
		tinyjson.Write(w, tricia)
	}

`go run main.go`

	❯ curl -i localhost:3000              -- INS --
	HTTP/1.1 200 OK
	Content-Length: 62
	Content-Type: application/json
	Date: Sun, 22 Dec 2013 13:05:49 GMT

	{"name":"Cylon Number Six","gender":"Female","is_human":false}