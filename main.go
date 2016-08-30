package main

import (
	"net/http"

	"github.com/rochford/urlshortenerapp/handler"
)

func main() {
	router := handler.Routes()
	http.ListenAndServe(":8080", router)
}
