package html

type PTag struct {
	ContainerTag
}

func P(elements ...Element) *PTag {
	return &PTag{
		ContainerTag: ContainerTag{
			tag:      "p",
			children: elements,
		},
	}
}
