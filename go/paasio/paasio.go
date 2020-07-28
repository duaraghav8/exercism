package paasio

import (
	"io"
	"sync/atomic"
)

// readCounter tracks read operations
type readCounter struct {
	reader    io.Reader
	opCount   int32
	bytesRead int64
}

func (r *readCounter) Read(p []byte) (n int, err error) {
	n, err = r.reader.Read(p)

	atomic.AddInt32(&r.opCount, 1)
	atomic.AddInt64(&r.bytesRead, int64(n))

	return n, err
}

func (r *readCounter) ReadCount() (n int64, nops int) {
	return r.bytesRead, int(r.opCount)
}

// NewReadCounter is a factory for ReadCounter
func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{reader: reader, bytesRead: 0, opCount: 0}
}

// writeCounter tracks write operations
type writeCounter struct {
	writer       io.Writer
	opCount      int32
	bytesWritten int64
}

func (w *writeCounter) Write(p []byte) (n int, err error) {
	n, err = w.writer.Write(p)

	atomic.AddInt32(&w.opCount, 1)
	atomic.AddInt64(&w.bytesWritten, int64(n))

	return n, err
}

func (w *writeCounter) WriteCount() (n int64, nops int) {
	return w.bytesWritten, int(w.opCount)
}

// NewWriteCounter is a factory for WriteCounter
func NewWriteCounter(writer io.Writer) WriteCounter {
	return &writeCounter{writer: writer, bytesWritten: 0, opCount: 0}
}

// readWriteCounter tracks all read-write operations
type readWriteCounter struct {
	rc ReadCounter
	wc WriteCounter
}

func (rw *readWriteCounter) Read(p []byte) (n int, err error) {
	return rw.rc.Read(p)
}

func (rw *readWriteCounter) ReadCount() (n int64, nops int) {
	return rw.rc.ReadCount()
}

func (rw *readWriteCounter) Write(p []byte) (n int, err error) {
	return rw.wc.Write(p)
}

func (rw *readWriteCounter) WriteCount() (n int64, nops int) {
	return rw.wc.WriteCount()
}

// NewReadWriteCounter is a factory for ReadWriteCounter
func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	return &readWriteCounter{
		rc: NewReadCounter(rw),
		wc: NewWriteCounter(rw),
	}
}
