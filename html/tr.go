package html

type TrTag struct {
	ContainerTag
}

func Tr(elements ...Element) *TrTag {
	return &TrTag{
		ContainerTag: ContainerTag{
			tag:      "tr",
			children: elements,
		},
	}
}
