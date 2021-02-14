// +build windows

package color

import (
	"syscall"
	"unsafe"
)

var fgColors = []uint16{
	0,
	0,
	foregroundRed,
	foregroundGreen,
	foregroundRed | foregroundGreen,
	foregroundBlue,
	foregroundRed | foregroundBlue,
	foregroundGreen | foregroundBlue,
	foregroundRed | foregroundGreen | foregroundBlue}

var bgColors = []uint16{
	0,
	0,
	backgroundRed,
	backgroundGreen,
	backgroundRed | backgroundGreen,
	backgroundBlue,
	backgroundRed | backgroundBlue,
	backgroundGreen | backgroundBlue,
	backgroundRed | backgroundGreen | backgroundBlue}

const (
	foregroundBlue      = uint16(0x0001)
	foregroundGreen     = uint16(0x0002)
	foregroundRed       = uint16(0x0004)
	foregroundIntensity = uint16(0x0008)
	backgroundBlue      = uint16(0x0010)
	backgroundGreen     = uint16(0x0020)
	backgroundRed       = uint16(0x0040)
	backgroundIntensity = uint16(0x0080)

	foregroundMask = foregroundBlue | foregroundGreen | foregroundRed | foregroundIntensity
	backgroundMask = backgroundBlue | backgroundGreen | backgroundRed | backgroundIntensity
)

var (
	kernel32 = syscall.NewLazyDLL("kernel32.dll")

	procGetStdHandle               = kernel32.NewProc("GetStdHandle")
	procSetConsoleTextAttribute    = kernel32.NewProc("SetConsoleTextAttribute")
	procGetConsoleScreenBufferInfo = kernel32.NewProc("GetConsoleScreenBufferInfo")

	hStdout        uintptr
	initScreenInfo *consoleScreenBufferInfo
)

func setConsoleTextAttribute(hConsoleOutput uintptr, wAttributes uint16) bool {
	ret, _, _ := procSetConsoleTextAttribute.Call(
		hConsoleOutput,
		uintptr(wAttributes))
	return ret != 0
}

type coord struct {
	X, Y int16
}

type smallRect struct {
	Left, Top, Right, Bottom int16
}

type consoleScreenBufferInfo struct {
	DwSize              coord
	DwCursorPosition    coord
	WAttributes         uint16
	SrWindow            smallRect
	DwMaximumWindowSize coord
}

func getConsoleScreenBufferInfo(hConsoleOutput uintptr) *consoleScreenBufferInfo {
	var csbi consoleScreenBufferInfo
	if ret, _, _ := procGetConsoleScreenBufferInfo.Call(hConsoleOutput, uintptr(unsafe.Pointer(&csbi))); ret == 0 {
		return nil
	}
	return &csbi
}

const (
	stdOutputHandle = uint32(-11 & 0xFFFFFFFF)
)

func init() {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")

	procGetStdHandle = kernel32.NewProc("GetStdHandle")

	hStdout, _, _ = procGetStdHandle.Call(uintptr(stdOutputHandle))

	initScreenInfo = getConsoleScreenBufferInfo(hStdout)

	_, _ = syscall.LoadDLL("")
}

func resetColor() {
	if initScreenInfo == nil { // No console info - Ex: stdout redirection
		return
	}
	setConsoleTextAttribute(hStdout, initScreenInfo.WAttributes)
}

func changeColor(fg Color, fgBright bool, bg Color, bgBright bool) {
	attr := uint16(0)
	if fg == None || bg == None {
		cbufinfo := getConsoleScreenBufferInfo(hStdout)
		if cbufinfo == nil { // No console info - Ex: stdout redirection
			return
		}
		attr = cbufinfo.WAttributes
	}
	if fg != None {
		attr = attr & ^foregroundMask | fgColors[fg]
		if fgBright {
			attr |= foregroundIntensity
		}
	}
	if bg != None {
		attr = attr & ^backgroundMask | bgColors[bg]
		if bgBright {
			attr |= backgroundIntensity
		}
	}
	setConsoleTextAttribute(hStdout, attr)
}
