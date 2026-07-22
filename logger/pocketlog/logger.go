package pocketlog

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// Logger é usado para fazer log de informações
type Logger struct {
	threshold Level
	output    io.Writer
}

// New retorna um logger que escreve a partir do nível solicitado
// Adiciona uma lista de configurarion functions para ajustá-lo a seu gosto
// O output padrão é Stdout
func New(threshold Level, options ...Option) *Logger {
	logger := &Logger{
		threshold: threshold,
		output:    os.Stdout,
	}

	for _, configFunc := range options {
		configFunc(logger)
	}

	return logger
}

// Debugf formata e exibe mensagem no nível debug
func (logger *Logger) Debugf(format string, args ...any) {
	logger.Logf(LevelDebug, format, args...)
}

// Infof formata e exibe mensagem no nível info
func (logger *Logger) Infof(format string, args ...any) {
	logger.Logf(LevelInfo, format, args...)
}

// Errorf formata e exibe mensagem no nível error
func (logger *Logger) Errorf(format string, args ...any) {
	logger.Logf(LevelError, format, args...)
}

// Logf verifica se o nível é compativel e formata e envia o log
// para o output
func (logger *Logger) Logf(level Level, format string, args ...any) {
	if logger.threshold > level {
		return
	}

	if logger.output == nil {
		logger.output = os.Stdout
	}

	_, _ = fmt.Fprintf(
		logger.output,
		"["+strings.ToUpper(level.Name())+"]: "+format+"\n",
		args...,
	)
}
