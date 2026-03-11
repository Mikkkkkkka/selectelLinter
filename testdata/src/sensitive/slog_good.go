package sensitive

import "log/slog"

func SlogGoodPassword() {
	slog.Info("user authenticated successfully")
}

func SlogGoodApi() {
	slog.Info("api request completed")
}

func SlogGoodToken() {
	slog.Info("token validated")
}
