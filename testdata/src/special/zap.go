package special

import "go.uber.org/zap"

func ZapBad() {
	logger := zap.NewExample()
	logger.Info("connection failed!!!") // want "log message must not contain special characters or emoji"
}

func ZapBadEmoji() {
	logger := zap.NewExample()
	logger.Info("connection failed 🚀") // want "log message must not contain special characters or emoji"
}

func ZapBadPunct() {
	logger := zap.NewExample()
	logger.Info("warning: something went wrong...") // want "log message must not contain special characters or emoji"
}

func ZapGood() {
	logger := zap.NewExample()
	logger.Info("connection failed")
}
