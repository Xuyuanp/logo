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
	"log"
	"testing"
)

type nilWriter int

func (w nilWriter) Write([]byte) (n int, err error) { return }

func BenchmarkLogo(b *testing.B) {
	var w nilWriter = 1
	l := NewLogo(LevelDebug, w, "", log.LstdFlags)

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
