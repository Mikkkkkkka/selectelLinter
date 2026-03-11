package sensitive

import "go.uber.org/zap"

func ZapBad() {
	logger := zap.NewExample()
	logger.Info("password reset started") // want "log message must not contain sensitive data"
}

func ZapBadSecret() {
	logger := zap.NewExample()
	logger.Info("secret key updated") // want "log message must not contain sensitive data"
}
