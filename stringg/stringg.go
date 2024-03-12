package stringg

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

func Capitalize(s string) string {
	if s == "" {
		return ""
	}

	r, size := utf8.DecodeRuneInString(s)

	return strings.ToUpper(string(r)) + s[size:]
}

func CapitalizeEachWord(input string) string {
	// Split the input into words using a regular expression to account for spaces and hyphens.
	re := regexp.MustCompile(`[\p{L}-]+`)
	words := re.FindAllString(input, -1)

	for i, word := range words {
		// Split further by hyphen to ensure words like "test-test" are handled correctly.
		parts := strings.Split(word, "-")
		for j, part := range parts {
			if len(part) > 0 {
				// Capitalize the first letter of each part.
				parts[j] = strings.ToUpper(part[:1]) + part[1:]
			}
		}
		// Join the parts back together, if they were split.
		words[i] = strings.Join(parts, "-")
	}

	// Join the words back into a single string.
	return strings.Join(words, " ")
}
