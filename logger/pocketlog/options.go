package pocketlog

import "io"

// Option define uma functional option para o logger.
type Option func(*Logger)

// WithOutput retorna uma função de configuração para definir o output dos logs
func WithOutput(output io.Writer) Option {
	return func(logger *Logger) {
		logger.output = output
	}
}
