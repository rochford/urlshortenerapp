package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/rochford/urlshortener"

	"github.com/julienschmidt/httprouter"
)

// errorPage function sends a HTML error page using errorMsg
func errorPage(httpCode int, errorMsg string, w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(httpCode)
	t, err := template.ParseFiles("public/error.html")
	if err != nil {
		bytes, _ := ioutil.ReadFile("public/error.html")
		fmt.Fprintln(w, string(bytes))
		return
	}
	t.Execute(w, errorMsg)
}

func createPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	bytes, err := ioutil.ReadFile("public/index.html")
	if err != nil {
		errorPage(500, err.Error(), w, r, ps)
		return
	}
	fmt.Fprintln(w, string(bytes))
}

func process(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	r.ParseForm()
	originalURL := r.FormValue("originalUrl")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()
	ctx = context.WithValue(ctx, "originalUrl", originalURL)

	shortURL, err := urlshortener.GenerateShortURL(ctx)

	if err != nil {
		errorPage(500, err.Error(), w, r, ps)
		return
	}

	t, err := template.ParseFiles("public/result.html")
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
	fmt.Println(url)
	w.Header().Set("Location", url)
	w.WriteHeader(302)
}

func main() {
	router := httprouter.New()
	router.GET("/", createPage)
	router.POST("/process", process)
	router.GET("/:id", resolveURL)

	http.ListenAndServe(":8080", router)
}
