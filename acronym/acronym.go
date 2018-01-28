package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(phrase string) string {
	re, _ := regexp.Compile("\\b(\\w)")
	result := re.FindAllString(strings.ToUpper(phrase), -1)
	return strings.Join(result, "")
}
