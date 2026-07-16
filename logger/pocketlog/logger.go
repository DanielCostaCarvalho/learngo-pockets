package pocketlog

import "fmt"

// Logger é usado para fazer log de informações
type Logger struct {
	threshold Level
}

// New retorna um logger que escreve a partir do nível solicitado
func New(threshold Level) *Logger {
	return &Logger{
		threshold: threshold,
	}
}

// Debugf formata e exibe mensagem no nível debug
func (logger *Logger) Debugf(format string, args ...any) {
	if logger.threshold > LevelDebug {
		return
	}

	fmt.Printf(format + "\n", args...)
}

// Infof formata e exibe mensagem no nível info
func (logger *Logger) Infof(format string, args ...any) {
	if logger.threshold > LevelInfo {
		return
	}

	fmt.Printf(format + "\n", args...)
}

// Errorf formata e exibe mensagem no nível error
func (logger *Logger) Errorf(format string, args ...any) {
	if logger.threshold > LevelError {
		return
	}

	fmt.Printf(format + "\n", args...)
}
