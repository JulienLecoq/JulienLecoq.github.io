package html

type ThTag struct {
	ContainerTag
}

func Th(elements ...Element) *ThTag {
	return &ThTag{
		ContainerTag: ContainerTag{
			tag:      "th",
			children: elements,
		},
	}
}
