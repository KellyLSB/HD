package main

import (
	"os"

	"git.hexxed.me/hd/mbr"
)

func main() {
	disk, err := mbr.OpenDisk(os.Args[1])
	if err != nil {
		panic(err)
	}

	disk.Worker()
	disk.ParseMBR()
	disk.Close()
}
