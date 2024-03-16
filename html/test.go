package html

import (
	"fmt"
)

func Test() {
	p1 := P()
	p2 := P()

	p3 := P()

	text1 := Text("Salut poto")
	text2 := Text("Salut poto")
	subSpan := Span(
		text1,
		text2,
	)
	span := Span(subSpan)

	div2 := Div(
		p3,
		span,
	)

	divStyle := map[string]string{
		"padding":          "16px",
		"background-color": "red",
	}

	div := Div(
		div2,
		p1,
		p2,
	).SetStyle(divStyle)

	fmt.Printf("Html content:\n\n%s", div)
}
