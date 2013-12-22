package tinyjson

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

type Response struct {
	Ok  bool
	Msg string
}

type Request struct {
	Name    string `json:"name"`
	Gender  string `json:"gender"`
	IsHuman bool   `json:"is_human"`
}

// helper -- thanks Martini!
func expect(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf(
			"Expected %v (type %v) - Got %v (type %v)",
			b,
			reflect.TypeOf(b),
			a,
			reflect.TypeOf(a),
		)
	}
}

func writeHandler(w http.ResponseWriter, r *http.Request) {
	res := Response{
		Ok:  true,
		Msg: "Success!",
	}
	Write(w, &res)
}

func TestWrite(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(writeHandler))
	defer ts.Close()
	var tresponse Response
	res, _ := http.Get(ts.URL)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &tresponse)

	// --- assertions ---
	expect(t, res.Header.Get("Content-Type"), "application/json")
	expect(t, tresponse.Ok, true)
	expect(t, tresponse.Msg, "Success!")
	// --- end assertions ---
}

func TestRead(t *testing.T) {
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var req Request
			Read(r, &req)

			// --- assertions ---
			expect(t, req.Name, "Cylon Number Six")
			expect(t, req.Gender, "Female")
			expect(t, req.IsHuman, false)
			// --- end assertions ---
		}))
	defer ts.Close()
	json := `{"name":"Cylon Number Six", "gender": "Female", "is_human": false}`
	body := strings.NewReader(json)
	http.Post(ts.URL, "application/json", body)
}

func TestGet(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(writeHandler))
	defer ts.Close()
	var tresponse Response
	Get(ts.URL, &tresponse)

	// --- assertions ---
	expect(t, tresponse.Ok, true)
	expect(t, tresponse.Msg, "Success!")
	// --- end assertions ---
}
