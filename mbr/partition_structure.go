package mbr

import "git.hexxed.me/hd/io"

// MBRNavigation is a int value that is used
// to reference various points in the MBR Partition Map.

var PartitionStructure = io.NamedPointers{
	"BootableFlag": {0x00, 1},
	"StartingCHS":  {0x01, 3},
	"Type":         {0x05, 1},
	"EndingCHS":    {0x06, 3},
	"StartingLBA":  {0x09, 4},
	"EndingLBA":    {0x0C, 4},
}
