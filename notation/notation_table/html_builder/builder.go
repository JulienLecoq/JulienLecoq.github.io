package html_builder

import (
	// "fmt"

	"juliano.com/passage-order/html"
	"juliano.com/passage-order/notation/graded_user"
	"juliano.com/passage-order/table"
)


func ToHtml(group *graded_user.GradedGroup) html.Element {
	TRs := append([]html.Element{headings()}, body(group)...)

	return html.Table(TRs...).SetStyle(table.TableStyle)
}

func String(group *graded_user.GradedGroup) string {
	html := ToHtml(group)
	return html.String()
}

func headings() html.Element {
	return html.Tr(
		html.Th().Text("Nom").SetStyle(table.ThStyle),
		html.Th().Text("Note").SetStyle(table.ThStyle),
	)
}

func body(group *graded_user.GradedGroup) []html.Element {
	TRs := []html.Element{}

	for index, user := range group.Users {
		tdStyle := table.TdStyle(index)

		nameTd := html.Td().Txt(user).SetStyle(tdStyle)
		noteTd := html.Td().Text(user.GradeFormatted()).SetStyle(tdStyle)

		tr := html.Tr(
			nameTd,
			noteTd,
		)

		TRs = append(TRs, tr)
	}

	return TRs
}
