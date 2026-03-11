package special

import "unicode"

func hasSpecialChars(message string) bool {
	for _, ch := range message {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) || unicode.IsSpace(ch) {
			continue
		}
		return true
	}
	return false
}
