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

	part := mbr.GetPartition(2)
	part.ParsePartition()

	disk.Close()
}
