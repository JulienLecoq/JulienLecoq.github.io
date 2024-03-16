package html

type TableTag struct {
	ContainerTag
}

func Table(elements ...Element) *TableTag {
	return &TableTag{
		ContainerTag: ContainerTag{
			tag:      "table",
			children: elements,
		},
	}
}
