package logo

import "fmt"

// LogLevel defines level of log message.
type LogLevel int

// Available LogLevel const, the higher the more importent.
const (
	LevelDebug LogLevel = iota
	LevelInfo
	LevelWarning
	LevelError
	LevelCritical
)

// String returns friendly level name.
func (level LogLevel) String() string {
	switch level {
	case LevelDebug:
		return "DBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarning:
		return "WARN"
	case LevelError:
		return "ERRO"
	case LevelCritical:
		return "CRIT"
	}
	return fmt.Sprintf("Level(%d)", level)
}
