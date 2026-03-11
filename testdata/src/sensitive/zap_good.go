package sensitive

import "go.uber.org/zap"

func ZapGoodPassword() {
	logger := zap.NewExample()
	logger.Info("user authenticated successfully")
}

func ZapGoodApi() {
	logger := zap.NewExample()
	logger.Info("api request completed")
}

func ZapGoodToken() {
	logger := zap.NewExample()
	logger.Info("token validated")
}
