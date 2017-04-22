FROM golang:1.8

# Install all the dependencies
RUN go get github.com/gorilla/mux
RUN go get github.com/gorilla/sessions
RUN go get golang.org/x/crypto/bcrypt
RUN go get gopkg.in/mgo.v2

# Copy and compile the code to a static binary
COPY . /go/src/github.com/gabrielalacchi/personal-site/
WORKDIR /go/src/github.com/gabrielalacchi/personal-site/

ENV GOOS=linux
ENV CGO_ENABLED=0
RUN go build -v -ldflags "-s" -a -installsuffix cgo -o /go/bin/personalsite