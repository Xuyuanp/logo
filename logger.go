package logo

type Logger func(level LogLevel, format string, args ...interface{})

func (l Logger) Log(level LogLevel, format string, args ...interface{}) {
	l(level, format, args...)
}

func Log(level LogLevel, format string, args ...interface{}) {
	DefaultLogo.Skip(2).Log(level, format, args...)
}

func Debug(format string, args ...interface{}) {
	DefaultLogo.Skip(2).Log(LevelDebug, format, args...)
}

func Info(format string, args ...interface{}) {
	DefaultLogo.Skip(2).Log(LevelInfo, format, args...)
}

func Warning(format string, args ...interface{}) {
	DefaultLogo.Skip(2).Log(LevelWarning, format, args...)
}

func Error(format string, args ...interface{}) {
	DefaultLogo.Skip(2).Log(LevelError, format, args...)
}

func Critical(format string, args ...interface{}) {
	DefaultLogo.Skip(2).Log(LevelCritical, format, args...)
}
