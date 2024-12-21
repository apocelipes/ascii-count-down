package render

import (
	"fmt"
	"os"
	"time"

	"github.com/apocelipes/ascii-count-down/pkg/char"
	"github.com/apocelipes/ascii-count-down/pkg/util"
)

type SharpCharRender struct {
	chars [][]string
}

func NewSharpCharRender() *SharpCharRender {
	return &SharpCharRender{
		chars: make([][]string, 0, 6),
	}
}

func (sr *SharpCharRender) RenderContent(duration time.Duration) {
	sr.chars = char.ConvertToChars(duration, char.SharpChars, sr.chars)
	// 等宽，不需要清除整行
	for i := 0; i < char.MaxSharpCharHeight(); i++ {
		fmt.Print(sr.chars[0][i])
		fmt.Print(" ")
		fmt.Print(sr.chars[1][i])
		fmt.Print("   ")
		fmt.Print(char.SharpChars[char.SharpColonIdx][i])
		fmt.Print("   ")
		fmt.Print(sr.chars[2][i])
		fmt.Print(" ")
		fmt.Print(sr.chars[3][i])
		fmt.Print("   ")
		fmt.Print(char.SharpChars[char.SharpColonIdx][i])
		fmt.Print("   ")
		fmt.Print(sr.chars[4][i])
		fmt.Print(" ")
		fmt.Print(sr.chars[5][i])
		fmt.Print("\n")
	}
	sr.chars = sr.chars[:0]
}

func (sr *SharpCharRender) RenderFlashing() {
	util.CursorDownForward(1, 20)
	fmt.Print("  ")
	util.CursorForward(23)
	fmt.Print("  ")
	util.CursorDownForward(4, 20)
	fmt.Print("  ")
	util.CursorForward(23)
	fmt.Print("  ")
	// move to bottom
	util.CursorDown(2)
}

func (sr *SharpCharRender) CanRender(_ time.Duration) error {
	// 00:00:00, 6 digits, 2 colons, digit space width 1
	maxWidth := char.MaxSharpCharWidth()*6 + 8*2 + 3
	maxHeight := char.MaxSharpCharHeight() + 2 // +2 for the prompt
	return util.CheckTerminal(int(os.Stdout.Fd()), maxWidth, maxHeight)
}

func (sr *SharpCharRender) RenderHeight() int {
	return char.MaxSharpCharHeight()
}
