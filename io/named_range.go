package io

import "fmt"

// NamedRange is a struct used to flexibly provide
// key value access to a ReadWriterAt struct.
type NamedRange struct {
	ReadWriterAt

	name               string
	ptrs               NamedPointers
	start, end, offset int64
}

// NewNamedRange creates a *NamedRange to flexibly provide
// key value access to a ReadWriterAt struct.
func NewNamedRange(rw ReadWriterAt, name string, offset ...int64) (nr *NamedRange) {
	nr = &NamedRange{ReadWriterAt: rw, name: name, ptrs: make(NamedPointers)}

	// arg 2 is not an offset value; rather it
	// is the length of the ReadWriterAt window.
	if len(offset) == 2 {
		nr.start = offset[0]
		nr.end = nr.start + offset[1]
	}

	return
}

// GetChild returns a pointer as a child *NamedRange
func (nr *NamedRange) GetChild(name string) *NamedRange {
	offset, length := nr.ptrs.GetPointerInfo(name)
	return NewNamedRange(nr.ReadWriterAt, name, nr.start+offset, int64(length))
}

// Name returns the name of the NamedRange
func (nr *NamedRange) Name() string {
	return nr.name
}

// Len returns the length of the NamedRange
func (nr *NamedRange) Len() int {
	if nr.start == 0 && nr.end == 0 {
		return -1
	}

	return int(nr.end - nr.start)
}

// Seek sets the offset or Read() and Write() to 0.
func (nr *NamedRange) Seek(n int) {
	nr.offset = int64(n)
}

// Reset resets the offset or Read() and Write() to 0.
func (nr *NamedRange) Reset() {
	nr.offset = 0
}

// Read fills the provided buffer with data from the ReadWriterAt
// until the length of the provided buffer has been filled.
func (nr *NamedRange) Read(buffer []byte) (n int, err error) {
	n, err = nr.ReadWriterAt.ReadAt(buffer, nr.start+nr.offset)
	nr.offset = nr.offset + int64(n)
	return
}

// Write takes the provided buffer and writes
// the value of the entire struct's ReadWriterAt at.
func (nr *NamedRange) Write(buffer []byte) (n int, err error) {
	n, err = nr.ReadWriterAt.WriteAt(buffer, nr.start+nr.offset)
	nr.offset = nr.offset + int64(n)
	return
}

// ReadAt fills the provided buffer with data from the ReadWriterAt
// until the length of the buffer has been filled; using the provided offset.
func (nr *NamedRange) ReadAt(buffer []byte, offset int64) (n int, err error) {
	nr.checkRange(len(buffer))

	return nr.ReadWriterAt.ReadAt(buffer, nr.start+offset)
}

// WriteAt writes the provided buffer to
// the ReadWriterAt; using the provided offset.
func (nr *NamedRange) WriteAt(buffer []byte, offset int64) (n int, err error) {
	nr.checkRange(len(buffer))

	return nr.ReadWriterAt.WriteAt(buffer, nr.start+offset)
}

// SetPointers bulk adds pointers to NamedPointers.
// this is useful for decoding bytes of saved information into key value stores.
func (nr *NamedRange) SetPointers(pointers NamedPointers) {
	nr.ptrs.SetPointers(pointers)
}

// ReadPointer returns the value of the named pointer
func (nr *NamedRange) ReadPointer(name string) ([]byte, int, error) {
	return nr.ptrs.ReadPointer(nr, name)
}

// ReadAllPointers returns all pointers and values as map[string][]byte.
func (nr *NamedRange) ReadAllPointers() (map[string][]byte, int, error) {
	return nr.ptrs.ReadAllPointers(nr)
}

// WritePointer writes the buffer to the pointer value
func (nr *NamedRange) WritePointer(name string, buffer []byte) (int, error) {
	return nr.ptrs.WritePointer(nr, name, buffer)
}

func (nr *NamedRange) checkRange(size int) {
	switch {
	case nr.Len() == -1:
		return
	case size > nr.Len():
		panic(fmt.Errorf("Cannot read/write from a larger window then provided by the data source."))
	}
}
