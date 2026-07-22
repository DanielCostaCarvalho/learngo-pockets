package pocketlog

// Level representa um nível de log.
type Level byte

const (
	// LevelDebug representa o menor nível de log.
	LevelDebug = iota
	// LevelInfo representa logs que contém informações.
	LevelInfo = iota
	// LevelError representa o maior nível de log, usado para erros.
	LevelError = iota
)

func (level *Level) Name() string {
	switch *level {
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelError:
		return "error"
	default:
		return ""
	}
}
