package routers

import (
	"github.com/gorilla/mux"
	"github.com/gabrielalacchi/personal-site/app/controller"
	"net/http"
	"github.com/gabrielalacchi/personal-site/app/shared"
	"log"
	"os"
	"path"
	"io"
	"html/template"
	"io/ioutil"
)

func AdminRouter() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/", controller.AdminHandler)
	r.HandleFunc("/login", controller.LoginHandler)
	r.HandleFunc("/logoff", logoffHandler)
	r.HandleFunc("/upload", uploadHandler)
	r.HandleFunc("/save", controller.SaveHandler)

	return r
}

func logoffHandler(res http.ResponseWriter, req *http.Request) {
	session, _ := shared.Store.Get(req, "session")
	session.Options.MaxAge = -1
	session.Save(req, res)
	http.Redirect(res, req, "/admin/login", http.StatusTemporaryRedirect)
}

type uploadData struct {
	Files []string
}

func uploadHandler(res http.ResponseWriter, req *http.Request) {

	session, _ := shared.Store.Get(req, "session")
	if session.Values["auth"] != "ok" {
		http.Redirect(res, req, "/admin/login", http.StatusUnauthorized)
		return
	}

	t, _ := template.New("upload").ParseFiles("app/templates/upload.html")

	if req.Method == "POST" {

		req.ParseMultipartForm(32 << 20)
		filename := req.FormValue("filename")
		file, _, err := req.FormFile("uploadfile")

		if err != nil {
			log.Println("[routers] [admin] Upload failed", err.Error())
			res.WriteHeader(500)
			return
		}
		defer file.Close()

		filename = path.Base(filename)

		filepath := path.Join("frontend", "www", "files", filename)

		stream, err := os.OpenFile(filepath, os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0777)
		if err != nil {
			log.Println("[routers] [admin] Upload failed", err.Error())
			res.WriteHeader(500)
			return
		}
		defer stream.Close()

		io.Copy(stream, file)

		files, _ := ioutil.ReadDir("frontend/www/files")

		var data uploadData

		for _, file := range files {
			if !file.IsDir() {
				data.Files = append(data.Files, file.Name())
			}
		}

		t.ExecuteTemplate(res, "upload.html", data)
	} else if req.Method == "GET" {
		files, _ := ioutil.ReadDir("frontend/www/files")

		var data uploadData

		for _, file := range files {
			if !file.IsDir() {
				data.Files = append(data.Files, file.Name())
			}
		}

		t.ExecuteTemplate(res, "upload.html", data)
	} else {
		res.WriteHeader(http.StatusMethodNotAllowed)
	}

}