// +build windows

package prutils

import (
	"syscall"
	"unsafe"
)

// COORD for win COORD
type COORD struct {
	X int16
	Y int16
}

// SmallRect is for win SMALL_RECT
type SmallRect struct {
	Left   int16
	Top    int16
	Right  int16
	Bottom int16
}

func printConsoleTitle(title string) (int, error) {
	handle, err := syscall.LoadLibrary("Kernel32.dll")
	if err != nil {
		return 0, err
	}
	proc, err := syscall.GetProcAddress(handle, "SetConsoleTitleW")
	if err != nil {
		return 0, err
	}
	ptr, _ := syscall.UTF16PtrFromString(title)
	r, _, err := syscall.Syscall(proc, 1, uintptr(unsafe.Pointer(ptr)), 0, 0)
	_ = syscall.FreeLibrary(handle)
	return int(r), err
}

// setConsoleRectWin implements rect change on win platform
func setConsoleRect(width, height int16) {
	var (
		kernel32DLL = syscall.NewLazyDLL("kernel32.dll")

		setConsoleScreenBufferSizeProc = kernel32DLL.NewProc("SetConsoleScreenBufferSize")
		setConsoleWindowInfoProc       = kernel32DLL.NewProc("SetConsoleWindowInfo")
		stdoutForRect, _               = syscall.GetStdHandle(syscall.STD_OUTPUT_HANDLE)
	)
	coord := COORD{
		X: width,
		Y: height,
	}
	rect := SmallRect{
		Left:   0,
		Top:    0,
		Right:  width - 1,
		Bottom: height - 1,
	}
	// SetConsoleScreenBufferSize
	_, _, _ = setConsoleScreenBufferSizeProc.Call(uintptr(stdoutForRect), coordToPointer(coord))
	use(coord)
	// SetConsoleWindowInfo
	// SetConsoleWindowInfo sets the size and position of the console screen buffer's window.
	// Note that the size and location must be within and no larger than the backing console screen buffer.
	// See https://msdn.microsoft.com/en-us/library/windows/desktop/ms686125(v=vs.85).aspx.
	_, _, _ = setConsoleWindowInfoProc.Call(uintptr(stdoutForRect), uintptr(boolToBOOL(IsAbsolute)), uintptr(unsafe.Pointer(&rect)))
	use(IsAbsolute)
	use(rect)
}

// use for avoiding pointer loss from garbage collector
func use(p interface{}) {}

// boolToBOOL for windows specific BOOL
func boolToBOOL(f bool) int32 {
	if f {
		return int32(1)
	} else {
		return int32(0)
	}
}

// coordToPointer for coord to uint pointer
func coordToPointer(c COORD) uintptr {
	// Note: This code assumes the two SHORTs are correctly laid out; the "cast" to uint32 is just to get a pointer to pass.
	return uintptr(*((*uint32)(unsafe.Pointer(&c))))
}