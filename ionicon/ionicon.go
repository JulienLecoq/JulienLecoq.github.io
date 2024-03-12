package ionicon

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Ionicon struct {
	app.Compo

	Name string
}

func (i *Ionicon) Render() app.UI {
	return app.Raw(fmt.Sprintf("<ion-icon name=\"%s\"></ion-icon>", i.Name))
}
