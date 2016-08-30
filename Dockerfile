# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD handler/ /go/src/github.com/rochford/urlshortenerapp/handler/
ADD . /go/src/github.com/rochford/urlshortenerapp/

# Dependencies
RUN go get github.com/rochford/urlshortener
RUN go get github.com/julienschmidt/httprouter

# Build the application inside the container.
RUN go install github.com/rochford/urlshortenerapp

# Run the command by default when the container starts.
ENTRYPOINT /go/bin/urlshortenerapp

# Document that the service listens on port 8080.
EXPOSE 8080
