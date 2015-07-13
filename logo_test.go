package logo

import (
	"log"
	"testing"
)

func TestLogo(t *testing.T) {
}

func BenchmarkNewLogo(b *testing.B) {
	benchmarkTask(b, func(i int) {
		NewLogo(LevelDebug)
	})
}

func BenchmarkLogo(b *testing.B) {
	var w nilWriter = 1
	l := NewLogo(LevelDebug, NewHandler(w, LevelDebug, log.LstdFlags))

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
