package english

import "go.uber.org/zap"

func ZapGood() {
	logger := zap.NewExample()
	logger.Info("connection failed")
}
