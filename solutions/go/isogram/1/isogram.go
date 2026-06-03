package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	count := map[string]int{}
	clean := strings.Map(func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return -1
		}
		return unicode.ToLower(r)
	}, word)
	for _, k := range clean {
		count[string(k)]++
	}
	for _, j := range count {
		if j > 1 {
			return false
		}
	}
	return true
}
