package logo

import (
	"log"
	"os"
	"testing"
)

func BenchmarkLogoFile(b *testing.B) {
	f, _ := OpenFile("/dev/null", 0644)
	l := New(LevelDebug, f, "", 0)

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		l.Debug("%d", i)
	}
}

func BenchmarkLogFile(b *testing.B) {
	f, _ := os.Open("/dev/null")
	l := log.New(f, "", 0)

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		l.Printf("%d", i)
	}
}
