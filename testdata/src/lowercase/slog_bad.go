package lowercase

import "log/slog"

func SlogBad() {
	slog.Info("Starting server") // want "log message should start with a lowercase letter"
}
