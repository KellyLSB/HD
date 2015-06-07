package mbr

import (
	"bytes"
	"fmt"
	"os"

	"git.hexxed.me/hd/io"
	"github.com/davecgh/go-spew/spew"
)

// MBR is a structure for assisting with parsing the
// Master Boot Record of a Disk using a io.NamedRange
type MBR struct {
	*io.NamedRange
}

// ReadMBR accepts a io.NamedRange and returns a MBR struct instance.
func ReadMBR(nr *io.NamedRange) (mbr *MBR) {
	mbr = new(MBR)
	mbr.NamedRange = nr
	mbr.SetPointers(MBRStructure)
	return
}

// GetPartition returns a Partition Record for each of the allowed
// four partitions available on the disk. (this is a limitation of MBR).
func (mbr *MBR) GetPartition(n int) (part *Partition) {
	switch n {
	case 1:
		return ReadPartition(mbr.GetChild("PartitionOne"))
	case 2:
		return ReadPartition(mbr.GetChild("PartitionTwo"))
	case 3:
		return ReadPartition(mbr.GetChild("PartitionThree"))
	case 4:
		return ReadPartition(mbr.GetChild("ParttionFour"))
	default:
		return nil
	}
}

func (mbr *MBR) CheckForErrors() {
	mbrMap, numread, err := mbr.NamedRange.ReadAllPointers()
	if err != nil {
		panic(err)
	}

	fmt.Printf("\nNumbytes read: %d\n", numread)

	if bytes.Equal(mbrMap["MBRSignature"], []byte{0xAA, 0x55}) {
		fmt.Fprintf(os.Stderr, "Err: MBR signature %#v is bad", mbrMap["DiskSignature"])
	}
}

func (mbr *MBR) ParseMBR() {
	mbrMap, numread, err := mbr.NamedRange.ReadAllPointers()
	if err != nil {
		panic(err)
	}

	fmt.Printf("\nNumbytes read: %d\n", numread)
	fmt.Printf("Buffer:\n%s\n", spew.Sdump(mbrMap))
}
