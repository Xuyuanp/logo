package logo

import "fmt"

// Formatter formats record into string.
type Formatter func(rec *Record) string

// Format method calls Formatter function.
func (fmter Formatter) Format(rec *Record) string {
	return fmter(rec)
}

// DefaultTimeFormat is the default time format string.
const DefaultTimeFormat = "2006-01-02 15:04:05"

// DefaultFormatter is a default formatter output record into
// "[2006-01-02 15:04:05 INFO main.go:10] hello world!" format.
var DefaultFormatter Formatter = func(rec *Record) string {
	return fmt.Sprintf("[%s %s %s:%d] %s",
		rec.Time.Format(DefaultTimeFormat),
		rec.Level,
		rec.SFile,
		rec.Line,
		rec.Message)
}
