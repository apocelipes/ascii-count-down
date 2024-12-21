package render

import (
	"fmt"
	"os"
	"time"

	"github.com/apocelipes/ascii-count-down/pkg/char"
	"github.com/apocelipes/ascii-count-down/pkg/util"
)

type ASCIIArtCharRender struct {
	chars [][]string
}

func NewASCIIArtCharRender() *ASCIIArtCharRender {
	return &ASCIIArtCharRender{
		chars: make([][]string, 0, 6),
	}
}

func (ar *ASCIIArtCharRender) RenderContent(duration time.Duration) {
	if len(ar.chars) > 0 {
		ar.chars = ar.chars[:0]
	}
	ar.chars = char.ConvertToChars(duration, char.ASCIIArtChars, ar.chars)
	for i := 0; i < char.MaxASCIIArtCharHeight(); i++ {
		util.CursorEraseEntireLine()
		fmt.Print(ar.chars[0][i])
		fmt.Print(" ")
		fmt.Print(ar.chars[1][i])
		fmt.Print("  ")
		fmt.Print(char.ASCIIArtChars[char.ASCIIArtColonIdx][i])
		fmt.Print("  ")
		fmt.Print(ar.chars[2][i])
		fmt.Print(" ")
		fmt.Print(ar.chars[3][i])
		fmt.Print("  ")
		fmt.Print(char.ASCIIArtChars[char.ASCIIArtColonIdx][i])
		fmt.Print("  ")
		fmt.Print(ar.chars[4][i])
		fmt.Print(" ")
		fmt.Print(ar.chars[5][i])
		fmt.Print("\n")
	}
}

func (ar *ASCIIArtCharRender) RenderFlashing() {
	util.CursorDownForward(1, 3+len(ar.chars[0][0])+1+len(ar.chars[1][0]))
	fmt.Print(" ")
	util.CursorForward(3 + len(ar.chars[2][0]) + 1 + len(ar.chars[3][0]) + 3)
	fmt.Print(" ")
	util.CursorDownForward(1, 2+len(ar.chars[0][0])+1+len(ar.chars[1][0]))
	fmt.Print("   ")
	util.CursorForward(2 + len(ar.chars[2][0]) + 1 + len(ar.chars[3][0]) + 2)
	fmt.Print("   ")

	util.CursorDownForward(2, 3+len(ar.chars[0][0])+1+len(ar.chars[1][0]))
	fmt.Print(" ")
	util.CursorForward(3 + len(ar.chars[2][0]) + 1 + len(ar.chars[3][0]) + 3)
	fmt.Print(" ")
	util.CursorDownForward(1, 2+len(ar.chars[0][0])+1+len(ar.chars[1][0]))
	fmt.Print("   ")
	util.CursorForward(2 + len(ar.chars[2][0]) + 1 + len(ar.chars[3][0]) + 2)
	fmt.Print("   ")
	// move to bottom
	util.CursorDown(1)
}

const asciiArtNonDigitTotalWidth = 7*2 + 3

func calcNumWidth(num int, frontZero bool) int {
	if !frontZero {
		return char.MaxASCIIArtCharWidth(5) + char.MaxASCIIArtCharWidth(9)
	}

	if num >= 10 {
		return char.MaxASCIIArtCharWidth(int(num)/10) + char.MaxASCIIArtCharWidth(9)
	}
	return char.MaxASCIIArtCharWidth(0) + char.MaxASCIIArtCharWidth(int(num)%10)
}

func asciiArtTotalWidth(d time.Duration) int {
	maxWidth := 0
	hours := d / time.Hour
	maxWidth += calcNumWidth(int(hours), true)
	allZero := hours == 0

	d -= hours * time.Hour
	minutes := d / time.Minute
	maxWidth += calcNumWidth(int(minutes), allZero)
	allZero = allZero && minutes == 0

	seconds := d / time.Second
	maxWidth += calcNumWidth(int(seconds), allZero)

	// 00:00:00, 6 digits, 2 colons, digit space width 1
	return maxWidth + asciiArtNonDigitTotalWidth
}

func (ar *ASCIIArtCharRender) CanRender(d time.Duration) error {
	maxWidth := asciiArtTotalWidth(d)
	maxHeight := char.MaxASCIIArtCharHeight() + 2 // +2 for the prompt
	return util.CheckTerminal(int(os.Stdout.Fd()), maxWidth, maxHeight)
}

func (ar *ASCIIArtCharRender) RenderHeight() int {
	return char.MaxASCIIArtCharHeight()
}
