package handler

var indexHTML = "<html>" +
	"  <head>" +
	"    <meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\" />" +
	"    <title>URL Shortener</title>" +
	"  </head>" +
	"  <body>" +
	"    <form action=\"/process\"" +
	"      method=\"post\"" +
	"      enctype=\"application/x-www-form-urlencoded\">" +
	"      <input type=\"text\" name=\"originalURL\" value=\"http://www.golang.org\"/>" +
	"      <input type=\"submit\"/>" +
	"    </form>" +
	"  </body>" +
	"</html>"
