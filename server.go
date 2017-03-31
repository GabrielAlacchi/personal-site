package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"log"
	"github.com/gabrielalacchi/personal-site/app/controller"
)

func main() {

	indexHandler := controller.IndexHandler
	staticHandler := http.FileServer(http.Dir("frontend/www"))

	r := mux.NewRouter()
	r.PathPrefix("/").HandlerFunc(func (res http.ResponseWriter, req *http.Request) {

		if req.URL.Path == "/" || req.URL.Path == "/index.html" {
			indexHandler(res, req)
		} else {
			staticHandler.ServeHTTP(res, req)
		}

	})

	srv := &http.Server{
		Handler: r,
		Addr: "localhost:3000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}