package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rochford/urlshortener"

	"github.com/julienschmidt/httprouter"
)

func createPage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	bytes, err := ioutil.ReadFile("public/index.html")
	if err != nil {
		fmt.Println("error")
		// TODO: fix this
		return
	}
	fmt.Fprintln(w, string(bytes))
}

func process(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	fmt.Println(r.FormValue("originalUrl"))

	shortURL := urlshortener.GenerateShortURL(r.FormValue("originalUrl"))

	fmt.Fprintln(w, r.Form["originalUrl"], shortURL)
}

func resolveURL(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	url := urlshortener.ResolveShortURL(ps.ByName("id"))
	if url == "" {
		url = "http://www.wikipedia.org"
		// TOOO 501 error and page
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
