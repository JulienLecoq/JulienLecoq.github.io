package html

import (
	"fmt"
	"strings"
)

type ContainerTag struct {
	id       string
	class    string
	tag      string
	children []Element
	style    StyleAttribute
}

func (e *ContainerTag) IsContainerTag() bool {
	return true
}

func (e *ContainerTag) IsSelfClosingTag() bool {
	return false
}

func (e *ContainerTag) TagName() string {
	return e.tag
}

func (e *ContainerTag) ID() string {
	return e.id
}

func (e *ContainerTag) SetID(id string) *ContainerTag {
	e.id = id
	return e
}

func (e *ContainerTag) Class() string {
	return e.class
}

func (e *ContainerTag) SetClass(class string) *ContainerTag {
	e.class = class
	return e
}

func (e *ContainerTag) Style() map[string]string {
	return e.style.value
}

func (e *ContainerTag) SetStyle(style map[string]string) *ContainerTag {
	e.style.value = style
	return e
}

func (e *ContainerTag) Text(text string) *ContainerTag {
	e.children = []Element{
		Text(text),
	}

	return e
}

func (e *ContainerTag) Txt(obj fmt.Stringer) *ContainerTag {
	e.children = []Element{
		Text(obj.String()),
	}

	return e
}

func (e *ContainerTag) StartTag() string {
	return fmt.Sprintf("<%s>", e.tag)
}

func (e *ContainerTag) StartTagWithAttributes() string {
	if e.style.IsEmpty() {
		return fmt.Sprintf("<%s>", e.tag)
	}

	return fmt.Sprintf("<%s %s>", e.tag, e.style)
}

func (e *ContainerTag) EndTag() string {
	return fmt.Sprintf("</%s>", e.tag)
}

func (e *ContainerTag) ChildTags() string {
	tags := ""

	for _, element := range e.children {
		tags += "\n\t" + element.String()
	}

	return tags
}

func (e *ContainerTag) String() string {
	return e.StringWithIndentation(0)
}

func (e *ContainerTag) StringWithIndentation(indentLevel int) string {
	indentation := strings.Repeat("\t", indentLevel)
	tags := ""

	for _, element := range e.children {
		// Now directly calling the new method on all child elements.
		tags += "\n" + element.StringWithIndentation(indentLevel+1)
	}

	if len(e.children) > 0 {
		return indentation + e.StartTagWithAttributes() + tags + "\n" + indentation + e.EndTag()
	}

	return indentation + e.StartTagWithAttributes() + e.EndTag()
}
