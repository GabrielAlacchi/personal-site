package controller

import (
	"net/http"
	"html/template"
	"log"
	"github.com/gabrielalacchi/personal-site/app/shared"
)

var (
	loginTmpl *template.Template
)

type loginData struct {
	Email string
	Failure bool
	Message string
}

func LoginHandler(res http.ResponseWriter, req *http.Request) {

	if req.Method == "GET" {
		loginPageGetHandler(res, req)
	} else if req.Method == "POST" {
		loginPagePostHandler(res, req)
	} else {
		res.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

}

func loginPageGetHandler(res http.ResponseWriter, req *http.Request) {
	var err error

	if loginTmpl == nil {
		loginTmpl, err = template.New("login").ParseFiles("app/templates/login.html")
		if err != nil {
			log.Fatal("[controller] [LoginPage] Cannot create login template")
		}
	}

	session, err := shared.Store.Get(req, "session")
	if session.Values["auth"] == "ok" {
		http.Redirect(res, req, "/admin/", http.StatusTemporaryRedirect)
		return
	}

	loginTmpl.ExecuteTemplate(res, "login.html", &loginData{
		Failure: false,
	})
}

func loginPagePostHandler(res http.ResponseWriter, req *http.Request) {
	var err error

	if loginTmpl == nil {
		loginTmpl, err = template.New("login").ParseFiles("app/templates/login.html")
		if err != nil {
			log.Fatal("[controller] [LoginPage] Cannot create login template")
		}
	}

	email := req.PostFormValue("email")
	pwd := req.PostFormValue("password")

	if auth, err := shared.Authenticate(email, pwd); err != nil {
		log.Println("[controller] [login] Auth error:", err.Error())
	} else if auth {
		session, _ := shared.Store.New(req, "session")
		session.Values["auth"] = "ok"
		session.Save(req, res)
		http.Redirect(res, req, "/", http.StatusTemporaryRedirect)
		return
	}

	loginTmpl.ExecuteTemplate(res, "login.html", &loginData{
		Email: email,
		Failure: true,
		Message: "How about not tampering with my admin panel bro. Where are your fucking manners?",
	})

}