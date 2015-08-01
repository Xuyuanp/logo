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
	"testing"
)

func fail(t *testing.T, msg string) {
	t.Errorf(msg)
}

func assertEqual(t *testing.T, actual, expected interface{}, args ...interface{}) {
	if actual != expected {
		msg := fmt.Sprint(args...)
		msg = fmt.Sprintf("Not equal(%s)! autual: %#v expected: %#v", msg, actual, expected)
		fail(t, msg)
	}
}

func assertTrue(t *testing.T, val bool, args ...interface{}) {
	if !val {
		msg := fmt.Sprint(args...)
		msg = fmt.Sprintf("Not true(%s)!", msg)
		fail(t, msg)
	}
}

func assertNotNil(t *testing.T, val interface{}, args ...interface{}) {
	if val == nil {
		msg := fmt.Sprint(args...)
		msg = fmt.Sprintf("Nil(%s)!", msg)
		fail(t, msg)
	}
}

func benchmarkTask(b *testing.B, task func(int)) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		task(i)
	}
}
