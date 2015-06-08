package main

import (
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

	part.SetType(0x83)
	h, s, c := part.GetStartingHSC()
	part.SetStartingHSC(h, s, c)
	//part.ParsePartition()

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
