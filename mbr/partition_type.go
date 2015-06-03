package mbr

type PartitionType int

const (
	Fat12               PartitionType = 0x01
	Fat16lt32           PartitionType = 0x04
	Extended            PartitionType = 0x05
	Fat16               PartitionType = 0x06
	NTFS                PartitionType = 0x07
	Win95Fat32          PartitionType = 0x0B
	Win95Fat32LBA       PartitionType = 0x0C
	Win95Fat16LBA       PartitionType = 0x0E
	Win95ExtendedLBA    PartitionType = 0x0F
	HiddenFat12         PartitionType = 0x11
	HiddenFAT16lt32     PartitionType = 0x14
	HiddenFAT16         PartitionType = 0x16
	HiddenNTFS          PartitionType = 0x17
	HiddenWIN95FAT32    PartitionType = 0x1B
	HiddenWIN95FAT32LBA PartitionType = 0x1C
	HiddenWIN95FAT16LBA PartitionType = 0x1E
	LinuxSwap           PartitionType = 0x82
	Linux               PartitionType = 0x83
	LinuxExtended       PartitionType = 0x85
	NTFSVolumeSet1      PartitionType = 0x86
	NTFSVolumeSet2      PartitionType = 0x87
	LinuxLVM            PartitionType = 0x8E
	BSDOS               PartitionType = 0x9f
	FreeBSD             PartitionType = 0xa5
	OpenBSD             PartitionType = 0xa6
	NetBSD              PartitionType = 0xa9
	BeOSfs              PartitionType = 0xeb
	EFIGPT              PartitionType = 0xee
	EFIFAT              PartitionType = 0xef
)
