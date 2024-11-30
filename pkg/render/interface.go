package render

import "time"

type TerminalRender interface {
	RenderContent(duration time.Duration)
	RenderFlashing()
	CanRender() error
	RenderHeight() int
}
