package sensitive

import "strings"

func containsSensitive(vars []string) bool {
	for _, name := range vars {
		lower := strings.ToLower(name)

		for _, kw := range sensitiveKeywords {
			if strings.Contains(lower, kw) {
				return true
			}
		}
	}
	return false
}

var sensitiveKeywords = []string{
	"password",
	"passwd",
	"pwd",
	"key",
	"api_key",
	"apikey",
	"token",
	"secret",
	"access_key",
	"private_key",
}
