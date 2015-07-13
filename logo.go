package logo

import "fmt"

// Logo struct
type Logo struct {
	Level    LogLevel
	handlers []Handler
}

var defaultLogo = &Logo{
	Level:    LevelDebug,
	handlers: []Handler{DefaultHandler},
}

// DefaultLogo return a global Logo instance.
func DefaultLogo() *Logo {
	return defaultLogo
}

// Output writes the output for a logging event.  The string s contains
// the text to print specified by the flags of the
// Logger.  A newline is appended if the last character of s is not
// already a newline.  depth is used to recover the PC and is
// provided for generality, although at the moment on all pre-defined
// paths it will be 2.
func (l *Logo) Output(depth int, level LogLevel, format string, args ...interface{}) {
	if level < l.Level {
		return
	}

	msg := fmt.Sprintf(format, args...)

	for _, handler := range l.handlers {
		handler.Handle(depth+1, level, msg)
	}
}

// Skip return a Logger log message by skip the provided depth.
// depth is 1 at most.
func (l *Logo) Skip(depth int) Logger {
	return func(level LogLevel, format string, args ...interface{}) {
		l.Output(depth+2, level, format, args...)
	}
}

// Log messages using this level.
func (l *Logo) Log(level LogLevel, format string, args ...interface{}) {
	l.Output(2, level, format, args...)
}

// Debug logs Debug level message.
func (l *Logo) Debug(format string, args ...interface{}) {
	l.Output(2, LevelDebug, format, args...)
}

// Info logs Info level message.
func (l *Logo) Info(format string, args ...interface{}) {
	l.Output(2, LevelInfo, format, args...)
}

// Warning logs warning level message.
func (l *Logo) Warning(format string, args ...interface{}) {
	l.Output(2, LevelWarning, format, args...)
}

// Error logs Error level message.
func (l *Logo) Error(format string, args ...interface{}) {
	l.Output(2, LevelError, format, args...)
}

// Critical log Critical level message.
func (l *Logo) Critical(format string, args ...interface{}) {
	l.Output(2, LevelCritical, format, args...)
}

// AddHandlers appends thes handlers to original handlers.
func (l *Logo) AddHandlers(handlers ...Handler) {
	l.handlers = append(l.handlers, handlers...)
}

// SetHandlers replaces original handlers by these.
func (l *Logo) SetHandlers(handlers ...Handler) {
	l.handlers = make([]Handler, len(handlers))
	copy(l.handlers, handlers)
}
