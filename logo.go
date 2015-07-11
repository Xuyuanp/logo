package logo

import "fmt"

// Logo struct
type Logo struct {
	handlers []Handler
}

// DefaultLogo with DefaultHandler.
var DefaultLogo = &Logo{
	handlers: []Handler{DefaultHandler},
}

func (l *Logo) log(depth int, level LogLevel, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	rec := NewRecord()
	defer recordPool.Put(rec)
	rec.Init(depth+1, level, msg)

	for _, handler := range l.handlers {
		handler.Handle(rec)
	}
}

// Skip return a Logger log message by skip the provided depth.
func (l *Logo) Skip(depth int) Logger {
	return func(level LogLevel, format string, args ...interface{}) {
		l.log(depth+2, level, format, args...)
	}
}

// Log messages using this level.
func (l *Logo) Log(level LogLevel, format string, args ...interface{}) {
	l.log(2, level, format, args...)
}

// Debug logs Debug level message.
func (l *Logo) Debug(format string, args ...interface{}) {
	l.log(2, LevelDebug, format, args...)
}

// Info logs Info level message.
func (l *Logo) Info(format string, args ...interface{}) {
	l.log(2, LevelInfo, format, args...)
}

// Warning logs warning level message.
func (l *Logo) Warning(format string, args ...interface{}) {
	l.log(2, LevelWarning, format, args...)
}

// Error logs Error level message.
func (l *Logo) Error(format string, args ...interface{}) {
	l.log(2, LevelError, format, args...)
}

// Critical log Critical level message.
func (l *Logo) Critical(format string, args ...interface{}) {
	l.log(2, LevelCritical, format, args...)
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
