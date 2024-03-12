package main

import (
	// "log"
	// "net/http"
	// "strings"

	"juliano.com/passage-order/root"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func main() {
	app.Route("/", &root.Root{})
	app.RunWhenOnBrowser()

	// generateDynamicWebsite()
	generateStaticWebsite()
}
