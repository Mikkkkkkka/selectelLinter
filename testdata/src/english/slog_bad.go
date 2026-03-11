package english

import "log/slog"

func SlogBad() {
	slog.Info("запуск сервера") // want "log message must be in English"
}
