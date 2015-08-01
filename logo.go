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

import "fmt"

// Logo function type
type Logo func(depth int, level LogLevel, msg string)

// NewLogo returns a new Logo function.
func NewLogo(level LogLevel, handlers ...Handler) Logo {
	return func(depth int, lvl LogLevel, msg string) {
		if lvl < level {
			return
		}

		for _, handler := range handlers {
			handler.Handle(depth+2, lvl, msg)
		}
	}
}

// Output writes the output for a logging event.  The string s contains
// the text to print specified by the flags of the
// Logger.  A newline is appended if the last character of s is not
// already a newline.  depth is used to recover the PC and is
// provided for generality, although at the moment on all pre-defined
// paths it will be 2.
func (l Logo) Output(depth int, level LogLevel, msg string) {
	l(depth, level, msg)
}

// Log messages using this level.
func (l Logo) Log(level LogLevel, format string, args ...interface{}) {
	l.Output(2, level, fmt.Sprintf(format, args...))
}

// Debug logs Debug level message.
func (l Logo) Debug(format string, args ...interface{}) {
	l.Output(2, LevelDebug, fmt.Sprintf(format, args...))
}

// Info logs Info level message.
func (l Logo) Info(format string, args ...interface{}) {
	l.Output(2, LevelInfo, fmt.Sprintf(format, args...))
}

// Warning logs warning level message.
func (l Logo) Warning(format string, args ...interface{}) {
	l.Output(2, LevelWarning, fmt.Sprintf(format, args...))
}

// Error logs Error level message.
func (l Logo) Error(format string, args ...interface{}) {
	l.Output(2, LevelError, fmt.Sprintf(format, args...))
}

// Critical log Critical level message.
func (l Logo) Critical(format string, args ...interface{}) {
	l.Output(2, LevelCritical, fmt.Sprintf(format, args...))
}
