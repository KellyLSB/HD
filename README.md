# HD Tools

A library for manipulating block devices, partition tables and partitions; coded from the ground up to be easy to understand, easy to use and provide (hopefully) a educational view into block storage and byte manipulation.

## Sample Output

    $ sudo ./hd_linux_amd64 /dev/mmcblk0
    Numbytes read: 512
    Buffer:
    (map[string][]uint8) (len=8) {
      (string) (len=12) "MBRSignature": ([]uint8) (len=2 cap=2) {
        00000000  55 aa                                             |U.|
      },
      (string) (len=12) "BootableCode": ([]uint8) (len=440 cap=440) {
        00000000  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
        00000010  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
        00000020  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
        00000030  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
        00000040  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
        00000050  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
        00000060  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
        00000070  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
        00000080  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
        00000090  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
        000000a0  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
        000000b0  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
        000000c0  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
        000000d0  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
        000000e0  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
        000000f0  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
        00000100  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
        00000110  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
        00000120  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
        00000130  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
        00000140  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
        00000150  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
        00000160  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
        00000170  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
        00000180  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
        00000190  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
        000001a0  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
        000001b0  00 00 00 00 00 00 00 00                           |........|
      },
      (string) (len=13) "DiskSignature": ([]uint8) (len=4 cap=4) {
        00000000  e1 37 a7 e0                                       |.7..|
      },
      (string) (len=5) "Nulls": ([]uint8) (len=2 cap=2) {
        00000000  00 00                                             |..|
      },
      (string) (len=12) "PartitionOne": ([]uint8) (len=16 cap=16) {
        00000000  00 00 01 20 0b 03 10 1f  00 08 00 00 00 00 01 00  |... ............|
      },
      (string) (len=12) "PartitionTwo": ([]uint8) (len=16 cap=16) {
        00000000  00 00 01 20 83 03 10 8f  00 08 01 00 00 1c 98 03  |... ............|
      },
      (string) (len=14) "PartitionThree": ([]uint8) (len=16 cap=16) {
        00000000  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
      },
      (string) (len=13) "PartitionFour": ([]uint8) (len=16 cap=16) {
        00000000  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
      }
    }

# Changelog

2015-06-01 - Initial Concept, MBR basic structure and named pointers - Kelly B.
2015-06-02 - Added queue, synchronicity and workers to Disk - Kelly B.
