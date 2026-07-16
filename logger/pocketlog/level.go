package pocketlog

// Representa um nível de log.
type Level byte

const (
	// Representa o menor nível de log
	LevelDebug = iota
	// Representa logs que contém informações
	LevelInfo = iota
	// Representa o maior nível de log, usado para erros
	LevelError = iota
)
