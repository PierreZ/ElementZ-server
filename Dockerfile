# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang
MAINTAINER PierreZ

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/pierrez/ElementZ-server

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go get github.com/martini-contrib/binding
RUN go get github.com/go-martini/martini
RUN go install github.com/github.com/pierrez/ElementZ-server

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/ElementZ-server

# Document that the service listens on port 8080.
EXPOSE 3000