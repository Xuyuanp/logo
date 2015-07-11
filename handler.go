package logo

import (
	"fmt"
	"io"
	"os"
)

// Handler handles the record and formats it by using its
// Formatter, and emit into its writer or something else.
type Handler func(rec *Record)

// Handle call Handler function
func (h Handler) Handle(rec *Record) {
	h(rec)
}

// HandlerFilter function
type HandlerFilter func(Handler) Handler

// Filter method filters wraps a handler to a new handler that
// filtered by its own rule.
func (hf HandlerFilter) Filter(handler Handler) Handler {
	return func(rec *Record) {
		hf(handler).Handle(rec)
	}
}

// LevelHandlerFilter returns a HandlerFilter that bypasses all
// the record whose level is lower than the provied one.
func LevelHandlerFilter(level LogLevel) HandlerFilter {
	return func(handler Handler) Handler {
		return func(rec *Record) {
			if rec.Level >= level {
				handler.Handle(rec)
			}
		}
	}
}

var DefaultHandler = WriterHandler(os.Stdout, DefaultFormatter)

func WriterHandler(w io.Writer, fmter Formatter) Handler {
	return func(rec *Record) {
		fmt.Fprintln(w, fmter.Format(rec))
	}
}
