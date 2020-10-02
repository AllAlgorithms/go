package vowelcount

import "unicode"

func Count(s string) int {
	c := 0
	for _, r := range s {
		if isVowel(r) {
			c += 1
		}
	}

	return c
}

func isVowel(r rune) bool {
	r = unicode.ToLower(r)
	switch r {
	case 'a', 'i', 'u', 'e', 'o':
		return true
	}
	return false
}
