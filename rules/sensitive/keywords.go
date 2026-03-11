package sensitive

import "strings"

func containsSensitive(message string) bool {
	lower := strings.ToLower(message)

	for _, kw := range sensitiveKeywords {
		if strings.Contains(lower, kw) {
			return true
		}
	}

	return false
}

var sensitiveKeywords = []string{
	"password",
	"passwd",
	"pwd",
	"api_key",
	"apikey",
	"api key",
	"token",
	"secret",
	"access_key",
	"access key",
	"private_key",
	"private key",
}
