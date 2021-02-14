package keyboard

import (
	"runtime"
	"time"
)

func ExampleNewKeyBonding() {
	kb, err := NewKeyBonding()
	if err != nil {
		panic(err)
	}

	// For linux, it is very important to wait 2 seconds
	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}

	// Select keys to be pressed
	kb.SetKeys(VirtualKey["VK_A"], VirtualKey["VK_B"])

	// Set shift to be pressed
	kb.HasSHIFT(true)

	// Press the selected keys
	err = kb.Launching()
	if err != nil {
		panic(err)
	}

	// Or you can use Press and Release
	kb.Press()
	time.Sleep(10 * time.Millisecond)
	kb.Release()

	// Here, the program will generate "ABAB" as if they were pressed on the keyboard.
	// Output:
}
