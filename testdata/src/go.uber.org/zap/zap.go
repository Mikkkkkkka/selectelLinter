package zap // Package zap stub for tests

type Logger struct{}

type SugaredLogger struct{}

func NewExample() *Logger {
	return &Logger{}
}

func (l *Logger) Info(_ string, _ ...any)   {}
func (l *Logger) Debug(_ string, _ ...any)  {}
func (l *Logger) Warn(_ string, _ ...any)   {}
func (l *Logger) Error(_ string, _ ...any)  {}
func (l *Logger) DPanic(_ string, _ ...any) {}
func (l *Logger) Panic(_ string, _ ...any)  {}
func (l *Logger) Fatal(_ string, _ ...any)  {}

func (l *SugaredLogger) Info(_ string, _ ...any)   {}
func (l *SugaredLogger) Debug(_ string, _ ...any)  {}
func (l *SugaredLogger) Warn(_ string, _ ...any)   {}
func (l *SugaredLogger) Error(_ string, _ ...any)  {}
func (l *SugaredLogger) DPanic(_ string, _ ...any) {}
func (l *SugaredLogger) Panic(_ string, _ ...any)  {}
func (l *SugaredLogger) Fatal(_ string, _ ...any)  {}
