package logo

import (
	"fmt"
	"log"
	"testing"
)

type nilWriter int

func (w nilWriter) Write([]byte) (n int, err error) { return }

func BenchmarkHandler(b *testing.B) {
	var w nilWriter = 1
	handler := NewHandler(w, LevelDebug, LstdFlags)

	benchmarkTask(b, func(i int) {
		handler(2, LevelDebug, fmt.Sprintf("%d", i))
	})
}

func BenchmarkLogOutput(b *testing.B) {
	var w nilWriter = 1
	l := log.New(w, "", log.LstdFlags)

	benchmarkTask(b, func(i int) {
		l.Output(2, fmt.Sprintf("%d", i))
	})
}
