package logo

import "testing"

type nilWriter int

func (w nilWriter) Write([]byte) (n int, err error) { return }

func BenchmarkHandler(b *testing.B) {
	rec := NewRecord()
	defer recordPool.Put(rec)
	rec.Init(1, LevelDebug, "test")

	var w nilWriter = 1
	handler := WriterHandler(w, DefaultFormatter)

	benchmarkTask(b, func(i int) {
		handler.Handle(rec)
	})
}
