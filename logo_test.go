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
	"bytes"
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"
)

type nilWriter int

func (w nilWriter) Write([]byte) (n int, err error) { return }

func TestOutput(t *testing.T) {
	teststr := "test"
	var buf bytes.Buffer
	l := New(LevelDebug, &buf, "", 0)
	l.Output(1, LevelDebug, teststr)
	assertEqual(t, buf.String(), teststr+"\n")
}

func TestAPI(t *testing.T) {
	var buf bytes.Buffer
	l := New(LevelDebug, &buf, "", 0)
	logMethods := [...]func(string, ...interface{}){
		l.Debug,
		l.Info,
		l.Warning,
		l.Error,
		l.Critical,
	}

	teststr := "test"
	for _, m := range logMethods {
		buf.Reset()
		m(teststr)
		assertEqual(t, buf.String(), teststr+"\n")
	}
}

func TestGroup(t *testing.T) {
	var buf1, buf2 bytes.Buffer
	l1 := New(LevelDebug, &buf1, "", 0)
	l2 := New(LevelError, &buf2, "", 0)
	l := Group(LevelInfo, l1, l2)

	teststr := "test"

	l.Log(LevelDebug, teststr)
	assertEqual(t, buf1.String(), "")
	assertEqual(t, buf2.String(), "")

	buf1.Reset()
	buf2.Reset()
	l.Log(LevelInfo, teststr)
	assertEqual(t, buf1.String(), teststr+"\n")
	assertEqual(t, buf2.String(), "")

	buf1.Reset()
	buf2.Reset()
	l.Log(LevelError, teststr)
	assertEqual(t, buf1.String(), teststr+"\n")
	assertEqual(t, buf2.String(), teststr+"\n")
}

func TestLevel(t *testing.T) {
	levels := [...]LogLevel{LevelDebug, LevelInfo, LevelWarning, LevelError, LevelCritical}
	teststr := "test"
	for _, level := range levels {
		var buf bytes.Buffer
		l := New(level, &buf, "", 0)

		for _, lvl := range levels {
			buf.Reset()
			l.Log(lvl, teststr)
			if lvl < level {
				assertEqual(t, buf.String(), "")
			} else {
				assertEqual(t, buf.String(), teststr+"\n")
			}
		}
	}
}

func TestFlagLevel(t *testing.T) {
	levels := [...]LogLevel{LevelDebug, LevelInfo, LevelWarning, LevelError, LevelCritical}
	var buf bytes.Buffer
	teststr := "test"
	l := New(LevelDebug, &buf, "", Llevel)

	for _, lvl := range levels {
		buf.Reset()
		l.Log(lvl, teststr)
		assertEqual(t, buf.String(), lvl.String()+" "+teststr+"\n")
	}
}

func TestFlagDate(t *testing.T) {
	teststr := "test"
	var buf bytes.Buffer
	l := New(LevelDebug, &buf, "", Ldate)
	now := time.Now()
	l.Debug(teststr)
	prefix := now.Format("2006/01/02")
	assertTrue(t, strings.HasPrefix(buf.String(), prefix), buf.String(), prefix)
}

func TestFlagTime(t *testing.T) {
	teststr := "test"
	var buf bytes.Buffer
	l := New(LevelDebug, &buf, "", Ltime|Lmicroseconds)
	l.Debug(teststr)
}

func TestFlagColor(t *testing.T) {
	teststr := "test"
	var buf bytes.Buffer
	l := New(LevelDebug, &buf, "", Lcolor)
	l.Debug(teststr)
}

func TestFlagFile(t *testing.T) {
	teststr := "test"
	var buf bytes.Buffer
	l := New(LevelDebug, &buf, "", Lshortfile)
	l.Debug(teststr)
	_, file, line, ok := runtime.Caller(0)
	if !ok {
		file = "???"
		line = 0
	}
	sfile := filepath.Base(file)
	prefix := fmt.Sprintf("%s:%d:", sfile, line-1)
	assertTrue(t, strings.HasPrefix(buf.String(), prefix), buf.String(), prefix)
}

func TestPrefix(t *testing.T) {
	teststr := "test"
	var buf bytes.Buffer
	l := New(LevelDebug, &buf, "prefix ", 0)
	l.Debug(teststr)

	assertEqual(t, buf.String(), "prefix "+teststr+"\n")
}

func TestGlobalAPI(t *testing.T) {
	apis := [...]func(string, ...interface{}){
		Debug, Info, Warning, Error, Critical,
	}
	var w nilWriter = 1
	std = New(LevelDebug, w, "", 0)
	for _, api := range apis {
		api("test")
	}
}

func BenchmarkLogo(b *testing.B) {
	var w nilWriter = 1
	l := New(LevelDebug, w, "", log.LstdFlags)

	benchmarkTask(b, func(i int) {
		l.Debug("%d", i)
	})
}

func BenchmarkLog(b *testing.B) {
	var w nilWriter = 1
	logger := log.New(w, "", log.LstdFlags)

	benchmarkTask(b, func(i int) {
		logger.Printf("%d", i)
	})
}
