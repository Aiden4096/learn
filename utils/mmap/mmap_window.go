package mmap

import (
	"golang.org/x/sys/windows"
	"os"
	"reflect"
	"unsafe"
)

func mmap_windows(len int, prot, hfile uintptr, off int64) ([]byte, error) {
	flProtect := uint32(windows.PAGE_READONLY)
	dwDesiredAccess := uint32(windows.FILE_MAP_READ)
	switch {
	case prot&COPY != 0:
		flProtect = windows.PAGE_WRITECOPY
		dwDesiredAccess = windows.FILE_MAP_COPY
	case prot&RDWR != 0:
		flProtect = windows.PAGE_READWRITE
		dwDesiredAccess = windows.FILE_MAP_WRITE
	}
	if prot&EXEC != 0 {
		flProtect <<= 4
		dwDesiredAccess |= windows.FILE_MAP_EXECUTE
	}

	maxSizeHigh := uint32((off + int64(len)) >> 32)
	maxSizeLow := uint32((off + int64(len)) & 0xFFFFFFFF)

	h, errno := windows.CreateFileMapping(windows.Handle(hfile), nil, flProtect, maxSizeHigh, maxSizeLow, nil)
	if h == 0 {
		return nil, os.NewSyscallError("CreateFileMapping", errno)
	}

	fileOffsetHigh := uint32(off >> 32)
	fileOffsetLow := uint32(off & 0xFFFFFFFF)
	addr, errno := windows.MapViewOfFile(h, dwDesiredAccess, fileOffsetHigh, fileOffsetLow, uintptr(len))
	if addr == 0 {
		windows.CloseHandle(windows.Handle(h))
		return nil, os.NewSyscallError("MapViewOfFile", errno)
	}
	m := make([]byte, len)
	head := (*reflect.SliceHeader)(unsafe.Pointer(&m))
	head.Data = addr
	head.Len = len
	head.Cap = head.Len
	return m, nil
}
