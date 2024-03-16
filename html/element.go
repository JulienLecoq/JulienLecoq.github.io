package html

type Element interface {
	TagName() string
	String() string
	StringWithIndentation(indentLevel int) string

	// ID() string
	// SetID(string) Element

	// Class() string
	// SetClass(string) Element

	// Style() map[string]string
	// SetStyle(map[string]string) Element

	IsContainerTag() bool
	IsSelfClosingTag() bool
}
