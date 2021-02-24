// +build !windows

package prutils

import (
	"fmt"
	"os"
	"strconv"
)

func printConsoleTitle(title string) (int, error)  {
	return fmt.Fprintf(os.Stdout, "\033]0;%s\007", title)
}

func setConsoleRect(width, height int16) {
	_, _ = fmt.Fprintf(os.Stdout, "\033[8;%s;%st", strconv.Itoa(int(height)), strconv.Itoa(int(width)))
}
