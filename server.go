package main

import (
	"github.com/go-martini/martini"
)

func main() {
	m := martini.Classic()
	m.Get("/", getHandler())

	css := martini.Static("frontend/css", martini.StaticOptions {
		Prefix: "/css",
	})
	m.Use(css)

	images := martini.Static("frontend/images", martini.StaticOptions{
		Prefix: "/images",
	})
	m.Use(images)

	libs := martini.Static("frontend/libs", martini.StaticOptions{
		Prefix: "/libs",
	})
	m.Use(libs)

	js := martini.Static("frontend/js", martini.StaticOptions{
		Prefix: "/js",
	})
	m.Use(js)

	favicon := martini.Static("frontend/favicon.ico", martini.StaticOptions{
		Prefix: "/favicon.ico",
	})
	m.Use(favicon)

	m.Run()
}