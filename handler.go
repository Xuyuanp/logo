package logo

import (
	"io"
	"os"
	"runtime"
	"sync"
	"time"
)

// These flags define which text to prefix to each log entry generated by the Logger.
const (
	// Bits or'ed together to control what's printed. There is no control over the
	// order they appear (the order listed here) or the format they present (as
	// described in the comments).  A colon appears after these items:
	//	2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message

	Ldate         = 1 << iota              // the date: 2009/01/23
	Ltime                                  // the time: 01:23:23
	Lmicroseconds                          // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile                              // full file name and line number: /a/b/c/d.go:23
	Lshortfile                             // final file name element and line number: d.go:23. overrides Llongfile
	Llevel                                 // the level name: DBUG|INFO|WARN...
	LstdFlags     = Ldate | Ltime | Llevel // initial values for the standard logger
	LfullFlags    = LstdFlags | Lshortfile
)

// Handler handles the record and formats it by using its
// Formatter, and emit into its writer or something else.
type Handler func(depth int, level LogLevel, msg string)

// Handle call Handler function
func (h Handler) Handle(depth int, level LogLevel, msg string) {
	h(depth, level, msg)
}

// DefaultHandler format message whose level is higher than LevelDebug
// with the LfullFlags and outputs message into stdout,
var DefaultHandler = NewHandler(os.Stdout, LevelDebug, LfullFlags)

// NewHandler return a new Handler.
func NewHandler(w io.Writer, level LogLevel, flag int) Handler {
	var mu sync.Mutex
	var buf []byte

	return func(depth int, lvl LogLevel, msg string) {
		// ignore low level.
		if lvl < level {
			return
		}
		now := time.Now()
		var file string
		var line int
		mu.Lock()
		defer mu.Unlock()
		if flag&(Lshortfile|Llongfile) != 0 {
			// release lock wihle getting caller info - it's expensive.
			mu.Unlock()
			var ok bool
			_, file, line, ok = runtime.Caller(depth + 1)
			if !ok {
				file = "???"
				line = 0
			}
			mu.Lock()
		}
		buf = buf[:0]
		formatHeader(&buf, flag, now, file, line, lvl)
		buf = append(buf, msg...)
		if len(msg) > 0 && msg[len(msg)-1] != '\n' {
			buf = append(buf, '\n')
		}
		w.Write(buf)
	}
}

// Stolen from log package.
func formatHeader(buf *[]byte, flag int, t time.Time, file string, line int, level LogLevel) {
	if flag&(Ldate|Ltime|Lmicroseconds) != 0 {
		if flag&Ldate != 0 {
			year, month, day := t.Date()
			itoa(buf, year, 4)
			*buf = append(*buf, '/')
			itoa(buf, int(month), 2)
			*buf = append(*buf, '/')
			itoa(buf, day, 2)
			*buf = append(*buf, ' ')
		}
		if flag&(Ltime|Lmicroseconds) != 0 {
			hour, min, sec := t.Clock()
			itoa(buf, hour, 2)
			*buf = append(*buf, ':')
			itoa(buf, min, 2)
			*buf = append(*buf, ':')
			itoa(buf, sec, 2)
			if flag&Lmicroseconds != 0 {
				*buf = append(*buf, '.')
				itoa(buf, t.Nanosecond()/1e3, 6)
			}
			*buf = append(*buf, ' ')
		}
	}
	if flag&Llevel != 0 {
		lvlName := level.String()
		*buf = append(*buf, lvlName...)
		*buf = append(*buf, ' ')
	}
	if flag&(Lshortfile|Llongfile) != 0 {
		if flag&Lshortfile != 0 {
			short := file
			for i := len(file) - 1; i > 0; i-- {
				if file[i] == '/' {
					short = file[i+1:]
					break
				}
			}
			file = short
		}
		*buf = append(*buf, file...)
		*buf = append(*buf, ':')
		itoa(buf, line, -1)
		*buf = append(*buf, ": "...)
	}
}

// Cheap integer to fixed-width decimal ASCII.  Give a negative width to avoid zero-padding.
// Knows the buffer has capacity.
// Stolen from log package.
func itoa(buf *[]byte, i int, wid int) {
	var u = uint(i)
	if u == 0 && wid <= 1 {
		*buf = append(*buf, '0')
		return
	}

	// Assemble decimal in reverse order.
	var b [32]byte
	bp := len(b)
	for ; u > 0 || wid > 0; u /= 10 {
		bp--
		wid--
		b[bp] = byte(u%10) + '0'
	}
	*buf = append(*buf, b[bp:]...)
}
