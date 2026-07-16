package pocketlog

type Logger struct {
	threshold Level
}

func New(threshold Level) *Logger {
	return &Logger{
		threshold: threshold,
	}
}

func (logger *Logger) Debugf(format string, args ...any) {
	//
}

func (logger *Logger) Infof(format string, args ...any) {
	//
}

func (logger *Logger) Errorf(format string, args ...any) {
	//
}
