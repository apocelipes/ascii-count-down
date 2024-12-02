package util

import (
	"errors"
	"fmt"
	"os"

	"golang.org/x/term"
)

func CursorUp(line int) {
	fmt.Printf("\033[%dF", line)
}

func CursorDown(line int) {
	fmt.Printf("\033[%dE", line)
}

func CursorForward(line int) {
	fmt.Printf("\033[%dC", line)
}

func CursorDownForward(line, char int) {
	fmt.Printf("\033[%dE\033[%dC", line, char)
}

func CursorEraseEntireLine() {
	fmt.Print("\033[2K")
}

func CursorHide() {
	fmt.Print("\033[?25l")
}

func CursorShow() {
	fmt.Print("\033[?25h")
}

func CheckTerminal(fd, maxWidth, maxHeight int) error {
	if !term.IsTerminal(int(os.Stdout.Fd())) {
		return errors.New("output should be a terminal")
	}
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return err
	}

	if maxWidth > width {
		return fmt.Errorf("no enough width, got: %d, want: %d", width, maxWidth)
	}
	if maxHeight > height {
		return fmt.Errorf("no enough height, got: %d, want: %d", height, maxHeight)
	}
	return nil
}

func PrintlnRed(s string) {
	fmt.Printf("\033[31m%s\033[0m\n", s)
}
