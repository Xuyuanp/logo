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
	"sync"
)

var (
	std  = NewStd(LevelDebug, "", LdefaultFlags)
	once sync.Once
)

// Std returns the default global Logger
func Std() Logger {
	return std
}

func setupGlobal(logger Logger) {
	std = logger
}

func SetupGlobal(logger Logger) {
	once.Do(func() {
		std = logger
	})
}

// Debug logs Debug level message.
func Debug(format string, args ...interface{}) {
	std.Output(2, LevelDebug, fmt.Sprintf(format, args...))
}

// Info logs Info level message
func Info(format string, args ...interface{}) {
	std.Output(2, LevelInfo, fmt.Sprintf(format, args...))
}

// Warning logs Warning level message.
func Warning(format string, args ...interface{}) {
	std.Output(2, LevelWarning, fmt.Sprintf(format, args...))
}

// Error logs Error level message.
func Error(format string, args ...interface{}) {
	std.Output(2, LevelError, fmt.Sprintf(format, args...))
}

// Critical logs Critical level message
func Critical(format string, args ...interface{}) {
	std.Output(2, LevelCritical, fmt.Sprintf(format, args...))
}
