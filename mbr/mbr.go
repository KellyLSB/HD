package mbr

import (
	"bytes"
	"fmt"
	"os"

	"git.hexxed.me/hd/io"
	"github.com/davecgh/go-spew/spew"
)

type MBR struct {
	*io.NamedRange
	disk *io.Disk
}

func OpenDisk(disk string) (mbr *MBR, err error) {
	mbr = new(MBR)

	mbr.disk, err = io.OpenDisk(disk)
	if err != nil {
		return
	}

	mbr.NamedRange = io.NewNamedRange(mbr.disk, "")
	mbr.SetPointers(MBRStructure)
	return
}

func (mbr *MBR) ParseMBR() {
	buffer, numread, err := mbr.NamedRange.ReadAllPointers()
	if err != nil {
		//panic(err)
	}

	if bytes.Equal(buffer["MBRSignature"], []byte{0xAA, 0x55}) {
		fmt.Fprintf(os.Stderr, "MBR signature %#v is bad", buffer["DiskSignature"])
	}

	fmt.Printf("\nNumbytes read: %d\n", numread)
	fmt.Printf("Buffer:\n%s\n", spew.Sdump(buffer))

	part := mbr.GetChild("PartitionTwo")
	part.SetPointers(PartitionStructure)
	fmt.Printf("\nMBR:\n%s\n", spew.Sdump(part))
	buffer, numread, err = part.ReadAllPointers()
	fmt.Printf("\nNumbytes read: %d\n", numread)
	fmt.Printf("Buffer:\n%s\n", spew.Sdump(buffer))
}

func (mbr *MBR) Worker() {
	mbr.disk.Worker()
}

func (mbr *MBR) Close() {
	mbr.disk.Close()
}
