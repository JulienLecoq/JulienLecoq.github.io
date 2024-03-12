package passage_order_table

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"juliano.com/passage-order/clipboard"
	"juliano.com/passage-order/component"
	"juliano.com/passage-order/popover"
	"juliano.com/passage-order/user"
)

func (p *PassageOrderTable) thClickItems() []popover.Item {
	return []popover.Item{
		popover.NewPopoverItem(
			"Randomiser",
			func(ctx app.Context, event app.Event) {
				p.Group.RandomizeUsers()
				p.Update()
			},
			nil,
		),
		popover.NewPopoverItem(
			"Trier par ordre alphabétique",
			func(ctx app.Context, event app.Event) {
				p.Group.SortAlphabetically()
				p.Update()
			},
			nil,
		),
		// popover.NewPopoverItem(
		// 	"Exporter",
		// 	func(ctx app.Context, event app.Event) {
		// 		html := component.HtmlOf(p)
		// 		clipboard.Paste(html)
		// 	},
		// 	nil,
		// ),
		popover.NewPopoverItem(
			"Copier les adresses e-mails",
			func(ctx app.Context, event app.Event) {
				emails := p.Group.EmailAdresses()
				clipboard.Paste(emails)
			},
			nil,
		),
		popover.NewPopoverItem(
			"Copier un template HTML pour un e-mail",
			func(ctx app.Context, event app.Event) {
				greetings := "<div>Bonjour à tous,<br /><br />"
				body := "Voici l'ordre de passage pour l'examen :</div><br />"
				table := component.HtmlOf(p)
				end := "<br /><div>Bien à vous,</div>"
				emailBody := greetings + body + table + end
				clipboard.Paste(emailBody)
			},
			nil,
		),
		popover.NewPopoverItem(
			"Copier le tableau en HTML",
			func(ctx app.Context, event app.Event) {
				html := component.HtmlOf(p)
				clipboard.Paste(html)
			},
			nil,
		),
	}
}

func (p *PassageOrderTable) tdClickItems(user *user.User, index int) []popover.Item {
	return []popover.Item{
		popover.NewPopoverItem(
			"Enlever",
			func(ctx app.Context, event app.Event) {
				p.Group.DeleteUser(index)
				p.Update()
			},
			map[string]string{
				"color":            "white",
				"background-color": "red",
			},
		),
	}
}
