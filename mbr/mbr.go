package mbr

import (
	"bytes"
	"fmt"
	"os"

	"git.hexxed.me/hd/io"
	"github.com/davecgh/go-spew/spew"
)

type MBR struct {
	io.Disk
}

func OpenDisk(disk string) (mbr *MBR, err error) {
	iodisk, err := io.OpenDisk(disk)
	if err != nil {
		return
	}

	mbr = new(MBR)
	mbr.Disk = *iodisk
	mbr.SetPointers(MBRStructure)
	return
}

func (mbr *MBR) ParseMBR() {
	buffer, numread, err := mbr.ReadAllPointers()
	if err != nil {
		panic(err)
	}

	if bytes.Equal(buffer["MBRSignature"], []byte{0xAA, 0x55}) {
		fmt.Fprintf(os.Stderr, "MBR signature %#v is bad", buffer["DiskSignature"])
	}

	fmt.Printf("Numbytes read: %d\n", numread)
	fmt.Printf("Buffer:\n%s\n", spew.Sdump(buffer))
}
