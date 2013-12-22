# tinyjson

This is a tiny Go package for reducing repetition and headache in handling JSON.

**Installation**: `go get github.com/bcarrell/tinyjson`

### Example

	package main

	import (
		"net/http"
		"strings"

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
		switch r.Method {
		case "GET": // demonstrates tinyjson.Write
			tricia := &person{
				Name:    "Cylon Number Six",
				Gender:  "Female",
				IsHuman: false,
			}
			tinyjson.Write(w, tricia)
		case "POST": // demonstrates tinyjson.Read
			var person person
			tinyjson.Read(r, &person)
			person.Name = strings.ToLower(person.Name)
			person.Gender = strings.ToLower(person.Gender)
			tinyjson.Write(w, &person)
		}
	}

<!-- separate -->

	go run main.go

<!-- separate -->

	❯ curl -i localhost:3000              -- INS --
	HTTP/1.1 200 OK
	Content-Length: 62
	Content-Type: application/json
	Date: Sun, 22 Dec 2013 13:05:49 GMT

	{"name":"Cylon Number Six","gender":"Female","is_human":false}

<!-- separate -->

	❯ curl -X POST -H "Content-Type: application/json" -d '{"name":"Gaius Baltar","gender":"Male","is_human":true}' http://localhost:3000

	{"name":"gaius baltar","gender":"male","is_human":true}