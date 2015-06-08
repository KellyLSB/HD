package mbr

import (
	"bytes"
	"fmt"

	"git.hexxed.me/hd/io"
	"github.com/davecgh/go-spew/spew"
)

// @IMPORTANT:
// 							- Fix setters of the

// Partition is a structure for
// manipulating a partition entry in the MBR tables.
type Partition struct {
	*io.NamedRange
}

// ReadPartition returns a new *Partition instance struct
// from the given *io.NamdRange; generally from *MBR.
func ReadPartition(nr *io.NamedRange) (part *Partition) {
	part = new(Partition)
	part.NamedRange = nr
	part.SetPointers(PartitionStructure)
	return
}

// CheckForErrors is a function reserved for
// running checks on the partition values as needed.
func (part *Partition) CheckForErrors() {
	_, numread, err := part.NamedRange.ReadAllPointers()
	if err != nil {
		panic(err)
	}

	fmt.Printf("\nNumbytes read: %d\n", numread)
}

// ParsePartition is used to dump all the pointers.
func (part *Partition) ParsePartition() {
	partMap, numread, err := part.NamedRange.ReadAllPointers()
	if err != nil {
		panic(err)
	}

	fmt.Printf("\nNumbytes read: %d\n", numread)
	fmt.Printf("Buffer:\n%s\n", spew.Sdump(partMap))
}

// IsBootable returns true or false depending on the state of BootableFlag.
func (part *Partition) IsBootable() bool {
	buffer, _, _ := part.NamedRange.ReadPointer("BootableFlag")
	return bytes.Equal(buffer, []byte{0x80})
}

// SetBootable sets the Partition as bootable.
func (part *Partition) SetBootable() {
	part.NamedRange.WritePointer("BootableFlag", []byte{0x80})
}

// UnSetBootable unsts the Partition as bootable.
func (part *Partition) UnSetBootable() {
	part.NamedRange.WritePointer("BootableFlag", []byte{0x00})
}

// GetStartingHSC returns the starting
// Head, Sector and Cylinder values for the partition.
func (part *Partition) GetStartingHSC() (head, sector byte, cylinder []byte) {
	return part.getHSC("Starting")
}

// SetStartingHSC sets the starting
// Cylinder, Head and Sector values for the parition.
func (part *Partition) SetStartingHSC(head, sector byte, cylinder []byte) {
	part.setHSC("Starting", head, sector, cylinder)
}

// GetType returns the PartitionType.
func (part *Partition) GetType() PartitionType {
	buffer, _, _ := part.NamedRange.ReadPointer("Type")
	return PartitionType(buffer[0])
}

// SetType sets the PartitionType.
func (part *Partition) SetType(pt PartitionType) {
	part.NamedRange.WritePointer("Type", []byte{byte(pt)})
}

// GetEndingHSC returns the ending
// Cylinder, Head and Sector values for the partition.
func (part *Partition) GetEndingHSC() (head, sector byte, cylinder []byte) {
	return part.getHSC("Ending")
}

// SetEndingHSC sets the ending
// Cylinder, Head and Sector values for the parition.
func (part *Partition) SetEndingHSC(head, sector byte, cylinder []byte) {
	part.setHSC("Ending", head, sector, cylinder)
}

func (part *Partition) getHSC(se string) (head, sector byte, cylinder []byte) {
	buffer, _, _ := part.NamedRange.ReadPointer(se + "CHS")
	return buffer[0], (buffer[1] | 0xC0) ^ 0xC0, []byte{buffer[1] >> 6, buffer[2]}
}

func (part *Partition) setHSC(se string, head, sector byte, cylinder []byte) {
	part.NamedRange.WritePointer(se+"CHS", []byte{
		head, sector | (cylinder[0] << 6), cylinder[1],
	})
}
