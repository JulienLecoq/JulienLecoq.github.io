package html

type SpanTag struct {
	ContainerTag
}

func Span(elements ...Element) *SpanTag {
	return &SpanTag{
		ContainerTag: ContainerTag{
			tag:      "span",
			children: elements,
		},
	}
}
