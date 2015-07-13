package logo

import (
	"log"
	"testing"
)

func TestLogo(t *testing.T) {
}

func BenchmarkLogo(b *testing.B) {
	var w nilWriter = 1
	handler := NewHandler(w, LevelDebug, log.LstdFlags)
	l := &Logo{}
	l.SetHandlers(handler)

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
