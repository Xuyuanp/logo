package logo

// Logger function to log message with level.
type Logger func(level LogLevel, format string, args ...interface{})

// Log mehtods calls Logger function.
func (l Logger) Log(level LogLevel, format string, args ...interface{}) {
	l(level, format, args...)
}

// Log calls DefaultLogo's Log method.
func Log(level LogLevel, format string, args ...interface{}) {
	DefaultLogo.Skip(2).Log(level, format, args...)
}

// Debug logs Debug level message.
func Debug(format string, args ...interface{}) {
	DefaultLogo.Skip(2).Log(LevelDebug, format, args...)
}

// Info logs Info level message
func Info(format string, args ...interface{}) {
	DefaultLogo.Skip(2).Log(LevelInfo, format, args...)
}

// Warning logs Warning level message.
func Warning(format string, args ...interface{}) {
	DefaultLogo.Skip(2).Log(LevelWarning, format, args...)
}

// Error logs Error level message.
func Error(format string, args ...interface{}) {
	DefaultLogo.Skip(2).Log(LevelError, format, args...)
}

// Critical logs Critical level message
func Critical(format string, args ...interface{}) {
	DefaultLogo.Skip(2).Log(LevelCritical, format, args...)
}
