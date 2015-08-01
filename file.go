/*
 * Copyright 2015 Xuyuan Pang
 * Author: Xuyuan Pang
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package logo

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
)

// RotatedFile is a os.File wrapper to log message. All of its methods
// are goroutine-safe.
type RotatedFile struct {
	name string
	file *os.File
	flag int
	mode os.FileMode
	mu   sync.Mutex
}

// OpenFile creates a new RotatedFile instance, and open the file.
func OpenFile(name string, mode os.FileMode) (*RotatedFile, error) {
	rf := &RotatedFile{
		name: name,
		flag: os.O_APPEND | os.O_WRONLY | os.O_CREATE,
		mode: mode,
	}
	return rf, rf.Open()
}

// Open opens file using the file name and mode, if the file is already opening,
// close it and reopen.
func (rf *RotatedFile) Open() error {
	rf.mu.Lock()
	defer rf.mu.Unlock()
	if rf.file != nil {
		rf.file.Close()
	}
	file, err := os.OpenFile(rf.name, rf.flag, rf.mode)
	if err != nil {
		return err
	}
	rf.file = file
	return nil
}

// Listen listens the provided signals, if any of these singals was
// received, reopen the file. This method is used for logrotate and
// it will block the current goroutine. If you don't call this method,
// it works as the same as normal *os.File.
func (rf *RotatedFile) Listen(sig ...os.Signal) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, sig...)
	for {
		<-c
		rf.Open()
	}
}

// Write writes data by using the inner *os.File.
func (rf *RotatedFile) Write(b []byte) (n int, err error) {
	rf.mu.Lock()
	if rf.file != nil {
		n, err := rf.file.Write(b)
		rf.mu.Unlock()
		return n, err
	}
	rf.mu.Unlock()
	return 0, fmt.Errorf("file not opened")
}

// Close close the inner *os.File if it's not nil.
func (rf *RotatedFile) Close() error {
	rf.mu.Lock()
	if rf.file == nil {
		rf.mu.Unlock()
		return nil
	}
	err := rf.file.Close()
	rf.file = nil
	rf.mu.Unlock()
	return err
}
