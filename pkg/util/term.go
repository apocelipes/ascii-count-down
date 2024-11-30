package util

import "fmt"

func CursorUp(line int) {
	fmt.Printf("\033[%dF", line)
}

func CursorDown(line int) {
	fmt.Printf("\033[%dE", line)
}

func CursorForward(line int) {
	fmt.Printf("\033[%dC", line)
}

func CursorEraseLine() {
	fmt.Print("\033[2K")
}
