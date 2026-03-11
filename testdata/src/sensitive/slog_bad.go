package sensitive

import (
	"log/slog"
)

func SlogBadPassword() {
	password := "123"
	slog.Info("user password " + password) // want "log message must not contain sensitive data"
}

func SlogBadAPIKey() {
	apiKey := "123"
	slog.Info("api key " + apiKey) // want "log message must not contain sensitive data"
}

func SlogBadToken() {
	token := "scary_token"
	slog.Info("token " + token) // want "log message must not contain sensitive data"
}
