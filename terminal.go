package main

import (
	"syscall"
	"unsafe"
)

type winsize struct {
	row    uint16
	col    uint16
	Xpixel uint16
	Ypixel uint16
}

func GetWidth() int {
	ws := &winsize{}

	syscall.Syscall(
		syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)),
	)

	return int(ws.col)
}
