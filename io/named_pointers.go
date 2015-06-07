package io

import (
	"fmt"
	goio "io"
)

type NamedPointers map[string][2]int64

func (np NamedPointers) AddPointer(
	name string, offset, length int64,
) {
	np[name] = [2]int64{offset, length}
}

func (np NamedPointers) SetPointers(ptrs NamedPointers) {
	for name, ranges := range ptrs {
		np.AddPointer(name, ranges[0], ranges[1])
	}
}

func (np NamedPointers) GetPointerInfo(name string) (offset int64, length int) {
	if np == nil || len(np[name]) != 2 {
		panic(fmt.Errorf("Cannot find pointer %s\n", name))
	}

	return np[name][0], int(np[name][1])
}

func (np NamedPointers) ReadPointer(
	reader goio.ReaderAt, name string,
) (
	buffer []byte, numbytes int, err error,
) {
	offset, length := np.GetPointerInfo(name)

	buffer = make([]byte, length)
	numbytes, err = reader.ReadAt(buffer, offset)
	if err != nil {
		panic(err)
	}

	return
}

func (np NamedPointers) ReadAllPointers(reader goio.ReaderAt) (
	buffers map[string][]byte, numbytes int, err error,
) {
	var pnumbytes int

	buffers = make(map[string][]byte)

	for name := range np {
		buffers[name], pnumbytes, err = np.ReadPointer(reader, name)
		if err != nil {
			panic(err)
		}

		numbytes = numbytes + pnumbytes
	}

	return
}

func (np NamedPointers) WritePointer(
	writer goio.WriterAt,
	name string,
	buffer []byte,
) (
	numbytes int,
	err error,
) {
	offset, length := np.GetPointerInfo(name)

	if len(buffer) > length {
		panic(fmt.Errorf(
			"Buffer length limit is %d; received length of %d\n",
			length, len(buffer),
		))
	}

	return writer.WriteAt(buffer, offset)
}
