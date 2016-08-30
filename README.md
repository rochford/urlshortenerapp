# urlshortenerapp

urlshortenerapp/
  Dockerfile - for this app. Run it like this:
    # docker run --publish 80:8080 --name urlshortenerapp  --rm urlshortenerapp
  main.go - contains minimal executable functionality. Majority of logic is in
    the handler to make it easier to test.
urlshortenerapp/handler
  handler.go - contains route handling logic.  extracts information from HTTP
    requests and forms http responses.

  static assets
    indexhtml.go - simple page with a html form to HTTP::POST a URL to shorten.
    resulthtml.go - resulting page including the shortened URL.
    errorhtml.go - page if cannot find shortURL or other errors
urlshortener/
  urlshortener.go - Stores data in a Map.

Design:
Each golang http request is handled in its own goroutine so shared access to
variables is not safe. E.g. http request handling routines using shared variables
https://nvisium.com/blog/2015/07/16/golang-security-and-concurrency/

A HTTP context is used to pass parameters and Cancel channels to HTTP request
processing functions (and any new goroutines).
See https://blog.golang.org/context
