package logo

import "testing"

func TestLogo(t *testing.T) {
}

func BenchmarkLogo(b *testing.B) {
	var w nilWriter = 1
	handler := WriterHandler(w, DefaultFormatter)
	l := &Logo{}
	l.SetHandlers(handler)

	benchmarkTask(b, func(i int) {
		l.Debug("test")
	})
}
