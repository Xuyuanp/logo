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
