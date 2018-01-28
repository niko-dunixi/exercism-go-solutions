package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate will take a string and return its acronym
func Abbreviate(phrase string) string {
	// 1 - Utilize a regex, at a given word boundary, take the first letter thereafter.
	// 2 - Convert the input to uppercase and then run the phrase through the regex
	// 3 - Join the strings without separation
	re, _ := regexp.Compile("\\b(\\w)")
	result := re.FindAllString(strings.ToUpper(phrase), -1)
	return strings.Join(result, "")
}
