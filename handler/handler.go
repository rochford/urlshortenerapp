package handler

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/rochford/urlshortener"

	"github.com/julienschmidt/httprouter"
)

// errorPage function sends a HTML error page using errorMsg
func errorPage(httpCode int, errorMsg string, w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(httpCode)
	tmpl := template.New("name")
	t, err := tmpl.Parse(errorHTML)
	if err != nil {
		bytes := bytes.NewBufferString(errorHTML)
		fmt.Fprintln(w, string(bytes.String()))
		return
	}
	t.Execute(w, errorMsg)
}

func createPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	bytes := bytes.NewBufferString(indexHTML)
	fmt.Fprintln(w, string(bytes.String()))
}

func process(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	r.ParseForm()
	originalURL := r.FormValue("originalURL")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()
	ctx = context.WithValue(ctx, "originalURL", originalURL)

	shortURL, err := urlshortener.GenerateShortURL(ctx)

	if err != nil {
		errorPage(500, err.Error(), w, r, ps)
		return
	}

	tmpl := template.New("name")
	t, err := tmpl.Parse(resultHTML)
	if err != nil {
		errorPage(500, err.Error(), w, r, ps)
		return
	}
	mp := make(map[string]string)
	mp["originalURL"] = originalURL
	mp["shortURL"] = shortURL
	t.Execute(w, mp)
}

func resolveURL(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()

	ctx = context.WithValue(ctx, "id", ps.ByName("id"))
	url, err := urlshortener.ResolveShortURL(ctx)

	if err != nil {
		errorPage(500, err.Error(), w, r, ps)
		return
	}
	if url == "" {
		errorPage(400, "Bad request message or unknown shortURL", w, r, ps)
		return
	}
	w.Header().Set("Location", url)
	w.WriteHeader(302)
}

// Routes function sets the HTTP API endpoints.
func Routes() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", createPage)
	router.POST("/process", process)
	router.GET("/:id", resolveURL)
	return router
}
