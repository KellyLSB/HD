package io

import (
	"os"
	"sync"
	"syscall"
)

type ByteSize float64

const (
	_           = iota // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

type OptIO uint

const (
	Write OptIO = 1 << iota
	Synchronous
)

type DiskO struct {
	buffer []byte
	offset int64
	num    int
	err    error
}

type DiskI struct {
	replyto chan<- DiskO
	buffer  []byte
	offset  int64
	options OptIO
}

type Disk struct {
	fd int

	fileinfo os.FileInfo

	offset  int64
	queue   chan DiskI
	workers sync.WaitGroup
	closed  bool
}

func OpenDisk(file string) (disk *Disk, err error) {
	disk = new(Disk)

	disk.fileinfo, err = os.Lstat(file)
	if err != nil {
		return
	}

	disk.fd, err = syscall.Open(file, syscall.O_RDWR, 0777)
	if err != nil {
		return
	}

	disk.queue = make(chan DiskI)

	return
}

func (disk *Disk) ReadAt(buffer []byte, offset int64) (int, error) {
	disko := disk.Push(buffer, offset, Synchronous)
	return int(disko.num), disko.err
}

func (disk *Disk) WriteAt(buffer []byte, offset int64) (int, error) {
	disko := disk.Push(buffer, offset, Synchronous)
	return int(disko.num), disko.err
}

func (disk *Disk) Read(buffer []byte) (int, error) {
	disko := disk.Push(buffer, disk.offset, Synchronous)
	disk.offset = disko.offset // Update offset
	return int(disko.num), disko.err
}

func (disk *Disk) Write(buffer []byte) (int, error) {
	disko := disk.Push(buffer, disk.offset, Synchronous)
	disk.offset = disko.offset // Update offset
	return int(disko.num), disko.err
}

func (disk *Disk) Seek(offset int64) {
	disk.offset = offset
}

func (disk *Disk) Offset() int64 {
	return disk.offset
}

func (disk *Disk) Reset() {
	disk.offset = 0x0000
}

func (disk *Disk) Close() error {
	close(disk.queue)
	disk.closed = true
	disk.workers.Wait()
	return syscall.Close(disk.fd)
}

func (disk *Disk) Push(buffer []byte, offset int64, options OptIO) *DiskO {
	if disk.queue == nil {
		disk.queue = make(chan DiskI)
	}

	replyto := make(chan DiskO, 1)

	disk.queue <- DiskI{
		replyto,
		buffer,
		offset,
		options,
	}

	// @TODO check if this was the
	// right way to do the bitwise check
	// I always get it confused......
	if options&^Synchronous < 1 {
		disko := <-replyto
		return &disko
	}

	return nil
}

func (disk *Disk) Queue() <-chan DiskI {
	return disk.queue
}

func (disk *Disk) Worker() {
	disk.workers.Add(1)

	go func(queue <-chan DiskI) {
		defer disk.workers.Done()
		var err error
		var num int

		for job := range disk.queue {
			if job.options&^Write < 1 {
				num, err = syscall.Pwrite(disk.fd, job.buffer, job.offset)
			} else {
				num, err = syscall.Pread(disk.fd, job.buffer, job.offset)
			}

			if job.options&^Synchronous < 1 {
				job.replyto <- DiskO{
					job.buffer,
					job.offset + int64(num),
					num, err,
				}
			}

			if disk.closed {
				break
			}
		}
	}(disk.queue)
}
