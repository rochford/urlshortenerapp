package handler

var errorHTML = "<!DOCTYPE html>" +
	"<html>" +
	"  <head>" +
	"    <meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\" />" +
	"    <title>URL Shortener - Error</title>" +
	"  </head>" +
	"  <body>" +
	"    An error was encountered while processing your request to shorten URL." +
	"    Error message: {{.}}." +
	"  </body>" +
	"</html>"
