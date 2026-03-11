package sensitive

import "log/slog"

func SlogBad() {
	slog.Info("user password: 123") // want "log message must not contain sensitive data"
}

func SlogBadAPIKey() {
	slog.Info("api_key=abcd") // want "log message must not contain sensitive data"
}

func SlogBadToken() {
	slog.Info("token: value") // want "log message must not contain sensitive data"
}
