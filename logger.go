package logo

// Logger function to log message with level.
type Logger func(level LogLevel, format string, args ...interface{})

// Log mehtods calls Logger function.
func (l Logger) Log(level LogLevel, format string, args ...interface{}) {
	l(level, format, args...)
}

var defaultLogger = defaultLogo.Skip(2)

// Log calls DefaultLogo's Log method.
func Log(level LogLevel, format string, args ...interface{}) {
	DefaultLogo().Output(2, level, format, args...)
}

// Debug logs Debug level message.
func Debug(format string, args ...interface{}) {
	defaultLogger.Log(LevelDebug, format, args...)
}

// Info logs Info level message
func Info(format string, args ...interface{}) {
	defaultLogger.Log(LevelInfo, format, args...)
}

// Warning logs Warning level message.
func Warning(format string, args ...interface{}) {
	defaultLogger.Log(LevelWarning, format, args...)
}

// Error logs Error level message.
func Error(format string, args ...interface{}) {
	defaultLogger.Log(LevelError, format, args...)
}

// Critical logs Critical level message
func Critical(format string, args ...interface{}) {
	defaultLogger.Log(LevelCritical, format, args...)
}
