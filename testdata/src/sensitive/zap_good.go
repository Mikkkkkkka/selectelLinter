package sensitive

import "go.uber.org/zap"

func ZapGood() {
	logger := zap.NewExample()
	logger.Info("request completed")
}
