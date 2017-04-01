package controller

import (
	"net/http"
	"github.com/gabrielalacchi/personal-site/app/shared"
	"html/template"
	"log"
	"github.com/gabrielalacchi/personal-site/app/model"
	"encoding/json"
)

var (
	adminTmpl *template.Template
)

func AdminHandler(res http.ResponseWriter, req *http.Request) {
	var err error

	session, _ := shared.Store.Get(req, "session")
	if session.Values["auth"] != "ok" {
		http.Redirect(res, req, "/admin/login", http.StatusTemporaryRedirect)
		return
	}

	if adminTmpl == nil {
		adminTmpl, err = template.New("index").ParseFiles("app/templates/admin.html")
		if err != nil {
			log.Fatal("[controller] [IndexHandler] Can't read template for index.html:", err.Error())
		}
	}

	var pagedata model.PageSource
	pagedata.Query()

	adminTmpl.ExecuteTemplate(res, "admin.html", pagedata)
}

func SaveHandler(res http.ResponseWriter, req *http.Request) {

	session, _ := shared.Store.Get(req, "session")
	if session.Values["auth"] != "ok" {
		http.Redirect(res, req, "/admin/login", http.StatusUnauthorized)
		return
	}

	if req.Method != "POST" {
		res.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var pagedata model.PageSource
	json.NewDecoder(req.Body).Decode(&pagedata)

	var pagedataQueried model.PageSource
	pagedataQueried.Query()

	pagedata.Id = pagedataQueried.Id
	err := pagedata.Update()

	if err != nil {
		log.Println("[controller] [admin] Save failed:", err.Error())
		res.WriteHeader(500)
		return
	}

	res.WriteHeader(200)
	return
}