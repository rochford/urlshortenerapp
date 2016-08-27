package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"

	"github.com/rochford/urlshortenerapp/handler"
)

var router *httprouter.Router

func init() {
	router = handler.Routes()
}

func TestHomePageWorking(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Errorf("http.NewRequest returned an error: ", err.Error())
	}

	rw := httptest.NewRecorder()
	router.ServeHTTP(rw, req)
	if rw.Code != 200 {
		t.Fatalf("HTTP response code must be 200 for the homepage", rw.Code)
	}
	t.Logf("All done")
}

// TODO: more tests...
