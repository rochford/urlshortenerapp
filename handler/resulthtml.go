package handler

var resultHTML = "<!DOCTYPE html>" +
	"<html>" +
	"  <head>" +
	"    <meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\" />" +
	"    <title>URL Shortener - Result</title>" +
	"  </head>" +
	"  <body>" +
	"    The short URL is <a href=\"/{{.shortURL}}\">{{.shortURL}}</a>" +
	"    <br>" +
	"    The short URL will redirect you to <a href=\"{{.originalURL}}\">{{.originalURL}}</a>" +
	"  </body>" +
	"</html>"
