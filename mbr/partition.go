package mbr

import (
	"fmt"

	"git.hexxed.me/hd/io"
	"github.com/davecgh/go-spew/spew"
)

type Partition struct {
	*io.NamedRange
}

func ReadPartition(nr *io.NamedRange) (part *Partition) {
	part = new(Partition)
	part.NamedRange = nr
	part.SetPointers(PartitionStructure)
	return
}

func (part *Partition) CheckForErrors() {
	_, numread, err := part.NamedRange.ReadAllPointers()
	if err != nil {
		panic(err)
	}

	fmt.Printf("\nNumbytes read: %d\n", numread)
}

func (part *Partition) ParsePartition() {
	partMap, numread, err := part.NamedRange.ReadAllPointers()
	if err != nil {
		panic(err)
	}

	fmt.Printf("\nNumbytes read: %d\n", numread)
	fmt.Printf("Buffer:\n%s\n", spew.Sdump(partMap))
}
