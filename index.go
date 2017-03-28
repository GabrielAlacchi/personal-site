package main

import (
	"net/http"
	"html/template"
	"encoding/json"
	"io/ioutil"
)

type Job struct {
	Company string
	Timeline string
	Title string
	Description string
}

type School struct {
	Name string
	Timeline string
	Major string
	DescriptionList []string
}

type Project struct {
	Name string
	Year int
	ImageURL string
	Description string
	LinkURL string
	LinkText string
	LinkIconClass string
	ShowLinkIcon bool
}

type Social struct {
	Link string
	IconClass string
}

type IndexFields struct {
	About string
	ProfilePicURL string
	ResumeURL string
	Experience []Job
	Education []School
	Projects []Project
	Skills []string
	SocialLinks []Social
	CopyrightYear int
}

func getHandler() http.HandlerFunc {
	t, err := template.New("index").ParseFiles("frontend/index.html")
	if err != nil {
		panic(err)
	}

	return func (res http.ResponseWriter, req *http.Request) {
		var index IndexFields
		bytes, err := ioutil.ReadFile("test_data.json")
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(bytes, &index)
		if err != nil {
			panic(err)
		}
		err = t.ExecuteTemplate(res, "index.html", index)
		if err != nil {
			panic(err.Error())
		}
	}
}