/*
 * Copyright 2015 Xuyuan Pang
 * Author: Xuyuan Pang
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package logo

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"
	"time"
)

// Logger function type
type Logger func(depth int, level LogLevel, msg string)

// New create a new Logger which outputs log message into io.Writer.
func New(lowest LogLevel, w io.Writer, prefix string, flag int) Logger {
	var mu sync.Mutex
	var buf []byte

	return func(depth int, lvl LogLevel, msg string) {
		// ignore low level.
		if lvl < lowest {
			return
		}
		now := time.Now()
		var file string
		var line int
		mu.Lock()
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
		formatHeader(&buf, prefix, flag, now, file, line, lvl)
		if flag&Lcolor != 0 {
			buf = append(buf, colorMap[lvl]...)
			buf = append(buf, msg...)
			buf = append(buf, colorReset...)
		} else {
			buf = append(buf, msg...)
		}
		if len(msg) > 0 && msg[len(msg)-1] != '\n' {
			buf = append(buf, '\n')
		}
		w.Write(buf)
		mu.Unlock()
	}
}

// NewStd creates a new Logger using stdout.
func NewStd(lowest LogLevel, prefix string, flags int) Logger {
	return New(lowest, os.Stdout, prefix, flags)
}

// Group combines multiple Logger interface into a new Logger which
// outputs all log message into each Logger of provided.
func Group(lowest LogLevel, Loggers ...Logger) Logger {
	return func(depth int, lvl LogLevel, msg string) {
		if lvl < lowest {
			return
		}
		for _, l := range Loggers {
			l.Output(depth+2, lvl, msg)
		}
	}
}

// Output writes the output for a logging event.  The string s contains
// the text to print specified by the flags of the
// Logger.  A newline is appended if the last character of s is not
// already a newline.  depth is used to recover the PC and is
// provided for generality, although at the moment on all pre-defined
// paths it will be 2.
func (l Logger) Output(depth int, level LogLevel, msg string) {
	l(depth, level, msg)
}

// Log messages using this level.
func (l Logger) Log(level LogLevel, format string, args ...interface{}) {
	l.Output(2, level, fmt.Sprintf(format, args...))
}

// Debug logs Debug level message.
func (l Logger) Debug(format string, args ...interface{}) {
	l.Output(2, LevelDebug, fmt.Sprintf(format, args...))
}

// Info logs Info level message.
func (l Logger) Info(format string, args ...interface{}) {
	l.Output(2, LevelInfo, fmt.Sprintf(format, args...))
}

// Warning logs warning level message.
func (l Logger) Warning(format string, args ...interface{}) {
	l.Output(2, LevelWarning, fmt.Sprintf(format, args...))
}

// Error logs Error level message.
func (l Logger) Error(format string, args ...interface{}) {
	l.Output(2, LevelError, fmt.Sprintf(format, args...))
}

// Critical log Critical level message.
func (l Logger) Critical(format string, args ...interface{}) {
	l.Output(2, LevelCritical, fmt.Sprintf(format, args...))
}
