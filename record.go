package logo

import (
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

var recordPool = sync.Pool{
	New: func() interface{} {
		return &Record{}
	},
}

type Record struct {
	Line    int
	Level   LogLevel
	Time    time.Time
	LFile   string
	SFile   string
	Message string
}

func NewRecord() *Record {
	rec := recordPool.Get().(*Record)
	rec.Reset()
	return rec
}

func (rec *Record) Reset() {
	rec.Line = 0
	rec.Level = LevelDebug
	rec.Message = ""
	rec.LFile = ""
	rec.SFile = ""
	rec.Time = time.Now()
}

func (rec *Record) Init(skip int, level LogLevel, message string) {
	_, lfile, line, ok := runtime.Caller(skip)
	if !ok {
		lfile = "??"
	}
	rec.Line = line
	rec.Level = level
	rec.LFile = lfile
	rec.SFile = filepath.Base(lfile)
	rec.Message = message
	rec.Time = time.Now()
}
