package html

type TdTag struct {
	ContainerTag
}

func Td(elements ...Element) *TdTag {
	return &TdTag{
		ContainerTag: ContainerTag{
			tag:      "td",
			children: elements,
		},
	}
}
