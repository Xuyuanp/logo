package logo

import "fmt"

type Logo struct {
	handlers []Handler
}

var DefaultLogo = &Logo{
	handlers: []Handler{DefaultHandler},
}

func (l *Logo) log(skip int, level LogLevel, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	rec := NewRecord()
	defer recordPool.Put(rec)
	rec.Init(skip+1, level, msg)

	for _, handler := range l.handlers {
		handler.Handle(rec)
	}
}

func (l *Logo) Skip(skip int) Logger {
	return func(level LogLevel, format string, args ...interface{}) {
		l.log(skip+2, level, format, args...)
	}
}

func (l *Logo) Log(level LogLevel, format string, args ...interface{}) {
	l.log(2, level, format, args...)
}

func (l *Logo) Debug(format string, args ...interface{}) {
	l.log(2, LevelDebug, format, args...)
}

func (l *Logo) Info(format string, args ...interface{}) {
	l.log(2, LevelInfo, format, args...)
}

func (l *Logo) Warning(format string, args ...interface{}) {
	l.log(2, LevelWarning, format, args...)
}

func (l *Logo) Error(format string, args ...interface{}) {
	l.log(2, LevelError, format, args...)
}

func (l *Logo) Critical(format string, args ...interface{}) {
	l.log(2, LevelCritical, format, args...)
}

func (l *Logo) AddHandlers(handlers ...Handler) {
	l.handlers = append(l.handlers, handlers...)
}

func (l *Logo) SetHandlers(handlers ...Handler) {
	l.handlers = make([]Handler, len(handlers))
	copy(l.handlers, handlers)
}
