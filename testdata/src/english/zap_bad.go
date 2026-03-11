package english

import "go.uber.org/zap"

func ZapBad() {
	logger := zap.NewExample()
	logger.Info("ошибка подключения") // want "log message must be in English"
}
