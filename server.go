package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"log"
	"github.com/gabrielalacchi/personal-site/app/controller"
	"github.com/gabrielalacchi/personal-site/app/routers"
)

func main() {

	indexHandler := controller.IndexHandler
	staticHandler := http.FileServer(http.Dir("frontend/www"))

	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)
	r.PathPrefix("/admin").Handler(http.StripPrefix("/admin", routers.AdminRouter()))
	r.NotFoundHandler = staticHandler

	srv := &http.Server{
		Handler: r,
		Addr: "localhost:3000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}