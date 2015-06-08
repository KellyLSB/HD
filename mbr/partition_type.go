package mbr

// PartitionType defines the proper hex value
// for a partition type.
type PartitionType byte

const (
	// Fat12 partition
	Fat12 PartitionType = 0x01
	// Fat16lt32 partition
	Fat16lt32 PartitionType = 0x04
	// Extended partition
	Extended PartitionType = 0x05
	// Fat16 partition
	Fat16 PartitionType = 0x06
	// NTFS partition
	NTFS PartitionType = 0x07
	// Win95Fat32 partition
	Win95Fat32 PartitionType = 0x0B
	// Win95Fat32LBA partition
	Win95Fat32LBA PartitionType = 0x0C
	// Win95Fat16LBA partition
	Win95Fat16LBA PartitionType = 0x0E
	// Win95ExtendedLBA partition
	Win95ExtendedLBA PartitionType = 0x0F
	// HiddenFat12 partition
	HiddenFat12 PartitionType = 0x11
	// HiddenFAT16lt32 partition
	HiddenFAT16lt32 PartitionType = 0x14
	// HiddenFAT16 partition
	HiddenFAT16 PartitionType = 0x16
	// HiddenNTFS partition
	HiddenNTFS PartitionType = 0x17
	// HiddenWIN95FAT32 partition
	HiddenWIN95FAT32 PartitionType = 0x1B
	// HiddenWIN95FAT32LBA partition
	HiddenWIN95FAT32LBA PartitionType = 0x1C
	// HiddenWIN95FAT16LBA partition
	HiddenWIN95FAT16LBA PartitionType = 0x1E
	// LinuxSwap partition
	LinuxSwap PartitionType = 0x82
	// Linux partition
	Linux PartitionType = 0x83
	// LinuxExtended partition
	LinuxExtended PartitionType = 0x85
	// NTFSVolumeSet1 partition
	NTFSVolumeSet1 PartitionType = 0x86
	// NTFSVolumeSet2 partition
	NTFSVolumeSet2 PartitionType = 0x87
	// LinuxLVM partition
	LinuxLVM PartitionType = 0x8E
	// BSDOS partition
	BSDOS PartitionType = 0x9f
	// FreeBSD partition
	FreeBSD PartitionType = 0xa5
	// OpenBSD partition
	OpenBSD PartitionType = 0xa6
	// NetBSD partition
	NetBSD PartitionType = 0xa9
	// BeOSfs partition
	BeOSfs PartitionType = 0xeb
	// EFIGPT partition
	EFIGPT PartitionType = 0xee
	// EFIFAT partition
	EFIFAT PartitionType = 0xef
)
