package main

import (
	"fmt"
	"os"

	"git.hexxed.me/hd/io"
	"git.hexxed.me/hd/mbr"
)

func main() {
	disk, err := io.OpenDisk(os.Args[1])
	if err != nil {
		panic(err)
	}

	disk.Worker()

	mbr := mbr.ReadMBR(io.NewNamedRange(disk, ""))
	mbr.ParseMBR()

	part := mbr.GetPartition(1)
	part.ParsePartition()

	fmt.Printf("\n%+v\n", part.GetType())
	part.GetStartingHSC()

	//part.SetStartingHSC(h, s, c)
	part.SetStartingHSC(0x3f, 0xFE, []byte{0xd3, 0x02})
	part.ParsePartition()

	// part = mbr.GetPartition(2)
	// part.ParsePartition()
	// fmt.Printf("\n%+v\n", part.GetType())
	//
	// part = mbr.GetPartition(3)
	// part.ParsePartition()
	// fmt.Printf("\n%+v\n", part.GetType())
	//
	// part = mbr.GetPartition(4)
	// part.ParsePartition()
	// fmt.Printf("\n%+v\n", part.GetType())

	disk.Close()
}
