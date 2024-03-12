package main

import (
	"log"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func generateStaticWebsite() {
	err := app.GenerateStaticWebsite(".", &app.Handler{
		Name:        "Ordre de passage",
		Description: "",
		Resources:   app.GitHubPages("passage-order"),
		Scripts: []string{
			"/web/main.js",
		},
		Styles: []string{
			"/web/main.css",
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
