package render

import (
	"fmt"
	"os"
	"time"

	"github.com/apocelips/ascii-count-down/pkg/char"
	"github.com/apocelips/ascii-count-down/pkg/util"
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

func (ar *ASCIIArtCharRender) CanRender() error {
	// 00:00:00, 6 digits, 2 colons, digit space width 1
	maxWidth := char.MaxASCIIArtCharWidth()*6 + 7*2 + 3
	maxHeight := char.MaxASCIIArtCharHeight()
	return util.CheckTerminal(int(os.Stdout.Fd()), maxWidth, maxHeight)
}

func (ar *ASCIIArtCharRender) RenderHeight() int {
	return char.MaxASCIIArtCharHeight()
}
