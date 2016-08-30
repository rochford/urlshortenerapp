package handler_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
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

func TestPostURL(t *testing.T) {
	expectedOriginalUrl := "http://www.golang.org"
	localhost := "The short URL is <a href=\"/"

	v := url.Values{}
	v.Set("originalURL", expectedOriginalUrl)
	postData := v.Encode()

	req, err := http.NewRequest("POST", "/process", bytes.NewBufferString(postData))
	if err != nil {
		t.Errorf("http.NewRequest returned an error: ", err.Error())
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(postData)))

	rw := httptest.NewRecorder()
	router.ServeHTTP(rw, req)
	if rw.Code != 200 {
		t.Fatalf("HTTP response code must be 200 for the homepage", rw.Code)
	}
	pos := strings.Index(rw.Body.String(), localhost)
	if pos < 0 {
		t.Fatalf("shortURL not found")
	}
	pos = pos + len(localhost)
	shortURL := rw.Body.String()[pos : pos+6]

	req, err = http.NewRequest("GET", "/"+shortURL, nil)
	if err != nil {
		t.Errorf("http.NewRequest returned an error: ", err.Error())
	}
	rw = httptest.NewRecorder()
	router.ServeHTTP(rw, req)
	if rw.Code != 302 {
		t.Fatalf("HTTP response code must be 200 for the homepage", rw.Code)
	}
	if rw.Header().Get("Location") != expectedOriginalUrl {
		t.Fatalf("Incorrect Location field in HTTP redirect",
			rw.Header().Get("Location"))
	}

	t.Logf("All done")
}

func TestErrorHTML(t *testing.T) {
	req, err := http.NewRequest("GET", "/NOTEXIST", nil)
	if err != nil {
		t.Errorf("http.NewRequest returned an error: ", err.Error())
	}
	rw := httptest.NewRecorder()
	router.ServeHTTP(rw, req)
	if rw.Code != 500 {
		t.Fatalf("HTTP response code must be 500 for unknown shortURL", rw.Code)
	}
}

// TODO: more tests...
