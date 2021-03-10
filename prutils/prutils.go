package prutils

import (
	"bufio"
	"fmt"
	"github.com/PrabhakarRai/goutils/clipboard"
	"github.com/PrabhakarRai/goutils/keyboard"
	"math/rand"
	"os"
	"runtime"
	"time"
)

// IsAbsolute is used to set position of cmd in windows
var IsAbsolute = true

// init seeds the randomInRange function
func init() {
	rand.Seed(time.Now().UnixNano())
}

// ClipboardCheck Tests if program is capable of clipboard read and write
func ClipboardCheck() {
	fmt.Println("[+] Running Clipboard check !")
	_ = clipboard.WriteAll("NT Clipboard Test")
	text, err := clipboard.ReadAll()
	if clipboard.Unsupported || text != "NT Clipboard Test" {
		fmt.Println("[!] Clipboard functionality not working !")
		fmt.Println("[!] Please refer to Readme file !")
		panic(err)
	}
}

// KeyboardCheck initializes new keyboard.Keybonding
func KeyboardCheck() *keyboard.KeyBonding {
	fmt.Println("[+] Running Key Press Simulation check !")
	keyBounding, err := keyboard.NewKeyBonding()
	if err != nil {
		fmt.Println("[!] Keypress functionality not working !")
		fmt.Println("[!] Please refer to Readme file !")
		panic(err)
	}
	// For linux, it is very important to wait 2 seconds
	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}
	return &keyBounding
}

// IntInputValidator Prints msg, stores int input in storage,
// validates if in lower upper range, breaks
func IntInputValidator(msg string, storage *int, lowerLimit int, upperLimit int) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s", msg)
		_, err := fmt.Fscan(reader, storage)
		if err == nil && (*storage >= lowerLimit && *storage <= upperLimit) {
			break
		}
		_, _ = reader.ReadString('\n')
		fmt.Println("[-] Please provide a valid numeric input.")
	}
}

// BoolInputValidator Prints msg, stores int input temporarily, extracts bool value, breaks
func BoolInputValidator(msg string, storage *bool) {
	var temp int
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s", msg)
		_, err := fmt.Fscan(reader, &temp)
		if err == nil && (temp >= 0 && temp <= 1) {
			switch temp {
			case 0:
				*storage = false
			case 1:
				*storage = true
			}
			break
		}
		_, _ = reader.ReadString('\n')
		fmt.Println("[-] Please provide a valid 0 or 1 input.")
	}
}

// RandomInRange return a int between upper and lower (both inclusive)
func RandomInRange(upper, lower int) int {
	if upper < lower {
		upper, lower = lower, upper
	}
	return (rand.Int() % (upper - lower + 1)) + lower
}

// PrintConsoleTitle sets the console title of the application
func PrintConsoleTitle(title string) {
	// Avoiding blank titles for further errors
	if title == "" {
		title = "~ By Madcap Hacker"
	}
	_, _ = printConsoleTitle(title)
}

// SetConsoleRect takes width and height and changes current rects
// of current console
func SetConsoleRect(width, height int16) {
	setConsoleRect(width, height)
}
