package main

import (
	"log"
	"net/http"
	"strings"

	"juliano.com/passage-order/root"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func main() {
	app.Route("/", &root.Root{})
	app.RunWhenOnBrowser()

	// Wrap the app.Handler to check for WASM file requests
	wasmHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check the request URL path to determine the correct Content-Type.
		if strings.HasSuffix(r.URL.Path, ".wasm") {
			w.Header().Set("Content-Type", "application/wasm")
		} else if strings.HasSuffix(r.URL.Path, ".js") {
			w.Header().Set("Content-Type", "text/javascript")
		}

		// Use the existing app.Handler to serve the request
		// This ensures that non-WASM files are served normally
		handler := &app.Handler{
			Name:        "Hello",
			Description: "An Hello World! example",
			Scripts: []string{
				"/web/main.js",
			},
			Styles: []string{
				"/web/main.css",
			},
		}

		handler.ServeHTTP(w, r)
	})

	// Use the custom handler for serving requests
	http.Handle("/", wasmHandler)

	// Start the server
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
