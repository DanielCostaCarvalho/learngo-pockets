package pocketlog

import (
	"fmt"
	"io"
	"os"
)

// Logger é usado para fazer log de informações
type Logger struct {
	threshold Level
	output    io.Writer
}

// New retorna um logger que escreve a partir do nível solicitado
// O output padrão é Stdout
func New(threshold Level, output io.Writer) *Logger {
	return &Logger{
		threshold: threshold,
		output:    output,
	}
}

// Debugf formata e exibe mensagem no nível debug
func (logger *Logger) Debugf(format string, args ...any) {
	logger.logfByLevel(LevelDebug, format, args...)
}

// Infof formata e exibe mensagem no nível info
func (logger *Logger) Infof(format string, args ...any) {
	logger.logfByLevel(LevelInfo, format, args...)
}

// Errorf formata e exibe mensagem no nível error
func (logger *Logger) Errorf(format string, args ...any) {
	logger.logfByLevel(LevelError, format, args...)
}

// logfByLevel verifica se o nível é compativel e formata e envia o log
// para o output
func (logger *Logger) logfByLevel(level Level, format string, args ...any) {
	if logger.threshold > level {
		return
	}

	if logger.output == nil {
		logger.output = os.Stdout
	}

	_, _ = fmt.Fprintf(logger.output, format+"\n", args...)
}
