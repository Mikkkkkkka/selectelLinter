package lowercase

import "go.uber.org/zap"

func ZapBad() {
	logger := zap.NewExample()
	logger.Info("Starting server") // want "log message should start with a lowercase letter"
}
