package controller

import (
	"html/template"
	"fmt"
	"net/http"
	"github.com/gabrielalacchi/personal-site/app/model"
	"log"
)

var (
	indexTmpl *template.Template
)

func IndexHandler(res http.ResponseWriter, req *http.Request) {
	var index model.PageSource
	var err error

	if indexTmpl == nil {
		indexTmpl, err = template.New("index").ParseFiles("app/templates/index.html")
		if err != nil {
			log.Fatal("[controller] [IndexHandler] Can't read template for index.html:", err.Error())
		}
	}

	// Fetch the index data
	err = index.Query()
	if err != nil {
		fmt.Println("[controller] [IndexHandler] Querying index data failed: ", err.Error())
		res.WriteHeader(500)
		return
	}

	err = indexTmpl.ExecuteTemplate(res, "index.html", index)
	if err != nil {
		fmt.Println("[controller] [IndexHandler] Template failed to execute: ", err.Error())
		res.WriteHeader(500)
		return
	}

}