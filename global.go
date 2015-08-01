package logo

import "fmt"

var defaultLogo = NewLogo(LevelDebug, defaultHandler)

// Debug logs Debug level message.
func Debug(format string, args ...interface{}) {
	defaultLogo.Output(2, LevelDebug, fmt.Sprintf(format, args...))
}

// Info logs Info level message
func Info(format string, args ...interface{}) {
	defaultLogo.Output(2, LevelInfo, fmt.Sprintf(format, args...))
}

// Warning logs Warning level message.
func Warning(format string, args ...interface{}) {
	defaultLogo.Output(2, LevelWarning, fmt.Sprintf(format, args...))
}

// Error logs Error level message.
func Error(format string, args ...interface{}) {
	defaultLogo.Output(2, LevelError, fmt.Sprintf(format, args...))
}

// Critical logs Critical level message
func Critical(format string, args ...interface{}) {
	defaultLogo.Output(2, LevelCritical, fmt.Sprintf(format, args...))
}
