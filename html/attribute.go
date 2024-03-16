package html

import (
	"fmt"
	"strings"
)

type StyleAttribute struct {
	value map[string]string
}

func (s StyleAttribute) String() string {
	k := "style"
	v := ""

	for key, value := range s.value {
		v += fmt.Sprintf("%s: %s; ", key, value)
	}

	// Remove the trailing "; " from the string
	v = strings.TrimSuffix(v, "; ")

	return fmt.Sprintf("%s=\"%s\"", k, v)
}

func (s *StyleAttribute) IsEmpty() bool {
	return s.value == nil || len(s.value) == 0
}
