package logo

import "testing"

func TestRecord(t *testing.T) {
	rec := NewRecord()
	assertNotNil(t, rec)

	rec.Init(1, LevelDebug, "test")

	assertEqual(t, rec.Line, 9)
	assertEqual(t, rec.SFile, "record_test.go")
	assertEqual(t, rec.Message, "test")
	assertEqual(t, rec.Level, LevelDebug)
}

func BenchmarkRecord(b *testing.B) {
	benchmarkTask(b, func(i int) {
		rec := NewRecord()
		rec.Init(1, LevelDebug, "test")
		defer recordPool.Put(rec)
	})
}
