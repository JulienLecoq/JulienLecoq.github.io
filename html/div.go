package html

type DivTag struct {
	ContainerTag
}

// func Div(elements ...Element) Element {
func Div(elements ...Element) *DivTag {
	return &DivTag{
		ContainerTag: ContainerTag{
			tag:      "div",
			children: elements,
		},
	}
}
