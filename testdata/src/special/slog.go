package special

import "log/slog"

func SlogBad() {
	slog.Info("server started! ") // want "log message must not contain special characters or emoji"
}

func SlogBadEmoji() {
	slog.Info("server started 🚀") // want "log message must not contain special characters or emoji"
}

func SlogBadPunct() {
	slog.Info("warning: something went wrong...") // want "log message must not contain special characters or emoji"
}

func SlogGood() {
	slog.Info("server started")
}
