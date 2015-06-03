package io

import (
	"fmt"
	"io"
)

type NamedPointers map[string][2]int64

func (np NamedPointers) AddPointer(name string, offset, length int) {
	np[name] = [2]int64{int64(offset), int64(length)}
}

func (np NamedPointers) GetPointerInfo(name string) (offset int64, length int) {
	return np[name][0], int(np[name][1])
}

func (np NamedPointers) GetPointersLen() (eoa int64) {
	for _, rng := range np {
		if eop := rng[0] + rng[1]; eop > eoa {
			eoa = eop
		}
	}
	return
}

func (np NamedPointers) SetPointers(ptrs NamedPointers) {
	for name, ranges := range ptrs {
		np[name] = ranges
	}
}

func (np NamedPointers) ReadPointer(
	reader io.ReaderAt, name string,
) (
	buffer []byte, numbytes int, err error,
) {
	buffer = make([]byte, int(np[name][1]))
	numbytes, err = reader.ReadAt(buffer, np[name][0])
	return
}

func (np NamedPointers) ReadAllPointers(reader io.ReaderAt) (
	buffers map[string][]byte, numbytes int, err error,
) {
	var pnumbytes int
	var errs []error

	buffers = make(map[string][]byte)

	for name := range np {
		buffers[name], pnumbytes, err = np.ReadPointer(reader, name)
		if err != nil {
			errs = append(errs, err)
		}

		numbytes = numbytes + pnumbytes
	}

	if len(errs) > 0 {
		err = errs[0]
		for _, e := range errs[1:] {
			err = fmt.Errorf("%s\n%s\n", err, e)
		}
	}

	return
}

func (np NamedPointers) WritePointer(
	writer io.WriterAt, name string, buffer []byte,
) (
	numbytes int, err error,
) {
	if len(buffer) > int(np[name][1]) {
		return 0, fmt.Errorf(
			"Buffer length limit is %d; received length of %d",
			np[name][1], len(buffer),
		)
	}

	return writer.WriteAt(buffer, np[name][0])
}
