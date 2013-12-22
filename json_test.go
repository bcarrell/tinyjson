package tinyjson

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type response struct {
	ok  bool
	msg string
}

func writeHandler(w http.ResponseWriter, r *http.Request) {

}

func TestWrite(t *testing.T) {
	ts := httptest.NewServer(writeHandler)
	defer ts.Close()

}
