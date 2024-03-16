package main

import (
	// "juliano.com/passage-order/global"

	"juliano.com/passage-order/nav_bar/routes"
	"juliano.com/passage-order/notation"
	"juliano.com/passage-order/passage_order"

	// "juliano.com/passage-order/popover"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func main() {
	// global.Popover = popover.NewPopover()

	app.Route(routes.Root, &passage_order.PassageOrderPage{})
	app.Route(routes.PassageOrder, &passage_order.PassageOrderPage{})
	app.Route(routes.Notation, &notation.NotationPage{})

	app.RunWhenOnBrowser()

	generateDynamicWebsite()
	// generateStaticWebsite()
}
