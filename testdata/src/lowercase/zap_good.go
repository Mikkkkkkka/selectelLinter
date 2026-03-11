package lowercase

import "go.uber.org/zap"

func ZapGood() {
	logger := zap.NewExample()
	logger.Info("starting server")
}
