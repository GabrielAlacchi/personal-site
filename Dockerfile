FROM golang:1.7

EXPOSE 3000

RUN go get github.com/gorilla/mux
RUN go get github.com/gorilla/sessions
RUN go get golang.org/x/crypto/bcrypt
RUN go get gopkg.in/mgo.v2

RUN mkdir -p /go/src/github.com/gabrielalacchi/personal-site

COPY . /go/src/github.com/gabrielalacchi/personal-site/

VOLUME /go/src/github.com/gabrielalacchi/personal-site/frontend/www/files

WORKDIR /go/src/github.com/gabrielalacchi/personal-site

RUN go install

ENTRYPOINT ["/go/bin/personal-site"]