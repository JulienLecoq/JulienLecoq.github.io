package main

import (
	"log"
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// Command to execute to run with this:
// GOARCH=wasm GOOS=js go build -o web/app.wasm && go build && ./passage-order
func generateDynamicWebsite() {
	http.Handle("/", &app.Handler{
		Name:        "Ordre de passage",
		Description: "",
		Scripts: []string{
			"/web/main.js",
		},
		Styles: []string{
			"/web/main.css",
		},
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
