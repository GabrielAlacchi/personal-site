FROM golang:1.7

EXPOSE 3000

RUN mkdir -p /go/src/github.com/gabrielalacchi/personal-site

COPY . /go/src/github.com/gabrielalacchi/personal-site/

VOLUME /go/src/github.com/gabrielalacchi/personal-site/frontend/www/files

WORKDIR /go/src/github.com/gabrielalacchi/personal-site

RUN go get github.com/gorilla/mux \
           github.com/gorilla/sessions \
           golang.org/x/crypto/bcrypt \
           gopkg.in/mgo.v2

RUN go install

ENV MONGO_URL mongodb://database:27017

ENTRYPOINT ["/go/bin/personal-site"]