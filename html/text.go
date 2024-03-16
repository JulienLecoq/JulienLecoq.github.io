package html

import "strings"

type TextTag struct {
	text string
}

func (t *TextTag) TagName() string {
	return ""
}

func (t *TextTag) String() string {
	return t.text
}

func (t *TextTag) StringWithIndentation(indentLevel int) string {
	indentation := strings.Repeat("\t", indentLevel)
	return indentation + t.String()
}

func (t *TextTag) IsContainerTag() bool {
	return false
}

func (t *TextTag) IsSelfClosingTag() bool {
	return false
}

func Text(text string) *TextTag {
	return &TextTag{
		text: text,
	}
}
