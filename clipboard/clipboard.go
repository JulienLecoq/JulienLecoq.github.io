package clipboard

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func Paste(text string) {
	app.Window().
		Call("copyToClipboard", text)
	// Get("document").
	// Call("getElementById", "YOUR_ID")
}
