package mbr

import "git.hexxed.me/hd/io"

// MBRNavigation is a int value that is used
// to reference various points in the MBR Partition Map.

var MBRStructure = io.NamedPointers{
	"BootableCode":   {0x0000, 440},
	"DiskSignature":  {0x01B8, 4},
	"Nulls":          {0x01BC, 2},
	"PartitionOne":   {0x01BE, 16},
	"PartitionTwo":   {0x01CE, 16},
	"PartitionThree": {0x01DE, 16},
	"PartitionFour":  {0x01EE, 16},
	"MBRSignature":   {0x01FE, 2},
}
