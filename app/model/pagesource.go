package model

import (
	"html/template"
	"time"
)

type Job struct {
	Company string `json:"Company" bson:"Company"`
	Timeline string `json:"Timeline" bson:"Timeline"`
	Title string `json:"Title" bson:"Title"`
	Description string `json:"Description" bson:"Description"`
}

type School struct {
	Name string `json:"Name" bson:"Name"`
	Timeline string `json:"Timeline" bson:"Timeline"`
	Major string `json:"Major" bson:"Major"`
	DescriptionList []string `json:"DescriptionList" bson:"DescriptionList"`
}

type Project struct {
	Name string `json:"Name" bson:"Name"`
	Year int `json:"Year" bson:"Year"`
	ImageURL string `json:"ImageURL" bson:"ImageURL"`
	Description string `json:"Description" bson:"Description"`
	Body template.HTML `json:"Body" bson:"Body"`
	LinkURL string `json:"LinkURL" bson:"LinkURL"`
	LinkText string `json:"LinkText" bson:"LinkText"`
	LinkIconClass string `json:"LinkIconClass" bson:"LinkIconClass"`
	ShowLinkIcon bool `json:"ShowLinkIcon" bson:"ShowLinkIcon"`
}

type Social struct {
	Link string `json:"Link" bson:"Link"`
	IconClass string `json:"IconClass" bson:"IconClass"`
}

type PageSource struct {
	About string `json:"About" bson:"About"`
	ProfilePicURL string `json:"ProfilePicURL" bson:"ProfilePicURL"`
	ResumeURL string `json:"ResumeURL" bson:"ResumeURL"`
	Experience []Job `json:"Experience" bson:"Experience"`
	Education []School `json:"Education" bson:"Education"`
	Projects []Project `json:"Projects" bson:"Projects"`
	Skills []string `json:"Skills" bson:"Skills"`
	SocialLinks []Social `json:"SocialLinks" bson:"SocialLinks"`
	CopyrightYear int `json:"CopyrightYear" bson:"CopyrightYear"`
}

func (index *PageSource) Query() error {
	sess := getSession()

	pagedata := sess.DB("alacchi-com").C("pagedata")
	err := pagedata.Find(nil).One(index)

	for i := 0; i < len(index.Projects); i++ {
		index.Projects[i].Body = template.HTML(index.Projects[i].Description)
	}

	index.CopyrightYear = time.Now().Year()
	return err
}
