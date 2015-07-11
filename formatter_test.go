package logo

import "testing"

func TestFormatter(t *testing.T) {
	rec := NewRecord()
	defer recordPool.Put(rec)
	rec.Init(1, LevelDebug, "test")

	t.Log(DefaultFormatter.Format(rec))
}

func BenchmarkFormatter(b *testing.B) {
	rec := NewRecord()
	defer recordPool.Put(rec)
	rec.Init(1, LevelDebug, "test")

	benchmarkTask(b, func(i int) {
		DefaultFormatter.Format(rec)
	})
}
