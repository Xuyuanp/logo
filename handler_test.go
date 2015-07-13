package logo

import (
	"fmt"
	"testing"
)

type nilWriter int

func (w nilWriter) Write([]byte) (n int, err error) { return }

func BenchmarkHandler(b *testing.B) {
	var w nilWriter = 1
	handler := NewHandler(w, LevelDebug, Ldate|Ltime|Lshortfile)

	benchmarkTask(b, func(i int) {
		handler.Handle(2, LevelDebug, fmt.Sprintf("%d", i))
	})
}
