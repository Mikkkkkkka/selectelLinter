package sensitive

import "log/slog"

func SlogGood() {
	slog.Info("user authenticated successfully")
}
