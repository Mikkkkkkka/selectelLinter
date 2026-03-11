package sensitive

import "go.uber.org/zap"

func ZapBadPassword() {
	logger := zap.NewExample()
	password := "123"
	logger.Info("user password " + password) // want "log message must not contain sensitive data"
}

func ZapBadApi() {
	logger := zap.NewExample()
	apiKey := "123"
	logger.Info("api key " + apiKey) // want "log message must not contain sensitive data"
}

func ZapBadToken() {
	logger := zap.NewExample()
	token := "scary_token"
	logger.Info("token " + token) // want "log message must not contain sensitive data"
}
