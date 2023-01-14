package mmap

import (
	"errors"
	"os"
)

const ANON = 1
const (
	// RDONLY maps the memory read-only.
	// Attempts to write to the MMap object will result in undefined behavior.
	RDONLY = 0
	// RDWR maps the memory as read-write. Writes to the MMap object will update the
	// underlying file.
	RDWR = 1 << iota
	// COPY maps the memory as copy-on-write. Writes to the MMap object will affect
	// memory, but the underlying file will remain unchanged.
	COPY
	// If EXEC is set, the mapped memory is marked as executable.
	EXEC
)

func Map(f *os.File, prot int) ([]byte, error) {
	return MmapOption(f, prot, -1, 0)
}
func MmapOption(f *os.File, prot, length int, off int64) ([]byte, error) {
	if off%int64(os.Getpagesize()) != 0 {
		return nil, errors.New("offset parameter must be a multiple of the system's page size")
	}
	if length < 0 {
		s, _ := f.Stat()
		length = int(s.Size())
	}
	return mmap_windows(length, uintptr(prot), f.Fd(), off)
}
