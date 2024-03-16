package notation_table

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"juliano.com/passage-order/clipboard"
	"juliano.com/passage-order/notation/graded_user"
	"juliano.com/passage-order/notation/notation_table/html_builder"
	"juliano.com/passage-order/popover"
)

func (n *NotationTable) thClickItems() []popover.Item {
	return []popover.Item{
		popover.NewPopoverItem(
			"Randomiser",
			func(ctx app.Context, event app.Event) {
				n.Group.RandomizeUsers()
				n.Update()
			},
			nil,
		),
		popover.NewPopoverItem(
			"Trier par ordre alphabétique",
			func(ctx app.Context, event app.Event) {
				n.Group.SortAlphabetically()
				n.Update()
			},
			nil,
		),
		popover.NewPopoverItem(
			"Copier les adresses e-mails",
			func(ctx app.Context, event app.Event) {
				emails := n.Group.EmailAdresses()
				clipboard.Paste(emails)
			},
			nil,
		),
		popover.NewPopoverItem(
			"Copier un template HTML pour un e-mail",
			func(ctx app.Context, event app.Event) {
				greetings := "<div>Bonjour à tous,<br /><br />"
				body := "Voici les notes de l'examen :</div><br />"
				table := html_builder.String(n.Group)
				end := "<br /><div>Bien à vous,</div>"
				emailBody := greetings + body + table + end
				clipboard.Paste(emailBody)
			},
			nil,
		),
		popover.NewPopoverItem(
			"Copier le tableau en HTML",
			func(ctx app.Context, event app.Event) {
				html := html_builder.String(n.Group)
				clipboard.Paste(html)
			},
			nil,
		),
	}
}

func (n *NotationTable) tdClickItems(_ *graded_user.GradedUser, index int) []popover.Item {
	return []popover.Item{
		popover.NewPopoverItem(
			"Enlever",
			func(ctx app.Context, event app.Event) {
				n.Group.DeleteUser(index)
				n.Update()
			},
			map[string]string{
				"color":            "white",
				"background-color": "red",
			},
		),
	}
}
